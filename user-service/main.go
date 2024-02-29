package main

import (
	"context"
	"log"
	http2 "net/http"
	"time"
	"user-service/internal/core/service"
	"user-service/internal/lib/db"
	"user-service/internal/repository"
	"user-service/internal/transport/http"
)

func main() {

	timeout := time.Second * 10

	ctx := context.Background()

	withTimeout, _ := context.WithTimeout(ctx, timeout)

	database := db.New(withTimeout)

	manager := repository.NewRepositoryManager(database)

	serv := service.NewAuthService(manager.UserRepository)

	router := http.InitRoutes(serv)

	if err := http2.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
