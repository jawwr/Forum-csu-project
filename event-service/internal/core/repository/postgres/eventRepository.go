package postgres

import (
	"context"
	"event-service/internal/core/interface/repository"
	"event-service/internal/core/repository/dbModel"
	"event-service/internal/lib/db"
	"fmt"
)

type _eventRepository struct {
	db *db.Db
}

func NewEventRepo(db *db.Db) repository.EventRepository {
	return _eventRepository{db}
}

func (eventRepository _eventRepository) SaveEvent(ctx context.Context, event dbModel.PostEvent) error {

	err := eventRepository.db.PgConn.QueryRow(ctx,
		`INSERT INTO public.events(title, post_id, user_id) values ($1,$2,$3)`, event.Title, event.PostId, event.UserId,
	)

	if err != nil {
		return fmt.Errorf("get post error: %s", err)
	}

	return nil
}

func (eventRepository _eventRepository) GetEventsByUserId(ctx context.Context, userId int) ([]dbModel.PostEvent, error) {
	var events []dbModel.PostEvent
	row, err := eventRepository.db.PgConn.Query(
		ctx,
		`SELECT id, title, post_id, user_id
			 FROM events
			 WHERE user_id = $1`,
		userId,
	)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var event dbModel.PostEvent
		if err := row.Scan(&event.Id, &event.Title, &event.PostId, &event.UserId); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
