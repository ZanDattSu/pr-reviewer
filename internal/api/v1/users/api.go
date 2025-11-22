package users

import (
	"context"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/shared/pkg/openapi/reviewer/v1"
)

type Api interface {
	UsersGetReviewGet(ctx context.Context, params reviewerV1.UsersGetReviewGetParams) (*reviewerV1.UsersGetReviewGetOK, error)
	UsersSetIsActivePost(ctx context.Context, req *reviewerV1.UsersSetIsActivePostReq) (reviewerV1.UsersSetIsActivePostRes, error)
}
