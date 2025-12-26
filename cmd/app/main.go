package main

import (
    "context"
	"fmt"
	"os"
    "time"

	"github.com/JustRussianGuy/GameStats/config"
	"github.com/JustRussianGuy/GameStats/internal/bootstrap"
    appredis "github.com/JustRussianGuy/GameStats/internal/redis"
    "github.com/JustRussianGuy/GameStats/internal/kafka"
)

func main() {
	cfg, err := config.LoadConfig(os.Getenv("configPath"))
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	// Redis FIRST
	bootstrap.InitRedis(cfg)

    ctx := context.Background()
    err = appredis.RDB.Set(ctx, "debug_key", "ok", 1*time.Minute).Err()
    if err != nil {
        fmt.Println("Redis SET error:", err)
    } else {
        fmt.Println("Redis SET ok")
    }

    val, err := appredis.RDB.Get(ctx, "debug_key").Result()
    if err != nil {
        fmt.Println("Redis GET error:", err)
    } else {
        fmt.Println("Redis GET debug_key =", val)
    }

	// PostgreSQL
	playerStorage := bootstrap.InitPGStorage(cfg)

	// Services
	gameStatsService := bootstrap.InitGameStatsService(playerStorage, cfg)

    brokers := []string{fmt.Sprintf("%s:%d", cfg.Kafka.Host, cfg.Kafka.Port)}
    producer := kafka.NewProducer(brokers, cfg.Kafka.PlayerEventsTopic)

	playerEventsProcessor := bootstrap.InitGameEventsProcessor(gameStatsService)
	kafkaConsumer := bootstrap.InitGameEventsConsumer(cfg, playerEventsProcessor)
	api := bootstrap.InitGameStatsAPI(gameStatsService, producer)

	bootstrap.AppRun(*api, kafkaConsumer)
}
