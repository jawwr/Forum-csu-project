package repository

import (
	"event-service/internal/core/interface/repository"
	"event-service/internal/core/repository/postgres"
	"event-service/internal/lib/db"
)

type RepositoryManager struct {
	repository.EventRepository
}

func NewRepositoryManager(db *db.Db) RepositoryManager {
	return RepositoryManager{
		postgres.NewEventRepo(db),
	}
}
