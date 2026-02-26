package main

import (
	"api-gateway-SiteZtta/config"
	_ "api-gateway-SiteZtta/docs"
	"api-gateway-SiteZtta/internal/app"
	"api-gateway-SiteZtta/pkg/logger"
	"context"
	"log/slog"
)

// @title		Api-gateway-SiteZtta
// @version		1.0
// @description	Service for routing requests to different API endpoints and gateway.

// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @schemes		http https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	ctx := context.TODO()
	cfg := config.MustLoad("local")
	log := logger.SetupLogger(cfg.Env)
	log.Info("downloaded congig", slog.String("cfgEnv", cfg.Env), slog.Any("cfg", cfg))
	app := app.New(cfg, log)
	if err := app.Run(ctx); err != nil {
		log.Error("failed to run app", logger.Err(err))
	}
}
