package pullrequest

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (s *prService) MergePullRequest(ctx context.Context, pullRequestID string) (model.PullRequest, error) {
	pr, err := s.prRepo.GetPRWithReviewers(ctx, pullRequestID)
	if err != nil {
		return model.PullRequest{}, err
	}

	updatePR, err := s.prRepo.UpdatePRStatus(ctx, pr, model.StatusMerged)
	if err != nil {
		return model.PullRequest{}, err
	}

	return updatePR, nil
}
