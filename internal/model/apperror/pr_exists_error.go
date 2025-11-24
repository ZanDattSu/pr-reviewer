package apperror

import "fmt"

type PRExistsError struct {
	pullRequestID string
}

func (e *PRExistsError) Error() string {
	return fmt.Sprintf("pull request %s already exists", e.pullRequestID)
}

func NewPRExistsError(pullRequestID string) *PRExistsError {
	return &PRExistsError{pullRequestID: pullRequestID}
}
