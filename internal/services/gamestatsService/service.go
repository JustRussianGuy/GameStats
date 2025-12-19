package gamestatsService

import (
	"context"

	"github.com/JustRussianGuy/GameStats/config"
)

type Service struct {
	storage       PlayerStatsStorage
	killPoints    int
	deathPenalty  int
}

func NewService(
	ctx context.Context,
	storage PlayerStatsStorage,
	cfg *config.Config,
) *Service {
	return &Service{
		storage:      storage,
		killPoints:   cfg.GameSettings.KillPoints,
		deathPenalty: cfg.GameSettings.DeathPenalty,
	}
}
