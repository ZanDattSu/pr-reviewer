package reviewer

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *reviewerRepository) AddReviewer(ctx context.Context, tx pgx.Tx, prID, reviewerID string) error {
	const query = `
        INSERT INTO pull_request_reviewers (pull_request_id, reviewer_id)
        VALUES ($1, $2);
    `

	_, err := tx.Exec(ctx, query, prID, reviewerID)
	return err
}
