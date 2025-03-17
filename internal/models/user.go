package models

type User struct {
	ID       uint64    `json:"id"`
	Contacts []Contact `json:"contacts"`
}
