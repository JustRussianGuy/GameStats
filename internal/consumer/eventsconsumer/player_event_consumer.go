package eventsconsumer

import "context"

type PlayerEventProcessor interface {
	Handle(ctx context.Context, event *PlayerEvent) error
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

