package bootstrap

import (
	"context"

	"github.com/JustRussianGuy/GameStats/config"
	"github.com/JustRussianGuy/GameStats/internal/services/gamestatsService"
	"github.com/JustRussianGuy/GameStats/internal/storage/pgstorage"
)

func InitGameStatsService(
	storage *pgstorage.PGstorage,
	cfg *config.Config,
) *gamestatsService.GameStatsService {

	return gamestatsService.NewGameStatsService(
		context.Background(),
		storage, // PGstorage должен реализовывать PlayerStatsStorage
		cfg,     // передаем весь конфиг
	)
}
