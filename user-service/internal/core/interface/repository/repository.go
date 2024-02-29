package repository

import (
	"context"
)

type UserRepository interface {
	GetUser(ctx context.Context, login, hashPassword string) (string, error)
	CreateUser(ctx context.Context, login, hashPassword string) (string, error)
}
