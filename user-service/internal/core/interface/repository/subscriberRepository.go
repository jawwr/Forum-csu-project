package repository

import (
	"context"
	dbModel "user-service/internal/core/repository/model"
)

type SubscriberRepository interface {
	Subscribe(ctx context.Context, userId int, subscriberId int) error
	GetAllSubscriber(ctx context.Context, userId int) ([]*dbModel.User, error)
}
