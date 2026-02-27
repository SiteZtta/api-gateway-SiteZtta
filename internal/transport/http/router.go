package http

import (
	"api-gateway-SiteZtta/config"
	"fmt"
	"log/slog"
	"net"
	"net/url"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	h   *Handler
	log *slog.Logger
}

func NewRouter(log *slog.Logger, cfg config.Config) *Router {
	return &Router{
		log: log,
		h:   NewHandler(cfg, log),
	}
}

func (r *Router) InitRoutes(cfg config.Config) *gin.Engine {
	router := gin.New()
	addCors(router, cfg)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.registerV1(router)
	return router
}

func (r *Router) registerV1(router *gin.Engine) {
	authV1 := router.Group("/auth/v1")
	{
		authV1.POST("/sign-up", r.h.signUpV1)
		authV1.POST("/sign-in", r.h.signInV1)
	}
	api := router.Group("/api/v1")
	{
		api.GET("/admin", r.h.userIdentity, r.h.adminIdentity, r.h.adminCabinetV1)
	}
}

func addCors(router *gin.Engine, cfg config.Config) {
	// UHttpsUI := "https://" + addrUI
	uHttpUI := &url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort(cfg.Clients.UiService.Host, fmt.Sprint(cfg.Clients.UiService.Port)),
	}
	uHttpsUI := &url.URL{
		Scheme: "https",
		Host:   net.JoinHostPort(cfg.Clients.UiService.Host, fmt.Sprint(cfg.Clients.UiService.Port)),
	}
	fmt.Printf("addrUI %s\n", uHttpUI)
	router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{uHttpUI.String(), uHttpsUI.String()},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", authorizationHeader},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
}
