package handlers

import (
	"github.com/MikeThesol/user-api/internal/services"
	"log/slog"
)

type Handler struct {
	service *services.Service
	log     *slog.Logger
}

func NewHandler(services *services.Service, log *slog.Logger) *Handler {
	return &Handler{service: services, log: log}
}

//func (h *Handler) InitRoutes() *gin.Engine {
//	router := gin.New()
//
//	api := router.Group("/api/v1/")
//	{
//
//	}
//	return nil
//}
