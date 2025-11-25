package pullrequest

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (s *SuiteService) TestReassignPullRequest_RandomCandidate() {
	s.ctx = context.Background()

	s.prRepo.
		On("GetPRWithReviewers", s.ctx, "pr-1001").
		Return(model.PullRequest{
			PullRequestID:     "pr-1001",
			PullRequestName:   "Add feature",
			AuthorID:          "u1",
			Status:            model.StatusOpen,
			AssignedReviewers: []string{"u2", "u3"},
		}, nil).
		Once()

	s.teamRepo.
		On("GetTeamActiveMembersWithoutUser", s.ctx, "u2").
		Return([]string{"u5", "u6"}, nil).
		Once()

	s.reviewerRepo.
		On("ReplaceReviewer", s.ctx, "pr-1001", "u2", mock.AnythingOfType("string")).
		Return(nil).
		Once()

	pr, newID, err := s.service.ReassignPullRequest(s.ctx, "pr-1001", "u2")

	s.Require().NoError(err)

	s.NotContains(pr.AssignedReviewers, "u2")

	s.Contains([]string{"u5", "u6"}, newID)

	s.Contains(pr.AssignedReviewers, newID)

	s.prRepo.AssertExpectations(s.T())
	s.teamRepo.AssertExpectations(s.T())
	s.reviewerRepo.AssertExpectations(s.T())
}

func (s *SuiteService) TestReassignPullRequest() {
	dbErr := errors.New("database error")

	tests := []struct {
		name          string
		prID          string
		oldReviewerID string
		setupMocks    func()
		expectedPR    model.PullRequest
		expectedNew   string
		expectedErr   error
		expectedType  error
	}{
		{
			name:          "успех — PR с 1 ревьювером, заменяем единственного",
			prID:          "pr-2001",
			oldReviewerID: "u10",
			setupMocks: func() {
				s.prRepo.
					On("GetPRWithReviewers", s.ctx, "pr-2001").
					Return(model.PullRequest{
						PullRequestID:     "pr-2001",
						PullRequestName:   "Fix crash",
						AuthorID:          "u5",
						Status:            model.StatusOpen,
						AssignedReviewers: []string{"u10"},
					}, nil).
					Once()

				s.teamRepo.
					On("GetTeamActiveMembersWithoutUser", s.ctx, "u10").
					Return([]string{"u12"}, nil).
					Once()

				s.reviewerRepo.
					On("ReplaceReviewer", s.ctx, "pr-2001", "u10", "u12").
					Return(nil).
					Once()
			},
			expectedNew: "u12",
			expectedPR: model.PullRequest{
				PullRequestID:     "pr-2001",
				PullRequestName:   "Fix crash",
				AuthorID:          "u5",
				Status:            model.StatusOpen,
				AssignedReviewers: []string{"u12"},
			},
		},

		{
			name:          "ошибка — PR не найден",
			prID:          "not-found",
			oldReviewerID: "u1",
			setupMocks: func() {
				s.prRepo.
					On("GetPRWithReviewers", s.ctx, "not-found").
					Return(model.PullRequest{}, apperror.NewPRNotFoundError("not-found")).
					Once()
			},
			expectedType: apperror.NewPRNotFoundError("not-found"),
		},

		{
			name:          "ошибка — PR уже MERGED",
			prID:          "pr-3001",
			oldReviewerID: "u4",
			setupMocks: func() {
				s.prRepo.
					On("GetPRWithReviewers", s.ctx, "pr-3001").
					Return(model.PullRequest{
						PullRequestID:     "pr-3001",
						Status:            model.StatusMerged,
						AssignedReviewers: []string{"u4", "u5"},
					}, nil).
					Once()
			},
			expectedType: apperror.NewPRMergedError("pr-3001"),
		},

		{
			name:          "ошибка — oldReviewerID не назначен",
			prID:          "pr-4001",
			oldReviewerID: "u99",
			setupMocks: func() {
				s.prRepo.
					On("GetPRWithReviewers", s.ctx, "pr-4001").
					Return(model.PullRequest{
						PullRequestID:     "pr-4001",
						Status:            model.StatusOpen,
						AssignedReviewers: []string{"u1", "u2"},
					}, nil).
					Once()
			},
			expectedType: apperror.NewNotAssignedError("u99"),
		},

		{
			name:          "ошибка — нет активных кандидатов в команде",
			prID:          "pr-5001",
			oldReviewerID: "u2",
			setupMocks: func() {
				s.prRepo.
					On("GetPRWithReviewers", s.ctx, "pr-5001").
					Return(model.PullRequest{
						PullRequestID:     "pr-5001",
						Status:            model.StatusOpen,
						AssignedReviewers: []string{"u2"},
					}, nil).
					Once()

				s.teamRepo.
					On("GetTeamActiveMembersWithoutUser", s.ctx, "u2").
					Return([]string{}, nil).
					Once()
			},
			expectedType: apperror.NewNoCandidateError("pr-5001"),
		},

		{
			name:          "ошибка — ReplaceReviewer возвращает ошибку базы",
			prID:          "pr-6001",
			oldReviewerID: "u5",
			setupMocks: func() {
				s.prRepo.
					On("GetPRWithReviewers", s.ctx, "pr-6001").
					Return(model.PullRequest{
						PullRequestID:     "pr-6001",
						Status:            model.StatusOpen,
						AssignedReviewers: []string{"u5", "u8"},
					}, nil).
					Once()

				s.teamRepo.
					On("GetTeamActiveMembersWithoutUser", s.ctx, "u5").
					Return([]string{"u7"}, nil).
					Once()

				s.reviewerRepo.
					On("ReplaceReviewer", s.ctx, "pr-6001", "u5", "u7").
					Return(dbErr).
					Once()
			},
			expectedErr: dbErr,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.setupMocks()

			pr, newID, err := s.service.ReassignPullRequest(s.ctx, tt.prID, tt.oldReviewerID)

			switch {
			case tt.expectedErr != nil:
				s.Error(err)
				s.Equal(tt.expectedErr.Error(), err.Error())

			case tt.expectedType != nil:
				s.Error(err)
				s.IsType(tt.expectedType, err)

			default:
				s.NoError(err)
				s.Equal(tt.expectedNew, newID)
				s.Equal(tt.expectedPR.AssignedReviewers, pr.AssignedReviewers)
			}

			s.prRepo.AssertExpectations(s.T())
			s.teamRepo.AssertExpectations(s.T())
			s.reviewerRepo.AssertExpectations(s.T())
		})
	}
}
