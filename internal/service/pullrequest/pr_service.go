package pullrequest

import (
	"github.com/ZanDattSu/pr-reviewer/internal/repository/pullrequest"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/reviewer"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/team"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/tx"
)

// Компиляторная проверка: убеждаемся, что *prService реализует интерфейс PRService.
var _ PRService = (*prService)(nil)

type prService struct {
	prRepo       pullrequest.PullRequestRepository
	reviewerRepo reviewer.ReviewerRepository
	teamRepo     team.TeamRepository
	tx           tx.Provider
}

func NewPrService(
	prRepo pullrequest.PullRequestRepository,
	reviewerRepo reviewer.ReviewerRepository,
	teamRepo team.TeamRepository,
	tx tx.Provider,
) *prService {
	return &prService{prRepo: prRepo, reviewerRepo: reviewerRepo, teamRepo: teamRepo, tx: tx}
}
