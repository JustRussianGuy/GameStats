package eventsconsumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/segmentio/kafka-go"
)

// Интерфейс обработчика игровых событий
type PlayerEventProcessor interface {
	Handle(ctx context.Context, event *models.GameEvent) error
}

// Консьюмер Kafka для игровых событий
type PlayerEventConsumer struct {
	playerEventProcessor PlayerEventProcessor
	kafkaBroker          []string
	topicName            string
}

// Конструктор
func NewPlayerEventConsumer(processor PlayerEventProcessor, kafkaBroker []string, topicName string) *PlayerEventConsumer {
	return &PlayerEventConsumer{
		playerEventProcessor: processor,
		kafkaBroker:          kafkaBroker,
		topicName:            topicName,
	}
}

// Метод запуска потребления сообщений
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

		var event models.GameEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			slog.Error("JSON unmarshal error", "error", err)
			continue
		}

		if err := c.playerEventProcessor.Handle(ctx, &event); err != nil {
			slog.Error("Handle error", "error", err)
		}
	}
}


