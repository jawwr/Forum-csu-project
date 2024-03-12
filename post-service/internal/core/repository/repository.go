package repository

import (
	"post-service/internal/core/interface/repository"
	"post-service/internal/core/repository/kafka"
	"post-service/internal/core/repository/postgres"
	"post-service/internal/lib/db"
)

type RepositoryManager struct {
	repository.PostRepository
	repository.EventRepository
}

func NewRepositoryManager(db *db.Db, host string) RepositoryManager {
	return RepositoryManager{
		postgres.NewPostRepo(db),
		kafka.NewEventRepo(host),
	}
}
