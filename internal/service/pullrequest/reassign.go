package pullrequest

import (
	"context"
	"slices"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (s *prService) ReassignPullRequest(ctx context.Context, pullRequestID, oldReviewerID string) (model.PullRequest, string, error) {
	pr, err := s.prRepo.GetPRWithReviewers(ctx, pullRequestID)
	if err != nil {
		return model.PullRequest{}, "", err
	}

	if pr.Status == model.StatusMerged {
		return model.PullRequest{}, "", apperror.NewPRMergedError(pullRequestID)
	}

	if !slices.Contains(pr.AssignedReviewers, oldReviewerID) {
		return model.PullRequest{}, "", apperror.NewNotAssignedError(oldReviewerID)
	}

	teamActiveMembers, err := s.teamRepo.GetTeamActiveMembersWithoutUser(ctx, oldReviewerID)
	if err != nil {
		return model.PullRequest{}, "", err
	}

	if len(teamActiveMembers) == 0 {
		return model.PullRequest{}, "", apperror.NewNoCandidateError(pullRequestID)
	}

	newReviewerID := pickReviewers(teamActiveMembers, 1)[0]

	if err := s.reviewerRepo.ReplaceReviewer(ctx, pullRequestID, oldReviewerID, newReviewerID); err != nil {
		return model.PullRequest{}, "", err
	}

	pr.AssignedReviewers = updatePRReviewers(pr, oldReviewerID, newReviewerID)

	return pr, newReviewerID, nil
}

func updatePRReviewers(pr model.PullRequest, oldReviewerID, newReviewerID string) []string {
	updated := make([]string, 0, len(pr.AssignedReviewers))

	for _, r := range pr.AssignedReviewers {
		if r == oldReviewerID {
			updated = append(updated, newReviewerID)
		} else {
			updated = append(updated, r)
		}
	}

	return updated
}
