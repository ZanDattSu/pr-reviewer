package reviewer

import (
	"context"
)

type ReviewerRepository interface {
	GetReviewersForPR(ctx context.Context, prID string) ([]string, error)
	ReplaceReviewer(ctx context.Context, prID, oldReviewerID, newUserID string) error
}
