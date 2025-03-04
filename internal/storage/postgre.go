package storage

import "user-api/internal/models"

type DB []models.User

func (d *DB) Create(us_rq models.User_req) bool {
	var u models.User
	u.Password = us_rq.Password
	u.Name = us_rq.Name
	u.Email = us_rq.Email
	*d = append(*d, u)
	return true
}

func (d *DB) Read(email string) models.User_req {
	var u models.User_req
	for i := 0; i < len(*d); i++ {
		if (*d)[i].Email == "" {
			u.Email = (*d)[i].Email
			u.Name = (*d)[i].Name
			u.Password = (*d)[i].Password
			return u
		}
	}
	return models.User_req{}
}
