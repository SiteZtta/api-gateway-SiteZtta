package dto

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserIdResponse struct {
	UserId string `json:"userId"`
}

type SignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
