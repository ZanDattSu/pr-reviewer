package pullrequest

import (
	"context"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (s *prService) CreatePullRequest(
	ctx context.Context,
	pullRequestID,
	pullRequestName,
	authorID string,
) (model.PullRequest, error) {
	prExists, err := s.prRepo.CheckPRExists(ctx, pullRequestID)
	if err != nil {
		return model.PullRequest{}, err
	}

	if prExists {
		return model.PullRequest{}, apperror.NewPRExistsError(pullRequestID)
	}

	userExists, err := s.userRepo.CheckUserExists(ctx, authorID)
	if err != nil {
		return model.PullRequest{}, err
	}

	if !userExists {
		return model.PullRequest{}, apperror.NewUserNotFoundError(authorID)
	}

	teamActiveMembers, err := s.teamRepo.GetTeamActiveMembersWithoutUser(ctx, authorID)
	if err != nil {
		return model.PullRequest{}, err
	}

	prReviewers := PickReviewers(teamActiveMembers, 2)

	pr := model.PullRequest{
		PullRequestID:     pullRequestID,
		PullRequestName:   pullRequestName,
		AuthorID:          authorID,
		Status:            model.StatusOpen,
		AssignedReviewers: prReviewers,
	}

	createdPR, err := s.prRepo.InsertPR(ctx, pr)
	if err != nil {
		return model.PullRequest{}, err
	}

	return createdPR, err
}
