package user

import (
	"context"
)

func (r *userRepository) CheckUserExists(ctx context.Context, userId string) (bool, error) {
	const q = `SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)`

	var exists bool
	err := r.pool.QueryRow(ctx, q, userId).Scan(&exists)
	return exists, err
}
