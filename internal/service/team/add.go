package team

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (s *teamService) AddTeam(ctx context.Context, team model.Team) (model.Team, error) {
	var result model.Team
	err := s.tm.Do(ctx, func(ctx context.Context) error {
		t, err := s.teamRepo.AddTeam(ctx, team)
		if err != nil {
			return err
		}
		result = t
		return nil
	})
	if err != nil {
		return model.Team{}, err
	}

	return result, nil
}
