package gamestats_api

import (
	"context"

	"github.com/JustRussianGuy/GameStats/internal/pb/gamestats_api"
)

func (g *GameStatsAPI) GetPlayerStats(
	ctx context.Context,
	req *gamestats_api.GetPlayerStatsRequest,
) (*gamestats_api.GetPlayerStatsResponse, error) {

	stats, err := g.service.GetPlayerStats(ctx, req.PlayerId)
	if err != nil {
		return &gamestats_api.GetPlayerStatsResponse{}, err
	}

	return &gamestats_api.GetPlayerStatsResponse{
		PlayerId: stats.PlayerID,
		Kills:    stats.Kills,
		Deaths:   stats.Deaths,
		Score:    stats.Score,
	}, nil
}
