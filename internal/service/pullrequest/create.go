package pullrequest

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

const (
	pgForeignKeyErr = "23503"
)

/*func (r *prRepository) CreatePullRequest(ctx context.Context, pullRequestID, pullRequestName, authorID string) (model.PullRequest, error) {
	pr := model.PullRequest{
		PullRequestID:   pullRequestID,
		PullRequestName: pullRequestName,
		AuthorID:        authorID,
		Status:          model.StatusOpen,
	}

	const insertQuery = `
						INSERT INTO pull_requests (pull_request_id, pull_request_name, author_id, status_id)
						VALUES ($1, $2, $3, $4)
						RETURNING created_at`
	repoStatus := converter.ServicePRStatusToRepo(pr.Status)

	var createdAt time.Time
	err := r.pool.QueryRow(ctx,
		insertQuery,
		pr.PullRequestID,
		pr.PullRequestName,
		pr.AuthorID,
		repoStatus,
	).Scan(&createdAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgForeignKeyErr:
				return model.PullRequest{}, apperror.NewUserNotFoundError(pr.AuthorID)
			}
		}
		return model.PullRequest{}, err
	}

	pr.CreatedAt = &createdAt
	return pr, nil
}
*/

func (s *prService) CreatePullRequest(ctx context.Context, pullRequestID, pullRequestName, authorID string) (model.PullRequest, error) {
	// TODO implement me
	panic("implement me")
}
