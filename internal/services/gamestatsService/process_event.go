package gamestatsService

import (
	"context"

	"github.com/JustRussianGuy/GameStats/internal/models"
)

func (s *GameStatsService) ProcessGameEvent(
	ctx context.Context,
	event *models.GameEvent,
) error {

	// Killer +1 kill
	if err := s.storage.IncrementKill(ctx, event.KillerID); err != nil {
		return err
	}

	// Victim +1 death
	if err := s.storage.IncrementDeath(ctx, event.VictimID); err != nil {
		return err
	}

	return nil
}
