package user

type Role int32

const (
	Unspecified Role = iota
	User
	Admin
)

func (r Role) String() string {
	switch r {
	case User:
		return "user"
	case Admin:
		return "admin"
	default:
		return "unknown"
	}
}
