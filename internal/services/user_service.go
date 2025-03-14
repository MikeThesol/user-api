package services

import (
	"github.com/MikeThesol/user-api/internal/models"
	"github.com/MikeThesol/user-api/internal/repository"
	"strings"
)

type UserService struct {
	repos *repository.Repository
}

func NewUserService(repos *repository.Repository) *UserService {
	return &UserService{repos: repos}
}

func (u *UserService) UpdateUserInfo(user *models.UserRequestForUpdate) error {
	err := u.repos.User.UpdateUserInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) GetUserByID(id int) (*models.UserResponse, error) {
	var user models.UserResponse
	userForDB, err := u.repos.User.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	dataString := string(userForDB.Photos)
	dataString = strings.Trim(dataString, "[]")
	dataString = strings.ReplaceAll(dataString, "\"", "")
	photos := strings.Split(dataString, ",")
	user = convertUserResponseForDBToUserResponse(userForDB, photos)
	return &user, nil
}

func convertUserResponseForDBToUserResponse(userForDB *models.UserResponseForDB, photos []string) models.UserResponse {
	return models.UserResponse{
		Name:      userForDB.Name,
		Age:       userForDB.Age,
		Gender:    userForDB.Gender,
		Interests: userForDB.Interests,
		Status:    userForDB.Status,
		Bio:       userForDB.Bio,
		Country:   userForDB.Country,
		City:      userForDB.City,
		Photos:    &photos,
	}
}
