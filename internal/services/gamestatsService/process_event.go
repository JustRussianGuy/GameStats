package gamestatsService

import (
	"context"
	"strconv"

	"github.com/JustRussianGuy/GameStats/internal/models"
)

func (s *GameStatsService) ProcessGameEvent(
	ctx context.Context,
	event *models.GameEvent,
) error {

	// Преобразуем uint64 ID в string для хранения в PostgreSQL
	killerID := strconv.FormatUint(event.KillerID, 10)
	victimID := strconv.FormatUint(event.VictimID, 10)

	// Killer +1 kill
	if err := s.storage.IncrementKill(ctx, killerID); err != nil {
		return err
	}

	// Victim +1 death
	if err := s.storage.IncrementDeath(ctx, victimID); err != nil {
		return err
	}

	return nil
}

