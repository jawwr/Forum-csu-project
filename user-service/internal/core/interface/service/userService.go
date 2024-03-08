package service

import (
	"context"
	"user-service/internal/transport/model"
)

type UserService interface {
	GetUserById(ctx context.Context, id int) (*model.UserResponse, error)
	GetAllUsers(ctx context.Context) ([]model.UserResponse, error)
	CreateUser(ctx context.Context, user *model.UserCredentials) error
}
