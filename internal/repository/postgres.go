package repository

import (
	"fmt"
	"github.com/MikeThesol/user-api/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	userTable  = "users"
	photoTable = "user_photo"
)

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DbName, cfg.Password, cfg.SslMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
