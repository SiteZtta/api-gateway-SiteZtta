package http

import (
	"api-gateway-SiteZtta/internal/transport/http/errorresponse"
	"api-gateway-SiteZtta/internal/transport/http/v1/dto"
	"api-gateway-SiteZtta/pkg/logger"
	"net/http"

	sitezttav2 "github.com/SiteZtta/protos-SiteZtta/gen/go/auth"
	"github.com/gin-gonic/gin"
)

// signUp 			godoc
// @Summary			SignUpV1
// @Tags 			auth
// @Description 	Create a new user
// @ID 		     	sign-up
// @Accept     		json
// @Produce      	json
// @param input body dto.SignUpRequest true "account info"
// @Success 200 {integer} integer "id"
// @Failure 400,404 {object} errorresponse.ErrorResponse
// @Failure 500 {object} errorresponse.ErrorResponse
// @Failure default {object} errorresponse.ErrorResponse
// @Router /auth/v1/sign-up/ [post]
func (h *Handler) signUpV1(c *gin.Context) {
	// validation
	var input dto.SignUpRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Error("error parsing input data:", logger.Err(err))
		errorresponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.validator.Struct(input); err != nil {
		h.log.Error("error validating input data:", logger.Err(err))
		errorresponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// Business logic
	req := &sitezttav2.SignUpRequest{
		UserName: input.Username,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: input.Password,
	}
	resp, err := h.AuthServiceClient.CreateUser(c, req)
	if err != nil {
		h.log.Error("error registering user", logger.Err(err))
		errorresponse.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	uidResp := &dto.UserIdResponse{UserId: resp.UserId}
	c.JSON(http.StatusOK, uidResp)
}

// signIn 			godoc
// @Summary			SignInV1
// @Tags 			auth
// @Description 	Sign in user in system
// @ID 		     	sign-in
// @Accept     		json
// @Produce      	json
// @param input body dto.SignInRequest true "account info"
// @Success 200 {integer} integer "id"
// @Failure 400,404 {object} errorresponse.ErrorResponse
// @Failure 500 {object} errorresponse.ErrorResponse
// @Failure default {object} errorresponse.ErrorResponse
// @Router /auth/v1/sign-in/ [post]
func (h *Handler) signInV1(c *gin.Context) {
	// validation
	var input dto.SignInRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Error("error parsing input data:", logger.Err(err))
		errorresponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.validator.Struct(input); err != nil {
		h.log.Error("error validating input data:", logger.Err(err))
		errorresponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// Business logic
	req := &sitezttav2.SignInRequest{
		Login:    input.Login,
		Password: input.Password,
	}
	resp, err := h.AuthServiceClient.GenerateToken(c, req)
	if err != nil {
		h.log.Error("error generating token", logger.Err(err))
		errorresponse.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	uidResp := &dto.TokenResponse{Token: resp.Token}
	c.JSON(http.StatusOK, uidResp)
}
