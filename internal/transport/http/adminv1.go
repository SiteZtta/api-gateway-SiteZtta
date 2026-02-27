package http

import (
	"api-gateway-SiteZtta/domain/user"
	"api-gateway-SiteZtta/internal/transport/http/errorresponse"
	"api-gateway-SiteZtta/internal/transport/http/v1/dto"
	"api-gateway-SiteZtta/pkg/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrGettingUser    = fmt.Errorf("error getting user")
	ErrGettingUserCtx = fmt.Errorf("error getting user context")
)

// adminCabinet godoc
// @Summary Admin personal cabinet for managing the site
// @Tags admin
// @Security ApiKeyAuth
// @Router /api/v1/admin [get]
func (h *Handler) adminCabinetV1(c *gin.Context) {
	fn := "api-gateway-SiteZtta.internal.transport.http.adminCabinetV1"
	log := h.log.With("fn", fn)
	authInfoCtx, ok := c.Get(userCtx)
	if !ok {
		log.Error("error getting auth info:", logger.Err(ErrGettingUserCtx))
		errorresponse.NewErrorResponse(c, http.StatusInternalServerError, ErrGettingUserCtx.Error())
		return
	}
	authInfo, ok := authInfoCtx.(*user.AuthInfo)
	if !ok {
		log.Error("authInfoCtx is not user.AuthInfo:", logger.Err(ErrGettingUser))
		errorresponse.NewErrorResponse(c, http.StatusInternalServerError, ErrGettingUser.Error())
		return
	}
	c.JSON(http.StatusOK, &dto.AuthInfoResponse{
		UserId:   authInfo.UserId,
		Role:     authInfo.Role,
		Username: authInfo.Username,
	})
}
