package apperror

import "fmt"

type UserInAnotherTeamError struct {
	userUUID string
}

func (e *UserInAnotherTeamError) Error() string {
	return fmt.Sprintf("user %s already in another team", e.userUUID)
}

func NewUserInAnotherTeamError(name string) *UserInAnotherTeamError {
	return &UserInAnotherTeamError{
		userUUID: name,
	}
}
