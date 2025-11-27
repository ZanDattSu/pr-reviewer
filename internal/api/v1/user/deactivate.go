package user

import (
	"context"
	"errors"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/api/apierrors"
	"github.com/ZanDattSu/pr-reviewer/internal/converter"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (h *userHandler) UsersDeactivatePost(ctx context.Context, req *reviewerV1.UsersDeactivatePostReq) (reviewerV1.UsersDeactivatePostRes, error) {
	reassignedPRs, err := h.userService.DeactivateUsersAndReassignPR(ctx, req.UserIds)
	if err != nil {
		var (
			prNF           *apperror.PRNotFoundError
			userNF         *apperror.UserNotFoundError
			mergedErr      *apperror.PRMergedError
			notAssignedErr *apperror.NotAssignedError
			noCandidateErr *apperror.NoCandidateError
		)

		switch {

		case errors.As(err, &prNF):

			return &reviewerV1.UsersDeactivatePostNotFound{
				Error: apierrors.NewPRNotFoundErr(prNF.PullRequestID()),
			}, nil

		case errors.As(err, &userNF):

			return &reviewerV1.UsersDeactivatePostNotFound{
				Error: apierrors.NewUserNotFoundErr(userNF.UserID()),
			}, nil

		case errors.As(err, &mergedErr):

			return &reviewerV1.UsersDeactivatePostConflict{
				Error: apierrors.NewPRMergedErr(mergedErr.PullRequestID()),
			}, nil

		case errors.As(err, &notAssignedErr):

			return &reviewerV1.UsersDeactivatePostConflict{
				Error: apierrors.NewReviewerNotAssignedErr(notAssignedErr.ReviewerID()),
			}, nil

		case errors.As(err, &noCandidateErr):

			return &reviewerV1.UsersDeactivatePostConflict{
				Error: apierrors.NewNoCandidateErr(),
			}, nil

		default:
			return nil, err

		}
	}

	items := converter.ServiceReassignedPRsToAPI(reassignedPRs)

	return &items, nil
}
