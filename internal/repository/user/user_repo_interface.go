package user

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

type UserRepository interface {
	UpdateUserStatus(ctx context.Context, userID string, isActive bool) (model.User, error)
	UserGetPRReviewer(ctx context.Context, userID string) ([]model.UserAssignedPR, error)
	CheckUserExists(ctx context.Context, userId string) (bool, error)
	GetUserStats(ctx context.Context, top int, onlyActive, onlyOpen bool) ([]model.UserStats, error)
	DeactivateUsersAndReassign(ctx context.Context, userIDs []string) (map[string]string, error)
}
