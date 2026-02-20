package user

type Role int

const (
	User Role = iota
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
