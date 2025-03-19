package models

type Role string

const (
	RoleProfile Role = "profile"
	RoleUser    Role = "user"
	RoleAdmin   Role = "admin"
)

type User struct {
	ID       uint64      `json:"id"`
	Role     ContactType `json:"role"`
	Contacts []Contact   `json:"contacts"`
}
