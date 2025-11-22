package v1

import (
	"context"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/shared/pkg/openapi/reviewer/v1"
)

type Api interface {
	HealthGet(ctx context.Context) (reviewerV1.HealthGetRes, error)

	PullRequestCreatePost(ctx context.Context, req *reviewerV1.PullRequestCreatePostReq) (reviewerV1.PullRequestCreatePostRes, error)
	PullRequestMergePost(ctx context.Context, req *reviewerV1.PullRequestMergePostReq) (reviewerV1.PullRequestMergePostRes, error)
	PullRequestReassignPost(ctx context.Context, req *reviewerV1.PullRequestReassignPostReq) (reviewerV1.PullRequestReassignPostRes, error)

	TeamAddPost(ctx context.Context, req *reviewerV1.Team) (reviewerV1.TeamAddPostRes, error)
	TeamGetGet(ctx context.Context, params reviewerV1.TeamGetGetParams) (reviewerV1.TeamGetGetRes, error)

	UsersGetReviewGet(ctx context.Context, params reviewerV1.UsersGetReviewGetParams) (*reviewerV1.UsersGetReviewGetOK, error)
	UsersSetIsActivePost(ctx context.Context, req *reviewerV1.UsersSetIsActivePostReq) (reviewerV1.UsersSetIsActivePostRes, error)
}
