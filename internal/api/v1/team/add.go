package team

import (
	"context"
	"errors"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/api/apierrors"
	"github.com/ZanDattSu/pr-reviewer/internal/converter"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (h *teamHandler) TeamAddPost(ctx context.Context, req *reviewerV1.Team) (reviewerV1.TeamAddPostRes, error) {
	team := reviewerV1.Team{
		TeamName: req.TeamName,
		Members:  req.Members,
	}

	createdTeam, err := h.teamService.AddTeam(ctx, converter.APIToServiceTeam(team))
	if err != nil {

		var teamExistsErr *apperror.TeamExistsError

		if errors.As(err, &teamExistsErr) {
			return &reviewerV1.ErrorResponse{
				Error: apierrors.NewTeamExistErr(teamExistsErr.TeamName()),
			}, nil
		}

		return nil, err

	}

	return &reviewerV1.TeamAddPostCreated{
		Team: reviewerV1.NewOptTeam(
			converter.ServiceTeamToAPI(createdTeam),
		),
	}, nil
}
