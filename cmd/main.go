package main

import (
	"log"
	"net"

	"go-grpc/cmd/config"
	"go-grpc/cmd/services"
	productPb "go-grpc/pb/product"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)


func main() {

	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to Listen %v", err.Error())
	}

	//connect ke database
	db := config.ConnectDatabase()

	grpcServer := grpc.NewServer()
	//connect productService to db
	productService := services.ProductService{DB: db}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server started at %v", netListen.Addr())
	if err :=grpcServer.Serve(netListen); err !=nil {
		log.Fatalf("Failed to serve %v", err.Error())
	}
}