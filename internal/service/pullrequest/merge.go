package pullrequest

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

/*func (r *prRepository) MergePullRequest(ctx context.Context, pullRequestID string) (model.PullRequest, error) {
	//UpdatePRStatusToMerged
	var pr repoModel.PullRequest

	const query = `
					UPDATE pull_requests pr
					SET status_id = 2,
						merged_at = COALESCE(merged_at, NOW())
					WHERE pr.pull_request_id = $1
					RETURNING pr.pull_request_id,
						pr.pull_request_name,
						pr.author_id,
						pr.status_id,
						pr.created_at,
						pr.merged_at`
	err := r.pool.QueryRow(ctx, query, pullRequestID).Scan(
		&pr.PullRequestID,
		&pr.PullRequestName,
		&pr.AuthorID,
		&pr.Status,
		&pr.CreatedAt,
		&pr.MergedAt,
	)
	// сервис чекуть ошибку
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.PullRequest{}, apperror.NewPRNotFoundError(pullRequestID)
		}
		return model.PullRequest{}, err
	}

	//GetReviewersForPR
	rows, err := r.pool.Query(ctx,
		`
			SELECT prr.reviewer_id
			FROM pull_request_reviewers prr
			WHERE pull_request_id = $1`,
		pullRequestID)
	if err != nil {
		return model.PullRequest{}, err
	}

	for rows.Next() {
		var reviewerID string
		if err = rows.Scan(&reviewerID); err != nil {
			return model.PullRequest{}, err
		}
		pr.AssignedReviewers = append(pr.AssignedReviewers, reviewerID)
	}

	return converter.RepoPRToService(pr), nil
}
*/

func (s *prService) MergePullRequest(ctx context.Context, pullRequestID string) (model.PullRequest, error) {
	// TODO implement me
	panic("implement me")
}
