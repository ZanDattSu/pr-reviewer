package pullrequest

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/converter"
	repoModel "github.com/ZanDattSu/pr-reviewer/internal/repository/model"
)

func (r *prRepository) GetPRWithReviewers(ctx context.Context, prID string) (model.PullRequest, []string, error) {
	const query = `
        SELECT 
            pr.pull_request_id,
            pr.pull_request_name,
            pr.author_id,
            pr.status_id,
            array_agg(prr.reviewer_id) FILTER (WHERE prr.reviewer_id IS NOT NULL)
        FROM pull_requests pr
        LEFT JOIN pull_request_reviewers prr 
              ON pr.pull_request_id = prr.pull_request_id
        WHERE pr.pull_request_id = $1
        GROUP BY pr.pull_request_id;
    `

	var pr repoModel.PullRequest
	var reviewers []string

	err := r.pool.QueryRow(ctx, query, prID).Scan(
		&pr.PullRequestID,
		&pr.PullRequestName,
		&pr.AuthorID,
		&pr.Status,
		&reviewers,
	)
	if err != nil {
		return model.PullRequest{}, nil, err
	}

	return converter.RepoPRToService(pr), reviewers, nil
}
