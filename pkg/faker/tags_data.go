package faker

import (
	"synergycommunity/internal/domain/entity"
	"time"
)

var tagsTitles = []string{
	"Бьютиблог",
	"Депра",
	"Дыбр",
	"Животные",
	"Игры",
	"Картинки и мемы",
	"Кулинария",
	"Левел-ап",
	"Обзоры",
	"Похудение",
	"Психология",
	"Сообщество",
	"Творчество",
	"Фандомное",
	"Фест",
	"Флэшмобы",
	"Холиварка",
	"Языки",
	"Suntek Store",
	"TPS6107х",
	"TSSOP20",
	"Timer",
	"UART",
	"UDK-32F107V",
	"UNIDK",
	"USART",
	"USB",
	"WAKE",
	"WH1602D",
	"WeiTol",
	"XC3S50AN",
	"XC95144",
	"Xilinx",
	"ZiBlog-Nano",
	"ZiChip",
	"zCurveTracer",
	"Автомобиль",
	"Архив",
	"Блок питания",
	"Векторник",
	"Вектроник",
	"Датчик",
	"ЖКИ",
	"Зарядное устройство",
	"ИК пульт",
	"ИК-пульт",
	"Книга",
	"Компаратор",
	"Коробочки",
	"ЛЛТ",
	"ЛУТ",
	"Лужение",
	"Макросы",
	"Маламут",
	"Новости",
	"Печатная палта",
	"Печатная плата",
	"Подделка",
	"Рабочее место",
	"Резисторы",
	"Репитер",
	"Рыбки",
	"СВОЙ КИТАЕЦ",
	"Сервер",
	"Станок",
	"Тенгель",
	"Тиски",
	"Ток потребления",
	"Ультразвук",
	"Флеш-память",
	"Фоторезист",
	"Фрезеровка",
	"ЧПУ",
	"Шитьё",
	"Энкодер",
}

func (f *F) GenerateTags(count int) []entity.Tag {
	now := time.Now().UTC()

	tags := make([]entity.Tag, count)

	for i := range tags {
		tags[i] = entity.Tag{
			Name:      f.randString(tagsTitles),
			CreatedAt: now,
			UpdatedAt: now,
		}
	}

	return tags
}
