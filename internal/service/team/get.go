package team

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (t *teamService) GetTeam(ctx context.Context, teamName string) (model.Team, error) {
	team, err := t.teamRepo.GetTeam(ctx, teamName)
	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}
