package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"log/slog"
	"user-service/internal/core/helper"
	"user-service/internal/core/interface/repository"
	"user-service/internal/core/interface/service"
	"user-service/internal/transport/model"
)

type _authService struct {
	repo         repository.UserRepository
	tokenService service.TokenService
}

func NewAuthService(repo repository.UserRepository, tokenService service.TokenService) service.AuthService {
	return _authService{
		repo:         repo,
		tokenService: tokenService,
	}
}

func (service _authService) Register(ctx context.Context, user model.UserCredentials) (string, error) {
	user.Password = generateHashPassword(user.Password)

	savedUser, err := service.repo.CreateUser(ctx, user)

	if err != nil {
		slog.Error(err.Error())
		return "", errors.New("не смогли создать пользователя")
	}

	token, err := service.tokenService.GenerateToken(ctx, *savedUser)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service _authService) Login(ctx context.Context, user model.UserCredentials) (string, error) {
	user.Password = generateHashPassword(user.Password)
	savedUser, err := service.repo.GetUserByCredentials(ctx, user)
	if err != nil {
		return "", err
	}
	token, err := service.tokenService.GenerateToken(ctx, *savedUser)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(helper.Salt)))
}
