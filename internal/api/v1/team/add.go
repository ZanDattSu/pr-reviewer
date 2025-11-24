package team

import (
	"context"
	"errors"
	"fmt"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/converter"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (t *teamHandler) TeamAddPost(ctx context.Context, req *reviewerV1.Team) (reviewerV1.TeamAddPostRes, error) {
	team := reviewerV1.Team{
		TeamName: req.TeamName,
		Members:  req.Members,
	}

	createdTeam, err := t.teamService.AddTeam(ctx, converter.APIToServiceTeam(team))
	if err != nil {
		var teamExistsErr *apperror.TeamExistsError
		if errors.As(err, &teamExistsErr) {
			return &reviewerV1.ErrorResponse{
				Error: reviewerV1.ErrorResponseError{
					Code:    reviewerV1.ErrorResponseErrorCodeTEAMEXISTS,
					Message: fmt.Sprintf("team %s already exists", teamExistsErr.TeamName()),
				},
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
