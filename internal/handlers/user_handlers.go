package handlers

import (
	"github.com/MikeThesol/user-api/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), h.log)
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUserInfo(c *gin.Context) {
	var inputUser models.UserRequestForUpdate

	if err := c.BindJSON(&inputUser); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), h.log)
		return
	}

	if err := h.service.User.UpdateUserInfo(&inputUser); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user was updated"})
}
