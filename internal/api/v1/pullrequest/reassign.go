package pullrequest

import (
	"context"
	"errors"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/api/apierrors"
	"github.com/ZanDattSu/pr-reviewer/internal/converter"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (h *prHandler) PullRequestReassignPost(ctx context.Context, req *reviewerV1.PullRequestReassignPostReq) (reviewerV1.PullRequestReassignPostRes, error) {
	pr, newReviewerID, err := h.prService.ReassignPullRequest(ctx, req.PullRequestID, req.OldReviewerID)
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

			return &reviewerV1.PullRequestReassignPostNotFound{
				Error: apierrors.NewPRNotFoundErr(prNF.PullRequestID()),
			}, nil

		case errors.As(err, &userNF):

			return &reviewerV1.PullRequestReassignPostNotFound{
				Error: apierrors.NewAuthorNotFoundErr(userNF.UserID()),
			}, nil

		case errors.As(err, &mergedErr):

			return &reviewerV1.PullRequestReassignPostConflict{
				Error: apierrors.NewPRMergedErr(mergedErr.PullRequestID()),
			}, nil

		case errors.As(err, &notAssignedErr):

			return &reviewerV1.PullRequestReassignPostNotFound{
				Error: apierrors.NewReviewerNotAssignedErr(notAssignedErr.ReviewerID()),
			}, nil

		case errors.As(err, &noCandidateErr):

			return &reviewerV1.PullRequestReassignPostNotFound{
				Error: apierrors.NewNoCandidateErr(),
			}, nil

		default:
			return nil, err

		}

	}

	return &reviewerV1.PullRequestReassignPostOK{
		Pr:         converter.ServicePRToAPI(pr),
		ReplacedBy: newReviewerID,
	}, nil
}
