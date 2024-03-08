package grpc

import (
	"context"
	"user-service/internal/core/interface/service"
	"user-service/internal/core/mapper"
	serviceManager "user-service/internal/core/service"
	pb "user-service/proto/generated/userService"
)

type _userServiceImpl struct {
	pb.UnimplementedUserServiceServer
	userService  service.UserService
	tokenService service.TokenService
}

func NewUserService(manager serviceManager.Manager) pb.UserServiceServer {
	return _userServiceImpl{
		userService:  manager.UserService,
		tokenService: manager.TokenService,
	}
}

func (s _userServiceImpl) GetUserByToken(ctx context.Context, token *pb.Token) (*pb.User, error) {
	savedToken, err := s.tokenService.GetToken(ctx, token.Token)
	if err != nil {
		return nil, err
	}
	user, err := s.userService.GetUserById(ctx, savedToken.UserId)
	if err != nil {
		return nil, err
	}
	return mapper.ToPbUser(user), nil
}
