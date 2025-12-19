package bootstrap

import (
	"fmt"

	"github.com/JustRussianGuy/GameStats/config"
	gameeventsconsumer "github.com/JustRussianGuy/GameStats/internal/consumer/eventsconsumer"
	gameeventsprocessor "github.com/JustRussianGuy/GameStats/internal/services/processors/game_events_processor"
)

func InitGameEventsConsumer(
	cfg *config.Config,
	processor *gameeventsprocessor.GameEventsProcessor,
) *gameeventsconsumer.GameEventsConsumer {

	brokers := []string{
		fmt.Sprintf("%s:%d", cfg.Kafka.Host, cfg.Kafka.Port),
	}

	return gameeventsconsumer.NewGameEventsConsumer(
		processor,
		brokers,
		cfg.Kafka.PlayerEventsTopic,
	)
}
