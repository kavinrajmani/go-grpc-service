package main

import (
	"context"
	"log"
	"net"

	"github.com/kavinrajmani/go-grpc-service/product"
	"google.golang.org/grpc"
)

type myProductServiceServer struct {
	product.UnimplementedProductServiceServer
}

func (s myProductServiceServer) CreateProduct(ctx context.Context, req *product.ProductRequest) (*product.ProductResponse, error) {

	name := req.Name

	return &product.ProductResponse{
		Id:          "200",
		Name:        name,
		Description: req.Description,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	service := &myProductServiceServer{}
	product.RegisterProductServiceServer(s, service)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
