package dbmodel

import "time"

type Tag struct {
	ID        int64     `db:"tag_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
