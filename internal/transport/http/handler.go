package http

import (
	"log/slog"
)

type Handler struct {
	// TODO: to inject services
	log *slog.Logger
}

func NewHandler() *Handler {
	return &Handler{}
}
