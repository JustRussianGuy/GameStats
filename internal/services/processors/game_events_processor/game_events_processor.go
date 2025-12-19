package game_events_processor

import (
	"context"

	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/JustRussianGuy/GameStats/internal/services/gamestatsService"
)

type GameStatsService interface {
	ProcessGameEvent(ctx context.Context, event *models.GameEvent) error
}

type GameEventsProcessor struct {
	service GameStatsService
}

func NewGameEventsProcessor(service GameStatsService) *GameEventsProcessor {
	return &GameEventsProcessor{
		service: service,
	}
}
