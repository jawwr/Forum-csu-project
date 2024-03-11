package service

import (
	"context"
	"io"
	"log"
	"post-service/internal/core/interface/service"
	"post-service/internal/core/model"
	"post-service/proto/generated/subscriberService"
)

type _subscriberService struct {
	subscriberClient subscriberService.SubscriberServiceClient
}

func NewSubscriberService(subscriberClient subscriberService.SubscriberServiceClient) service.SubscriberService {
	return _subscriberService{subscriberClient: subscriberClient}
}

func (s _subscriberService) GetAllUserSubscribers(ctx context.Context, userId int) ([]model.User, error) {
	subscribers, err := s.subscriberClient.GetAllUserSubscribers(ctx, &subscriberService.UserRequest{Id: int32(userId)})
	if err != nil {
		return nil, err
	}

	var users []model.User
	for {
		resp, err := subscribers.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Ошибка при получении сообщения: %v", err)
		}

		users = append(users, model.User{
			Id:    int(resp.Id),
			Login: resp.Login,
		})
	}

	return users, nil
}
