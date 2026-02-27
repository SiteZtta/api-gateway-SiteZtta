package http

import (
	"api-gateway-SiteZtta/domain/user"
	"api-gateway-SiteZtta/internal/transport/http/errorresponse"
	"api-gateway-SiteZtta/pkg/logger"
	"fmt"
	"net/http"
	"strings"

	sitezttav1 "github.com/SiteZtta/protos-SiteZtta/gen/go/auth"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "AuthInfo"
)

func (h *Handler) userIdentity(c *gin.Context) {
	const fn = "api-gateway-SiteZtta.internal.transport.http.middleware.userIdentity"
	log := h.log.With("fn", fn)
	authInfo, err := h.getAuthInfo(c)
	if err != nil {
		log.Error("error getting authInfo", logger.Err(err))
		errorresponse.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if authInfo == nil {
		log.Error("authInfo is empty")
		errorresponse.NewErrorResponse(c, http.StatusUnauthorized, "Authentication is required")
		return
	}
	c.Set(userCtx, authInfo)
	c.Next()
}

// without repeateble parsing of token
func (h *Handler) adminIdentity(c *gin.Context) {
	const fn = "api-gateway-SiteZtta.internal.transport.http.middleware.adminIdentity"
	log := h.log.With("fn", fn)
	claims, _ := c.Get(userCtx)
	if claims == nil {
		log.Error("authInfo is empty")
		errorresponse.NewErrorResponse(c, http.StatusUnauthorized, "Authentication is required")
		return
	}
	authInfo, ok := claims.(*user.AuthInfo)
	if !ok {
		log.Error("token claims is not user.AuthInfo")
		errorresponse.NewErrorResponse(c, http.StatusUnauthorized, "Authentication is required")
		return
	}
	if authInfo.Role != user.Admin {
		log.Error("token claims is not admin")
		errorresponse.NewErrorResponse(c, http.StatusForbidden, "Access denied: Admin access only")
		return
	}
	c.Next()
}

func (h *Handler) getAuthInfo(c *gin.Context) (*user.AuthInfo, error) {
	token := c.GetHeader(authorizationHeader)
	fmt.Printf("token: %s\n", token)
	if token == "" {
		return nil, fmt.Errorf("missing authorization header")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	req := &sitezttav1.TokenRequest{Token: token}
	// Business logic
	resp, err := h.AuthServiceClient.ValidateToken(c, req)
	if err != nil {
		return nil, err
	}
	fmt.Printf("protobuf authInfo: %v\n", resp)
	authInfo := &user.AuthInfo{
		UserId: resp.UserId,
		Role:   user.Role(resp.Role),
	}
	fmt.Printf("authInfo: %v\n", authInfo)
	return authInfo, nil
}

// @Summary Test auth
// @Security ApiKeyAuth
// @Router /auth/test [get]
func (h *Handler) testAuth(c *gin.Context) {
	user := c.MustGet(userCtx)
	c.JSON(200, user)
}
