package team

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

type TeamService interface {
	AddTeam(ctx context.Context, team model.Team) (model.Team, error)
	GetTeam(ctx context.Context, teamName string) (model.Team, error)
}
