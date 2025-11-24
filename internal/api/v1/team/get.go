package team

import (
	"context"
	"errors"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/converter"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (t *teamHandler) TeamGetGet(ctx context.Context, params reviewerV1.TeamGetGetParams) (reviewerV1.TeamGetGetRes, error) {
	teamModel, err := t.teamService.GetTeam(ctx, params.TeamName)
	team := converter.ServiceTeamToAPI(teamModel)

	if err != nil {
		var teamNotFoundError *apperror.TeamNotFoundError
		if errors.As(err, &teamNotFoundError) {
			return &reviewerV1.ErrorResponse{Error: reviewerV1.ErrorResponseError{
				Code:    reviewerV1.ErrorResponseErrorCodeNOTFOUND,
				Message: "team not found",
			}}, nil
		}
		return nil, err
	}

	return &reviewerV1.Team{
		TeamName: team.TeamName,
		Members:  team.Members,
	}, nil
}
