package team

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (t *teamService) AddTeam(ctx context.Context, team model.Team) (model.Team, error) {
	return t.teamRepo.AddTeam(ctx, team)
}
