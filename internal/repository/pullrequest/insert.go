package pullrequest

import (
	"context"
	"time"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/converter"
)

func (r *prRepository) InsertPR(ctx context.Context, pr model.PullRequest) (time.Time, error) {
	repoPR := converter.ServicePRToRepo(pr)
	var createdAt time.Time

	const q = `
        INSERT INTO pull_requests (pull_request_id, pull_request_name, author_id, status_id)
        VALUES ($1, $2, $3, $4)
        RETURNING created_at;
    `
	err := r.pool.QueryRow(ctx, q,
		repoPR.PullRequestID,
		repoPR.PullRequestName,
		repoPR.AuthorID,
		repoPR.Status,
	).Scan(&createdAt)

	return createdAt, err
}
