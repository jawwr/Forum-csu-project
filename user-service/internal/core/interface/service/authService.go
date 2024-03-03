package service

import (
	"context"
	"user-service/internal/transport/model"
)

type AuthService interface {
	Register(ctx context.Context, user model.UserCredentials) (string, error)
	Login(ctx context.Context, user model.UserCredentials) (string, error)
}
