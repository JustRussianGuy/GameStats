package gamestats_api

import (
	"context"

	"github.com/JustRussianGuy/GameStats/internal/pb/gamestats_api"
)

func (g *GameStatsAPI) GetPlayerStats(
	ctx context.Context,
	req *gamestats_api.PlayerRequest,
) (*gamestats_api.PlayerStats, error) {

	stats, err := g.service.GetPlayerStats(ctx, req.PlayerId)
	if err != nil {
		return nil, err
	}

	return &gamestats_api.PlayerStats{
		PlayerId: stats.PlayerID,
		Kills:    stats.Kills,
		Deaths:   stats.Deaths,
		Score:    stats.Score,
	}, nil
}

