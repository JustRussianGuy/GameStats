package gamestatsService

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/JustRussianGuy/GameStats/internal/redis"
)

func (s *GameStatsService) GetLeaderboard(ctx context.Context, limit int) ([]*models.PlayerStats, error) {
	if limit <= 0 {
		limit = 10 // дефолтное значение
	}
	fmt.Println("GetLeaderboard called, limit =", limit)

	key := fmt.Sprintf("leaderboard:%d", limit)

	// 1. Пробуем Redis
	cached, err := redis.RDB.Get(context.Background(), key).Result()
	if err == nil && cached != "" {
		var stats []*models.PlayerStats
		if err := json.Unmarshal([]byte(cached), &stats); err == nil && len(stats) > 0 {
			fmt.Println("Cache hit for key:", key)
			return stats, nil
		}
	}

	// 2. Берём из PostgreSQL
	stats, err := s.storage.GetLeaderboard(ctx, limit)
	if err != nil {
		return nil, err
	}
	fmt.Println("Storage returned", len(stats), "players")

	// 3. Кладём в Redis с TTL 30 секунд
	data, _ := json.Marshal(stats)
	err = redis.RDB.Set(context.Background(), key, data, 30*time.Second).Err()
	if err != nil {
		fmt.Println("Redis SET error:", err)
	} else {
		fmt.Println("Redis cached key:", key)
	}

	return stats, nil
}
