package user

import (
	"context"
	"fmt"
	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (s *userService) DeactivateUsersAndReassignPR(ctx context.Context, userIDs []string) ([]model.ReassignedPR, error) {
	reassignedPRs := make([]model.ReassignedPR, 0)

	err := s.tm.Do(ctx, func(ctx context.Context) error {
		deactivated, err := s.userRepo.DeactivateUsers(ctx, userIDs)
		if err != nil {
			return fmt.Errorf("failed to deactivate users: %w", err)
		}

		if len(deactivated) == 0 {
			return fmt.Errorf("no users were deactivated")
		}

		openPRs, err := s.prRepo.FindOpenPRsWithReviewers(ctx, deactivated)
		if err != nil {
			return fmt.Errorf("failed to find open PRs: %w", err)
		}

		if len(openPRs) == 0 {
			return nil
		}

		for _, openPR := range openPRs {
			pr, newReviewerID, err := s.prService.ReassignPullRequest(ctx, openPR.PRID, openPR.OldReviewer)
			if err != nil {
				return fmt.Errorf("failed to reassign PR %s: %w", openPR.PRID, err)
			}

			reassignedPRs = append(reassignedPRs, model.ReassignedPR{
				PullRequestID: pr,
				ReplacedBy:    newReviewerID,
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return reassignedPRs, nil
}
