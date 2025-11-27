package pullrequest

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (r *prRepository) FindOpenPRsWithReviewers(ctx context.Context, reviewerIDs []string) ([]model.OpenPR, error) {
	const q = `
        SELECT prr.pull_request_id, prr.reviewer_id
        FROM pull_request_reviewers prr
        JOIN pull_requests pr
          ON pr.pull_request_id = prr.pull_request_id
        WHERE prr.reviewer_id = ANY($1)
          AND pr.status_id = 1
    `

	conn := r.getter.DefaultTrOrDB(ctx, r.pool)
	rows, err := conn.Query(ctx, q, reviewerIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.OpenPR
	for rows.Next() {
		var pr model.OpenPR
		if err := rows.Scan(&pr.PRID, &pr.OldReviewer); err != nil {
			return nil, err
		}
		list = append(list, pr)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
