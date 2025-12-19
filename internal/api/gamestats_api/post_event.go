package gamestats_api

import (
	"context"
	"time"
	"strconv"

	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/JustRussianGuy/GameStats/internal/pb/gamestats_api"
)

func (g *GameStatsAPI) PostEvent(
	ctx context.Context,
	req *gamestats_api.PlayerEvent,
) (*gamestats_api.AddEventResponse, error) {

	// Преобразуем uint64 -> string для доменной модели
	killerID := strconv.FormatUint(req.PlayerId, 10)
	victimID := strconv.FormatUint(req.VictimId, 10)

	event := &models.GameEvent{
		KillerID:   killerID,
		VictimID:   victimID,
		OccurredAt: time.Now(),
	}

	err := g.service.ProcessKillEvent(ctx, event)
	if err != nil {
		return &gamestats_api.AddEventResponse{Success: false}, err
	}

	return &gamestats_api.AddEventResponse{Success: true}, nil
}

