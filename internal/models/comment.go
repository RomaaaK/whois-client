package models

import "time"

type Comment struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
