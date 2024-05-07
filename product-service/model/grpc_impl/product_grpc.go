package grpc_impl

import (
	"context"

	pb "github.com/sainath123112/e-commerce-backend/product-service/proto"
	"github.com/sainath123112/e-commerce-backend/product-service/service"
)

type ProductServiceServer struct {
	pb.UnimplementedProductServiceServer
}

func (s *ProductServiceServer) ValidateProduct(ctx context.Context, in *pb.ValidateProductRequest) (*pb.ValidateProductResponse, error) {
	isProductValid, err := service.ValidateProductById(in.ProductId)
	return &pb.ValidateProductResponse{IsProductValid: isProductValid}, err
}
