package handlers

import (
	"github.com/MikeThesol/user-api/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Register(c *gin.Context) {
	var inputUser models.User

	if err := c.BindJSON(&inputUser); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), h.log)
		return
	}

	id, err := h.service.Register.CreateUser(inputUser)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
