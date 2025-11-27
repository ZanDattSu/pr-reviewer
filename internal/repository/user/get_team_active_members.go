package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (r *userRepository) GetTeamActiveMembers(ctx context.Context, userID string) ([]string, error) {
	const q = `
        SELECT user_id
        FROM users
        WHERE team_id = (
            SELECT team_id FROM users WHERE user_id = $1
        )
        AND is_active = TRUE
        AND user_id <> $1;
    `

	conn := r.getter.DefaultTrOrDB(ctx, r.pool)
	rows, err := conn.Query(ctx, q, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.NewUserNotFoundError(userID)
		}
		return nil, err
	}
	defer rows.Close()

	var members []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		members = append(members, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}
