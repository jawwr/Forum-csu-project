package main

import (
	"context"
	"log"
	http2 "net/http"
	"post-service/internal/core/repository"
	"post-service/internal/core/service"
	"post-service/internal/lib/db"
	"post-service/internal/transport/router"
	"time"
)

func main() {
	timeout := time.Second * 10

	ctx := context.Background()

	withTimeout, _ := context.WithTimeout(ctx, timeout)

	database := db.New(withTimeout)

	manager := repository.NewRepositoryManager(database)

	postServ := service.NewPostService(manager.PostRepository)

	routes := router.InitRoutes(postServ)

	if err := http2.ListenAndServe(":8081", routes); err != nil {
		log.Fatal(err)
	}
}
