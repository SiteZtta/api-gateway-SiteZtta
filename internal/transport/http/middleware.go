package http

import (
	"api-gateway-SiteZtta/domain/user"
	"api-gateway-SiteZtta/internal/transport/http/errorresponse"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "TokenClaims"
)

func (h *Handler) userIdentity(c *gin.Context) {
	const fn = "api-gateway-SiteZtta.internal.transport.http.middleware.userIdentity"
	log := h.log.With("fn", fn)
	tokeClaims := h.getTokenClaims(c)
	if tokeClaims.UserId == 0 {
		log.Error("token claims is empty")
		errorresponse.NewErrorResponse(c, http.StatusUnauthorized, "Authentication is required")
		return
	}
	c.Set(userCtx, tokeClaims)
	c.Next()
}

// without repeateble parsing of token
func (h *Handler) adminIdentity(c *gin.Context) {
	const fn = "api-gateway-SiteZtta.internal.transport.http.middleware.adminIdentity"
	log := h.log.With("fn", fn)
	claims, _ := c.Get(userCtx)
	if claims == nil {
		log.Error("token claims is empty")
		errorresponse.NewErrorResponse(c, http.StatusUnauthorized, "Authentication is required")
		return
	}
	tokenClaims, ok := claims.(user.AuthInfo)
	if !ok {
		log.Error("token claims is not user.TokenClaims")
		errorresponse.NewErrorResponse(c, http.StatusUnauthorized, "Authentication is required")
		return
	}
	if tokenClaims.Role != user.Admin {
		log.Error("token claims is not admin")
		errorresponse.NewErrorResponse(c, http.StatusForbidden, "Access denied: Admin access only")
		return
	}
	c.Set(userCtx, tokenClaims)
	c.Next()
}

func (h *Handler) getTokenClaims(c *gin.Context) user.AuthInfo {
	// TODO: to implement grpc method func (IAuthorization) ParseToken(token string) (auth.TokenClaims, error)
	return user.AuthInfo{}
}
