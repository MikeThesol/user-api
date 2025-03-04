package app

import (
	"log/slog"
	"time"
	grpcapp "user-api/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, userServiceAddr string, tokenTTL time.Duration) *App {
	//authCon, err := grpc.NewClient(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Error("failed to connect to user service", slog.String("error", err.Error()))
	//	panic(err)
	//}
	//
	//authClient := auth.NewAuthClient(authCon)
	//
	//userService := // TODO
	return &App{}
}
