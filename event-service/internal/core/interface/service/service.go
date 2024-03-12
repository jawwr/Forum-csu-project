package service

import (
	"context"
	"event-service/internal/core/model"
)

type EventService interface {
	GetUserEvents(ctx context.Context, userId int) ([]model.EventResponse, error)
	SaveEvent(ctx context.Context, event model.PostEvent) error
}

type UserService interface {
	GetUserByToken(ctx context.Context, token string) (model.User, error)
	GetUserById(ctx context.Context, userId int) (model.User, error)
}
