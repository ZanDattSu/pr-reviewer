package team

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ZanDattSu/pr-reviewer/internal/repository/mocks"
)

type SuiteService struct {
	suite.Suite

	ctx context.Context //nolint:containedctx

	teamRepo *mocks.TeamRepository

	service *teamService
}

func (s *SuiteService) SetupTest() {
	s.ctx = context.Background()

	s.teamRepo = mocks.NewTeamRepository(s.T())

	s.service = NewTeamService(s.teamRepo)
}

func (s *SuiteService) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(SuiteService))
}
