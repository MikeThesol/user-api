package models

import "mime/multipart"

type UserPhoto struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Photo    string `json:"photo"`
	IsAvatar bool   `json:"is_avatar"`
}

type UserPhotoRequest struct {
	UserID int             `json:"user_id"`
	Photos []*PhotoRequest `json:"photo"`
}

type PhotoRequest struct {
	File     *multipart.FileHeader `json:"files"`
	IsAvatar bool                  `json:"is_avatar"`
}

type UserPhotoResponse struct {
	Photo  []*string `json:"photo"`
	Avatar string    `json:"avatar"`
}
