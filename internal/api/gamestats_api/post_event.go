package gamestats_api

import (
	"context"
	"time"
	"strconv"

	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/JustRussianGuy/GameStats/internal/pb/gamestats_api"
)

func (g *GameStatsAPI) AddEvent(
	ctx context.Context,
	req *gamestats_api.AddEventRequest,
) (*gamestats_api.AddEventResponse, error) {

	// Преобразуем string -> uint64
	killerID, err := strconv.ParseUint(req.KillerId, 10, 64)
	if err != nil {
		return &gamestats_api.AddEventResponse{Success: false}, err
	}

	victimID, err := strconv.ParseUint(req.VictimId, 10, 64)
	if err != nil {
		return &gamestats_api.AddEventResponse{Success: false}, err
	}

	event := &models.GameEvent{
		KillerID:   killerID,
		VictimID:   victimID,
		OccurredAt: time.Now(),
	}

	err = g.service.ProcessKillEvent(ctx, event)
	if err != nil {
		return &gamestats_api.AddEventResponse{Success: false}, err
	}

	if err := g.kafkaProducer.ProduceEvent(ctx, event); err != nil {
		return &gamestats_api.AddEventResponse{Success: false}, err
	}

	return &gamestats_api.AddEventResponse{Success: true}, nil
}


