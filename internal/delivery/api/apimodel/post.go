package apimodel

import (
	"synergycommunity/internal/domain/entity"
	"time"
)

type Post struct {
	ID        int64        `json:"id"`
	UserID    int64        `json:"user_id"`
	GroupID   int64        `json:"group_id"`
	Title     string       `json:"name"`
	Text      string       `json:"text"`
	Tags      []entity.Tag `json:"tags"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type AddPostInput struct {
	GroupID int64   `json:"group_id"`
	Name    string  `json:"name" validate:"min=3,max=64"`
	Text    string  `json:"text" validate:"min=8,max=2000"`
	Tags    []int64 `json:"tags"`
}

func (p AddPostInput) ToEntity() entity.Post {
	return entity.Post{
		GroupID: p.GroupID,
		Title:   p.Name,
		Text:    p.Text,
		Tags:    Tags(p.Tags).ToEntity(),
	}
}

type EditPostInput struct {
	ID   int64   `json:"id"`
	Name string  `json:"name" validate:"min=3,max=64"`
	Text string  `json:"text" validate:"min=8,max=2000"`
	Tags []int64 `json:"tags"`
}

func (p EditPostInput) ToEntity() entity.Post {
	return entity.Post{
		ID:    p.ID,
		Title: p.Name,
		Text:  p.Text,
		Tags:  Tags(p.Tags).ToEntity(),
	}
}

type PostsWithPagination struct {
	Items []Post `json:"items"`
	Pagination
}
