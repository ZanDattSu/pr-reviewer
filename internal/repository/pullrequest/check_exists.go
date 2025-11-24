package pullrequest

import (
	"context"
	"errors"

	pgx "github.com/jackc/pgx/v5"
)

func (r *prRepository) CheckPRExists(ctx context.Context, id string) (bool, error) {
	const q = `SELECT 1 FROM pull_requests WHERE pull_request_id = $1`

	var exists int
	err := r.pool.QueryRow(ctx, q, id).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
