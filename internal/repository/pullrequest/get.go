package pullrequest

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/converter"
	repoModel "github.com/ZanDattSu/pr-reviewer/internal/repository/model"
)

func (r *prRepository) GetPR(ctx context.Context, prID string) (model.PullRequest, error) {
	const query = `
        SELECT pull_request_id,
               pull_request_name,
               author_id,
               status_id,
               created_at,
               merged_at
        FROM pull_requests
        WHERE pull_request_id = $1
    `

	var pr repoModel.PullRequest

	err := r.pool.QueryRow(ctx, query, prID).Scan(
		&pr.PullRequestID,
		&pr.PullRequestName,
		&pr.AuthorID,
		&pr.Status,
		&pr.CreatedAt,
		&pr.MergedAt,
	)
	if err != nil {
		return model.PullRequest{}, err
	}

	return converter.RepoPRToService(pr), nil
}
