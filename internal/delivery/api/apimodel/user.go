package apimodel

import (
	"synergycommunity/internal/domain/entity"
	"time"
)

type UsersWithPagination struct {
	Items []entity.User `json:"items"`
	Pagination
}

type UserInput struct {
	ID        int64  `json:"id"`
	Nickname  string `json:"nickname" validate:"min=3,max=64"`
	FirstName string `json:"first_name" validate:"min=3,max=64"`
	LastName  string `json:"last_name" validate:"min=3,max=64"`
}

type UserRole struct {
	UserID    int64     `json:"user_id"`
	GroupID   int64     `json:"group_id"`
	RoleID    int64     `json:"role_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

type UserSubscriptionInput struct {
	Model   entity.Model `json:"model"`
	ModelID int64        `json:"model_id"`
}

func (r UserSubscriptionInput) ToEntity(userID int64) entity.Subscription {
	return entity.Subscription{
		Model:   r.Model,
		ModelID: r.ModelID,
		UserID:  userID,
	}
}
