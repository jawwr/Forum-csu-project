package service

import (
	"context"
	"event-service/internal/core/interface/repository"
	"event-service/internal/core/interface/service"
	"event-service/internal/core/model"
	"event-service/internal/core/repository/dbModel"
)

type _eventService struct {
	repo repository.EventRepository
}

func (eventService _eventService) GetUserEvents(ctx context.Context, userId int) ([]model.EventResponse, error) {
	events, err := eventService.repo.GetEventsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	var postEvents []model.EventResponse

	for _, event := range events {
		postEvents = append(postEvents, model.EventResponse{
			Id:     event.Id,
			Title:  event.Title,
			PostId: event.PostId,
		})
	}

	return postEvents, nil
}

func (eventService _eventService) SaveEvent(ctx context.Context, event model.PostEvent) error {
	for _, user := range event.SubscriberIds {
		var dbEvent = dbModel.PostEvent{
			Title:  event.Title,
			PostId: event.PostId,
			UserId: user,
		}

		if err := eventService.repo.SaveEvent(ctx, dbEvent); err != nil {
			return err
		}

	}

	return nil
}

func NewEventService(kafka repository.EventRepository) service.EventService {
	return _eventService{repo: kafka}
}
