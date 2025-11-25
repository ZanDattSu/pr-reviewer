package pullrequest

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/converter"
	repoModel "github.com/ZanDattSu/pr-reviewer/internal/repository/model"
)

func (r *prRepository) GetPRWithReviewers(ctx context.Context, prID string) (model.PullRequest, error) {
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
		if errors.Is(err, pgx.ErrNoRows) {
			return model.PullRequest{}, apperror.NewPRNotFoundError(prID)
		}
		return model.PullRequest{}, err
	}

	pr.AssignedReviewers = reviewers

	return converter.RepoPRToService(pr), nil
}
