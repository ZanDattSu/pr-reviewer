package pullrequest

import (
	"context"
	"errors"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/api/apierrors"
	"github.com/ZanDattSu/pr-reviewer/internal/converter"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (h *prHandler) PullRequestCreatePost(ctx context.Context, req *reviewerV1.PullRequestCreatePostReq) (reviewerV1.PullRequestCreatePostRes, error) {
	pr, err := h.prService.CreatePullRequest(
		ctx,
		req.PullRequestID,
		req.PullRequestName,
		req.AuthorID,
	)
	if err != nil {

		var (
			prExistsErr *apperror.PRExistsError
			userNF      *apperror.UserNotFoundError
			teamNF      *apperror.TeamNotFoundError
		)

		switch {

		case errors.As(err, &prExistsErr):

			return &reviewerV1.PullRequestCreatePostConflict{
				Error: apierrors.NewPrExistErr(prExistsErr.PullRequestID()),
			}, nil

		case errors.As(err, &userNF):

			return &reviewerV1.PullRequestCreatePostNotFound{
				Error: apierrors.NewAuthorNotFoundErr(userNF.UserID()),
			}, nil

		case errors.As(err, &teamNF):

			return &reviewerV1.PullRequestCreatePostNotFound{
				Error: apierrors.NewTeamNotFoundErr(teamNF.TeamName()),
			}, nil

		default:
			return nil, err
		}

	}

	return &reviewerV1.PullRequestCreatePostCreated{
		Pr: reviewerV1.NewOptPullRequest(
			converter.ServicePRToAPI(pr),
		),
	}, nil
}
