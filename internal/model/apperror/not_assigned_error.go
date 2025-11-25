package apperror

import "fmt"

type NotAssignedError struct {
	reviewerID string
}

func (e *NotAssignedError) ReviewerID() string {
	return e.reviewerID
}

func NewNotAssignedError(reviewerID string) *NotAssignedError {
	return &NotAssignedError{reviewerID: reviewerID}
}

func (e *NotAssignedError) Error() string {
	return fmt.Sprintf("reviewer %s is not assigned to this PR", e.reviewerID)
}
