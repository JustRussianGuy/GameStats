package bootstrap

import (
	playereventsprocessor "github.com/JustRussianGuy/GameStats/internal/services/processors/player_events_processor"
	"github.com/JustRussianGuy/GameStats/internal/services/gamestatsService"
)

func InitPlayerEventsProcessor(
	service *gamestats.GameStatsService,
) *playereventsprocessor.PlayerEventsProcessor {

	return playereventsprocessor.NewPlayerEventsProcessor(service)
}
