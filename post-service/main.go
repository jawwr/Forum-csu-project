package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	http2 "net/http"
	"os"
	"post-service/internal/core/repository"
	"post-service/internal/core/service"
	"post-service/internal/lib/db"
	"post-service/internal/transport/router"
	"post-service/proto/generated/subscriberService"
	pb "post-service/proto/generated/userService"
	"time"
)

func main() {
	conn, err := grpc.Dial(os.Getenv("USER_SERVICE_GRPC_HOST")+
		":"+os.Getenv("USER_SERVICE_GRPC_PORT"),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userServiceClient := pb.NewUserServiceClient(conn)

	timeout := time.Second * 10

	ctx := context.Background()

	withTimeout, _ := context.WithTimeout(ctx, timeout)

	database := db.New(withTimeout)

	manager := repository.NewRepositoryManager(database,
		os.Getenv("KAFKA_HOST")+":"+os.Getenv("KAFKA_PORT"))

	subscriberClient := subscriberService.NewSubscriberServiceClient(conn)
	subscriberService1 := service.NewSubscriberService(subscriberClient)
	postServ := service.NewPostService(manager.PostRepository, manager.EventRepository, subscriberService1)
	grpcService := service.NewUserGrpcService(userServiceClient)

	routes := router.InitRoutes(postServ, grpcService)

	if err := http2.ListenAndServe(":8080", routes); err != nil {
		log.Fatal(err)
	}
}
