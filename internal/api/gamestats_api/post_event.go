package gamestats_api

import (
	"context"
	"time"

	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/JustRussianGuy/GameStats/internal/pb/gamestats_api"
)

func (g *GameStatsAPI) PostEvent(
	ctx context.Context,
	req *gamestats_api.PostEventRequest,
) (*gamestats_api.PostEventResponse, error) {

	event := &models.GameEvent{
		KillerID:   req.KillerId,
		VictimID:  req.VictimId,
		OccurredAt: time.Now(),
	}

	err := g.service.ProcessKillEvent(ctx, event)
	if err != nil {
		return &gamestats_api.PostEventResponse{}, err
	}

	return &gamestats_api.PostEventResponse{
		Status: "ok",
	}, nil
}
