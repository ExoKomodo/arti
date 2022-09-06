package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/rs/zerolog"
)

type Server struct {
	Log    zerolog.Logger
	Routes http.Handler
	Ctx    context.Context
}

func New() Server {
	s := Server{}
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger := httplog.NewLogger("leadster", httplog.Options{
		// JSON: true,
		Concise: true,
		// Tags: map[string]string{
		// 	"version": "v1.0-81aa4244d9fc8076a",
		// 	"env":     "dev",
		// },
	})
	s.Log = logger
	s.Log.Info().Int("pid", os.Getpid()).Int("uid", os.Getuid()).Int("gid", os.Getgid()).Msg("Server started")

	timeout := time.Duration(10) * time.Second  // TODO
	s.Routes = service(s.Log)
	srv := &http.Server{
		Addr:    "127.0.0.1:5555",
		Handler: s.Routes,
		// ReadHeaderTimeout is here as well
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		IdleTimeout:  timeout,
	}
	

	shutdownCtx := s.gracefullShutdown(srv)
	s.Ctx = shutdownCtx

	go func() {
		// if err := srv.ListenAndServeTLS(certFile, keyFile); err != nil {
		if err := srv.ListenAndServe(); err != nil {
			s.Log.Error().Err(err).Msg("Server stopped")
			s.triggerShutdown(shutdownCtx, srv)
		}
	}()

	return s
}

func (s *Server) gracefullShutdown(server *http.Server) context.Context {
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		// Shutdown signal with grace period of 30 seconds TODO
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				s.Log.Fatal().Msgf("graceful shutdown timed out.. forcing exit.")
			}
		}()

		s.triggerShutdown(shutdownCtx, server)
		serverStopCtx()
	}()
	return serverCtx
}

func (s *Server) triggerShutdown(ctx context.Context, server *http.Server) {
	err := server.Shutdown(ctx)
	if err != nil {
		s.Log.Error().Stack().Msgf("error shutting down server (%s): %v", server.Addr, err)
		err = server.Close()
		if err != nil {
			s.Log.Error().Stack().Msgf("error closing server (%s): %v", server.Addr, err)
		}
	}
}

func service(logger zerolog.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/ping"))
	r.Use(httplog.RequestLogger(logger))

	r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		oplog := httplog.LogEntry(r.Context())
		w.Header().Add("Content-Type", "text/plain")
		oplog.Info().Msg("info here")
		w.Write([]byte("info here"))
	})

	return r
}
