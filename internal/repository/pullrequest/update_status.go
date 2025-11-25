package pullrequest

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/converter"
)

func (r *prRepository) UpdatePRStatus(ctx context.Context, pr model.PullRequest, status model.Status) (model.PullRequest, error) {
	const q = `
        UPDATE pull_requests
        SET status_id = $2,
            merged_at = COALESCE(merged_at, NOW())
        WHERE pull_request_id = $1
        RETURNING status_id, merged_at;
    `

	pr.Status = status

	repoPR := converter.ServicePRToRepo(pr)

	err := r.pool.QueryRow(ctx, q, pr.PullRequestID, repoPR.Status).Scan(&repoPR.Status, &repoPR.MergedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.PullRequest{}, apperror.NewPRNotFoundError(pr.PullRequestID)
		}
		return model.PullRequest{}, err

	}

	return converter.RepoPRToService(repoPR), nil
}
