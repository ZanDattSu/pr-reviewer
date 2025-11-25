package v1

import (
	"context"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/api/v1/health"
	"github.com/ZanDattSu/pr-reviewer/internal/api/v1/pullrequest"
	"github.com/ZanDattSu/pr-reviewer/internal/api/v1/team"
	"github.com/ZanDattSu/pr-reviewer/internal/api/v1/user"
)

type mainHandler struct {
	HealthApi      health.HealthApi
	PullRequestApi pullrequest.PRApi
	TeamApi        team.TeamApi
	UsersApi       user.UserApi
}

func NewApi(healthApi health.HealthApi, pullRequestApi pullrequest.PRApi, teamApi team.TeamApi, usersApi user.UserApi) *mainHandler {
	return &mainHandler{
		HealthApi:      healthApi,
		PullRequestApi: pullRequestApi,
		TeamApi:        teamApi,
		UsersApi:       usersApi,
	}
}

func (a *mainHandler) HealthGet(ctx context.Context) (reviewerV1.HealthGetRes, error) {
	return a.HealthApi.HealthGet(ctx)
}

func (a *mainHandler) TeamAddPost(ctx context.Context, req *reviewerV1.Team) (reviewerV1.TeamAddPostRes, error) {
	return a.TeamApi.TeamAddPost(ctx, req)
}

func (a *mainHandler) TeamGetGet(ctx context.Context, params reviewerV1.TeamGetGetParams) (reviewerV1.TeamGetGetRes, error) {
	return a.TeamApi.TeamGetGet(ctx, params)
}

func (a *mainHandler) UsersDeactivatePost(ctx context.Context, req *reviewerV1.UsersDeactivatePostReq) (reviewerV1.UsersDeactivatePostRes, error) {
	return a.UsersApi.UsersDeactivatePost(ctx, req)
}

func (a *mainHandler) UsersSetIsActivePost(ctx context.Context, req *reviewerV1.UsersSetIsActivePostReq) (reviewerV1.UsersSetIsActivePostRes, error) {
	return a.UsersApi.UsersSetIsActivePost(ctx, req)
}

func (a *mainHandler) UsersGetReviewGet(ctx context.Context, params reviewerV1.UsersGetReviewGetParams) (reviewerV1.UsersGetReviewGetRes, error) {
	return a.UsersApi.UsersGetReviewGet(ctx, params)
}

func (a *mainHandler) UsersStatsGet(ctx context.Context, params reviewerV1.UsersStatsGetParams) (reviewerV1.UsersStatsGetRes, error) {
	return a.UsersApi.UsersStatsGet(ctx, params)
}

func (a *mainHandler) PullRequestCreatePost(ctx context.Context, req *reviewerV1.PullRequestCreatePostReq) (reviewerV1.PullRequestCreatePostRes, error) {
	return a.PullRequestApi.PullRequestCreatePost(ctx, req)
}

func (a *mainHandler) PullRequestMergePost(ctx context.Context, req *reviewerV1.PullRequestMergePostReq) (reviewerV1.PullRequestMergePostRes, error) {
	return a.PullRequestApi.PullRequestMergePost(ctx, req)
}

func (a *mainHandler) PullRequestReassignPost(ctx context.Context, req *reviewerV1.PullRequestReassignPostReq) (reviewerV1.PullRequestReassignPostRes, error) {
	return a.PullRequestApi.PullRequestReassignPost(ctx, req)
}
