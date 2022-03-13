package dbmodel

import (
	"encoding/json"
	"github.com/gocraft/dbr"
	"synergycommunity/internal/domain/entity"
	"time"
)

type Group struct {
	ID          int64        `db:"group_id"`
	UserID      int64        `db:"user_id"`
	Name        string       `db:"name"`
	Slug        string       `db:"slug"`
	Description string       `db:"description"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	DeletedAt   dbr.NullTime `db:"deleted_at"`
}

type GroupWithTags struct {
	Group
	Tags Tags `db:"tags"`
}

type Tags []entity.Tag

func (g *Tags) Scan(v interface{}) error {
	if v != nil {
		return json.Unmarshal(v.([]uint8), g)
	}

	return nil
}
