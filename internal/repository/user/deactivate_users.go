package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *userRepository) DeactivateUsers(ctx context.Context, userIDs []string) ([]string, error) {
	const q = `
        UPDATE users
        SET is_active = FALSE
        WHERE user_id = ANY($1)
        RETURNING user_id
    `

	conn := r.getter.DefaultTrOrDB(ctx, r.pool)
	rows, err := conn.Query(ctx, q, userIDs)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []string{}, fmt.Errorf("no users were deactivated")
		}
		return nil, err
	}
	defer rows.Close()

	deactivated := make([]string, 0)
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		deactivated = append(deactivated, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return deactivated, nil
}
