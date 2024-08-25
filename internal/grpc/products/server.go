package productsgrpc

import (
	"context"
	"github.com/Sanchir01/microservice_sandjma_products/internal/domain/models"
	sandjmav1 "github.com/Sanchir01/protos_files_job/pkg/gen/golang/products"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

type productsServerApi struct {
	sandjmav1.UnimplementedProductServer
	products Products
}

type Products interface {
	AllProducts(ctx context.Context) ([]models.Product, error)
}

func NewProductsServerApi(gRPC *grpc.Server, products Products) {
	sandjmav1.RegisterProductServer(gRPC, &productsServerApi{products: products})
}

func (s *productsServerApi) GetAllProducts(ctx context.Context, _ *emptypb.Empty) (*sandjmav1.GetAllProductsResponse, error) {
	products, err := s.products.AllProducts(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	slog.Warn("sda", products)
	return &sandjmav1.GetAllProductsResponse{Products: nil}, nil
}
