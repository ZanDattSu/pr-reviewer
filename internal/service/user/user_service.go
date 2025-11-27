package user

import (
	"github.com/avito-tech/go-transaction-manager/trm"

	"github.com/ZanDattSu/pr-reviewer/internal/repository/pullrequest"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/user"
	prService "github.com/ZanDattSu/pr-reviewer/internal/service/pullrequest"
)

// Компиляторная проверка: убеждаемся, что *userService реализует интерфейс UserService.
var _ UserService = (*userService)(nil)

type userService struct {
	prService prService.PRService
	userRepo  user.UserRepository
	prRepo    pullrequest.PullRequestRepository
	tm        trm.Manager
}

func NewUserService(
	prService prService.PRService,
	userRepo user.UserRepository,
	prRepo pullrequest.PullRequestRepository,
	tm trm.Manager,
) *userService {
	return &userService{
		prService: prService,
		userRepo:  userRepo,
		prRepo:    prRepo,
		tm:        tm,
	}
}
