package faker

import (
	"math/rand"
	"strconv"
	"synergycommunity/internal/domain/entity"
	"time"
)

var groupsTitlesRu = []string{
	"Программирование как наказание",
	"Рукоделие руками",
	"Советы трейдерам",
	"Медики - пьем и лечим",
	"Физика - наука о физике",
	"Аферисты",
	"Фиксики чинят",
	"Синергия онлайн",
	"Клубничка",
	"Порно девяностых",
	"Огороды и хозяйство",
	"Травничество",
	"Наука и техника",
	"Зоология",
	"Все о котах",
	"Все о собаках",
	"Брачные истории",
	"Подслушано в Синергии",
	"Все о политике",
	"Все о политике",
	"Тимлиды хотят спать",
}

func (f *F) GenerateGroups(count int, users []entity.User, tags []entity.Tag) []entity.Group {
	now := time.Now().UTC()

	groups := make([]entity.Group, count)

	for i := range groups {
		needCountTags := rand.Intn(6) + 1

		if needCountTags > len(tags) {
			needCountTags = rand.Intn(len(tags))
		}

		start := rand.Intn(len(tags) - needCountTags)

		newTags := tags[start+1 : start+needCountTags]

		groups[i] = entity.Group{
			UserID:      users[rand.Intn(len(users))].ID,
			Name:        f.randString(groupsTitlesRu),
			Slug:        strconv.Itoa(rand.Int()),
			Description: f.randString(postsTextsRu),
			Tags:        newTags,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
	}

	return groups
}
