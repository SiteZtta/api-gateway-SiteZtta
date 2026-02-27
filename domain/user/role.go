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
		return "USER"
	case Admin:
		return "ADMIN"
	default:
		return "UNSPECIFIED"
	}
}
