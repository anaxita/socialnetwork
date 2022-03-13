// Package entity contains base models for different entities (regardless of their DB implementations)
package entity

import (
	"time"
)

type User struct {
	ID        int64     `json:"id" db:"id"`
	Nickname  string    `json:"nickname" db:"nickname"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserRole struct {
	UserID    int64
	GroupID   int64
	RoleID    int64
	ExpiresAt time.Time
}
