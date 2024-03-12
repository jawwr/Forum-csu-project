package kafka

import (
	"context"
	"encoding/json"
	"event-service/internal/core/interface/service"
	"event-service/internal/core/model"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"log/slog"
	"os"
)

func Listen(service service.EventService) {
	ctx := context.Background()
	conn, err := kafka.DialLeader(ctx, "tcp", os.Getenv("KAFKA_HOST")+":"+os.Getenv("KAFKA_PORT"), "post", 0)

	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := conn.ReadMessage(10e6)

		if err != nil {
			slog.Error(fmt.Sprint("error receiving message: ", err.Error()))
			continue
		}
		var message model.PostEvent

		err = json.Unmarshal(msg.Value, &message)

		if err != nil {
			slog.Error("error decoding message")
			continue
		}

		if err != nil {
			slog.Error(fmt.Sprint("error creating post: ", err.Error()))
		}

		err = service.SaveEvent(ctx, message)

		if err != nil {
			slog.Error("error handling message")
		}

		conn.Offset()
	}
}
