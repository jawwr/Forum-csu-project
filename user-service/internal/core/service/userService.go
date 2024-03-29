package service

import (
	"context"
	"user-service/internal/core/interface/repository"
	"user-service/internal/core/interface/service"
	"user-service/internal/core/mapper"
	"user-service/internal/transport/model"
)

type _userService struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return _userService{repository: userRepository}
}

func (service _userService) GetUserById(ctx context.Context, id int) (*model.UserResponse, error) {
	savedUser, err := service.repository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.FromUserDb(savedUser), nil
}

func (service _userService) GetAllUsers(ctx context.Context) ([]model.UserResponse, error) {
	savedUsers, err := service.repository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]model.UserResponse, 0)
	for _, user := range savedUsers {
		dto := mapper.FromUserDb(user)
		users = append(users, *dto)
	}
	return users, nil
}

func (service _userService) CreateUser(ctx context.Context, user *model.UserCredentials) error {
	_, err := service.repository.CreateUser(ctx, *user)
	if err != nil {
		return err
	}
	return nil
}
