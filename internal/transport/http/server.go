package http

import (
	"api-gateway-SiteZtta/cfg"
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
)

type Server struct {
	HttpServer *http.Server
	log        *slog.Logger
}

func NewServer(config cfg.Config, log *slog.Logger) *Server {
	server := &Server{
		log: log,
		HttpServer: &http.Server{
			Addr:        net.JoinHostPort(config.HttpServer.Address, fmt.Sprint(config.HttpServer.Port)),
			ReadTimeout: config.HttpServer.Timeout,
			IdleTimeout: config.HttpServer.IdleTimeout,
			Handler:     NewRouter(log).InitRoutes(),
		},
	}
	return server
}

func (s *Server) Start() error {
	return s.HttpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}
