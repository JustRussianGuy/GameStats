package bootstrap

import (
	
	"github.com/JustRussianGuy/GameStats/internal/redis"
	"github.com/JustRussianGuy/GameStats/config"
)

func InitRedis(cfg *config.Config) {
	redis.InitRedis(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.DB)
}
