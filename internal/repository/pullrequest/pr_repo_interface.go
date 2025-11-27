package pullrequest

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

type PullRequestRepository interface {
	GetPRWithReviewers(ctx context.Context, prID string) (model.PullRequest, error)
	CheckPRExists(ctx context.Context, prID string) (bool, error)
	InsertPR(ctx context.Context, pr model.PullRequest) (model.PullRequest, error)
	UpdatePRStatus(ctx context.Context, pr model.PullRequest, status model.Status) (model.PullRequest, error)
	FindOpenPRsWithReviewers(ctx context.Context, reviewerIDs []string) ([]model.OpenPR, error)
}
