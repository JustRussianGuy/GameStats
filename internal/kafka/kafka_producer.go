package kafka

import (
    "context"
    "encoding/json"
    "github.com/JustRussianGuy/GameStats/internal/models"
    "github.com/segmentio/kafka-go"
)

type Producer struct {
    writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
    return &Producer{
        writer: &kafka.Writer{
            Addr:     kafka.TCP(brokers...),
            Topic:    topic,
            Balancer: &kafka.LeastBytes{},
        },
    }
}

func (p *Producer) ProduceEvent(ctx context.Context, event *models.GameEvent) error {
    data, _ := json.Marshal(event)
    return p.writer.WriteMessages(ctx, kafka.Message{
        Value: data,
    })
}
