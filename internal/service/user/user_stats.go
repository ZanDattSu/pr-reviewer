package user

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (s *userService) GetUserStats(ctx context.Context, top int, onlyActive, onlyOpen bool) ([]model.UserStats, error) {
	userStats, err := s.userRepo.GetUserStats(ctx, top, onlyActive, onlyOpen)
	if err != nil {
		return nil, err
	}

	return userStats, nil
}
