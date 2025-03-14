package services

import (
	"github.com/MikeThesol/user-api/internal/models"
	"github.com/MikeThesol/user-api/internal/repository"
)

type Register interface {
	CreateUser(user models.User) (int, error)
}

type User interface {
	GetUserByID(id int) (*models.UserResponse, error)
	UpdateUserInfo(user *models.UserRequestForUpdate) error
}

type Photo interface {
	CreatePhoto(photoReq *models.UserPhotoRequest) ([]int, error)
}

type Service struct {
	Register
	User
	Photo
}

func NewService(repos *repository.Repository, path string) *Service {
	return &Service{
		Register: NewRegisterService(repos),
		User:     NewUserService(repos),
		Photo:    NewPhotoService(repos, path),
	}
}
