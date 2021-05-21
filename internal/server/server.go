package server

import (
	"context"
	"golang-core/config"
	"net/http"
)


type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server{
	return &Server{
		httpServer: &http.Server{
			Addr: ":"+cfg.HTTP.Port,
			Handler: handler,
			ReadTimeout: cfg.HTTP.ReadTimeout,
			WriteTimeout: cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes,
		},
	}
}

func (s *Server) Run() error{
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context)	error{
	return s.httpServer.Shutdown(ctx)
}