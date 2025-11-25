package pullrequest

import (
	"context"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/pullrequest"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/pullrequest/mocks"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/reviewer"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/team"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/user"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuiteService struct {
	suite.Suite

	ctx context.Context //nolint:containedctx

	prRepo       pullrequest.PullRequestRepository
	reviewerRepo reviewer.ReviewerRepository
	teamRepo     team.TeamRepository
	userRepo     user.UserRepository
}

func (s *SuiteService) SetupTest() {
	s.ctx = context.Background()

	s.prRepo = mocks.NewPullRequestRepository(s.T())
	s.reviewerRepo = mocks.New
}

func (s *SuiteService) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(SuiteService))
}
