package models

type ContactType string

const (
	ContactTypeFullName      ContactType = "fullName"
	ContactTypePhone         ContactType = "phone"
	ContactTypeInstagramNick ContactType = "instagramNick"
)

type Contact struct {
	ID    uint64      `json:"id"`
	Type  ContactType `json:"type"`
	Value string      `json:"value"`
}
