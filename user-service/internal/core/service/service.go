package service

import (
	"user-service/internal/core/interface/service"
	"user-service/internal/core/repository"
)

type Manager struct {
	service.UserService
	service.TokenService
	service.AuthService
}

func NewManager(manager repository.Manager) Manager {
	tokenService := NewTokenService(manager.TokenRepository)
	return Manager{
		UserService:  NewUserService(manager.UserRepository),
		TokenService: tokenService,
		AuthService:  NewAuthService(manager.UserRepository, tokenService),
	}
}
