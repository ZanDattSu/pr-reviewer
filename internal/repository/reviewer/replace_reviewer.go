package reviewer

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

const (
	pgForeignKeyErr = "23503"
)

func (r *reviewerRepository) ReplaceReviewer(ctx context.Context, prID, oldReviewerID, newReviewerID string) error {
	const q = `
        UPDATE pull_request_reviewers
        SET reviewer_id = $3
        WHERE pull_request_id = $1 AND reviewer_id = $2
    `

	conn := r.getter.DefaultTrOrDB(ctx, r.pool)
	cmdTag, err := conn.Exec(ctx, q, prID, oldReviewerID, newReviewerID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgForeignKeyErr {
			return apperror.NewUserNotFoundError(newReviewerID)
		}
		return err

	}

	if cmdTag.RowsAffected() == 0 {
		return apperror.NewNotAssignedError(oldReviewerID)
	}

	return err
}
