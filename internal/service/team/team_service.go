package team

import (
	"github.com/ZanDattSu/pr-reviewer/internal/repository/team"
)

// Компиляторная проверка: убеждаемся, что *teamService реализует интерфейс TeamService.
var _ TeamService = (*teamService)(nil)

type teamService struct {
	teamRepo team.TeamRepository
}

func NewTeamService(teamRepo team.TeamRepository) *teamService {
	return &teamService{teamRepo: teamRepo}
}
