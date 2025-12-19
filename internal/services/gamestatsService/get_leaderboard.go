package gamestatsService

import (
	"context"
	//"encoding/json"
	//"time"

	"github.com/JustRussianGuy/GameStats/internal/models"
	//"github.com/JustRussianGuy/GameStats/internal/redis"
)

func (s *GameStatsService) GetLeaderboard(ctx context.Context, limit int) ([]*models.PlayerStats, error) {
	// Пытаемся получить из Redis
	/*
	cached, err := redis.RDB.Get(ctx, "leaderboard").Result()
	if err == nil {
		var stats []*models.PlayerStats
		json.Unmarshal([]byte(cached), &stats)
		return stats, nil
	}
	*/
	// Если нет в кэше, берём из storage
	stats, err := s.storage.GetLeaderboard(ctx, limit)
	if err != nil {
		return nil, err
	}
	/*
	// Сохраняем в Redis
	data, _ := json.Marshal(stats)
	redis.RDB.Set(ctx, "leaderboard", data, 5*time.Second)
	*/
	return stats, nil
}
