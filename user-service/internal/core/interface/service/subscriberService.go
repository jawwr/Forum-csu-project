package service

import (
	"context"
	modelDto "user-service/internal/transport/model"
)

type SubscriberService interface {
	Subscribe(ctx context.Context, userId int, subscriberId int) error
	GetAllSubscribers(ctx context.Context, userId int) ([]modelDto.UserResponse, error)
}
