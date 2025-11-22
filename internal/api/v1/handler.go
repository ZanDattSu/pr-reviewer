package v1

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/api/v1/health"
	"github.com/ZanDattSu/pr-reviewer/internal/api/v1/pullrequest"
	"github.com/ZanDattSu/pr-reviewer/internal/api/v1/team"
	"github.com/ZanDattSu/pr-reviewer/internal/api/v1/users"
	reviewerV1 "github.com/ZanDattSu/pr-reviewer/shared/pkg/openapi/reviewer/v1"
)

type api struct {
	HealthApi      health.Api
	PullRequestApi pullrequest.Api
	TeamApi        team.Api
	UsersApi       users.Api
}

func NewApi(healthApi health.Api, pullRequestApi pullrequest.Api, teamApi team.Api, usersApi users.Api) *api {
	return &api{
		HealthApi:      healthApi,
		PullRequestApi: pullRequestApi,
		TeamApi:        teamApi,
		UsersApi:       usersApi,
	}
}

func (a *api) HealthGet(ctx context.Context) (reviewerV1.HealthGetRes, error) {
	return a.HealthApi.HealthGet(ctx)
}

func (a *api) PullRequestCreatePost(ctx context.Context, req *reviewerV1.PullRequestCreatePostReq) (reviewerV1.PullRequestCreatePostRes, error) {
	return a.PullRequestApi.PullRequestCreatePost(ctx, req)
}

func (a *api) PullRequestMergePost(ctx context.Context, req *reviewerV1.PullRequestMergePostReq) (reviewerV1.PullRequestMergePostRes, error) {
	return a.PullRequestApi.PullRequestMergePost(ctx, req)
}

func (a *api) PullRequestReassignPost(ctx context.Context, req *reviewerV1.PullRequestReassignPostReq) (reviewerV1.PullRequestReassignPostRes, error) {
	return a.PullRequestApi.PullRequestReassignPost(ctx, req)
}

func (a *api) TeamAddPost(ctx context.Context, req *reviewerV1.Team) (reviewerV1.TeamAddPostRes, error) {
	return a.TeamApi.TeamAddPost(ctx, req)
}

func (a *api) TeamGetGet(ctx context.Context, params reviewerV1.TeamGetGetParams) (reviewerV1.TeamGetGetRes, error) {
	return a.TeamApi.TeamGetGet(ctx, params)
}

func (a *api) UsersGetReviewGet(ctx context.Context, params reviewerV1.UsersGetReviewGetParams) (*reviewerV1.UsersGetReviewGetOK, error) {
	return a.UsersApi.UsersGetReviewGet(ctx, params)
}

func (a *api) UsersSetIsActivePost(ctx context.Context, req *reviewerV1.UsersSetIsActivePostReq) (reviewerV1.UsersSetIsActivePostRes, error) {
	return a.UsersApi.UsersSetIsActivePost(ctx, req)
}
