package user

import (
	"context"
	"errors"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/api/apierrors"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (h *userHandler) UsersDeactivatePost(ctx context.Context, req *reviewerV1.UsersDeactivatePostReq) (reviewerV1.UsersDeactivatePostRes, error) {
	deactivateResult, err := h.userService.DeactivateUsers(ctx, req.UserIds)
	if err != nil {
		var (
			userNF         *apperror.UserNotFoundError
			noCandidateErr *apperror.NoCandidateError
		)

		switch {
		case errors.As(err, &userNF):
			return &reviewerV1.UsersDeactivatePostNotFound{
				Error: apierrors.NewUserNotFoundErr(""),
			}, nil

		case errors.As(err, &noCandidateErr):
			return &reviewerV1.UsersDeactivatePostConflict{
				Error: apierrors.NewNoCandidateErr(),
			}, nil

		default:
			return nil, err
		}
	}

	items := make([]reviewerV1.DeactivateResultItem, 0, len(deactivateResult))

	for _, r := range deactivateResult {
		items = append(items, reviewerV1.DeactivateResultItem{
			PullRequestID: r.PullRequestID,
			ReplacedBy:    r.ReplacedBy,
		})
	}

	return &reviewerV1.UsersDeactivatePostOK{
		Results: items,
	}, nil
}
