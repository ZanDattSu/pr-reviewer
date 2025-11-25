package user

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (s *userService) DeactivateUsers(ctx context.Context, userIDs []string) ([]model.DeactivateResult, error) {
	if len(userIDs) == 0 {
		return []model.DeactivateResult{}, nil
	}

	mapping, err := s.userRepo.DeactivateUsersAndReassign(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	results := make([]model.DeactivateResult, 0, len(mapping))
	for prID, newReviewer := range mapping {
		results = append(results, model.DeactivateResult{
			PullRequestID: prID,
			ReplacedBy:    newReviewer,
		})
	}

	return results, nil
}
