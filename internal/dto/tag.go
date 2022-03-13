package dto

import (
	"math"
	"synergycommunity/internal/delivery/api/apimodel"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/infrastructure/dbmodel"
)

// func TagToRest(e entity.Tag) apimodel.Tag {
// 	return apimodel.Tag{
// 		ID:        e.ID,
// 		Name:      e.Name,
// 		CreatedAt: e.CreatedAt,
// 		UpdatedAt: e.UpdatedAt,
// 	}
// }
//
// func TagsToRest(e []entity.Tag) []apimodel.Tag {
// 	r := make([]apimodel.Tag, len(e))
//
// 	for i, v := range e {
// 		r[i] = apimodel.Tag{
// 			ID:        v.ID,
// 			Name:      v.Name,
// 			CreatedAt: v.CreatedAt,
// 			UpdatedAt: v.UpdatedAt,
// 		}
// 	}
//
// 	return r
// }

func TagsToDB(e []entity.Tag) []dbmodel.Tag {
	d := make([]dbmodel.Tag, len(e))

	for i, v := range e {
		d[i] = dbmodel.Tag{
			ID:        v.ID,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return d
}

func TagsFromRest(r []int64) []entity.Tag {
	e := make([]entity.Tag, len(r))

	for i, v := range r {
		e[i] = entity.Tag{
			ID: v,
		}
	}

	return e
}

func TagsFromDB(d []dbmodel.Tag) []entity.Tag {
	e := make([]entity.Tag, len(d))

	for i, v := range d {
		e[i] = entity.Tag{
			ID:        v.ID,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return e
}

func AddTagToEntity(r apimodel.AddTagInput) entity.Tag {
	return entity.Tag{
		Name: r.Name,
	}
}

func EditTagToEntity(t apimodel.EditTagInput) entity.Tag {
	return entity.Tag{
		ID:   t.ID,
		Name: t.Name,
	}
}

func TagFromDB(d dbmodel.Tag) entity.Tag {
	return entity.Tag{
		ID:        d.ID,
		Name:      d.Name,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func TagToDB(e entity.Tag) dbmodel.Tag {
	return dbmodel.Tag{
		ID:        e.ID,
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func TagsWithPaginationToRest(e []entity.Tag, o apimodel.OptionsInput, count int64) apimodel.TagsWithPagination {
	data := apimodel.TagsWithPagination{
		Items: e,
		Pagination: apimodel.Pagination{
			Page:       o.Page,
			CountItems: count,
			CountPages: int64(math.Ceil(float64(count) / float64(o.Limit))),
		},
	}

	return data
}
