package pullrequest

import (
	"github.com/avito-tech/go-transaction-manager/trm"

	"github.com/ZanDattSu/pr-reviewer/internal/repository/pullrequest"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/reviewer"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/user"
)

// Компиляторная проверка: убеждаемся, что *prService реализует интерфейс PRService.
var _ PRService = (*prService)(nil)

type prService struct {
	prRepo       pullrequest.PullRequestRepository
	reviewerRepo reviewer.ReviewerRepository
	userRepo     user.UserRepository
	tm           trm.Manager
}

func NewPrService(
	prRepo pullrequest.PullRequestRepository,
	reviewerRepo reviewer.ReviewerRepository,
	userRepo user.UserRepository,
	tm trm.Manager,
) *prService {
	return &prService{
		prRepo:       prRepo,
		reviewerRepo: reviewerRepo,
		userRepo:     userRepo,
		tm:           tm,
	}
}
