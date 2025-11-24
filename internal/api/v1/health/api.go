package health

import (
	"context"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
)

type HealthApi interface {
	HealthGet(ctx context.Context) (reviewerV1.HealthGetRes, error)
}
