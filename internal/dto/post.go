package dto

import (
	"math"
	"synergycommunity/internal/delivery/api/apimodel"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/infrastructure/dbmodel"
)

func PostToRest(p entity.Post) apimodel.Post {
	return apimodel.Post{
		ID:        p.ID,
		UserID:    p.UserID,
		GroupID:   p.GroupID,
		Title:     p.Title,
		Text:      p.Text,
		Tags:      p.Tags,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func PostsToRest(p []entity.Post) []apimodel.Post {
	posts := make([]apimodel.Post, len(p))

	for i, v := range p {
		posts[i] = apimodel.Post{
			ID:        v.ID,
			UserID:    v.UserID,
			GroupID:   v.GroupID,
			Title:     v.Title,
			Text:      v.Text,
			Tags:      v.Tags,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return posts
}

func PostsWithPaginationToRest(
	p []entity.Post, o apimodel.OptionsInput, count int64,
) apimodel.PostsWithPagination {
	posts := make([]apimodel.Post, len(p))

	for i, v := range p {
		posts[i] = apimodel.Post{
			ID:        v.ID,
			UserID:    v.UserID,
			GroupID:   v.GroupID,
			Title:     v.Title,
			Text:      v.Text,
			Tags:      v.Tags,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	data := apimodel.PostsWithPagination{
		Items: PostsToRest(p),
		Pagination: apimodel.Pagination{
			Page:       o.Page,
			CountItems: count,
			CountPages: int64(math.Ceil(float64(count) / float64(o.Limit))),
		},
	}

	return data
}

func PostsToDB(p []entity.Post) []dbmodel.Post {
	posts := make([]dbmodel.Post, len(p))

	for i, v := range p {
		posts[i] = dbmodel.Post{
			ID:        v.ID,
			UserID:    v.UserID,
			GroupID:   v.GroupID,
			Name:      v.Title,
			Text:      v.Text,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return posts
}

func PostsFromDB(p []dbmodel.Post) []entity.Post {
	posts := make([]entity.Post, len(p))

	for i, v := range p {
		e := entity.Post{
			ID:        v.ID,
			UserID:    v.UserID,
			GroupID:   v.GroupID,
			Title:     v.Name,
			Text:      v.Text,
			Tags:      v.Tags,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		posts[i] = e
	}

	return posts
}

func PostFromDB(p dbmodel.Post, t []dbmodel.Tag) entity.Post {
	return entity.Post{
		ID:        p.ID,
		UserID:    p.UserID,
		GroupID:   p.GroupID,
		Title:     p.Name,
		Text:      p.Text,
		Tags:      TagsFromDB(t),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func PostToDB(p entity.Post) (dbmodel.Post, []dbmodel.Tag) {
	return dbmodel.Post{
		ID:        p.ID,
		UserID:    p.UserID,
		GroupID:   p.GroupID,
		Name:      p.Title,
		Text:      p.Text,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, TagsToDB(p.Tags)
}
