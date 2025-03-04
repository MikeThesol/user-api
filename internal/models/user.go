package models

type Gender string

const (
	FEMALE Gender = "Женщина"
	MALE   Gender = "Мужчина"
)

type Interests string

const (
	InterestSports       Interests = "Спорт"
	InterestMusic        Interests = "Музыка"
	InterestTravel       Interests = "Путешествия"
	InterestReading      Interests = "Чтение"
	InterestCooking      Interests = "Кулинария"
	InterestTechnology   Interests = "Технологии"
	InterestArt          Interests = "Искусство"
	InterestPhotography  Interests = "Фотография"
	InterestGaming       Interests = "Игры"
	InterestFitness      Interests = "Фитнес"
	InterestFashion      Interests = "Мода"
	InterestNature       Interests = "Природа"
	InterestMovies       Interests = "Фильмы"
	InterestTheater      Interests = "Театр"
	InterestVolunteering Interests = "Волонтерство"
	InterestMeditation   Interests = "Медитация"
	InterestPets         Interests = "Домашние животные"
	InterestCollecting   Interests = "Коллекционирование"
	InterestHistory      Interests = "История"
	InterestScience      Interests = "Наука"
	InterestPolitics     Interests = "Политика"
	InterestFinance      Interests = "Финансы"
	InterestGardening    Interests = "Садоводство"
	InterestHiking       Interests = "Походы"
	InterestFishing      Interests = "Рыбалка"
	InterestCamping      Interests = "Кемпинг"
	InterestWriting      Interests = "Письмо"
	InterestBlogging     Interests = "Ведение блога"
	InterestLanguages    Interests = "Языки"
	InterestAstronomy    Interests = "Астрономия"
	InterestBoardGames   Interests = "Настольные игры"
	InterestBaking       Interests = "Выпечка"
)

type Status string

const (
	LookingForARelationship Status = "Ищу отношения"
	LookingForAFriendship   Status = "Ищу дружбу"
	LookingForDate          Status = "Схожу на свидание"
	NotLooking              Status = "Не ищу ничего"
	ItIsComplicated         Status = "Все сложно"
)

type User struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Email        string      `json:"email"`
	Password     []byte      `json:"password"`
	Age          uint8       `json:"age"`
	Gender       Gender      `json:"gender"`
	Interests    []Interests `json:"interests"`
	Status       Status      `json:"status"`
	Bio          string      `json:"bio"`
	County       string      `json:"county"`
	City         string      `json:"city"`
	LikedUsers   []int       `json:"liked_users"`
	YourLikes    []int       `json:"your_likes"`
	YourPartners []int       `json:"your_partners"`
}

type UserDTO struct {
	Name      string      `json:"name"`
	Age       uint8       `json:"age"`
	Gender    Gender      `json:"gender"`
	Interests []Interests `json:"interests"`
	Status    Status      `json:"status"`
	Bio       string      `json:"bio"`
	City      string      `json:"city"`
}

type UserPhoto struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Photo    []byte `json:"photo"`
	IsAvatar bool   `json:"is_avatar"`
}

type User_req struct {
	Name     string
	Email    string
	Password []byte
}
