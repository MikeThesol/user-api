package repository

import (
	"fmt"
	"github.com/MikeThesol/user-api/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) GetUserByID(id int) (*models.UserResponseForDB, error) {
	var user models.UserResponseForDB

	query := fmt.Sprintf(`SELECT
	u.name, u.age, u.gender,
		u.interests, u.status, u.bio,
		u.country, u.city,
		COALESCE(JSON_AGG(up.photo), '[]') AS photos
	FROM %s u
	LEFT JOIN %s up ON u.id = up.user_id
	WHERE u.id = $1
	GROUP BY u.id`, userTable, photoTable)

	if err := u.db.Get(&user, query, id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserPostgres) UpdateUserInfo(user *models.UserRequestForUpdate) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1, age = $2, gender = $3, interests = $4, status = $5, bio = $6, country = $7, city = $8 WHERE id = %d", userTable, user.ID)
	result, err := u.db.Exec(query, user.Name, user.Age, user.Gender, pq.Array(user.Interests), user.Status, user.Bio, user.Country, user.City)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	_ = rows
	return nil
}
