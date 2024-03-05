package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	http2 "net/http"
	"time"
	grpc2 "user-service/internal/core/grpc"
	"user-service/internal/core/repository"
	"user-service/internal/core/service"
	"user-service/internal/lib/db"
	"user-service/internal/transport/router"

	pb "user-service/proto"
)

func main() {
	timeout := time.Second * 10

	ctx := context.Background()

	withTimeout, _ := context.WithTimeout(ctx, timeout)

	database := db.New(withTimeout)

	repositoryManager := repository.NewManager(database)
	serviceManager := service.NewManager(repositoryManager)

	routes := router.InitRoutes(serviceManager)

	go startGrpc(serviceManager)

	if err := http2.ListenAndServe(":8080", routes); err != nil {
		log.Fatal(err)
	}
}

func startGrpc(manager service.Manager) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":9090"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, grpc2.NewUserService(manager))
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
