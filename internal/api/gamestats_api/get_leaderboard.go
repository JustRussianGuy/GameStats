package gamestats_api

import (
	"context"

	"github.com/JustRussianGuy/GameStats/internal/pb/gamestats_api"
)

func (g *GameStatsAPI) GetLeaderboard(
	ctx context.Context,
	req *gamestats_api.GetLeaderboardRequest,
) (*gamestats_api.GetLeaderboardResponse, error) {

	players, err := g.service.GetLeaderboard(ctx, int(req.Limit))
	if err != nil {
		return &gamestats_api.GetLeaderboardResponse{}, err
	}

	resp := make([]*gamestats_api.LeaderboardEntry, 0, len(players))
	for _, p := range players {
		resp = append(resp, &gamestats_api.LeaderboardEntry{
			PlayerId: p.PlayerID,
			Score:    p.Score,
		})
	}

	return &gamestats_api.GetLeaderboardResponse{
		Players: resp,
	}, nil
}
