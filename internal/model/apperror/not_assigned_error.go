package apperror

import "fmt"

type ReviewerNotAssignedError struct {
	reviewerUUID string
}

func NewReviewerNotAssignedError(reviewerUUID string) *ReviewerNotAssignedError {
	return &ReviewerNotAssignedError{reviewerUUID: reviewerUUID}
}

func (e *ReviewerNotAssignedError) Error() string {
	return fmt.Sprintf("reviewer %s is not assigned to this PR", e.reviewerUUID)
}
