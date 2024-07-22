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

func main() {
	CreateGRPCServer()
	CreateHTTPServer()
}

func CreateHTTPServer() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	conn, err := grpc.NewClient("0.0.0.0:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	err = product.RegisterProductHandler(ctx, mux, conn)
	if err != nil {
		log.Fatal(err)
	}

	user.RegisterUserHandler(ctx, mux, conn)
	if err != nil {
		log.Fatal(err)
	}

	if err == nil {
		gwServer := &http.Server{
			Addr:    ":8090",
			Handler: mux,
		}
		log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
		gwServer.ListenAndServe()
	}
}

func CreateGRPCServer() {
	server := grpc.NewServer()
	RegisterHdlr(server)

	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		lis, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Fatalln(server.Serve(lis))
	}()
}

func RegisterHdlr(s *grpc.Server) {
	proHdlr := product.NewHdlr()
	userHdlr := user.NewHdlr()
	product.RegisterProductServer(s, proHdlr)
	user.RegisterUserServer(s, userHdlr)
}
