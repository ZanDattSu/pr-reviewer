package apperror

import "fmt"

type TeamNotFoundError struct {
	teamName string
}

func (e *TeamNotFoundError) Error() string {
	return fmt.Sprintf("team %s not found", e.teamName)
}

func NewTeamNotFoundError(name string) *TeamNotFoundError {
	return &TeamNotFoundError{
		teamName: name,
	}
}

func (e *TeamNotFoundError) TeamName() string {
	return e.teamName
}
