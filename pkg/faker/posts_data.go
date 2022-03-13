package faker

import (
	"math/rand"
	"synergycommunity/internal/domain/entity"
	"time"
)

var postsTitlesRu = []string{
	"Объясните, что такое OOП?",
	"В чем разница между #import и #include?",
	"Что такое статический элемент?",
	"Векторы в C++",
	"В чем разница ежду структурой и классом?",
	"Использование точки в C++",
	"Чем new() отличается от maloc()?",
	"Что такое this?",
	"В чем разница меду массивом и списком?",
	"Что такое динамическая и статическая типизация?",
	"Что подразумевается под делегатом?",
	"Метод Mutator. Что это?",
	"Объясните, что такое одиочное и множественное наследование?",
	"Может ли встроенная функция быть рекурсивной?",
	"может вызвать встроенную функцию из самой себя компилятор не сможет сгенерировать встроенный код, т. к. не сможет определить глубину рекурсии во время компиляции. Компилятор с хорошим оптимизатором может встроить рекурсивные вызовы до некоторой фиксированной глубины во время компиляции, а также вставлять не рекурсивные вызовы для случаев, когда в рантайме превышается фактическая глубина.",
	"Объясните, что такое инкапсуляция?",
	"Что такое абстракция? Чем она отлиается от инкапсуляции?",
	"Что такое встроенные функции? Каков синтаксис для определеия?",
	"В чем разница между указателем и ссылкой в C++?",
	"Почему C++ является языком программирования среднео уровня?",
	"Объясните, что такое полиморфизм?",
	"Что такое определение класса?",
	"Что означают ключевые слова voatile и mutable?",
	"Что такое виртуальная функция?",
	"Что подразумевается под перегрузкой функций и операторов?",
	"Что такое переопределение функции?",
}

