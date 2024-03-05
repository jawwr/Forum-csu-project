package repository

import (
	"context"
	"user-service/internal/core/repository/model"
)

type TokenRepository interface {
	GetToken(ctx context.Context, token string) (*model.Token, error)
	GetValidTokenByUserId(ctx context.Context, userId int) ([]model.Token, error)
	SaveToken(ctx context.Context, token model.Token) error
}
