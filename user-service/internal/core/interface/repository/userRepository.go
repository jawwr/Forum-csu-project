package repository

import (
	"context"
	modelDb "user-service/internal/core/repository/model"
	"user-service/internal/transport/model"
)

type UserRepository interface {
	GetUserByCredentials(ctx context.Context, userDto model.UserCredentials) (*modelDb.User, error)
	GetUserById(ctx context.Context, userId int) (*modelDb.User, error)
	CreateUser(ctx context.Context, userDto model.UserCredentials) (*modelDb.User, error)
}
