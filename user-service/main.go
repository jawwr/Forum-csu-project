package main

import (
	"context"
	"log"
	http2 "net/http"
	"time"
	"user-service/internal/core/repository"
	"user-service/internal/core/service"
	"user-service/internal/lib/db"
	"user-service/internal/transport/router"
)

func main() {
	timeout := time.Second * 10

	ctx := context.Background()

	withTimeout, _ := context.WithTimeout(ctx, timeout)

	database := db.New(withTimeout)

	repositoryManager := repository.NewManager(database)
	serviceManager := service.NewManager(repositoryManager)

	routes := router.InitRoutes(serviceManager)

	if err := http2.ListenAndServe(":8080", routes); err != nil {
		log.Fatal(err)
	}
}
