package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kavinrajmani/go-grpc-service/product"
	"github.com/kavinrajmani/go-grpc-service/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type myProductServiceServer struct {
	product.UnimplementedProductServer
}

type myUserServiceServer struct {
	user.UnimplementedUserServer
}

func (s myUserServiceServer) GetUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	return &user.UserResponse{
		Id:    req.Id,
		Name:  "Devdarshan Kavinraj",
		Email: "devdarshan@gmail.com",
	}, nil
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
	log.Println("main function entered")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("step 1")

	server := grpc.NewServer()
	productServer := &myProductServiceServer{}
	userServer := &myUserServiceServer{}

	log.Println("step 2")

	product.RegisterProductServer(server, productServer)
	user.RegisterUserServer(server, userServer)

	log.Println("step 3")

	// if err := server.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(server.Serve(lis))
	}()

	ctx := context.Background()
	mux := runtime.NewServeMux()
	conn, err := grpc.NewClient("0.0.0.0:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	err = product.RegisterProductHandler(ctx, mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	user.RegisterUserHandler(ctx, mux, conn)

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	gwServer.ListenAndServe()

}
