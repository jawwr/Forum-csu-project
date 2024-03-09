package mapper

import (
	modelDb "user-service/internal/core/repository/model"
	modelDto "user-service/internal/transport/model"
	pbSubscriberService "user-service/proto/generated/subscriberService"
	pbUserService "user-service/proto/generated/userService"
)

func FromUserDb(user *modelDb.User) *modelDto.UserResponse {
	return &modelDto.UserResponse{
		Id:    user.Id,
		Login: user.Login,
	}
}

func ToPbUser(user *modelDto.UserResponse) *pbUserService.User {
	return &pbUserService.User{
		Id:    int32(user.Id),
		Login: user.Login,
	}
}

func ToPbSubscriber(user *modelDto.UserResponse) *pbSubscriberService.SubscriberResponse {
	return &pbSubscriberService.SubscriberResponse{
		Id:    int32(user.Id),
		Login: user.Login,
	}
}
