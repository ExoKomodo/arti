package oidc

import (
	"arti/server"

	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var (
	clientID = "XXX.apps.googleusercontent.com"
)

type OIDC struct {
	Config   oauth2.Config
	Server   server.Server
	verifier *oidc.IDTokenVerifier
	provider *oidc.Provider
}

func New(srv server.Server) *OIDC {
	paths := []string{"auth.google.clientId", "auth.google.clientSecret"}
	server.RequiredConfig(paths, srv.Log)

	provider, err := oidc.NewProvider(srv.Ctx, "https://accounts.google.com")
	if err != nil {
		srv.Log.Fatal().Err(err).Msg("Could not create provider")
	}

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}

	// Configure an OpenID Connect aware OAuth2 client.
	oauth2Config := oauth2.Config{
		ClientID:     viper.GetString("auth.google.clientId"),
		ClientSecret: viper.GetString("auth.google.clientSecret"),
		RedirectURL:  "http://127.0.0.1:5555/auth/google/callback",

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}
	o := &OIDC{Config: oauth2Config, Server: srv, verifier: provider.Verifier(oidcConfig), provider: provider}
	o.registerURLs()
	return o
}

func (o *OIDC) registerURLs() {
	o.Server.Routes.Get("/auth", o.routeAuth)
	o.Server.Routes.Get("/auth/{provider}/callback", o.routeCallback)
}

func (o *OIDC) routeCallback(w http.ResponseWriter, r *http.Request) {
	state, err := r.Cookie("state")
	sv := "empty"
	if err == nil {
		sv = state.Value
	}
	/*if err != nil {
		http.Error(w, "state not found", http.StatusBadRequest)
		return
	}
	if r.URL.Query().Get("state") != state.Value {
		http.Error(w, "state did not match", http.StatusBadRequest)
		return
	}*/
	o.Server.Log.Debug().Str("cookie", sv).Str("query", r.URL.Query().Get("state")).Msg("State")

	oauth2Token, err := o.Config.Exchange(o.Server.Ctx, r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}
	idToken, err := o.verifier.Verify(o.Server.Ctx, rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	nonce, err := r.Cookie("nonce")
	nv := "empty"
	if err == nil {
		nv = nonce.Value
	}
	/*if err != nil {
		http.Error(w, "nonce not found", http.StatusBadRequest)
		return
	}
	if idToken.Nonce != nonce.Value {
		http.Error(w, "nonce did not match", http.StatusBadRequest)
		return
	}*/
	o.Server.Log.Debug().Str("cookie", nv).Str("query", r.URL.Query().Get("nonce")).Msg("Nonce")

	userInfo, err := o.provider.UserInfo(o.Server.Ctx, oauth2.StaticTokenSource(oauth2Token))
	if err != nil {
		http.Error(w, "Failed to get userinfo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// oauth2Token.AccessToken = "*REDACTED*"

	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
		UserInfo      *oidc.UserInfo
	}{oauth2Token, new(json.RawMessage), userInfo}

	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (o *OIDC) routeAuth(w http.ResponseWriter, r *http.Request) {
	state, err := randString(16)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	nonce, err := randString(16)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	setCallbackCookie(w, r, "state", state)
	setCallbackCookie(w, r, "nonce", nonce)

	http.Redirect(w, r, o.Config.AuthCodeURL(state, oidc.Nonce(nonce)), http.StatusFound)
}

func randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func setCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	c := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}
