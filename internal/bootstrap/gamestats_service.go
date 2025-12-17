package bootstrap

import (
	"context"

	"github.com/JustRussianGuy/GameStats/config"
	"github.com/JustRussianGuy/GameStats/internal/services/gamestats"
	"github.com/JustRussianGuy/GameStats/internal/storage/pgstorage"
)

func InitGameStatsService(
	storage *pgstorage.PGstorage,
	cfg *config.Config,
) *gamestats.GameStatsService {

	return gamestats.NewGameStatsService(
		context.Background(),
		storage,
		cfg.GameSettings.KillPoints,
		cfg.GameSettings.DeathPenalty,
	)
}
