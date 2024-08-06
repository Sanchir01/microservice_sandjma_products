package productsgrpc

import (
	"context"
	sandjmav1 "github.com/Sanchir01/protos_files_job/pkg/gen/golang/products"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productsServerApi struct {
	sandjmav1.UnimplementedProductServer
	products Products
}

type Products interface {
	GetAllProducts(ctx context.Context, req *sandjmav1.Empty) (*sandjmav1.GetAllResponse, error)
}

func NewProductsServerApi(gRPC *grpc.Server, products Products) {
	sandjmav1.RegisterProductServer(gRPC, &productsServerApi{products: products})
}

func (s *productsServerApi) GetAllProducts(ctx context.Context, req *sandjmav1.Empty) (*sandjmav1.GetAllResponse, error) {
	products, err := s.products.GetAllProducts(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &sandjmav1.GetAllResponse{Products: products.Products}, nil
}
