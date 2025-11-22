package pullrequest

import (
	"context"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/shared/pkg/openapi/reviewer/v1"
)

type Api interface {
	PullRequestCreatePost(ctx context.Context, req *reviewerV1.PullRequestCreatePostReq) (reviewerV1.PullRequestCreatePostRes, error)
	PullRequestMergePost(ctx context.Context, req *reviewerV1.PullRequestMergePostReq) (reviewerV1.PullRequestMergePostRes, error)
	PullRequestReassignPost(ctx context.Context, req *reviewerV1.PullRequestReassignPostReq) (reviewerV1.PullRequestReassignPostRes, error)
}
