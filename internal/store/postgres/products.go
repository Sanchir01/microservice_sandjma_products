package postgres

import (
	"context"
	"github.com/Sanchir01/microservice_sandjma_products/internal/domain/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"log/slog"
	"time"
)

type ProductPostgresStorage struct {
	db *sqlx.DB
}

func NewProductPostgresStorage(db *sqlx.DB) *ProductPostgresStorage {
	return &ProductPostgresStorage{db: db}
}

func (db *ProductPostgresStorage) AllProducts(ctx context.Context) ([]models.Product, error) {
	const op = "store.postgres.AllProducts"

	log := slog.With(
		slog.String("op", op),
		slog.String("postgres", "postgres AllCategories"))
	log.Info("AllProductsDB")

	conn, err := db.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	var products []productdb
	if err := conn.SelectContext(ctx, &products, "SELECT * FROM products"); err != nil {
		return nil, err
	}

	return lo.Map(products, func(product productdb, _ int) models.Product { return models.Product(product) }), nil
}

type productdb struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Images      []string  `json:"images"`
	CategoryID  uuid.UUID `json:"category_id"`
	Description string    `json:"description"`
	Version     uint      `json:"version"`
}
