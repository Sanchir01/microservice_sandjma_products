package grpcapp

import (
	"fmt"
	productsgrpc "github.com/Sanchir01/microservice_sandjma_products/internal/grpc/products"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type GrpcApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewGrpcApp(log *slog.Logger, port int) *GrpcApp {
	gRPCServer := grpc.NewServer()
	productsgrpc.NewProductsServerApi(gRPCServer)
	return &GrpcApp{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (g *GrpcApp) MustRun() {
	if err := g.Run(); err != nil {
		g.log.Error("error starting server", err.Error())
		panic(err)
	}

}

func (g *GrpcApp) Run() error {
	const op = "grpcapp.GrpcApp.Run"

	log := g.log.With(slog.String("op", op), slog.Int("port", g.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("starting gRPC server")
	if err := g.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (g *GrpcApp) Stop() {
	const op = "grpcapp.GrpcApp.Stop"
	g.log.With(slog.String("op", op), slog.Int("port", g.port))
	g.gRPCServer.GracefulStop()
}
