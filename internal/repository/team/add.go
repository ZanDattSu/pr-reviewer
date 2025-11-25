package team

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/converter"
	repoModel "github.com/ZanDattSu/pr-reviewer/internal/repository/model"
	"github.com/ZanDattSu/pr-reviewer/pkg/logger"
)

const PgSqlUniqueViolationErr = "23505"

func (r *teamRepository) AddTeam(ctx context.Context, team model.Team) (model.Team, error) {
	repoTeam := converter.ServiceTeamToRepo(team)

	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return model.Team{}, err
	}

	defer func(tx pgx.Tx, ctx context.Context) {
		if rbErr := tx.Rollback(ctx); rbErr != nil && !errors.Is(rbErr, pgx.ErrTxClosed) {
			logger.Warn(ctx, "rollback failed", zap.Error(rbErr))
		}
	}(tx, ctx)

	teamID, err := r.insertTeam(ctx, tx, repoTeam.TeamName)
	if err != nil {
		return model.Team{}, err
	}

	err = r.upsertTeamMembers(ctx, tx, teamID, repoTeam.Members)
	if err != nil {
		return model.Team{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return model.Team{}, err
	}

	return converter.RepoTeamToService(repoTeam), nil
}

func (r *teamRepository) insertTeam(ctx context.Context, tx pgx.Tx, teamName string) (string, error) {
	const q = `
        INSERT INTO teams (team_name)
        VALUES ($1)
        RETURNING team_id
    `

	var teamID string

	err := tx.QueryRow(ctx, q, teamName).Scan(&teamID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == PgSqlUniqueViolationErr {
			return "", apperror.NewTeamExistsError(teamName)
		}
		return "", err
	}

	return teamID, err
}

func (r *teamRepository) upsertTeamMembers(
	ctx context.Context,
	tx pgx.Tx,
	teamID string,
	members []repoModel.TeamMember,
) error {
	// Возвращаем user_id чтобы проверить pgx.ErrNoRows и вернуть ошибку UserInAnotherTeam
	const q = `
        INSERT INTO users (user_id, username, team_id, is_active)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (user_id) DO UPDATE 
            SET username = EXCLUDED.username,
                is_active = EXCLUDED.is_active
        WHERE users.team_id = EXCLUDED.team_id
        RETURNING user_id
    `

	batch := &pgx.Batch{}

	for _, m := range members {
		batch.Queue(q, m.UserID, m.Username, teamID, m.IsActive)
	}

	batchRes := tx.SendBatch(ctx, batch)

	defer func(batchRes pgx.BatchResults) {
		if brCerr := batchRes.Close(); brCerr != nil {
			logger.Warn(ctx, "Batch close error", zap.Error(brCerr))
		}
	}(batchRes)

	for _, member := range members {
		var userID string
		if err := batchRes.QueryRow().Scan(&userID); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return apperror.NewUserInAnotherTeamError(member.UserID)
			}
			return err
		}
	}

	return nil
}
