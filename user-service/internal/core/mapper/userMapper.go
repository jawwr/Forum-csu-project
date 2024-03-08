package mapper

import (
	modelDb "user-service/internal/core/repository/model"
	modelDto "user-service/internal/transport/model"
	pb "user-service/proto/generated/userService"
)

func FromUserDb(user *modelDb.User) *modelDto.UserResponse {
	return &modelDto.UserResponse{
		Id:    user.Id,
		Login: user.Login,
	}
}

func ToPbUser(user *modelDto.UserResponse) *pb.User {
	return &pb.User{
		Id:    int32(user.Id),
		Login: user.Login,
	}
}
