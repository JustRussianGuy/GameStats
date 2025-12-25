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

	playerID, err := strconv.ParseUint(req.PlayerId, 10, 64)
	if err != nil {
		return nil, err
	}

	stats, err := g.service.GetPlayerStats(ctx, playerID)
	if err != nil {
		return nil, err
	}

	return &gamestats_api.PlayerStats{
		PlayerId: stats.PlayerID,
		Kills:    int64(stats.Kills),
		Deaths:   int64(stats.Deaths),
		Score:    int64(stats.Score),
	}, nil
}
