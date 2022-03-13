package apimodel

import "synergycommunity/internal/domain/entity"

type Tags []int64

func (t Tags) ToEntity() []entity.Tag {
	tags := make([]entity.Tag, len(t))

	for i, v := range t {
		tags[i] = entity.Tag{
			ID: v,
		}
	}

	return tags
}

type AddTagInput struct {
	Name string `json:"name" validate:"min=3,max=64"`
}

type EditTagInput struct {
	ID   int64  `json:"id"`
	Name string `json:"name" validate:"min=3,max=64"`
}

type TagsWithPagination struct {
	Items []entity.Tag `json:"items"`
	Pagination
}
