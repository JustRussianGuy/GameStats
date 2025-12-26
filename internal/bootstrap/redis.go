package bootstrap

import (
	"github.com/JustRussianGuy/GameStats/config"
	"github.com/JustRussianGuy/GameStats/internal/redis"
)

func InitRedis(cfg *config.Config) {
	redis.InitRedis(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.DB)
}
