package apperror

import "fmt"

type UserInAnotherTeamError struct {
	userID string
}

func (e *UserInAnotherTeamError) Error() string {
	return fmt.Sprintf("user %s already in another team", e.userID)
}

func NewUserInAnotherTeamError(name string) *UserInAnotherTeamError {
	return &UserInAnotherTeamError{
		userID: name,
	}
}
