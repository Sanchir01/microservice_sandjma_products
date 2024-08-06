package products

import (
	"context"
	"github.com/Sanchir01/microservice_sandjma_products/internal/domain/models"
	"log/slog"
)

type Products struct {
	log            *slog.Logger
	appProvider    AppProvider
	serviceProduct ServiceProduct
}

type ServiceProduct interface {
	Products(ctx context.Context) ([]models.Product, error)
}

type AppProvider interface {
	App(ctx context.Context, appId int64) (models.App, error)
}

func NewProducts(log *slog.Logger, appProvider AppProvider, serviceProduct ServiceProduct) Products {
	return Products{
		log:            log,
		appProvider:    appProvider,
		serviceProduct: serviceProduct,
	}
}
func (s *Products) GetAllProducts() ([]models.Product, error) {
	const op = "products.Products.GetAllProducts"

	log := s.log.With(slog.String("op", op), slog.String("method", "GetAllProducts"))
	log.Info("getting all products")
	products, err := s.serviceProduct.Products(context.Background())
	if err != nil {
		s.log.Error("failed to get products", err)
		return nil, err
	}

	return products, nil
}
