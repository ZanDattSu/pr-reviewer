package user

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/converter"
	repoModel "github.com/ZanDattSu/pr-reviewer/internal/repository/model"
)

func (r *userRepository) UpdateUserIsActive(
	ctx context.Context,
	userID string,
	isActive bool,
) (model.User, error) {
	const q = `
        UPDATE users u
        SET is_active = $2
        FROM teams t
        WHERE u.user_id = $1
          AND t.team_id = u.team_id
        RETURNING u.user_id, u.username, t.team_name, u.is_active
    `

	var u repoModel.User
	err := r.pool.QueryRow(ctx, q, userID, isActive).Scan(
		&u.UserID,
		&u.Username,
		&u.TeamName,
		&u.IsActive,
	)

	return converter.RepoUserToService(u), err
}

/*func (r *userRepository) UserSetIsActive(ctx context.Context, userID string, isActive bool) (model.User, error) {
	var user repoModel.User
	const query = `UPDATE users u
				   SET is_active = $2
				   FROM teams t
				   WHERE u.user_id = $1
				   AND t.team_uuid = u.team_uuid
				   RETURNING u.user_id, u.username, t.team_name, u.is_active`
	err := r.pool.QueryRow(ctx, query, userID, isActive).Scan(
		&user.UserID,
		&user.Username,
		&user.TeamName,
		&user.IsActive,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, apperror.NewUserNotFoundError(userID)
		}
		return model.User{}, err
	}

	return converter.RepoUserToService(user), nil
}*/
