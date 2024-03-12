package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"post-service/internal/core/interface/repository"
	"post-service/internal/core/model"
)

type _eventRepository struct {
	conn *kafka.Conn
}

func NewEventRepo(host string) repository.EventRepository {
	conn, err := kafka.DialLeader(context.Background(), "tcp", host, "post", 0)

	if err != nil {
		log.Fatal(err)
	}

	return _eventRepository{conn: conn}

}

func (repo _eventRepository) SendNewPostEvent(ctx context.Context, event model.PostEvent) error {
	bytes, err := json.Marshal(event)

	if err != nil {
		return errors.New("error marshaling message")
	}

	_, err = repo.conn.WriteMessages(
		kafka.Message{
			Topic: "post",
			Value: bytes,
		})

	if err != nil {
		return fmt.Errorf("error sending message %s", err.Error())
	}

	return err
}
