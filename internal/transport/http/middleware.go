package http

import "github.com/gin-gonic/gin"

const (
	authorizationHeader = "Authorization"
	userCtx             = "TokenClaims"
)

func (h *Handler) userIdentity(c *gin.Context) {
}

func (h *Handler) adminIdentity(c *gin.Context) {
}

func (h *Handler) getTokenClaims(c *gin.Context) {
}
