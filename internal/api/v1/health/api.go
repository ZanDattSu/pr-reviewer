package health

import (
	"context"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/shared/pkg/openapi/reviewer/v1"
)

type Api interface {
	HealthGet(ctx context.Context) (reviewerV1.HealthGetRes, error)
}
