package pullrequest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ZanDattSu/pr-reviewer/internal/repository/mocks"
	serviceMocks "github.com/ZanDattSu/pr-reviewer/internal/service/mocks"
)

type SuiteService struct {
	suite.Suite

	ctx context.Context //nolint:containedctx

	prRepo       *mocks.PullRequestRepository
	reviewerRepo *mocks.ReviewerRepository
	userRepo     *mocks.UserRepository
	trm          *serviceMocks.TransactionManager

	service *prService
}

func (s *SuiteService) SetupTest() {
	s.ctx = context.Background()

	s.prRepo = mocks.NewPullRequestRepository(s.T())
	s.reviewerRepo = mocks.NewReviewerRepository(s.T())
	s.userRepo = mocks.NewUserRepository(s.T())
	s.trm = serviceMocks.NewTransactionManager(s.T())

	s.service = NewPrService(s.prRepo, s.reviewerRepo, s.userRepo, s.trm)
}

func (s *SuiteService) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(SuiteService))
}
