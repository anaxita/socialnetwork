package dto

import (
	"math"
	"synergycommunity/internal/delivery/api/apimodel"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/infrastructure/dbmodel"

	"github.com/gocraft/dbr"
)

func UsersWithPaginationToRest(e []entity.User, o apimodel.OptionsInput, count int64) apimodel.UsersWithPagination {
	data := apimodel.UsersWithPagination{
		Items: e,
		Pagination: apimodel.Pagination{
			Page:       o.Page,
			CountItems: count,
			CountPages: int64(math.Ceil(float64(count) / float64(o.Limit))),
		},
	}

	return data
}

func UsersFromDB(u []dbmodel.User) []entity.User {
	users := make([]entity.User, len(u))

	for i, v := range u {
		users[i] = entity.User{
			ID:        v.ID,
			Nickname:  v.Nickname,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return users
}

func UserFromRest(u apimodel.UserInput) entity.User {
	return entity.User{
		ID:        u.ID,
		Nickname:  u.Nickname,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

func UserFromDB(u dbmodel.User) entity.User {
	return entity.User{
		ID:        u.ID,
		Nickname:  u.Nickname,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func UserToDB(u entity.User) dbmodel.User {
	return dbmodel.User{
		ID:        u.ID,
		Nickname:  u.Nickname,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func UserRoleFromRest(u apimodel.UserRole) entity.UserRole {
	return entity.UserRole{
		UserID:    u.UserID,
		GroupID:   u.GroupID,
		RoleID:    u.RoleID,
		ExpiresAt: u.ExpiresAt,
	}
}

func UserRoleToDB(u entity.UserRole) dbmodel.UserRole {
	t := dbr.NullTime{
		Time:  u.ExpiresAt,
		Valid: !u.ExpiresAt.IsZero(),
	}

	return dbmodel.UserRole{
		UserID:    u.UserID,
		GroupID:   u.GroupID,
		RoleID:    u.RoleID,
		ExpiresAt: t,
	}
}
