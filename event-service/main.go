package main

import (
	"context"
	"event-service/internal/core/repository"
	"event-service/internal/core/service"
	"event-service/internal/lib/db"
	"event-service/internal/transport/kafka"
	"event-service/internal/transport/router"
	pb "event-service/proto/generated/userService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	http2 "net/http"
	"os"
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

	manager := repository.NewRepositoryManager(database)

	eventService := service.NewEventService(manager.EventRepository)
	userService := service.NewUserGrpcService(userServiceClient)

	routes := router.InitRoutes(eventService, userService)

	go kafka.Listen(eventService)

	if err := http2.ListenAndServe(":8080", routes); err != nil {
		log.Fatal(err)
	}
}
