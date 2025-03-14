package repository

import (
	"fmt"
	"github.com/MikeThesol/user-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type RegPostgres struct {
	db *sqlx.DB
}

func NewRegPostgres(db *sqlx.DB) *RegPostgres {
	return &RegPostgres{db: db}
}

func (r *RegPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, password, age, gender) values ($1, $2, $3, $4, $5) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.Age, user.Gender)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
