package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/converter"
	repoModel "github.com/ZanDattSu/pr-reviewer/internal/repository/model"
)

func (r *userRepository) GetUserStats(
	ctx context.Context,
	top int,
	onlyActive bool,
	onlyOpen bool,
) ([]model.UserStats, error) {
	const q = `
		SELECT 
			u.user_id,
			COUNT(pr.pull_request_id) AS total_pr
		FROM users u
		LEFT JOIN pull_requests pr 
			   ON pr.author_id = u.user_id
		WHERE 
			($1 = FALSE OR u.is_active = TRUE)
			AND ($2 = FALSE OR pr.status_id = $4)
		GROUP BY u.user_id
		ORDER BY total_pr DESC
		LIMIT CASE WHEN $3 > 0 THEN $3 END
	`

	rows, err := r.pool.Query(ctx, q, onlyActive, onlyOpen, top, repoModel.StatusOpen)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.NewNoDataError()
		}
		return nil, err
	}
	defer rows.Close()

	var stats []repoModel.UserStats

	for rows.Next() {
		var s repoModel.UserStats
		if err := rows.Scan(&s.UserID, &s.TotalPR); err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}

	return converter.RepoUserStatsToService(stats), nil
}
