package repository

import (
	"fmt"
	"github.com/MikeThesol/user-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type PhotoPostgres struct {
	db *sqlx.DB
}

func NewPhotoPostgres(db *sqlx.DB) *PhotoPostgres {
	return &PhotoPostgres{db: db}
}

func (p *PhotoPostgres) CreatePhoto(photo *models.UserPhoto) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, photo, is_avatar) VALUES ($1, $2, $3) RETURNING id", photoTable)
	row := p.db.QueryRow(query, photo.UserID, photo.Photo, photo.IsAvatar)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
