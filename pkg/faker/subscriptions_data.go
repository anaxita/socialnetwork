package faker

import (
	"math/rand"
	"synergycommunity/internal/domain/entity"
)

func (f *F) GenerateSubscriptions(count int, users []entity.User, groups []entity.Group, tags []entity.Tag) []entity.Subscription {
	subs := make([]entity.Subscription, count)

	models := []entity.Model{entity.ModelGroup, entity.ModelTag, entity.ModelUser}

	for i := range subs {
		m := models[rand.Intn(len(models))]

		var mID int64

		switch m {
		case entity.ModelTag:
			mID = tags[rand.Intn(len(tags))].ID
		case entity.ModelGroup:
			mID = groups[rand.Intn(len(groups))].ID
		case entity.ModelUser:
			mID = users[rand.Intn(len(users))].ID
		}

		subs[i] = entity.Subscription{
			Model:   m,
			ModelID: mID,
			UserID:  users[rand.Intn(len(users))].ID,
		}
	}

	return subs
}
