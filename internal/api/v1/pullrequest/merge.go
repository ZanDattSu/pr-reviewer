package pullrequest

import (
	"context"
	"errors"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/api/apierrors"
	"github.com/ZanDattSu/pr-reviewer/internal/converter"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (h *prHandler) PullRequestMergePost(ctx context.Context, req *reviewerV1.PullRequestMergePostReq) (reviewerV1.PullRequestMergePostRes, error) {
	pr, err := h.prService.MergePullRequest(ctx, req.PullRequestID)
	if err != nil {

		var prNF *apperror.PRNotFoundError

		if errors.As(err, &prNF) {
			return &reviewerV1.ErrorResponse{
				Error: apierrors.NewPRNotFoundErr(prNF.PullRequestID()),
			}, nil
		}

		return nil, err

	}

	return &reviewerV1.PullRequestMergePostOK{
		Pr: reviewerV1.NewOptPullRequest(
			converter.ServicePRToAPI(pr),
		),
	}, nil
}
