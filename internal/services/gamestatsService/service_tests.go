package gamestatsService

import (
	"context"
	"errors"
	"testing"

	"github.com/JustRussianGuy/GameStats/config"
	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/JustRussianGuy/GameStats/internal/services/gamestatsService/mocks"
	"github.com/stretchr/testify/suite"
)

type GameStatsServiceSuite struct {
	suite.Suite
	ctx     context.Context
	storage *mocks.PlayerStatsStorage
	service *GameStatsService
}

func (s *GameStatsServiceSuite) SetupTest() {
	s.ctx = context.Background()
	s.storage = mocks.NewPlayerStatsStorage(s.T())

	cfg := &config.Config{
		GameSettings: config.GameSettings{
			KillPoints:   1,
			DeathPenalty: -1,
		},
	}

	s.service = NewGameStatsService(s.ctx, s.storage, cfg)
}

func (s *GameStatsServiceSuite) TestProcessGameEventSuccess() {
	event := &models.GameEvent{
		KillerID: "player-1",
		VictimID: "player-2",
	}

	s.storage.EXPECT().
		IncrementKill(s.ctx, "player-1").
		Return(nil)

	s.storage.EXPECT().
		IncrementDeath(s.ctx, "player-2").
		Return(nil)

	err := s.service.ProcessGameEvent(s.ctx, event)
	s.NoError(err)
}

func (s *GameStatsServiceSuite) TestProcessGameEventKillError() {
	event := &models.GameEvent{
		KillerID: "player-1",
		VictimID: "player-2",
	}

	wantErr := errors.New("db error")

	s.storage.EXPECT().
		IncrementKill(s.ctx, "player-1").
		Return(wantErr)

	err := s.service.ProcessGameEvent(s.ctx, event)
	s.ErrorIs(err, wantErr)
}

func (s *GameStatsServiceSuite) TestGetPlayerStats() {
	stats := &models.PlayerStats{
		PlayerID: "player-1",
		Kills:    10,
		Deaths:   3,
		Score:    7,
	}

	s.storage.EXPECT().
		GetPlayerStats(s.ctx, "player-1").
		Return(stats, nil)

	res, err := s.service.GetPlayerStats(s.ctx, "player-1")

	s.NoError(err)
	s.Equal(stats, res)
}

func (s *GameStatsServiceSuite) TestGetLeaderboardDefaultLimit() {
	leaderboard := []*models.PlayerStats{
		{PlayerID: "p1", Score: 10},
		{PlayerID: "p2", Score: 7},
	}

	s.storage.EXPECT().
		GetLeaderboard(s.ctx, 10).
		Return(leaderboard, nil)

	res, err := s.service.GetLeaderboard(s.ctx, 0)

	s.NoError(err)
	s.Len(res, 2)
}

func TestGameStatsServiceSuite(t *testing.T) {
	suite.Run(t, new(GameStatsServiceSuite))
}
