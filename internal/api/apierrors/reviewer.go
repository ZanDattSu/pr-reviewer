package apierrors

import (
	"fmt"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
)

func NewPrExistErr(prID string) reviewerV1.ErrorResponseError {
	return reviewerV1.ErrorResponseError{
		Code:    reviewerV1.ErrorResponseErrorCodePREXISTS,
		Message: fmt.Sprintf("PR %s already exists", prID),
	}
}

func NewTeamExistErr(teamName string) reviewerV1.ErrorResponseError {
	return reviewerV1.ErrorResponseError{
		Code:    reviewerV1.ErrorResponseErrorCodeTEAMEXISTS,
		Message: fmt.Sprintf("team %s already exists", teamName),
	}
}

func NewAuthorNotFoundErr(userID string) reviewerV1.ErrorResponseError {
	return reviewerV1.ErrorResponseError{
		Code:    reviewerV1.ErrorResponseErrorCodeNOTFOUND,
		Message: fmt.Sprintf("author %s not found", userID),
	}
}

func NewUserNotFoundErr(userID string) reviewerV1.ErrorResponseError {
	return reviewerV1.ErrorResponseError{
		Code:    reviewerV1.ErrorResponseErrorCodeNOTFOUND,
		Message: fmt.Sprintf("user %s not found", userID),
	}
}

func NewNoStatsError() reviewerV1.ErrorResponseError {
	return reviewerV1.ErrorResponseError{
		Code:    reviewerV1.ErrorResponseErrorCodeNOTFOUND,
		Message: "no stats available",
	}
}

func NewTeamNotFoundErr(teamName string) reviewerV1.ErrorResponseError {
	return reviewerV1.ErrorResponseError{
		Code:    reviewerV1.ErrorResponseErrorCodeNOTFOUND,
		Message: fmt.Sprintf("team %s not found", teamName),
	}
}

func NewReviewerNotAssignedErr(userID string) reviewerV1.ErrorResponseError {
	return reviewerV1.ErrorResponseError{
		Code:    reviewerV1.ErrorResponseErrorCodeNOTASSIGNED,
		Message: fmt.Sprintf("reviewer %s is not assigned to this PR", userID),
	}
}

func NewPRNotFoundErr(prID string) reviewerV1.ErrorResponseError {
	return reviewerV1.ErrorResponseError{
		Code:    reviewerV1.ErrorResponseErrorCodeNOTFOUND,
		Message: fmt.Sprintf("PR %s not found", prID),
	}
}

func NewPRMergedErr(prID string) reviewerV1.ErrorResponseError {
	return reviewerV1.ErrorResponseError{
		Code:    reviewerV1.ErrorResponseErrorCodePRMERGED,
		Message: fmt.Sprintf("cannot reassign on merged PR %s", prID),
	}
}

func NewNoCandidateErr() reviewerV1.ErrorResponseError {
	return reviewerV1.ErrorResponseError{
		Code:    reviewerV1.ErrorResponseErrorCodeNOCANDIDATE,
		Message: "no active replacement candidate in team",
	}
}
