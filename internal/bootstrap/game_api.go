package bootstrap

import (
	server "github.com/JustRussianGuy/GameStats/internal/api/gamestats_api"
	"github.com/JustRussianGuy/GameStats/internal/services/gamestatsService"
)

func InitGameStatsAPI(
	service *gamestats.GameStatsService,
) *gamestats_api.GameStatsAPI {

	return gamestats_api.NewGameStatsAPI(service)
}
