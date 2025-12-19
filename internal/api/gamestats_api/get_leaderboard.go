package gamestats_api

import (
	"context"
	"strconv"

	"github.com/JustRussianGuy/GameStats/internal/pb/gamestats_api"
)

func (g *GameStatsAPI) GetLeaderboard(
	ctx context.Context,
	req *gamestats_api.LeaderboardRequest,
) (*gamestats_api.LeaderboardResponse, error) {

	players, err := g.service.GetLeaderboard(ctx, int(req.Limit))
	if err != nil {
		return nil, err
	}

	resp := make([]*gamestats_api.PlayerStats, 0, len(players))
	for _, p := range players {
		resp = append(resp, &gamestats_api.PlayerStats{
			PlayerId: strconv.FormatUint(p.PlayerID, 10), // ← ВАЖНО
			Kills:    int64(p.Kills),
			Deaths:   int64(p.Deaths),
			Score:    int64(p.Score),
		})
	}

	return &gamestats_api.LeaderboardResponse{
		Players: resp,
	}, nil
}
