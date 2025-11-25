package apperror

import "fmt"

type PRNotFoundError struct {
	pullRequestID string
}

func (e *PRNotFoundError) PullRequestID() string {
	return e.pullRequestID
}

func (e *PRNotFoundError) Error() string {
	return fmt.Sprintf("pull request %s not found", e.pullRequestID)
}

func NewPRNotFoundError(name string) *PRNotFoundError {
	return &PRNotFoundError{
		pullRequestID: name,
	}
}
