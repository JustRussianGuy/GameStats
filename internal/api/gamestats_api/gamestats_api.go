package gamestats_api

import (
	"context"

	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/JustRussianGuy/GameStats/internal/pb/gamestats_api"
)

type gameStatsService interface {
	ProcessKillEvent(ctx context.Context, event *models.GameEvent) error
	GetPlayerStats(ctx context.Context, playerID uint64) (*models.PlayerStats, error)
	GetLeaderboard(ctx context.Context, limit int) ([]*models.PlayerStats, error)
}

// Реализация gRPC сервера
type GameStatsAPI struct {
	gamestats_api.UnimplementedGameStatsServiceServer
	service gameStatsService
}

func NewGameStatsAPI(service gameStatsService) *GameStatsAPI {
	return &GameStatsAPI{
		service: service,
	}
}

