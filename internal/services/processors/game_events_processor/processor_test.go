package game_events_processor

import (
	"context"
	"testing"

	"github.com/JustRussianGuy/GameStats/internal/models"
	"github.com/JustRussianGuy/GameStats/internal/services/gamestatsService/mocks"
	"github.com/stretchr/testify/suite"
)

type GameEventsProcessorSuite struct {
	suite.Suite
	ctx       context.Context
	service   *mocks.Service
	processor *GameEventsProcessor
}

func (s *GameEventsProcessorSuite) SetupTest() {
	s.ctx = context.Background()
	s.service = mocks.NewService(s.T())
	s.processor = NewGameEventsProcessor(s.service)
}

func (s *GameEventsProcessorSuite) TestHandleSuccess() {
	event := &models.GameEvent{
		KillerID:  "p1",
		VictimID: "p2",
	}

	s.service.EXPECT().ProcessGameEvent(s.ctx, event).Return(nil)

	err := s.processor.Handle(s.ctx, event)
	s.NoError(err)
}

func TestGameEventsProcessorSuite(t *testing.T) {
	suite.Run(t, new(GameEventsProcessorSuite))
}
