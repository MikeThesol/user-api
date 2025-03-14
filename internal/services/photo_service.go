package services

import (
	"fmt"
	"github.com/MikeThesol/user-api/internal/models"
	"github.com/MikeThesol/user-api/internal/repository"
	"io"
	"os"
	"path/filepath"
)

type PhotoService struct {
	repos        *repository.Repository
	pathForPhoto string
}

func NewPhotoService(repos *repository.Repository, path string) *PhotoService {
	return &PhotoService{repos: repos, pathForPhoto: path}
}

func (p *PhotoService) CreatePhoto(photoReq *models.UserPhotoRequest) ([]int, error) {
	photoReq.Photos[0].IsAvatar = true
	var arrayOfID []int
	for _, photo := range photoReq.Photos {
		dstPath := filepath.Join(p.pathForPhoto, photo.File.Filename)

		src, err := photo.File.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		dst, err := os.Create(dstPath)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return nil, err
		}

		fileUrl := fmt.Sprintf("http://localhost:8080/photos/%s", photo.File.Filename)

		photoForSave := models.UserPhoto{
			UserID:   photoReq.UserID,
			Photo:    fileUrl,
			IsAvatar: photo.IsAvatar,
		}
		id, err := p.repos.CreatePhoto(&photoForSave)
		if err != nil {
			return nil, err
		}
		arrayOfID = append(arrayOfID, id)

	}
	return arrayOfID, nil
}
