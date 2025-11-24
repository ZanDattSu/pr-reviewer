package reviewer

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type ReviewerRepository interface {
	GetReviewersForPR(ctx context.Context, prID string) ([]string, error)
	AddReviewer(ctx context.Context, tx pgx.Tx, prID, reviewerID string) error
	DeleteReviewer(ctx context.Context, tx pgx.Tx, prID, reviewerID string) error
}
