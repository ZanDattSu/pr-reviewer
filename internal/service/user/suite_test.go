package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ZanDattSu/pr-reviewer/internal/repository/mocks"
)

type SuiteService struct {
	suite.Suite

	ctx context.Context //nolint:containedctx

	userRepo *mocks.UserRepository
	service  *userService
}

func (s *SuiteService) SetupTest() {
	s.ctx = context.Background()

	s.userRepo = mocks.NewUserRepository(s.T())
	s.service = NewUserService(s.userRepo)
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(SuiteService))
}
