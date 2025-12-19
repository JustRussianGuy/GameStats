package bootstrap

import (
	gameeventsprocessor "github.com/JustRussianGuy/GameStats/internal/services/processors/game_events_processor"
	"github.com/JustRussianGuy/GameStats/internal/services/gamestatsService"
)

func InitGameEventsProcessor(
	service *gamestatsService.GameStatsService,
) *gameeventsprocessor.GameEventsProcessor {

	return gameeventsprocessor.NewGameEventsProcessor(service)
}
