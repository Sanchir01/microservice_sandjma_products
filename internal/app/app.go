package app

import (
	grpcapp "github.com/Sanchir01/microservice_sandjma_products/internal/app/grpc"
	productservice "github.com/Sanchir01/microservice_sandjma_products/internal/services/products"
	"github.com/Sanchir01/microservice_sandjma_products/internal/store/postgres"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.GrpcApp
}

func NewApp(log *slog.Logger, grpcPort int, db *sqlx.DB) *App {

	store := postgres.NewProductPostgresStorage(db)

	productService := productservice.NewProducts(log, store)

	grpcApp := grpcapp.NewGrpcApp(log, grpcPort, productService)
	return &App{GRPCSrv: grpcApp}
}
