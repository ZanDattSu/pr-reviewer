package team

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

const PgSqlUniqueViolationErr = "23505"

func (t *teamService) AddTeam(ctx context.Context, team model.Team) (model.Team, error) {
	createdTeam, err := t.teamRepo.AddTeam(ctx, team)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == PgSqlUniqueViolationErr {
			return model.Team{}, apperror.NewTeamExistsError(team.TeamName)
		}

		return model.Team{}, err
	}
	return createdTeam, nil
}
