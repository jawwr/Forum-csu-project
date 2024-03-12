package repository

import (
	"context"
	"event-service/internal/core/repository/dbModel"
)

type EventRepository interface {
	GetEventsByUserId(ctx context.Context, userId int) ([]dbModel.PostEvent, error)
	SaveEvent(ctx context.Context, event dbModel.PostEvent) error
}
