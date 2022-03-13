package entity

import "time"

type Post struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	GroupID   int64     `json:"group_id" db:"group_id"`
	Title     string    `json:"name" db:"name"`
	Text      string    `json:"text" db:"text"`
	Tags      []Tag     `json:"tags" db:"tags"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" json:"updated_at"`
}
