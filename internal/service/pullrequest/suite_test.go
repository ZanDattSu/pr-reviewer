package pullrequest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ZanDattSu/pr-reviewer/internal/repository/mocks"
)

type SuiteService struct {
	suite.Suite

	ctx context.Context //nolint:containedctx

	prRepo       *mocks.PullRequestRepository
	reviewerRepo *mocks.ReviewerRepository
	teamRepo     *mocks.TeamRepository
	userRepo     *mocks.UserRepository

	service *prService
}

func (s *SuiteService) SetupTest() {
	s.ctx = context.Background()

	s.prRepo = mocks.NewPullRequestRepository(s.T())
	s.reviewerRepo = mocks.NewReviewerRepository(s.T())
	s.teamRepo = mocks.NewTeamRepository(s.T())
	s.userRepo = mocks.NewUserRepository(s.T())

	s.service = NewPrService(s.prRepo, s.reviewerRepo, s.teamRepo, s.userRepo)
}

func (s *SuiteService) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(SuiteService))
}
