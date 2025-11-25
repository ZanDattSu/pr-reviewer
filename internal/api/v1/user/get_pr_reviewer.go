package user

import (
	"context"
	"errors"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/api/apierrors"
	"github.com/ZanDattSu/pr-reviewer/internal/converter"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (h *userHandler) UsersGetReviewGet(ctx context.Context, params reviewerV1.UsersGetReviewGetParams) (reviewerV1.UsersGetReviewGetRes, error) {
	userAssignedPRs, err := h.userService.UserGetPRReviewer(ctx, params.UserID)
	if err != nil {

		var userNF *apperror.UserNotFoundError

		if errors.As(err, &userNF) {
			return &reviewerV1.ErrorResponse{
				Error: apierrors.NewUserNotFoundErr(userNF.UserID()),
			}, nil
		}

		return nil, err

	}

	return &reviewerV1.UsersGetReviewGetOK{
		UserID:       params.UserID,
		PullRequests: converter.ServiceUserAssignedPRsToAPI(userAssignedPRs),
	}, nil
}
