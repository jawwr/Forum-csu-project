package repository

import (
	"post-service/internal/core/interface/repository"
	"post-service/internal/core/repository/postgres"
	"post-service/internal/lib/db"
)

type RepositoryManager struct {
	repository.PostRepository
}

func NewRepositoryManager(db *db.Db) RepositoryManager {
	return RepositoryManager{
		postgres.NewPostRepo(db),
	}
}
