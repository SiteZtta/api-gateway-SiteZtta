package http

import (
	"api-gateway-SiteZtta/config"
	"api-gateway-SiteZtta/internal/clients/auth-service/grpc"
	"fmt"
	"log/slog"
	"net"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	AuthServiceClient *grpc.Client
	log               *slog.Logger
	validator         *validator.Validate
}

func NewHandler(cfg config.Config, log *slog.Logger) *Handler {
	addr := net.JoinHostPort(cfg.Clients.AuthService.Host, fmt.Sprint(cfg.Clients.AuthService.Port))
	log.Info("creating grpc client")
	authClient, err := grpc.New(addr, log, cfg.Clients.AuthService.Timeout, cfg.Clients.AuthService.RetriesCount)
	if err != nil {
		log.Error("failed to create auth service client", "error", err)
		return nil
	}
	return &Handler{AuthServiceClient: authClient, log: log, validator: validator.New()}
}
