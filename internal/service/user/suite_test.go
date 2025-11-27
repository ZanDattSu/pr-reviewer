package user

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

	prService *serviceMocks.PRService
	userRepo  *mocks.UserRepository
	prRepo    *mocks.PullRequestRepository
	trm       *serviceMocks.TransactionManager

	service *userService
}

func (s *SuiteService) SetupTest() {
	s.ctx = context.Background()

	s.prService = serviceMocks.NewPRService(s.T())
	s.userRepo = mocks.NewUserRepository(s.T())
	s.prRepo = mocks.NewPullRequestRepository(s.T())
	s.trm = serviceMocks.NewTransactionManager(s.T())

	s.service = NewUserService(
		s.prService,
		s.userRepo,
		s.prRepo,
		s.trm,
	)
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(SuiteService))
}
