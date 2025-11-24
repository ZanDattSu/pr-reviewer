package reviewer

import "context"

func (r *reviewerRepository) GetReviewersForPR(ctx context.Context, prID string) ([]string, error) {
	const q = `
        SELECT reviewer_id
        FROM pull_request_reviewers
        WHERE pull_request_id = $1;
    `
	rows, err := r.pool.Query(ctx, q, prID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]string, 0)

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		res = append(res, id)
	}

	return res, nil
}
