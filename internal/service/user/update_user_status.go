package user

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (u *userService) UpdateUserStatus(ctx context.Context, userID string, isActive bool) (model.User, error) {
	return u.userRepo.UpdateUserStatus(ctx, userID, isActive)
}
