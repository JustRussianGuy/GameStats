package bootstrap

import (
    "github.com/JustRussianGuy/GameStats/internal/api/gamestats_api"
    "github.com/JustRussianGuy/GameStats/internal/kafka"
    "github.com/JustRussianGuy/GameStats/internal/services/gamestatsService"
)

// InitGameStatsAPI инициализирует gRPC API для GameStats
func InitGameStatsAPI(
    service *gamestatsService.GameStatsService,
    producer *kafka.Producer,
) *gamestats_api.GameStatsAPI {
    return gamestats_api.NewGameStatsAPI(service, producer)
}
