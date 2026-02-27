package http

import (
	"api-gateway-SiteZtta/config"
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

func NewServer(config config.Config, log *slog.Logger) *Server {
	server := &Server{
		log: log,
		HttpServer: &http.Server{
			Addr:        net.JoinHostPort(config.HttpServer.Host, fmt.Sprint(config.HttpServer.Port)),
			ReadTimeout: config.HttpServer.Timeout,
			IdleTimeout: config.HttpServer.IdleTimeout,
			Handler:     NewRouter(log, config).InitRoutes(config),
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
