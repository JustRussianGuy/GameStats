package bootstrap

import (
	"fmt"

	"github.com/JustRussianGuy/GameStats/config"
	playereventsconsumer "github.com/JustRussianGuy/GameStats/internal/consumer/eventsconsumer"
	playereventsprocessor "github.com/JustRussianGuy/GameStats/internal/services/processors/game_events_processor"
)

func InitPlayerEventsConsumer(
	cfg *config.Config,
	processor *playereventsprocessor.PlayerEventsProcessor,
) *playereventsconsumer.PlayerEventsConsumer {

	brokers := []string{
		fmt.Sprintf("%s:%d", cfg.Kafka.Host, cfg.Kafka.Port),
	}

	return playereventsconsumer.NewPlayerEventsConsumer(
		processor,
		brokers,
		cfg.Kafka.PlayerEventsTopic,
	)
}
