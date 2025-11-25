package pullrequest

import (
	"context"
)

func (r *prRepository) CheckPRExists(ctx context.Context, id string) (bool, error) {
	const q = `SELECT EXISTS(SELECT 1 FROM pull_requests WHERE pull_request_id = $1)`

	var exists bool
	err := r.pool.QueryRow(ctx, q, id).Scan(&exists)
	return exists, err
}
