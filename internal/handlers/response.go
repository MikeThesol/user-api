package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string, log *slog.Logger) {
	log.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
