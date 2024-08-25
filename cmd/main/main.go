package main

import (
	"github.com/Sanchir01/microservice_sandjma_products/internal/app"
	"github.com/Sanchir01/microservice_sandjma_products/internal/config"
	"github.com/Sanchir01/microservice_sandjma_products/pkg/db/connect"
	"github.com/Sanchir01/microservice_sandjma_products/pkg/lib/logger/handlers/slogpretty"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var (
	development = "development"
	production  = "production"
)

func main() {
	cfg := config.MustLoad()

	lg := setupLogger(cfg.Env)

	db := connect.PostgresMain(cfg, lg)
	defer db.Close()

	application := app.NewApp(lg, cfg.GRPC.Port, db)

	go func() {
		application.GRPCSrv.MustRun()
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	lg.Info("stoppping application", slog.String("signal", sign.String()))

	application.GRPCSrv.Stop()
}

func setupLogger(env string) *slog.Logger {
	var lg *slog.Logger
	switch env {
	case development:
		lg = setupPrettySlog()
	case production:
		lg = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return lg
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

func setupSlog() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}
