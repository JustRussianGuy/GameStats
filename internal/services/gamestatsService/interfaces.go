package gamestatsService

import (
	"context"

	"github.com/JustRussianGuy/GameStats/internal/models"
)

type PlayerStatsStorage interface {
	IncrementKill(ctx context.Context, playerID string) error
	IncrementDeath(ctx context.Context, playerID string) error
	GetPlayerStats(ctx context.Context, playerID string) (*models.PlayerStats, error)
	GetLeaderboard(ctx context.Context, limit int) ([]*models.PlayerStats, error)
}
