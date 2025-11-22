package team

import (
	"context"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/shared/pkg/openapi/reviewer/v1"
)

type Api interface {
	TeamAddPost(ctx context.Context, req *reviewerV1.Team) (reviewerV1.TeamAddPostRes, error)
	TeamGetGet(ctx context.Context, params reviewerV1.TeamGetGetParams) (reviewerV1.TeamGetGetRes, error)
}
