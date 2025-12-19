package gamestats_api

import (
	"context"
	"time"

	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/JustRussianGuy/GameStats/internal/pb/gamestats_api"
)

func (g *GameStatsAPI) PostEvent(
	ctx context.Context,
	req *gamestats_api.PlayerEvent,
) (*gamestats_api.AddEventResponse, error) {

	event := &models.GameEvent{
		KillerID:   req.PlayerId,
		VictimID:   req.VictimId,
		OccurredAt: time.Now(),
	}

	err := g.service.ProcessKillEvent(ctx, event)
	if err != nil {
		return &gamestats_api.AddEventResponse{Success: false}, err
	}

	return &gamestats_api.AddEventResponse{Success: true}, nil
}

