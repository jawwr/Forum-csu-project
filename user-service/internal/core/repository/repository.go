package repository

import (
	"user-service/internal/core/interface/repository"
	"user-service/internal/core/repository/postgres"
	"user-service/internal/lib/db"
)

type Manager struct {
	repository.UserRepository
	repository.TokenRepository
	repository.SubscriberRepository
}

func NewManager(db *db.Db) Manager {
	return Manager{
		postgres.NewUserRepository(db),
		postgres.NewTokenRepository(db),
		postgres.NewSubscriberRepository(db),
	}
}
