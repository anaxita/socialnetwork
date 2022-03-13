package dbmodel

import (
	"github.com/gocraft/dbr"
	"time"
)

type Post struct {
	ID        int64        `db:"post_id"`
	UserID    int64        `db:"user_id"`
	GroupID   int64        `db:"group_id"`
	Name      string       `db:"name"`
	Text      string       `db:"text"`
	Tags      Tags         `db:"tags"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt dbr.NullTime `db:"deleted_at"`
}
