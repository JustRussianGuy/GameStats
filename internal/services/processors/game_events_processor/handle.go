package game_events_processor

import (
	"context"

	"github.com/JustRussianGuy/GameStats/internal/models"
)

func (p *GameEventsProcessor) Handle(ctx context.Context, event *models.GameEvent) error {
	return p.service.ProcessGameEvent(ctx, event)
}

