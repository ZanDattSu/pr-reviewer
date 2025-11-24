package team

import (
	"github.com/ZanDattSu/pr-reviewer/internal/service/team"
)

// Компиляторная проверка: убеждаемся, что *teamHandler реализует интерфейс TeamApi.
var _ TeamApi = (*teamHandler)(nil)

type teamHandler struct {
	teamService team.TeamService
}

func NewTeamHandler(teamService team.TeamService) *teamHandler {
	return &teamHandler{teamService: teamService}
}
