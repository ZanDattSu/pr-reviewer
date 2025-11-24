package team

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/converter"
	repoModel "github.com/ZanDattSu/pr-reviewer/internal/repository/model"
)

func (r *teamRepository) GetTeam(ctx context.Context, teamName string) (model.Team, error) {
	teamRow, err := r.getTeamRow(ctx, teamName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Team{}, apperror.NewTeamNotFoundError(teamName)
		}
		return model.Team{}, err
	}

	members, err := r.getTeamMembers(ctx, teamName)
	if err != nil {
		return model.Team{}, err
	}

	teamRow.Members = members
	return converter.RepoTeamToService(teamRow), nil
}

func (r *teamRepository) getTeamRow(ctx context.Context, teamName string) (repoModel.Team, error) {
	const q = `SELECT team_name FROM teams WHERE team_name = $1`

	var t repoModel.Team
	err := r.pool.QueryRow(ctx, q, teamName).Scan(
		&t.TeamName,
	)

	return t, err
}

func (r *teamRepository) getTeamMembers(ctx context.Context, teamName string) ([]repoModel.TeamMember, error) {
	const q = `
        SELECT u.user_id, u.username, u.is_active
        FROM users u
        JOIN teams t ON u.team_id = t.team_id
        WHERE t.team_name = $1
    `

	rows, err := r.pool.Query(ctx, q, teamName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []repoModel.TeamMember

	for rows.Next() {
		var m repoModel.TeamMember
		if err := rows.Scan(&m.UserUUID, &m.Username, &m.IsActive); err != nil {
			return nil, err
		}
		members = append(members, m)
	}

	return members, nil
}
