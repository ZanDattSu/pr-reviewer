package apperror

import "fmt"

type NoCandidateError struct {
	pullRequestID string
}

func (e *NoCandidateError) Error() string {
	return fmt.Sprintf("no candidate for pull request %s", e.pullRequestID)
}

func NewNoCandidateError(pullRequestID string) *NoCandidateError {
	return &NoCandidateError{pullRequestID: pullRequestID}
}
