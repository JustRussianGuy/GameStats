package gamestatsService

import (
	"context"

	"github.com/JustRussianGuy/GameStats/internal/models"
)

func (s *GameStatsService) GetPlayerStats(ctx context.Context, playerID string) (*models.PlayerStats, error) {
	return s.storage.GetPlayerStats(ctx, playerID)
}
