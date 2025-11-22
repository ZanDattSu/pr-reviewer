package health

import (
	"context"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/shared/pkg/openapi/reviewer/v1"
)

type healthHandler struct{}

func NewHealthHandler() *healthHandler {
	return &healthHandler{}
}

func (a *healthHandler) HealthGet(_ context.Context) (reviewerV1.HealthGetRes, error) {
	return &reviewerV1.HealthResponse{
		Status:  "ok",
		Service: "reviewer-service",
	}, nil
}
