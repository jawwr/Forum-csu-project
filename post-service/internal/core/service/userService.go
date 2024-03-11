package service

import (
	"context"
	"post-service/internal/core/interface/service"
	"post-service/internal/core/model"
	"post-service/proto/generated/userService"
)

type _userGrpcService struct {
	userClient userService.UserServiceClient
}

func NewUserGrpcService(userClient userService.UserServiceClient) service.UserService {
	return _userGrpcService{userClient: userClient}
}

func (s _userGrpcService) GetUserByToken(ctx context.Context, token string) (model.User, error) {
	user, err := s.userClient.GetUserByToken(ctx, &userService.Token{Token: token})
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		Id:    int(user.Id),
		Login: user.Login,
	}, nil
}

func (s _userGrpcService) GetUserById(ctx context.Context, userId int) (model.User, error) {
	user, err := s.userClient.GetUserById(ctx)
	if err != nil {
		return model.User{}, err
	}

	if err := user.Send(&userService.UserRequest{Id: int32(userId)}); err != nil {
		return model.User{}, err
	}

	resp, err := user.Recv()
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		Id:    int(resp.Id),
		Login: resp.Login,
	}, nil
}
