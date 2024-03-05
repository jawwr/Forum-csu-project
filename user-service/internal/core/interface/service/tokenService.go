package service

import (
	"context"
	modelDb "user-service/internal/core/repository/model"
)

type TokenService interface {
	GetToken(ctx context.Context, token string) (*modelDb.Token, error)
	SaveToken(ctx context.Context, token modelDb.Token) error
	GenerateToken(ctx context.Context, user modelDb.User) (string, error)
}
