package repository

import (
	"github.com/MikeThesol/user-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type Registration interface {
	CreateUser(user models.User) (int, error)
}

type User interface {
	GetUserByID(id int) (*models.UserResponseForDB, error)
	UpdateUserInfo(user *models.UserRequestForUpdate) error
}

type Photo interface {
	CreatePhoto(photo *models.UserPhoto) (int, error)
}

type Repository struct {
	Registration
	User
	Photo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Registration: NewRegPostgres(db),
		User:         NewUserPostgres(db),
		Photo:        NewPhotoPostgres(db),
	}
}
