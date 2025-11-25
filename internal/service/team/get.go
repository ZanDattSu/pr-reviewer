package team

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (s *teamService) GetTeam(ctx context.Context, teamName string) (model.Team, error) {
	return s.teamRepo.GetTeam(ctx, teamName)
}
