package postgres

import (
	"context"
	"user-service/internal/core/interface/repository"
	"user-service/internal/core/repository/model"
	"user-service/internal/lib/db"
)

type _subscriberRepository struct {
	*db.Db
}

func NewSubscriberRepository(db *db.Db) repository.SubscriberRepository {
	return _subscriberRepository{db}
}

func (repo _subscriberRepository) Subscribe(ctx context.Context, userId int, subscriberId int) error {
	_, err := repo.PgConn.Query(
		ctx,
		`INSERT INTO subscribers(user_id, subscriber_id)
			 VALUES ($1, $2)`,
		userId, subscriberId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo _subscriberRepository) GetAllSubscriber(ctx context.Context, userId int) ([]*model.User, error) {
	var subscribers []*model.User
	rows, err := repo.PgConn.Query(
		ctx,
		`SELECT users.*
			 FROM subscribers
			 JOIN users
			 ON subscribers.subscriber_id = users.id
			 WHERE subscribers.user_id = $1`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Login, &user.Password); err != nil {
			return nil, err
		}
		subscribers = append(subscribers, &user)
	}
	return subscribers, nil
}
