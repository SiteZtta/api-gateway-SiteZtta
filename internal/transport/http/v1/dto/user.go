package dto

import "api-gateway-SiteZtta/domain/user"

type SignUpRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,e164"`
	Password string `json:"password" validate:"required,min=8"`
}

type SignInRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserIdResponse struct {
	UserId int64 `json:"userId"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type AuthInfoResponse struct {
	UserId   int64     `json:"userId"`
	Role     user.Role `json:"role"`
	Username string    `json:"username"`
}
