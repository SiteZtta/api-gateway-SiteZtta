package user

type AuthInfo struct {
	UserId int64 `json:"userId"`
	Role   Role  `json:"role"`
}
