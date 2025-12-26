package gamestatsService

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JustRussianGuy/GameStats/internal/models"
	appredis "github.com/JustRussianGuy/GameStats/internal/redis"
)

func (s *GameStatsService) ProcessGameEvent(
	ctx context.Context,
	event *models.GameEvent,
) error {

	killerID := strconv.FormatUint(event.KillerID, 10)
	victimID := strconv.FormatUint(event.VictimID, 10)

	if err := s.storage.IncrementKill(ctx, killerID); err != nil {
		return err
	}
	if err := s.storage.IncrementDeath(ctx, victimID); err != nil {
		return err
	}

	// --- Cache invalidation ---
	err := appredis.InvalidateByPattern(ctx, "leaderboard:*")
	if err != nil {
		fmt.Println("[Redis] invalidate leaderboard error:", err)
	} else {
		fmt.Println("[Redis] invalidate leaderboard") // <- добавляем здесь
	}

	_ = appredis.RDB.Del(ctx, "player:"+killerID).Err()
	_ = appredis.RDB.Del(ctx, "player:"+victimID).Err()

	return nil
}

func (s *GameStatsService) ProcessKillEvent(
	ctx context.Context,
	event *models.GameEvent,
) error {
	return s.ProcessGameEvent(ctx, event)
}

