package team

import (
	"context"
)

func (r *teamRepository) GetTeamActiveMembersWithoutUser(ctx context.Context, userID string) ([]string, error) {
	const query = `
        SELECT user_id
        FROM users
        WHERE team_id = (
            SELECT team_id FROM users WHERE user_id = $1
        )
        AND is_active = TRUE
        AND user_id <> $1;
    `

	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members := make([]string, 0)
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		members = append(members, id)
	}

	return members, nil
}
