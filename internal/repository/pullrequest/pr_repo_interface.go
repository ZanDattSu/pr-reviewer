package pullrequest

import (
	"context"
	"time"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

type PullRequestRepository interface {
	GetPR(ctx context.Context, prID string) (model.PullRequest, error)
	GetPRWithReviewers(ctx context.Context, prID string) (model.PullRequest, []string, error)
	CheckPRExists(ctx context.Context, prID string) (bool, error)
	InsertPR(ctx context.Context, pr model.PullRequest) (time.Time, error)
	UpdatePRStatus(ctx context.Context, prID string, status model.Status) (model.PullRequest, error)
}
