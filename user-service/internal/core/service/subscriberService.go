package service

import (
	"context"
	"user-service/internal/core/interface/repository"
	"user-service/internal/core/interface/service"
	"user-service/internal/core/mapper"
	"user-service/internal/transport/model"
)

type _subscriberService struct {
	repository repository.SubscriberRepository
}

func NewSubscriberService(repository repository.SubscriberRepository) service.SubscriberService {
	return _subscriberService{repository}
}

func (service _subscriberService) Subscribe(ctx context.Context, userId int, subscriberId int) error {
	return service.repository.Subscribe(ctx, userId, subscriberId)
}

func (service _subscriberService) GetAllSubscribers(ctx context.Context, userId int) ([]model.UserResponse, error) {
	savedSubscribers, err := service.repository.GetAllSubscriber(ctx, userId)
	if err != nil {
		return nil, err
	}
	subscribers := make([]model.UserResponse, 0)
	for _, subscriber := range savedSubscribers {
		user := mapper.FromUserDb(subscriber)
		subscribers = append(subscribers, *user)
	}
	return subscribers, nil
}
