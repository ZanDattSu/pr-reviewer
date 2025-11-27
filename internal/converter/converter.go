package converter

import (
	api "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	service "github.com/ZanDattSu/pr-reviewer/internal/model"
)

// TEAM

// service => api

func ServiceTeamToAPI(t service.Team) api.Team {
	return api.Team{
		TeamName: t.TeamName,
		Members:  ServiceTeamMembersToAPI(t.Members),
	}
}

func ServiceTeamMemberToAPI(m service.TeamMember) api.TeamMember {
	return api.TeamMember{
		UserID:   m.UserID,
		Username: m.Username,
		IsActive: m.IsActive,
	}
}

func ServiceTeamMembersToAPI(ms []service.TeamMember) []api.TeamMember {
	out := make([]api.TeamMember, 0, len(ms))

	for _, m := range ms {
		out = append(out, ServiceTeamMemberToAPI(m))
	}

	return out
}

// api => service

func APIToServiceTeam(t api.Team) service.Team {
	return service.Team{
		TeamName: t.TeamName,
		Members:  APIToServiceTeamMembers(t.Members),
	}
}

func APIToServiceTeamMember(m api.TeamMember) service.TeamMember {
	return service.TeamMember{
		UserID:   m.UserID,
		Username: m.Username,
		IsActive: m.IsActive,
	}
}

func APIToServiceTeamMembers(ms []api.TeamMember) []service.TeamMember {
	out := make([]service.TeamMember, 0, len(ms))

	for _, m := range ms {
		out = append(out, APIToServiceTeamMember(m))
	}

	return out
}

// USER

func ServiceUserToAPI(u service.User) api.User {
	return api.User{
		UserID:   u.UserID,
		Username: u.Username,
		TeamName: u.TeamName,
		IsActive: u.IsActive,
	}
}

func APIToServiceUser(u api.User) service.User {
	return service.User{
		UserID:   u.UserID,
		Username: u.Username,
		TeamName: u.TeamName,
		IsActive: u.IsActive,
	}
}

// PULL REQUEST

func ServicePRToAPI(pr service.PullRequest) api.PullRequest {
	reviewers := make([]string, 0, 2)

	for _, r := range pr.AssignedReviewers {
		if r != "" {
			reviewers = append(reviewers, r)
		}
	}

	apiPR := api.PullRequest{
		PullRequestID:     pr.PullRequestID,
		PullRequestName:   pr.PullRequestName,
		AuthorID:          pr.AuthorID,
		Status:            api.PullRequestStatus(pr.Status),
		AssignedReviewers: reviewers,
	}

	if pr.CreatedAt != nil {
		apiPR.CreatedAt = api.NewOptNilDateTime(*pr.CreatedAt)
	}

	if pr.MergedAt != nil {
		apiPR.MergedAt = api.NewOptNilDateTime(*pr.MergedAt)
	}

	return apiPR
}

// USER ASSIGNED PR

// service => api

func ServiceUserAssignedPRsToUsersGetReview(
	userID string,
	prs []service.UserAssignedPR,
) api.UsersGetReviewGetOK {
	return api.UsersGetReviewGetOK{
		UserID:       userID,
		PullRequests: ServiceUserAssignedPRsToAPI(prs),
	}
}

func ServiceUserAssignedPRToAPI(pr service.UserAssignedPR) api.PullRequestShort {
	return api.PullRequestShort{
		PullRequestID:   pr.PullRequestID,
		PullRequestName: pr.PullRequestName,
		AuthorID:        pr.AuthorID,
		Status:          api.PullRequestShortStatus(pr.Status),
	}
}

func ServiceUserAssignedPRsToAPI(prs []service.UserAssignedPR) []api.PullRequestShort {
	out := make([]api.PullRequestShort, 0, len(prs))

	for _, pr := range prs {
		out = append(out, ServiceUserAssignedPRToAPI(pr))
	}

	return out
}

// api => service

func APIToServiceUserAssignedPR(pr api.PullRequestShort) service.UserAssignedPR {
	return service.UserAssignedPR{
		PullRequestID:   pr.PullRequestID,
		PullRequestName: pr.PullRequestName,
		AuthorID:        pr.AuthorID,
		Status:          service.Status(pr.Status),
	}
}

func APIToServiceUserAssignedPRs(prs []api.PullRequestShort) []service.UserAssignedPR {
	out := make([]service.UserAssignedPR, 0, len(prs))

	for _, pr := range prs {
		out = append(out, APIToServiceUserAssignedPR(pr))
	}

	return out
}

func ServiceReassignedPRsToAPI(results []service.ReassignedPR) api.UsersDeactivatePostOK {
	items := make([]api.DeactivateResultItem, 0, len(results))

	for _, result := range results {
		items = append(items, ServiceReassignedPRToAPI(result.PullRequestID, result.ReplacedBy))
	}

	return api.UsersDeactivatePostOK{
		Results: items,
	}
}

func ServiceReassignedPRToAPI(pr service.PullRequest, newReviewerID string) api.DeactivateResultItem {
	apiPR := api.DeactivateResultItemPullRequest{
		PullRequestID:     pr.PullRequestID,
		PullRequestName:   pr.PullRequestName,
		AuthorID:          pr.AuthorID,
		Status:            toDeactivateResultItemStatusFromModel(pr.Status),
		AssignedReviewers: pr.AssignedReviewers,
	}

	if pr.CreatedAt != nil {
		apiPR.SetCreatedAt(api.NewOptNilDateTime(*pr.CreatedAt))
	} else {
		apiPR.CreatedAt.SetToNull()
	}

	if pr.MergedAt != nil {
		apiPR.SetMergedAt(api.NewOptNilDateTime(*pr.MergedAt))
	} else {
		apiPR.MergedAt.SetToNull()
	}

	return api.DeactivateResultItem{
		PullRequest: apiPR,
		ReplacedBy:  newReviewerID,
	}
}

func toDeactivateResultItemStatusFromModel(status service.Status) api.DeactivateResultItemPullRequestStatus {
	switch status {
	case service.StatusOpen:
		return api.DeactivateResultItemPullRequestStatusOPEN
	case service.StatusMerged:
		return api.DeactivateResultItemPullRequestStatusMERGED
	default:
		return api.DeactivateResultItemPullRequestStatusOPEN
	}
}
