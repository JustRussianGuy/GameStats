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
type GameEventsProcessor interface {
	Handle(ctx context.Context, event *models.GameEvent) error
}

// Консьюмер Kafka для игровых событий
type GameEventsConsumer struct {
	playerEventProcessor GameEventsProcessor
	kafkaBroker          []string
	topicName            string
}

// Конструктор
func NewGameEventsConsumer(processor GameEventsProcessor, kafkaBroker []string, topicName string) *GameEventsConsumer {
	return &GameEventsConsumer{
		playerEventProcessor: processor,
		kafkaBroker:          kafkaBroker,
		topicName:            topicName,
	}
}

// Метод запуска потребления сообщений
func (c *GameEventsConsumer) Consume(ctx context.Context) {
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


