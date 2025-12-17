package main

import (
    "fmt"
    "os"

    "github.com/JustRussianGuy/GameStats/config"
    "github.com/JustRussianGuy/GameStats/internal/bootstrap"
)

func main() {

    cfg, err := config.LoadConfig(os.Getenv("configPath"))
    if err != nil {
        panic(fmt.Sprintf("failed to load config: %v", err))
    }

    // PostgreSQL storage
    playerStorage := bootstrap.InitPostgresPlayerStorage(cfg)

    // Main business service
    gameStatsService := bootstrap.InitGameStatsService(playerStorage, cfg)

    // Processor for player events (kill/death)
    playerEventsProcessor := bootstrap.InitPlayerEventsProcessor(gameStatsService)

    // Kafka consumer for game events
    kafkaConsumer := bootstrap.InitPlayerEventsConsumer(cfg, playerEventsProcessor)

    // API (POST /events, GET /stats/{id}, GET /leaderboard)
    api := bootstrap.InitGameStatsAPI(gameStatsService)

    // Run HTTP API + Kafka consumer
    bootstrap.AppRun(*api, kafkaConsumer)
}