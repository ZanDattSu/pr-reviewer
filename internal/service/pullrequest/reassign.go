package pullrequest

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

/*
проверяет, что PR существует

проверяет, что oldReviewer назначен на PR

находит активных пользователей его команды, исключая старого

выбирает случайного

меняет ревьювера в таблице pull_request_reviewers

возвращает обновлённый PR
*/

/*func (r *prRepository) ReassignPullRequest(ctx context.Context, pullRequestID, oldUserID string) (model.PullRequest, string, error) {
	const getPRQuery = `
			SELECT pr.pull_request_id,
			       pr.pull_request_name,
			       pr.author_id,
			       pr.status_id,
			       array_agg(prr.reviewer_id)
			FROM pull_requests pr
			    LEFT JOIN pull_request_reviewers prr ON pr.pull_request_id = prr.pull_request_id
			WHERE pr.pull_request_id = $1
			GROUP BY pr.pull_request_id`

	var pr repoModel.PullRequest
	var oldReviewers []string
	err := r.pool.QueryRow(ctx, getPRQuery, pullRequestID).Scan(
		&pr.PullRequestID,
		&pr.PullRequestName,
		&pr.AuthorID,
		&pr.Status,
		&oldReviewers,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.PullRequest{}, "", apperror.NewPRNotFoundError(pullRequestID)
		}
		return model.PullRequest{}, "", err
	}

	isOldReviewerAssigned := false
	for _, oldReviewer := range oldReviewers {
		if oldReviewer == oldUserID {
			isOldReviewerAssigned = true
			break
		}
	}
	if !isOldReviewerAssigned {
		return model.PullRequest{}, "", apperror.NewReviewerNotAssignedError(oldUserID)
	}

	const query = `
					SELECT u.user_id
					FROM users u
					LEFT JOIN pull_request_reviewers prr
					    ON u.user_id = prr.reviewer_id AND prr.pull_request_id = $2
					WHERE u.team_uuid = (SELECT team_uuid FROM users WHERE user_id = $1)
					AND u.is_active IS TRUE
					AND u.user_id <> $1
					AND prr.reviewer_id IS NULL`
	rows, err := r.pool.Query(ctx, query, oldUserID, pr.PullRequestID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.PullRequest{}, "", apperror.NewNoCandidateError(pullRequestID)
		}
		return model.PullRequest{}, "", err
	}

	var prCandidate []string
	for rows.Next() {
		var userID string
		if err := rows.Scan(&userID); err != nil {
			return model.PullRequest{}, "", err
		}
		prCandidate = append(prCandidate, userID)
	}
}*/

func (s *prService) ReassignPullRequest(ctx context.Context, pullRequestID, oldUserID string) (model.PullRequest, string, error) {
	// TODO implement me
	panic("implement me")
}
