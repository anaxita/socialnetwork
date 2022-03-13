// Package dbmodel contains database specific entity models
package dbmodel

import (
	"time"

	"github.com/gocraft/dbr"
)

type User struct {
	ID        int64        `db:"user_id"`
	Nickname  string       `db:"nickname"`
	FirstName string       `db:"first_name"`
	LastName  string       `db:"last_name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt dbr.NullTime `db:"deleted_at"`
}

type UserRole struct {
	UserID    int64        `db:"user_id"`
	GroupID   int64        `db:"group_id"`
	RoleID    int64        `db:"role_id"`
	ExpiresAt dbr.NullTime `db:"expires_at"`
}

type Subscriptions struct {
	Groups []GroupWithTags `db:"groups"`
	Tags   []Tag           `db:"tags"`
	Users  []User          `db:"users"`
}
