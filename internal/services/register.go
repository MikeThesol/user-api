package services

import (
	"crypto/sha256"
	"fmt"
	"github.com/MikeThesol/user-api/internal/models"
	"github.com/MikeThesol/user-api/internal/repository"
)

const salt = "hrienfevgek"

type RegisterService struct {
	repo *repository.Repository
}

func NewRegisterService(repo *repository.Repository) *RegisterService {
	return &RegisterService{repo: repo}
}

func (r *RegisterService) CreateUser(user models.User) (int, error) {
	user.Password = GeneratePasswordHash(user.Password)
	return r.repo.CreateUser(user)
}

func GeneratePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
