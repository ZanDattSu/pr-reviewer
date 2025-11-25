package user

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

type UserService interface {
	UpdateUserStatus(ctx context.Context, userID string, isActive bool) (model.User, error)
	UserGetPRReviewer(ctx context.Context, userID string) ([]model.UserAssignedPR, error)
}
