package apperror

import "fmt"

type TeamExistsError struct {
	teamName string
}

func (e *TeamExistsError) Error() string {
	return fmt.Sprintf("team %s already exists", e.teamName)
}

func NewTeamExistsError(name string) *TeamExistsError {
	return &TeamExistsError{
		teamName: name,
	}
}

func (e *TeamExistsError) TeamName() string {
	return e.teamName
}
