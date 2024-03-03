package mapper

import (
	modelDb "user-service/internal/core/repository/model"
	modelDto "user-service/internal/transport/model"
)

func FromUserDb(user *modelDb.User) *modelDto.UserResponse {
	return &modelDto.UserResponse{
		Id:    user.Id,
		Login: user.Login,
	}
}
