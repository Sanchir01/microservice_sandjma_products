package productservice

import (
	"context"
	"github.com/Sanchir01/microservice_sandjma_products/internal/domain/models"
	"log/slog"
)

type Products struct {
	log            *slog.Logger
	serviceProduct ServiceProduct
}

type ServiceProduct interface {
	AllProducts(ctx context.Context) ([]models.Product, error)
}

func New(log *slog.Logger, serviceProduct ServiceProduct) *Products {
	return &Products{
		log:            log,
		serviceProduct: serviceProduct,
	}
}
func (s *Products) AllProducts(ctx context.Context) ([]models.Product, error) {
	const op = "products.Products.GetAllProducts"

	log := s.log.With(slog.String("op", op), slog.String("method", "GetAllProducts"))
	log.Info("getting all products")

	products, err := s.serviceProduct.AllProducts(ctx)
	if err != nil {
		s.log.Error("failed to get products", err.Error())
		return nil, err
	}

	return products, nil
}
