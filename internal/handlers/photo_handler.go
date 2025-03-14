package handlers

import (
	"github.com/MikeThesol/user-api/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreatePhoto(c *gin.Context) {
	var photo models.UserPhotoRequest
	id, err := strconv.Atoi(c.PostForm("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), h.log)
		return
	}
	photo.UserID = id

	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), h.log)
		return
	}

	files := form.File["photos"]

	if len(files) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "Files haven't been downloaded", h.log)
		return
	}
	var photoReqArray []*models.PhotoRequest

	for _, file := range files {
		ph := models.PhotoRequest{file, false}
		photoReqArray = append(photoReqArray, &ph)
	}

	photo.Photos = photoReqArray

	arrayOfID, err := h.service.Photo.CreatePhoto(&photo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": arrayOfID})
}
