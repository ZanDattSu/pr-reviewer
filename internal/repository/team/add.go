package team

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/converter"
	repoModel "github.com/ZanDattSu/pr-reviewer/internal/repository/model"
	"github.com/ZanDattSu/pr-reviewer/pkg/logger"
)

func (r *teamRepository) AddTeam(ctx context.Context, team model.Team) (model.Team, error) {
	repoTeam := converter.ServiceTeamToRepo(team)

	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return model.Team{}, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			logger.Warn(ctx, "Rollback", zap.Error(err))
		}
	}(tx, ctx)

	teamUUID, err := r.insertTeam(ctx, tx, repoTeam.TeamName)
	if err != nil {
		return model.Team{}, err
	}

	err = r.upsertTeamMembers(ctx, tx, teamUUID, repoTeam.Members)
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
	return teamID, err
}

func (r *teamRepository) upsertTeamMembers(
	ctx context.Context,
	tx pgx.Tx,
	teamUUID string,
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
		batch.Queue(q, m.UserUUID, m.Username, teamUUID, m.IsActive)
	}

	batchRes := tx.SendBatch(ctx, batch)
	defer func(batchRes pgx.BatchResults) {
		err := batchRes.Close()
		if err != nil {
			logger.Warn(ctx, "Batch close error", zap.Error(err))
		}
	}(batchRes)

	for _, member := range members {
		var userUUID string
		if err := batchRes.QueryRow().Scan(&userUUID); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return apperror.NewUserInAnotherTeamError(member.UserUUID)
			}
			return err
		}
	}

	return nil
}
