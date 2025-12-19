package gamestatsService

import (
	"context"

	"github.com/JustRussianGuy/GameStats/internal/models"
)

func (s *Service) GetLeaderboard(
	ctx context.Context,
	limit int,
) ([]*models.PlayerStats, error) {

	if limit <= 0 {
		limit = 10
	}

	return s.storage.GetLeaderboard(ctx, limit)
}
