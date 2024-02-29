package repository

import (
	"user-service/internal/core/interface/repository"
	"user-service/internal/lib/db"
	"user-service/internal/repository/postgres"
)

type RepositoryManager struct {
	repository.UserRepository
}

func NewRepositoryManager(db *db.Db) RepositoryManager {
	return RepositoryManager{
		postgres.NewRepo(db),
	}
}
