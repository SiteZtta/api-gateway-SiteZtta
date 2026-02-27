package http

import (
	"api-gateway-SiteZtta/config"
	"log/slog"

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

func (r *Router) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.registerV1(router)
	router.GET("/auth/test", r.h.userIdentity, r.h.adminIdentity, r.h.testAuth)
	return router
}

func (r *Router) registerV1(router *gin.Engine) {
	authV1 := router.Group("/auth/v1")
	{
		authV1.POST("/sign-up", r.h.signUpV1)
		authV1.POST("/sign-in", r.h.signInV1)
	}
}
