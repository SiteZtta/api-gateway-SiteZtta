package app

import (
	"api-gateway-SiteZtta/config"
	TransHttp "api-gateway-SiteZtta/internal/transport/http"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

type App struct {
	Server *TransHttp.Server
	log    *slog.Logger
}

func New(cfg config.Config, log *slog.Logger) *App {
	return &App{
		Server: TransHttp.NewServer(cfg, log),
		log:    log,
	}
}

func (a *App) Run(ctx context.Context) error {
	errch := make(chan error, 1)

	go func() {
		errch <- a.Server.Start()
	}()

	select {
	case err := <-errch:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // parent ctx is already cancelled
		defer cancel()
		// Shutdown gracefully with a 5-second timeout (server stops acepting new connections and waits for 5 seconds for existing connections to finish)
		if err := a.Server.Shutdown(shutdownCtx); err != nil {
			return err
		}
		if err := <-errch; err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	}
}
