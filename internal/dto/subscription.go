package dto

import (
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/infrastructure/dbmodel"
)

func SubscriptionsFromDB(d dbmodel.Subscriptions) entity.UserSubscriptions {
	return entity.UserSubscriptions{
		Users:  UsersFromDB(d.Users),
		Groups: GroupsFromDB(d.Groups),
		Tags:   TagsFromDB(d.Tags),
	}
}
