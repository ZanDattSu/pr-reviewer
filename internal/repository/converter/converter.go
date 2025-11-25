package converter

import (
	"github.com/ZanDattSu/pr-reviewer/internal/model"
	repoModel "github.com/ZanDattSu/pr-reviewer/internal/repository/model"
)

// TEAM

// repo => service

func RepoTeamToService(team repoModel.Team) model.Team {
	return model.Team{
		TeamName: team.TeamName,
		Members:  RepoTeamMembersToService(team.Members),
	}
}

func RepoTeamMemberToService(m repoModel.TeamMember) model.TeamMember {
	return model.TeamMember{
		UserID:   m.UserID,
		Username: m.Username,
		IsActive: m.IsActive,
	}
}

func RepoTeamMembersToService(ms []repoModel.TeamMember) []model.TeamMember {
	out := make([]model.TeamMember, 0, len(ms))

	for _, m := range ms {
		out = append(out, RepoTeamMemberToService(m))
	}

	return out
}

// service => repo

func ServiceTeamToRepo(team model.Team) repoModel.Team {
	return repoModel.Team{
		TeamName: team.TeamName,
		Members:  ServiceTeamMembersToRepo(team.Members),
	}
}

func ServiceTeamMemberToRepo(m model.TeamMember) repoModel.TeamMember {
	return repoModel.TeamMember{
		UserID:   m.UserID,
		Username: m.Username,
		IsActive: m.IsActive,
	}
}

func ServiceTeamMembersToRepo(ms []model.TeamMember) []repoModel.TeamMember {
	out := make([]repoModel.TeamMember, 0, len(ms))

	for _, m := range ms {
		out = append(out, ServiceTeamMemberToRepo(m))
	}

	return out
}

// USER

func RepoUserToService(u repoModel.User) model.User {
	return model.User{
		UserID:   u.UserID,
		Username: u.Username,
		TeamName: u.TeamName,
		IsActive: u.IsActive,
	}
}

func ServiceUserToRepo(u model.User) repoModel.User {
	return repoModel.User{
		UserID:   u.UserID,
		Username: u.Username,
		TeamName: u.TeamName,
		IsActive: u.IsActive,
	}
}

// PULL REQUEST

func RepoPRToService(pr repoModel.PullRequest) model.PullRequest {
	return model.PullRequest{
		PullRequestID:     pr.PullRequestID,
		PullRequestName:   pr.PullRequestName,
		AuthorID:          pr.AuthorID,
		Status:            RepoPRStatusToService(pr.Status),
		AssignedReviewers: pr.AssignedReviewers,
		CreatedAt:         pr.CreatedAt,
		MergedAt:          pr.MergedAt,
	}
}

func RepoPRStatusToService(s repoModel.Status) model.Status {
	switch s {

	case repoModel.StatusOpen:

		return model.StatusOpen
	case repoModel.StatusMerged:

		return model.StatusMerged
	default:

		return model.StatusUnknown

	}
}

func ServicePRToRepo(pr model.PullRequest) repoModel.PullRequest {
	return repoModel.PullRequest{
		PullRequestID:     pr.PullRequestID,
		PullRequestName:   pr.PullRequestName,
		AuthorID:          pr.AuthorID,
		Status:            ServicePRStatusToRepo(pr.Status),
		AssignedReviewers: pr.AssignedReviewers,
		CreatedAt:         pr.CreatedAt,
		MergedAt:          pr.MergedAt,
	}
}

func ServicePRStatusToRepo(s model.Status) repoModel.Status {
	switch s {

	case model.StatusOpen:

		return repoModel.StatusOpen
	case model.StatusMerged:

		return repoModel.StatusMerged
	default:

		return repoModel.StatusUnknown

	}
}

// USER ASSIGNED PR

// repo => service

func RepoUserAssignedPRToService(pr repoModel.UserAssignedPR) model.UserAssignedPR {
	return model.UserAssignedPR{
		PullRequestID:   pr.PullRequestID,
		PullRequestName: pr.PullRequestName,
		AuthorID:        pr.AuthorID,
		Status:          RepoPRStatusToService(pr.Status),
	}
}

func RepoUserAssignedPRsToService(prs []repoModel.UserAssignedPR) []model.UserAssignedPR {
	out := make([]model.UserAssignedPR, 0, len(prs))

	for _, pr := range prs {
		out = append(out, RepoUserAssignedPRToService(pr))
	}

	return out
}

// service => repo

func ServiceUserAssignedPRToRepo(pr model.UserAssignedPR) repoModel.UserAssignedPR {
	return repoModel.UserAssignedPR{
		PullRequestID:   pr.PullRequestID,
		PullRequestName: pr.PullRequestName,
		AuthorID:        pr.AuthorID,
		Status:          ServicePRStatusToRepo(pr.Status),
	}
}

func ServiceUserAssignedPRsRepo(prs []model.UserAssignedPR) []repoModel.UserAssignedPR {
	out := make([]repoModel.UserAssignedPR, 0, len(prs))

	for _, pr := range prs {
		out = append(out, ServiceUserAssignedPRToRepo(pr))
	}

	return out
}
