package reviewer

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *reviewerRepository) DeleteReviewer(ctx context.Context, tx pgx.Tx, prID, reviewerID string) error {
	const query = `
        DELETE FROM pull_request_reviewers
        WHERE pull_request_id = $1 
          AND reviewer_id = $2;
    `

	_, err := tx.Exec(ctx, query, prID, reviewerID)
	return err
}
