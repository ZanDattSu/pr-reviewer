package pullrequest

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

type PRService interface {
	CreatePullRequest(ctx context.Context, pullRequestID, pullRequestName, authorID string) (model.PullRequest, error)
	MergePullRequest(ctx context.Context, pullRequestID string) (model.PullRequest, error)
	ReassignPullRequest(ctx context.Context, pullRequestID, oldReviewerID string) (model.PullRequest, string, error)
}
