package apperror

import "fmt"

type PRNotFoundError struct {
	pullRequestUUID string
}

func (e *PRNotFoundError) Error() string {
	return fmt.Sprintf("pull request %s not found", e.pullRequestUUID)
}

func NewPRNotFoundError(name string) *PRNotFoundError {
	return &PRNotFoundError{
		pullRequestUUID: name,
	}
}
