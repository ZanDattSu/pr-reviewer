package pullrequest

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/samber/lo"
	"go.uber.org/zap"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/converter"
	"github.com/ZanDattSu/pr-reviewer/pkg/logger"
)

func (r *prRepository) InsertPR(ctx context.Context, pr model.PullRequest) (model.PullRequest, error) {
	repoPR := converter.ServicePRToRepo(pr)

	var createdAt time.Time

	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return model.PullRequest{}, err
	}

	defer func(tx pgx.Tx, ctx context.Context) {
		if rbErr := tx.Rollback(ctx); rbErr != nil && !errors.Is(rbErr, pgx.ErrTxClosed) {
			logger.Error(ctx, "tx rollback failed:", zap.Error(rbErr))
		}
	}(tx, ctx)

	const insertPR = `

        INSERT INTO pull_requests (pull_request_id, pull_request_name, author_id, status_id)

        VALUES ($1, $2, $3, $4)

        RETURNING created_at;

    `

	err = tx.QueryRow(ctx, insertPR,

		repoPR.PullRequestID,

		repoPR.PullRequestName,

		repoPR.AuthorID,

		repoPR.Status,
	).Scan(&createdAt)
	if err != nil {
		return model.PullRequest{}, err
	}

	if len(pr.AssignedReviewers) != 0 {

		err := updatePRReviewers(ctx, tx, pr)
		if err != nil {
			return model.PullRequest{}, err
		}

	}

	if err := tx.Commit(ctx); err != nil {
		return model.PullRequest{}, err
	}

	pr.CreatedAt = lo.ToPtr(createdAt)

	return pr, err
}

func updatePRReviewers(ctx context.Context, tx pgx.Tx, pr model.PullRequest) error {
	const insertPRReviewers = `

		INSERT INTO pull_request_reviewers (pull_request_id, reviewer_id)

		VALUES ($1, $2)

		`

	for _, reviewer := range pr.AssignedReviewers {

		_, err := tx.Exec(ctx, insertPRReviewers, pr.PullRequestID, reviewer)
		if err != nil {
			return err
		}

	}

	return nil
}
