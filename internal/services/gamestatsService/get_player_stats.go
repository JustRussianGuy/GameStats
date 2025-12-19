package gamestatsService

import (
	"context"
	"strconv"

	"github.com/JustRussianGuy/GameStats/internal/models"
)

func (s *GameStatsService) GetPlayerStats(ctx context.Context, playerID uint64) (*models.PlayerStats, error) {
	playerIDStr := strconv.FormatUint(playerID, 10)
	return s.storage.GetPlayerStats(ctx, playerIDStr)
}

