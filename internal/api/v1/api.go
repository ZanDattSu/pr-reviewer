package v1

import (
	"context"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/shared/pkg/openapi/reviewer/v1"
)

type api struct{}

func NewApi() *api {
	return &api{}
}

func (a *api) HealthGet(_ context.Context) (reviewerV1.HealthGetRes, error) {
	return &reviewerV1.HealthResponse{
		Status:  "ok",
		Service: "reviewer-service",
	}, nil
}

func (a *api) PullRequestCreatePost(ctx context.Context, req *reviewerV1.PullRequestCreatePostReq) (reviewerV1.PullRequestCreatePostRes, error) {
	// TODO implement me
	panic("implement me")
}

func (a *api) PullRequestMergePost(ctx context.Context, req *reviewerV1.PullRequestMergePostReq) (reviewerV1.PullRequestMergePostRes, error) {
	// TODO implement me
	panic("implement me")
}

func (a *api) PullRequestReassignPost(ctx context.Context, req *reviewerV1.PullRequestReassignPostReq) (reviewerV1.PullRequestReassignPostRes, error) {
	// TODO implement me
	panic("implement me")
}

func (a *api) TeamAddPost(ctx context.Context, req *reviewerV1.Team) (reviewerV1.TeamAddPostRes, error) {
	// TODO implement me
	panic("implement me")
}

func (a *api) TeamGetGet(ctx context.Context, params reviewerV1.TeamGetGetParams) (reviewerV1.TeamGetGetRes, error) {
	// TODO implement me
	panic("implement me")
}

func (a *api) UsersGetReviewGet(ctx context.Context, params reviewerV1.UsersGetReviewGetParams) (*reviewerV1.UsersGetReviewGetOK, error) {
	// TODO implement me
	panic("implement me")
}

func (a *api) UsersSetIsActivePost(ctx context.Context, req *reviewerV1.UsersSetIsActivePostReq) (reviewerV1.UsersSetIsActivePostRes, error) {
	// TODO implement me
	panic("implement me")
}
