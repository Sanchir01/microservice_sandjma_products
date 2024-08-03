package productsgrpc

import (
	"context"
	sandjmav1 "github.com/Sanchir01/protos_files_job/pkg/gen/products"
	"google.golang.org/grpc"
)

type productsServerApi struct {
	sandjmav1.UnimplementedProductServer
}

func NewProductsServerApi(gRPC *grpc.Server) {
	sandjmav1.RegisterProductServer(gRPC, &productsServerApi{})
}

func (s *productsServerApi) GetAllProducts(ctx context.Context, req *sandjmav1.Empty) (*sandjmav1.GetAllResponse, error) {
	return &sandjmav1.GetAllResponse{Id: "123344512", Name: "test", Price: 213, Version: 1}, nil
}
