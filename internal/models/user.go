package models

type Gender string

const (
	FEMALE Gender = "Женщина"
	MALE   Gender = "Мужчина"
)

// Псевдоним для интересов пользователя
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

// Псевдоним для статуса отношений
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
	Name         string      `json:"name" binding:"required"`
	Email        string      `json:"email" binding:"required"`
	Password     string      `json:"password" binding:"required"`
	Age          uint8       `json:"age" binding:"required"`
	Gender       Gender      `json:"gender" binding:"required"`
	Interests    []Interests `json:"interests"`
	Status       Status      `json:"status"`
	Bio          string      `json:"bio"`
	County       string      `json:"county"`
	City         string      `json:"city"`
	LikedUsers   []int       `json:"liked_users"`
	YourLikes    []int       `json:"your_likes"`
	YourPartners []int       `json:"your_partners"`
}

// Модель которая будет попадаться в ленте
type UserResponse struct {
	Name      string       `json:"name"`
	Age       uint8        `json:"age"`
	Gender    Gender       `json:"gender"`
	Interests *[]Interests `json:"interests"`
	Status    *string      `json:"status"`
	Bio       *string      `json:"bio"`
	Country   *string      `json:"country"`
	City      *string      `json:"city"`
	Photos    *[]string    `json:"photos"`
}

type UserResponseForDB struct {
	Name      string       `json:"name"`
	Age       uint8        `json:"age"`
	Gender    Gender       `json:"gender"`
	Interests *[]Interests `json:"interests"`
	Status    *string      `json:"status"`
	Bio       *string      `json:"bio"`
	Country   *string      `json:"country"`
	City      *string      `json:"city"`
	Photos    []uint8      `json:"photos"`
}

type UserRequestForUpdate struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Age       uint8       `json:"age"`
	Gender    Gender      `json:"gender"`
	Interests []Interests `json:"interests"`
	Status    Status      `json:"status"`
	Bio       string      `json:"bio"`
	Country   string      `json:"country"`
	City      string      `json:"city"`
}
