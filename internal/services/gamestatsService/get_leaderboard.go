package gamestatsService

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/JustRussianGuy/GameStats/internal/models"
	appredis "github.com/JustRussianGuy/GameStats/internal/redis"
	goredis "github.com/redis/go-redis/v9"
)

func (s *GameStatsService) GetLeaderboard(
	ctx context.Context,
	limit int,
) ([]*models.PlayerStats, error) {

	if limit <= 0 {
		limit = 10
	}

	key := fmt.Sprintf("leaderboard:%d", limit)

	// --- Redis GET ---
	cached, err := appredis.RDB.Get(ctx, key).Result()
	if err == nil {
		var stats []*models.PlayerStats
		if err := json.Unmarshal([]byte(cached), &stats); err == nil {
			fmt.Println("[Redis] cache hit:", key)
			return stats, nil
		}
	} else if err == goredis.Nil {
		fmt.Println("[Redis] cache miss:", key)
	} else {
		fmt.Println("[Redis] GET error:", err)
	}

	// --- PostgreSQL ---
	stats, err := s.storage.GetLeaderboard(ctx, limit)
	if err != nil {
		return nil, err
	}

	// --- Redis SET ---
	data, err := json.Marshal(stats)
	if err == nil {
		if err := appredis.RDB.Set(
			ctx,
			key,
			data,
			5*time.Minute,
		).Err(); err == nil {
			fmt.Println("[Redis] cached:", key)
		}
	}

	return stats, nil
}

