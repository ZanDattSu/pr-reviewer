package user

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (u *userService) UserGetPRReviewer(ctx context.Context, userID string) ([]model.UserAssignedPR, error) {
	exists, err := u.userRepo.CheckUserExists(ctx, userID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, apperror.NewUserNotFoundError(userID)
	}

	return u.userRepo.UserGetPRReviewer(ctx, userID)
}
