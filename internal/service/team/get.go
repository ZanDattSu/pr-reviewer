package team

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (t *teamService) GetTeam(ctx context.Context, teamName string) (model.Team, error) {
	return t.teamRepo.GetTeam(ctx, teamName)
}
