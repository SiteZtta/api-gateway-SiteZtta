package app

import (
	"api-gateway-SiteZtta/cfg"
	"api-gateway-SiteZtta/internal/transport/http"
	"context"
	"errors"
	"log/slog"
	HTTP "net/http"
	"time"
)

type App struct {
	Server *http.Server
	log    *slog.Logger
}

func New(cfg cfg.Config, log *slog.Logger) *App {
	return &App{
		Server: http.New(cfg, log),
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
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := a.Server.Shutdown(shutdownCtx); err != nil {
			return err
		}

		err := <-errch
		if err != nil && !errors.Is(err, HTTP.ErrServerClosed) {
			return err
		}

		return nil
	}
}
