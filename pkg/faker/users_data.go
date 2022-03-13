package faker

import (
	"math/rand"
	"strconv"
	"synergycommunity/internal/domain/entity"
	"time"
)

var namesRu = []string{
	"Артём",
	"Александр",
	"Максим",
	"Даниил",
	"Дмитрий",
	"Иван",
	"Кирилл",
	"Никита",
	"Михаил",
	"Егор",
	"Матвей",
	"Андрей",
	"Илья",
	"Алексей",
	"Роман",
	"Сергей",
	"Владислав",
	"Ярослав",
	"Тимофей",
	"Арсений",
	"Денис",
	"Владимир",
	"Павел",
	"Глеб",
	"Константин",
	"Богдан",
	"Евгений",
	"Николай",
	"Степан",
	"Захар",
}

var surnamesRu = []string{
	"Иванов",
	"Смирнов",
	"Кузенецов",
	"Попов",
	"Васильев",
	"Петров",
	"Соколов",
	"Михайлов",
	"Новиков",
	"Федоров",
	"Морозов",
	"Волков",
	"Алекссеев",
	"Лебедев",
	"Семенов",
	"Егоров",
	"Павлов",
	"Козлов",
	"Степанов",
	"Николаев",
	"Орлов",
	"Андреев",
	"Макаров",
	"Никитин",
	"Захаров",
	"Зайцев",
	"Соловьев",
	"Борисов",
	"Яковлев",
	"Григорьев",
}

var nicknamesEn = []string{
	"Abrahym",
	"Achilles",
	"Addison",
	"Alberto",
	"Aldridge",
	"Alexander",
	"Andrew",
	"Artemis",
	"Augustus",
	"Bartholomew",
	"Baxter",
	"Beckett",
	"Brandon",
	"Callahan",
	"Cameron",
	"Carlisle",
	"Casey",
	"Christopher",
	"Creighton",
	"Daniel",
	"Dimitri",
	"Edward",
	"Elliott",
	"Fernando",
	"Frederick",
	"Gregory",
	"Isaiah",
	"Irving",
	"Jacob",
	"Jeremiah",
	"Kenndy",
	"Lennox",
	"Leonardo",
	"Lucas",
	"Maximilian",
	"Nathaniel",
	"Patrick",
	"Pierson",
	"Richard",
	"Sullivan",
	"Theodore",
	"Thomas",
	"Warren",
	"Wesley",
	"William",
	"Zachary",
}

func (f *F) GenerateUsers(count int) []entity.User {
	now := time.Now().UTC()

	users := make([]entity.User, count)

	for i := range users {
		users[i] = entity.User{
			Nickname:  f.randString(namesRu) + strconv.Itoa(rand.Int()),
			FirstName: f.randString(surnamesRu),
			LastName:  f.randString(nicknamesEn),
			// Rating:    rand.Int63n(10001),
			CreatedAt: now,
			UpdatedAt: now,
		}
	}

	return users
}
