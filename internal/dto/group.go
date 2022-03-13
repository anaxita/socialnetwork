package dto

import (
	"math"
	"synergycommunity/internal/delivery/api/apimodel"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/infrastructure/dbmodel"
)

// func GroupToRest(g entity.Group) apimodel.Group {
// 	return apimodel.Group{
// 		ID:          g.ID,
// 		UserID:      g.UserID,
// 		Name:        g.Name,
// 		Slug:        g.Slug,
// 		Description: g.Description,
// 		Tags:        g.Tags,
// 		CreatedAt:   g.CreatedAt,
// 		UpdatedAt:   g.UpdatedAt,
// 	}
// }
//
// func GroupsToRest(g []entity.Group) []apimodel.Group {
// 	e := make([]apimodel.Group, len(g))
//
// 	for i, v := range g {
// 		e[i] = apimodel.Group{
// 			ID:          v.ID,
// 			UserID:      v.UserID,
// 			Name:        v.Name,
// 			Slug:        v.Slug,
// 			Description: v.Description,
// 			Tags:        v.Tags,
// 			CreatedAt:   v.CreatedAt,
// 			UpdatedAt:   v.UpdatedAt,
// 		}
// 	}
//
// 	return e
// }

func GroupsWithPaginationToRest(g []entity.Group, o apimodel.OptionsInput, count int) apimodel.GroupsWithPagination {
	// e := make([]apimodel.Group, len(g))
	//
	// for i, v := range g {
	// 	e[i] = apimodel.Group{
	// 		ID:          v.ID,
	// 		UserID:      v.UserID,
	// 		Name:        v.Name,
	// 		Slug:        v.Slug,
	// 		Description: v.Description,
	// 		Tags:        v.Tags,
	// 		CreatedAt:   v.CreatedAt,
	// 		UpdatedAt:   v.UpdatedAt,
	// 	}
	// }

	data := apimodel.GroupsWithPagination{
		Items: g,
		Pagination: apimodel.Pagination{
			Page:       o.Page,
			CountItems: int64(count),
			CountPages: int64(math.Ceil(float64(count) / float64(o.Limit))),
		},
	}

	return data
}

func GroupsToDB(g []entity.Group) []dbmodel.Group {
	e := make([]dbmodel.Group, len(g))

	for i, v := range g {
		e[i] = dbmodel.Group{
			ID:          v.ID,
			UserID:      v.UserID,
			Name:        v.Name,
			Slug:        v.Slug,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
	}

	return e
}

// func GroupsFromRest(g []apimodel.Group) []entity.Group {
// 	e := make([]entity.Group, len(g))
//
// 	for i, v := range g {
// 		tags := make([]entity.Tag, len(v.Tags))
//
// 		for ti, vi := range v.Tags {
// 			tags[ti] = entity.Tag{
// 				ID:        vi.ID,
// 				Name:      vi.Name,
// 				CreatedAt: vi.CreatedAt,
// 				UpdatedAt: vi.UpdatedAt,
// 			}
// 		}
//
// 		e[i] = entity.Group{
// 			ID:          v.ID,
// 			UserID:      v.UserID,
// 			Name:        v.Name,
// 			Slug:        v.Slug,
// 			Description: v.Description,
// 			Tags:        tags,
// 			CreatedAt:   v.CreatedAt,
// 			UpdatedAt:   v.UpdatedAt,
// 		}
// 	}
//
// 	return e
// }

func GroupsFromDB(g []dbmodel.GroupWithTags) []entity.Group {
	e := make([]entity.Group, len(g))

	for i, v := range g {
		e[i] = entity.Group{
			ID:          v.ID,
			UserID:      v.UserID,
			Name:        v.Name,
			Slug:        v.Slug,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			Tags:        v.Tags,
		}
	}

	return e
}

// func AddGroupInputToEntity(g apimodel.AddGroupInput) entity.Group {
// 	tags := make([]entity.Tag, len(g.Tags))
//
// 	for i, v := range g.Tags {
// 		tags[i] = entity.Tag{
// 			ID:   v.ID,
// 			Name: v.Name,
// 		}
// 	}
//
// 	return entity.Group{
// 		Name:        g.Name,
// 		Description: g.Description,
// 		Tags:        tags,
// 	}
// }

// func EditGroupInputToEntity(g apimodel.EditGroupInput) entity.Group {
// 	tags := make([]entity.Tag, len(g.Tags))
//
// 	for i, v := range g.Tags {
// 		tags[i] = entity.Tag{
// 			ID:   v.ID,
// 			Name: v.Name,
// 		}
// 	}
//
// 	return entity.Group{
// 		ID:          g.ID,
// 		UserID:      g.UserID,
// 		Name:        g.Name,
// 		Slug:        g.Slug,
// 		Description: g.Description,
// 		Tags:        tags,
// 	}
// }

func GroupFromDB(g dbmodel.GroupWithTags) entity.Group {
	e := entity.Group{
		ID:          g.ID,
		UserID:      g.UserID,
		Name:        g.Name,
		Slug:        g.Slug,
		Tags:        g.Tags,
		Description: g.Description,
		CreatedAt:   g.CreatedAt,
		UpdatedAt:   g.UpdatedAt,
	}

	return e
}

func GroupToDB(g entity.Group) dbmodel.GroupWithTags {
	return dbmodel.GroupWithTags{
		Group: dbmodel.Group{
			ID:          g.ID,
			UserID:      g.UserID,
			Name:        g.Name,
			Slug:        g.Slug,
			Description: g.Description,
			CreatedAt:   g.CreatedAt,
			UpdatedAt:   g.UpdatedAt,
		},
		Tags: g.Tags,
	}
}
