package main

import (
	"github.com/MikeThesol/user-api/internal/config"
	"github.com/MikeThesol/user-api/internal/handlers"
	"github.com/MikeThesol/user-api/internal/repository"
	"github.com/MikeThesol/user-api/internal/services"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

// TODO: доделать логирование

// Главная функция
func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("Start server")

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Error("failed to connect with DB", slog.String("error", err.Error()))
		return
	}

	repos := repository.NewRepository(db)
	service := services.NewService(repos, cfg.PathForPhoto)
	handler := handlers.NewHandler(service, log)

	// TODO: Сделать функцию где регестрирует все endpoints
	router := gin.New()
	router.POST("/reg", handler.Register)
	router.GET("/user/:id", handler.GetUserByID)
	router.POST("/photo/create", handler.CreatePhoto)
	router.PUT("/user/update", handler.UpdateUserInfo)
	router.Static("/photos", "/home/thesol/Desktop/Golang_projects/user-api/photos")
	router.Run(":8080")
}

// Фуннкция для инициализации логера
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
