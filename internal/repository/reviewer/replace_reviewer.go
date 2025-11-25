package reviewer

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"

	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
	"github.com/ZanDattSu/pr-reviewer/pkg/logger"
)

const (
	pgForeignKeyErr = "23503"
)

func (r *reviewerRepository) ReplaceReviewer(ctx context.Context, prID, oldReviewerID, newUserID string) error {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return err
	}

	defer func(tx pgx.Tx, ctx context.Context) {
		if rbErr := tx.Rollback(ctx); rbErr != nil && !errors.Is(rbErr, pgx.ErrTxClosed) {
			logger.Error(ctx, "tx rollback failed!", zap.Error(rbErr))
		}
	}(tx, ctx)

	const deleteQuery = `
		DELETE FROM pull_request_reviewers
		WHERE pull_request_id = $1 AND reviewer_id = $2
	`

	cmdTag, err := tx.Exec(ctx, deleteQuery, prID, oldReviewerID)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return apperror.NewNotAssignedError(oldReviewerID)
	}

	const insertQuery = `	
		INSERT INTO pull_request_reviewers (pull_request_id, reviewer_id)
		VALUES ($1, $2)
		`

	_, err = tx.Exec(ctx, insertQuery, prID, newUserID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgForeignKeyErr {
			return apperror.NewUserNotFoundError(newUserID)
		}
		return err

	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
