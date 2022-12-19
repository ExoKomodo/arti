package api

import (
	"fmt"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
)

type ApiMethod struct {
	Route   ApiRoute
	Handler func(http.ResponseWriter, *http.Request)
	Method  string
}

type ApiRoute = string

type Controller interface {
	Routes() []*ApiMethod
	BaseUrl() string
}

func Register(controller Controller, router *chi.Mux) error {
	for _, method := range controller.Routes() {
		route := path.Join("/api", controller.BaseUrl(), method.Route)
		switch method.Method {
		case http.MethodDelete:
			router.Delete(route, method.Handler)
		case http.MethodGet:
			router.Get(route, method.Handler)
		case http.MethodPost:
			router.Post(route, method.Handler)
		case http.MethodPatch:
			router.Patch(route, method.Handler)
		default:
			return fmt.Errorf("failed to register route: %s %s", method.Method, method.Route)
		}
	}
	return nil
}

type ArtifactsController struct{}

type RootController struct{}
