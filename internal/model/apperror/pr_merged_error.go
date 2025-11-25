package apperror

import "fmt"

type PRMergedError struct {
	prID string
}

func (e *PRMergedError) Error() string {
	return fmt.Sprintf("cannot reassign on merged PR %s", e.prID)
}

func (e *PRMergedError) PullRequestID() string {
	return e.prID
}

func NewPRMergedError(prID string) *PRMergedError {
	return &PRMergedError{
		prID: prID,
	}
}
