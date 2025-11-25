package user

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"github.com/ZanDattSu/pr-reviewer/pkg/logger"
)

func (r *userRepository) DeactivateUsersAndReassign(ctx context.Context, userIDs []string) (map[string]string, error) {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return nil, err
	}

	defer func(tx pgx.Tx, ctx context.Context) {
		if rbErr := tx.Rollback(ctx); rbErr != nil && !errors.Is(rbErr, pgx.ErrTxClosed) {
			logger.Warn(ctx, "rollback failed", zap.Error(rbErr))
		}
	}(tx, ctx)

	defer func(tx pgx.Tx, ctx context.Context) {
		if commitErr := tx.Commit(ctx); commitErr != nil && !errors.Is(commitErr, pgx.ErrTxClosed) {
			logger.Warn(ctx, "commit failed", zap.Error(commitErr))
		}
	}(tx, ctx)

	deactivated, err := r.deactivateUsers(ctx, tx, userIDs)
	if err != nil {
		return nil, err
	}
	if len(deactivated) == 0 {
		return nil, fmt.Errorf("no users deactivated")
	}

	openPRs, err := r.findOpenPRsWithUsers(ctx, tx, deactivated)
	if err != nil {
		return nil, err
	}

	if len(openPRs) == 0 {
		return map[string]string{}, nil
	}

	result := make(map[string]string)
	for _, item := range openPRs {
		newID, err := r.reassignSingleReviewer(ctx, tx, item.prID, item.oldReviewer)
		if err != nil {
			return nil, err
		}
		result[item.prID] = newID
	}

	return result, nil
}

func (r *userRepository) deactivateUsers(
	ctx context.Context,
	tx pgx.Tx,
	userIDs []string,
) ([]string, error) {
	const q = `
        UPDATE users
        SET is_active = FALSE
        WHERE user_id = ANY($1)
        RETURNING user_id
    `

	rows, err := tx.Query(ctx, q, userIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		res = append(res, id)
	}
	return res, nil
}

type openPR struct {
	prID        string
	oldReviewer string
}

func (r *userRepository) findOpenPRsWithUsers(
	ctx context.Context,
	tx pgx.Tx,
	userIDs []string,
) ([]openPR, error) {
	const q = `
        SELECT prr.pull_request_id, prr.reviewer_id
        FROM pull_request_reviewers prr
        JOIN pull_requests pr
          ON pr.pull_request_id = prr.pull_request_id
        WHERE prr.reviewer_id = ANY($1)
          AND pr.status_id = 1
    `

	rows, err := tx.Query(ctx, q, userIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []openPR
	for rows.Next() {
		var o openPR
		if err := rows.Scan(&o.prID, &o.oldReviewer); err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, nil
}

func (r *userRepository) reassignSingleReviewer(
	ctx context.Context,
	tx pgx.Tx,
	prID, oldReviewer string,
) (string, error) {
	candidates, err := r.getTeamActiveMembersWithoutUserTx(ctx, tx, oldReviewer)
	if err != nil {
		return "", err
	}
	if len(candidates) == 0 {
		return "", fmt.Errorf("no candidates for %s", oldReviewer)
	}

	newReviewer := pickOne(candidates)

	if err := r.replaceReviewer(ctx, tx, prID, oldReviewer, newReviewer); err != nil {
		return "", err
	}

	return newReviewer, nil
}

func (r *userRepository) replaceReviewer(
	ctx context.Context,
	tx pgx.Tx,
	prID, oldID, newID string,
) error {
	const q = `
        UPDATE pull_request_reviewers
        SET reviewer_id = $3
        WHERE pull_request_id = $1 AND reviewer_id = $2
    `
	_, err := tx.Exec(ctx, q, prID, oldID, newID)
	return err
}

func (r *userRepository) getTeamActiveMembersWithoutUserTx(
	ctx context.Context,
	tx pgx.Tx,
	userID string,
) ([]string, error) {
	const q = `
        SELECT u2.user_id
        FROM users u
        JOIN users u2 ON u.team_id = u2.team_id
        WHERE u.user_id = $1
          AND u2.user_id <> $1
          AND u2.is_active = TRUE
    `

	rows, err := tx.Query(ctx, q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		list = append(list, id)
	}
	return list, nil
}

func pickOne(list []string) string {
	if len(list) == 1 {
		return list[0]
	}
	return list[rand.Intn(len(list))] //nolint:gosec
}
