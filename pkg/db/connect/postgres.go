package connect

import (
	"fmt"
	"github.com/Sanchir01/microservice_sandjma_products/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

func PostgresMain(cfg *config.Config, lg *slog.Logger) *sqlx.DB {
	postgresString := fmt.Sprintf(
		"user=%s dbname=%s sslmode=%s password=%s port=%s host=%s",
		cfg.DB.User, cfg.DB.Database, cfg.DB.SSL, os.Getenv("PASSWORD_POSTGRES"),
		cfg.DB.Port, cfg.DB.Host,
	)

	db, err := sqlx.Open("postgres", postgresString)

	if err != nil {
		lg.Error("sqlx.Connect error", slog.String("error", err.Error()))
	}
	return db
}
