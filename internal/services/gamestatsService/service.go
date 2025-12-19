package gamestatsService

import (
	"context"

	"github.com/JustRussianGuy/GameStats/config"
)

type GameStatsService struct {
	storage      PlayerStatsStorage
	killPoints   int
	deathPenalty int
}

func NewGameStatsService(
	ctx context.Context,
	storage PlayerStatsStorage,
	cfg *config.Config,
) *GameStatsService {
	return &GameStatsService{
		storage:      storage,
		killPoints:   cfg.GameSettings.KillPoints,
		deathPenalty: cfg.GameSettings.DeathPenalty,
	}
}