var postsTextsRu = []string{
	"Функция становится переопределенной, если производный класс унаследует ее от базового и определит ее у",
	"себя. Если функция существует в двух классах (базовый и производный), то при вызове будет выполнена переопределенная функция, а функция базового класса будет проигнорирована.",
	"Перегруженные функции – это функции не только с одинаковыми именами, но и с разными типами и количеств",
	"ом передаваемых параметров. Приведем несколько классов, в которых можно перегружать, например, арифметические операторы: Complex Number, Fracional Number, Big Integer.",
	"Самый популярный вопрос на собеседовании по C++. Объектно-ориентированное программирование (ООП) – это",
	"тип программирования, в котором программист указывает тип данных для операций, функций и структур данных, используемых в коде, а также создат отношения между объектами.",
	"В ООП структура данных является объектом, в состав которого могут входить как данные, так и функции. О",
	"бъектно-ориентированное программирование в основном направлено на реализацию реальных сущностей. К ним относятся: абстракция, инкапсуляция, наследование, полиморфизм и т. д.",
	"Распространенный вопрос на собеседовании по C++. Оператор #include используется в C++ для включения ис",
	"ходного файла или импорта заголовков файлов, содержащих объявление функций и прочих конструкций, которые будут совместно использоваться в программе. Оператор #import – это «майкрософт-специфичный» оператор, используемый в бинарных библиотеках, таких как DLLили Lib. Он очень похож на #include, ведь загружает все определения функций и заголовка из файла DLL, а разработчик может использовать заголовки так же, как в случае с #include.",
	"Оператор #include позволяет подключать один и тот же файл несколько раз, а #importгарантирует, что пре",
	"процессор включает файл только один раз.",
	"Static – ключевое слово, используемое для придания элементу уникальных характеристик. Для хранения ста",
	"тических элементов, память выделяется только один раз в течение всего жизненного цикла программы. Такие элементы похожи на глобальные переменные, за исключением области видимости (public, private). С их помощью можно ограничить ее использование.",
	"Объявленные методы с тем же именем и параметрами не могут быть перегружены, если любое из них является",
	"статическим, плюс статическая функция не может быть объявлена как const, volatileили const volatile.",
	"Векторы – это своего рода контейнеры для структур данных, которые представляют собой массивы, способны",
	"е изменять свой размер. Точно так же, как массивы, векторы используют смежные ячейки хранения для своих элементов, это означает, что элементы векторов могут быть доступны с помощью смещений так же эффективно, как и в массивах. Но в отличие от массивов, их размер может изменяться динамически, а хранилищем автоматически управляет сам контейнер.",
	"В C++ класс является расширением структуры, используемой в ЯП. Класс – это пользовательский тип данных",
	", связывающий данные и зависимые функции в одном блоке. Структура и класс в C++ сильно отличаются, поскольку структура имеет ограниченную функциональность и возможности по сравнению с классом. Структура также является пользовательским типом данных с конкретным шаблоном и может содержать как методы, так и классы. Эти два понятия различаются по назначению: класс используется для абстракции данных и дальнейшего наследования, а структура обычно предназначена для группировки данных.",
	"Точка – это чаще всего ссылка на метод или свойство объекта в ООП. Связь между объектом, атрибутами и ",
	"методами обозначается точкой («.«), установленной между ними. Как точка, так и оператор «->«, используются для ссылки на отдельные члены классов, структуры и объединения. Оператор точка применяется к фактическому объекту, определенному в классе.",
	"New() является препроцессором, в то время как malloc() – методом. Пользователю нет необходимости выдел",
	"ять память при использовании «new«, а в malloc() для выделения памяти необходимо использовать функцию sizeof(). «new» инициализирует новую память в 0, в то время как malloc() сохраняет случайное значение в новой выделенной ячейке памяти.",
	"Один из самых популярных вопросов на собеседовании по C++. Ключевое слово this передается всем нестати",
	"ческим методам как скрытый аргумент и доступен в виде локальной переменной внутри всех нестатических методов. Оператор this является постоянным указателем, который хранит в памяти ссылку на текущий объект. Он недоступен в статических функциях, т. к. они могут вызываться без какого-либо объекта (используя имя класса).",
	"Массив – это набор однородных элементов, а список – разноро",
	"дных.",
	"Распределение памяти массива всегда статическое и непрерывное, а в списке все это динамическое и рандо",
	"мное.",
	"В случае с массивами пользователю не нужно управлять выделением памяти, а при использовании списков пр",
	"идется, из-за ее динамичности.",
	"Собеседование п",
	"о C++",
	"Статически типизированные языки – это языки, в которых проверка типа совершается во время компиляции, ",
	"а в динамически типизированных – в рантайме. Поскольку C++ является статически типизированным языком, пользователь должен сообщить компилятору, с каким типом объекта он работает во время компиляции.",
	"Делегат – это объект, действующий от имени, или в паре с другим объектом, обнаружившим событие во врем",
	"я выполнения программы. Зачастую это просто указатель на функцию, использующую обратные вызовы.",
	"Делегаты могут быть сохранены пользователем. Как правило, они сохраняются автоматически, поэтому можно",
	"избежать лишних циклов сохранения и не производить запись повторно.",
	"Функция доступа создает элемент типа protected или private для внешнего использования, но она не дает ",
	"разрешения на его редактирование или изменение. Изменение protected-элемента данных всегда требует вызова функции-мутатора. Мутатор обеспечивают прямой доступ к защищенным данным, поэтому при создании функции мутатора и аксессора нужно быть очень внимательным.",
	"Наследование позволяет определить класс, имеющий определенный набор характеристик (например, методы и ",
	"переменные экземпляра), а затем создать другие классы, производные от этого класса. Производный класс наследует все функции родительского класса и обычно добавляет некоторые собственные функции.",
	"Множественное наследование является особенностью C++, где один класс может наследовать объекты и метод",
	"ы нескольких классов. Конструкторы наследуемых классов вызываются в том же порядке, в котором наследуются базовые классы.",
	"Среди вопросов на собеседовании по C++ на этом»сыпятся» почти все новички. Инкапсуляция – это механизм",
	"объединения используемых данных и функций для сокрытия деталей реализации от пользователя. При этом пользователь может выполнять только ограниченный набор операций со скрытыми членами класса, используя внутренние методы. Эта ООП концепция часто применяется для сокрытия внутреннего представления или состояния объекта от «внешнего мира».",
	"Абстракция – это механизм предоставления только интерфейсов, сокрытия сведений о реализации и “показ” ",
	"необходимых деталей функционала. Инкапсуляцию можно понимать как сокрытие свойств и методов от внешнего мира. Класс является лучшим примером инкапсуляции в C++.",
	"Встроенная функция – это функция, объявляемая с ключевым словом inline. Всякий раз, когда вызывается в",
	"строенная функция, ее полный код подставляется в место вызова. Компилятор выполняет эту подстановку во время компиляции. Встроенная функция может повысить эффективность кода.",
	"Синтаксис для определения фун",
	"Санкции это плохо:",
	"Указатель может быть переназначен n-раз, в то время как ссылка не может быть переназначена после бинда",
	"Указатели могут указывать в NULL, тогда как ссылка всегда ссылается на объект. Программист не может получить адрес ссылки, как это возможно с указателями, но можно взять адрес объекта, на который указывает ссылка, и выполнить действия с ним.",
	"Почти весь код на С будет успешно работать на C++, но это не дает повода рассматривать его как надмнож",
	"ество C. С++ можно назвать языком среднего уровня, т. к. он обладает чертами ЯП низкого и высокого уровней одновременно.",
	"Еще один из важных вопросов на собеседовании по C++. Полиморфизм – это способность функции работать с ",
	"разными типами данных. Обычно в ЯП речь идет о двух типах полиморфизма:",
	"Полиморфизм подтипов – вызывающий код использует объект, опираясь только на его интерфейс, не зная при",
	"этом фактического типа.",
	"Параметрический полиморфизм позволяет шаблону, определенному в классе конкретного типа, быть определен",
	"ным в другом типе.",
	"Класс описывает поведение и свойства, общие для любого конкретного объекта. Это пользовательский тип д",
	"анных, содержащий собственные данные и функции, к которым можно получить доступ и использовать, создав экземпляр этого класса. Определение класса означает описание схемы элементов для типа данных или объекта. Оно начинается с ключевого слова class, за которым следует имя и тело класса, заключенное в фигурные скобки. Определение класса всегда должно сопровождаться точкой с запятой или списком объявлений.",
	"Вопросы на собеседовании п",
	"о C++ детям",
	"volatile сообщает компилятору, что переменная может измениться без его ведома. Эти переменные не кэшир",
	"уются компилятором и поэтому всегда считываются из памяти.",
	"mutable может использоваться для переменных класса. Такие переменные могут изменяться изнутри функций ",
	"класса.",
	"Виртуальная функция – это метод, который используется в рантайме для замены реализованного функционала",
	", предоставляемого базовым классом. Виртуальные функции всегда используются с наследованием и вызываются в соответствии с типом объекта, на который указывает или ссылается объект, а не в соответствии с типом указателя или ссылки.",
	"Ключевое слово virtual используется для создания виртуального ме",
	"Тогда это всё.",
	"Распространенный вопрос на собеседовании по C++. C++ позволяет указывать несколько определений функций",
	"или операторов в одной области видимости для нормальной работы в пользовательских классах. Это называется перегрузкой функций и оператор",
}

func (f *F) GeneratePosts(count int, users []entity.User, groups []entity.Group, tags []entity.Tag) []entity.Post {

	posts := make([]entity.Post, count)

	for i := range posts {
		createdAt := time.Time{}.AddDate(rand.Intn(13)+2010, rand.Intn(13)+1, rand.Intn(31)+1)
		updatedAt := createdAt.AddDate(rand.Intn(3), rand.Intn(6), rand.Intn(7))

		needCountTags := rand.Intn(6) + 1

		if needCountTags > len(tags) {
			needCountTags = rand.Intn(len(tags))
		}

		start := rand.Intn(len(tags) - needCountTags)

		newTags := tags[start+1 : start+needCountTags]

		posts[i] = entity.Post{
			UserID:    users[rand.Intn(len(users))].ID,
			GroupID:   groups[rand.Intn(len(groups))].ID,
			Title:     f.randString(postsTitlesRu),
			Text:      f.randString(postsTextsRu),
			Tags:      newTags,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
	}

	return posts
}
