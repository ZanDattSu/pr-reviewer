package user

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/converter"
	repoModel "github.com/ZanDattSu/pr-reviewer/internal/repository/model"
)

func (r *userRepository) GetPRReviewer(ctx context.Context, userID string) ([]model.UserAssignedPR, error) {
	const q = `
        SELECT pr.pull_request_id,
               pr.pull_request_name,
               pr.author_id,
               prs.id
        FROM pull_request_reviewers prr
        JOIN pull_requests pr ON prr.pull_request_id = pr.pull_request_id
        JOIN pull_request_status prs ON pr.status_id = prs.id
        WHERE prr.reviewer_id = $1
    `

	rows, err := r.pool.Query(ctx, q, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	userPRs := make([]repoModel.UserAssignedPR, 0)

	for rows.Next() {
		var pr repoModel.UserAssignedPR
		err := rows.Scan(
			&pr.PullRequestID,
			&pr.PullRequestName,
			&pr.AuthorID,
			&pr.Status,
		)
		if err != nil {
			return nil, err
		}
		userPRs = append(userPRs, pr)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return converter.RepoUserAssignedPRsToService(userPRs), nil
}
