package http

import "github.com/gin-gonic/gin"

// signUp 			godoc
// @Summary			SignUpV1
// @Tags 			auth
// @Description 	Create a new user
// @ID 		     	sign-up
// @Accept     		json
// @Produce      	json
// @param input body v1.SignUpRequst true "account info"
// @Success 200 {integer} integer "id"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/v1/sign-up/ [post]
func (h *Handler) signUpV1(c *gin.Context) {
}

// signIn 			godoc
// @Summary			SignInV1
// @Tags 			auth
// @Description 	Sign in user in system
// @ID 		     	sign-in
// @Accept     		json
// @Produce      	json
// @param input body v1.SignInRequst true "account info"
// @Success 200 {integer} integer "id"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/v1/sign-in/ [post]
func (h *Handler) signInV1(c *gin.Context) {
}
