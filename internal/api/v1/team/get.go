package team

import (
	"context"
	"errors"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/api/apierrors"
	"github.com/ZanDattSu/pr-reviewer/internal/converter"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (h *teamHandler) TeamGetGet(ctx context.Context, params reviewerV1.TeamGetGetParams) (reviewerV1.TeamGetGetRes, error) {
	teamModel, err := h.teamService.GetTeam(ctx, params.TeamName)

	team := converter.ServiceTeamToAPI(teamModel)

	if err != nil {

		var teamNF *apperror.TeamNotFoundError

		if errors.As(err, &teamNF) {
			return &reviewerV1.ErrorResponse{
				Error: apierrors.NewTeamNotFoundErr(teamNF.TeamName()),
			}, nil
		}

		return nil, err

	}

	return &reviewerV1.Team{
		TeamName: team.TeamName,
		Members:  team.Members,
	}, nil
}
