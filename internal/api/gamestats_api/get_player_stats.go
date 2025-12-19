package gamestats_api

import (
	"context"
	"strconv"

	"github.com/JustRussianGuy/GameStats/internal/pb/gamestats_api"
)

func (g *GameStatsAPI) GetPlayerStats(
	ctx context.Context,
	req *gamestats_api.PlayerRequest,
) (*gamestats_api.PlayerStats, error) {

	// Преобразуем string -> uint64
	playerID, err := strconv.ParseUint(req.PlayerId, 10, 64)
	if err != nil {
		return nil, err
	}

	stats, err := g.service.GetPlayerStats(ctx, playerID)
	if err != nil {
		return nil, err
	}

	return &gamestats_api.PlayerStats{
		PlayerId: strconv.FormatUint(stats.PlayerID, 10), // uint64 -> string
		Kills:    int64(stats.Kills),                     // uint64 -> int64
		Deaths:   int64(stats.Deaths),                    // uint64 -> int64
		Score:    int64(stats.Score),
	}, nil
}

