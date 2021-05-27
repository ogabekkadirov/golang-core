package server

import (
	"context"
	"net/http"
	"os"
)


type Server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler) *Server{
	return &Server{
		httpServer: &http.Server{
			Addr: os.Getenv("PORT"),
			Handler: handler,
			// ReadTimeout: cfg.HTTP.ReadTimeout,
			// WriteTimeout: cfg.HTTP.WriteTimeout,
			// MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes,
		},
	}
}

func (s *Server) Run() error{
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context)	error{
	return s.httpServer.Shutdown(ctx)
}