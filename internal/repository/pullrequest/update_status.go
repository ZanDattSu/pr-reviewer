package pullrequest

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/converter"
	repoModel "github.com/ZanDattSu/pr-reviewer/internal/repository/model"
)

func (r *prRepository) UpdatePRStatus(ctx context.Context, prID string, status model.Status) (model.PullRequest, error) {
	const q = `
        UPDATE pull_requests
        SET status_id = $2,
            merged_at = COALESCE(merged_at, NOW())
        WHERE pull_request_id = $1
        RETURNING pull_request_id, pull_request_name, author_id, status_id, created_at, merged_at;
    `

	var pr repoModel.PullRequest

	err := r.pool.QueryRow(ctx, q, prID, converter.ServicePRStatusToRepo(status)).Scan(
		&pr.PullRequestID,
		&pr.PullRequestName,
		&pr.AuthorID,
		&pr.Status,
		&pr.CreatedAt,
		&pr.MergedAt,
	)
	return converter.RepoPRToService(pr), err
}
