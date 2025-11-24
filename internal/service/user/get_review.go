package user

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (u *userService) UserGetReview(ctx context.Context, userID string) ([]model.UserAssignedPR, error) {
	return u.userRepo.UserGetReview(ctx, userID)
}
