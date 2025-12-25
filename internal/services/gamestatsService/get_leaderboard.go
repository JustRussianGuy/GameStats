package gamestatsService

import (
	"context"
	"encoding/json"
	"time"
	"fmt"

	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/JustRussianGuy/GameStats/internal/redis"
)

func (s *GameStatsService) GetLeaderboard(ctx context.Context, limit int) ([]*models.PlayerStats, error) {
	fmt.Println("GetLeaderboard called, limit =", limit)
	key := fmt.Sprintf("leaderboard:%d", limit)

	// 1. Пробуем Redis
	cached, err := redis.RDB.Get(ctx, key).Result()
	if err == nil && cached != "" {
		var stats []*models.PlayerStats
		if err := json.Unmarshal([]byte(cached), &stats); err == nil && len(stats) > 0 {
			return stats, nil
		}
	}

	// 2. Берём из PostgreSQL
	stats, err := s.storage.GetLeaderboard(ctx, limit)
	if err != nil {
		return nil, err
	}

	// 3. Кладём в Redis
	data, _ := json.Marshal(stats)
	redis.RDB.Set(ctx, key, data, 10*time.Second)

	return stats, nil
}
