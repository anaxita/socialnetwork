package apimodel

import (
	"synergycommunity/internal/domain/entity"
	"time"
)

type AddGroupInput struct {
	Name        string  `json:"name" validate:"min=3,max=64"`
	Description string  `json:"description" validate:"min=10,max=255"`
	Tags        []int64 `json:"tags"`
}

func (g AddGroupInput) ToEntity() entity.Group {
	tags := make([]entity.Tag, len(g.Tags))

	for i, v := range g.Tags {
		tags[i] = entity.Tag{
			ID: v,
		}
	}

	return entity.Group{
		Name:        g.Name,
		Description: g.Description,
		Tags:        tags,
	}
}

type EditGroupInput struct {
	ID          int64   `json:"id"`
	UserID      int64   `json:"user_id" `
	Name        string  `json:"name" validate:"min=3,max=64"`
	Slug        string  `json:"slug" validate:"min=3,max=64"`
	Description string  `json:"description" validate:"min=10,max=255"`
	Tags        []int64 `json:"tags"`
}

func (g EditGroupInput) ToEntity() entity.Group {
	tags := make([]entity.Tag, len(g.Tags))

	for i, v := range g.Tags {
		tags[i] = entity.Tag{
			ID: v,
		}
	}

	return entity.Group{
		ID:          g.ID,
		UserID:      g.UserID,
		Name:        g.Name,
		Slug:        g.Slug,
		Description: g.Description,
		Tags:        tags,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

type GroupsWithPagination struct {
	Items []entity.Group `json:"items"`
	Pagination
}
