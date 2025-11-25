package team

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (s *teamService) AddTeam(ctx context.Context, team model.Team) (model.Team, error) {
	return s.teamRepo.AddTeam(ctx, team)
}
