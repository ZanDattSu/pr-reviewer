package team

import (
	"github.com/avito-tech/go-transaction-manager/trm"

	"github.com/ZanDattSu/pr-reviewer/internal/repository/team"
)

// Компиляторная проверка: убеждаемся, что *teamService реализует интерфейс TeamService.
var _ TeamService = (*teamService)(nil)

type teamService struct {
	teamRepo team.TeamRepository
	tm       trm.Manager
}

func NewTeamService(teamRepo team.TeamRepository, tm trm.Manager) *teamService {
	return &teamService{
		teamRepo: teamRepo,
		tm:       tm,
	}
}
