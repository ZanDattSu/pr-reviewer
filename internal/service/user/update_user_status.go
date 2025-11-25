package user

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (s *userService) UpdateUserStatus(ctx context.Context, userID string, isActive bool) (model.User, error) {
	return s.userRepo.UpdateUserStatus(ctx, userID, isActive)
}
