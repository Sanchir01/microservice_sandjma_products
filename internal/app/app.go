package app

import (
	grpcapp "github.com/Sanchir01/microservice_sandjma_products/internal/app/grpc"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.GrpcApp
}

func NewApp(log *slog.Logger, grpcPort int) *App {

	// TODO: init products storage

	// TODO: init products server

	grpcApp := grpcapp.New(log, grpcPort)
	return &App{GRPCSrv: grpcApp}
}
