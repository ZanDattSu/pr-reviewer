package user

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

type UserService interface {
	UserSetIsActive(ctx context.Context, userID string, isActive bool) (model.User, error)
	UserGetReview(ctx context.Context, userID string) ([]model.UserAssignedPR, error)
}
