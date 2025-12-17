package eventsconsumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/segmentio/kafka-go"
)

type PlayerEventProcessor interface {
	Handle(ctx context.Context, event *models.PlayerEvent) error
}

type PlayerEventConsumer struct {
	playerEventProcessor PlayerEventProcessor
	kafkaBroker          []string
	topicName            string
}

func NewPlayerEventConsumer(processor PlayerEventProcessor, kafkaBroker []string, topicName string) *PlayerEventConsumer {
	return &PlayerEventConsumer{
		playerEventProcessor: processor,
		kafkaBroker:          kafkaBroker,
		topicName:            topicName,
	}
}

func (c *PlayerEventConsumer) Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           c.kafkaBroker,
		GroupID:           "GameStats_group",
		Topic:             c.topicName,
		HeartbeatInterval: 3 * time.Second,
		SessionTimeout:    30 * time.Second,
	})
	defer r.Close()

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			slog.Error("PlayerEventConsumer.consume error", "error", err.Error())
			continue
		}

		var event *models.PlayerEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			slog.Error("json unmarshal error", "error", err)
			continue
		}

		if err := c.playerEventProcessor.Handle(ctx, event); err != nil {
			slog.Error("Handle error", "error", err)
		}
	}
}
