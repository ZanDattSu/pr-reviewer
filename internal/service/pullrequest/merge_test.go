package pullrequest

import (
	"errors"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/mocks"
)

func (s *SuiteService) TestMergePullRequest() {
	dbErr := errors.New("database error")

	tests := []struct {
		name              string
		pullRequestID     string
		setupMocks        func()
		expectedPR        model.PullRequest
		expectedError     error
		expectedErrorType error
	}{
		{
			name:          "успешный merge — первый вызов, устанавливаем MERGED и mergedAt",
			pullRequestID: "pr-1001",
			setupMocks: func() {
				s.prRepo.(*mocks.PullRequestRepository).
					On("GetPRWithReviewers", s.ctx, "pr-1001").
					Return(model.PullRequest{
						PullRequestID:     "pr-1001",
						PullRequestName:   "Add search",
						AuthorID:          "u1",
						Status:            model.StatusOpen,
						AssignedReviewers: []string{"u2", "u3"},
					}, nil).
					Once()

				s.prRepo.(*mocks.PullRequestRepository).
					On("UpdatePRStatus", s.ctx, mock.MatchedBy(func(pr model.PullRequest) bool {
						return pr.PullRequestID == "pr-1001"
					}), model.StatusMerged).
					Return(model.PullRequest{
						PullRequestID:     "pr-1001",
						PullRequestName:   "Add search",
						AuthorID:          "u1",
						Status:            model.StatusMerged,
						AssignedReviewers: []string{"u2", "u3"},
						MergedAt:          lo.ToPtr(time.Now()),
					}, nil).
					Once()
			},
			expectedPR: model.PullRequest{
				PullRequestID:   "pr-1001",
				PullRequestName: "Add search",
				AuthorID:        "u1",
				Status:          model.StatusMerged,
			},
			expectedError: nil,
		},

		{
			name:          "повторный merge — операция идемпотентна, mergedAt НЕ меняется",
			pullRequestID: "pr-1002",
			setupMocks: func() {
				oldMergeTime := time.Now().Add(-time.Hour)

				s.prRepo.(*mocks.PullRequestRepository).
					On("GetPRWithReviewers", s.ctx, "pr-1002").
					Return(model.PullRequest{
						PullRequestID:     "pr-1002",
						PullRequestName:   "Fix bug",
						AuthorID:          "u2",
						Status:            model.StatusMerged,
						AssignedReviewers: []string{"u1"},
						MergedAt:          lo.ToPtr(oldMergeTime),
					}, nil).
					Once()

				// Update должен вернуть тот же mergedAt из-за COALESCE(merged_at, NOW())
				s.prRepo.(*mocks.PullRequestRepository).
					On("UpdatePRStatus", s.ctx, mock.Anything, model.StatusMerged).
					Return(model.PullRequest{
						PullRequestID:     "pr-1002",
						PullRequestName:   "Fix bug",
						AuthorID:          "u2",
						Status:            model.StatusMerged,
						AssignedReviewers: []string{"u1"},
						MergedAt:          lo.ToPtr(oldMergeTime),
					}, nil).
					Once()
			},
			expectedPR: model.PullRequest{
				PullRequestID:   "pr-1002",
				PullRequestName: "Fix bug",
				AuthorID:        "u2",
				Status:          model.StatusMerged,
			},
			expectedError: nil,
		},

		{
			name:          "ошибка: PR не найден",
			pullRequestID: "pr-9999",
			setupMocks: func() {
				s.prRepo.(*mocks.PullRequestRepository).
					On("GetPRWithReviewers", s.ctx, "pr-9999").
					Return(model.PullRequest{}, apperror.NewPRNotFoundError("pr-9999")).
					Once()
			},
			expectedPR:        model.PullRequest{},
			expectedErrorType: apperror.NewPRNotFoundError("pr-9999"),
		},

		{
			name:          "ошибка в GetPRWithReviewers — внутренняя ошибка",
			pullRequestID: "pr-1003",
			setupMocks: func() {
				s.prRepo.(*mocks.PullRequestRepository).
					On("GetPRWithReviewers", s.ctx, "pr-1003").
					Return(model.PullRequest{}, dbErr).
					Once()
			},
			expectedPR:    model.PullRequest{},
			expectedError: dbErr,
		},

		{
			name:          "ошибка в UpdatePRStatus — например, потеря соединения",
			pullRequestID: "pr-1004",
			setupMocks: func() {
				s.prRepo.(*mocks.PullRequestRepository).
					On("GetPRWithReviewers", s.ctx, "pr-1004").
					Return(model.PullRequest{
						PullRequestID: "pr-1004",
						Status:        model.StatusOpen,
					}, nil).
					Once()

				s.prRepo.(*mocks.PullRequestRepository).
					On("UpdatePRStatus", s.ctx, mock.Anything, model.StatusMerged).
					Return(model.PullRequest{}, dbErr).
					Once()
			},
			expectedPR:    model.PullRequest{},
			expectedError: dbErr,
		},

		{
			name:          "merge работает корректно, даже если у PR нет reviewers",
			pullRequestID: "pr-1005",
			setupMocks: func() {
				s.prRepo.(*mocks.PullRequestRepository).
					On("GetPRWithReviewers", s.ctx, "pr-1005").
					Return(model.PullRequest{
						PullRequestID:     "pr-1005",
						PullRequestName:   "Empty reviewers",
						AuthorID:          "u3",
						Status:            model.StatusOpen,
						AssignedReviewers: []string{},
					}, nil).
					Once()

				s.prRepo.(*mocks.PullRequestRepository).
					On("UpdatePRStatus", s.ctx, mock.Anything, model.StatusMerged).
					Return(model.PullRequest{
						PullRequestID:     "pr-1005",
						PullRequestName:   "Empty reviewers",
						AuthorID:          "u3",
						Status:            model.StatusMerged,
						AssignedReviewers: []string{},
						MergedAt:          lo.ToPtr(time.Now()),
					}, nil).
					Once()
			},
			expectedPR: model.PullRequest{
				PullRequestID: "pr-1005",
				Status:        model.StatusMerged,
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.setupMocks()

			actualPR, err := s.service.MergePullRequest(s.ctx, tt.pullRequestID)

			//nolint:gocritic
			if tt.expectedError != nil {
				s.Error(err)
				s.Equal(tt.expectedError.Error(), err.Error())
			} else if tt.expectedErrorType != nil {
				s.Error(err)
				s.IsType(tt.expectedErrorType, err)
			} else {
				s.NoError(err)
				s.Equal(tt.expectedPR.Status, actualPR.Status)
				s.Equal(tt.expectedPR.PullRequestID, actualPR.PullRequestID)
			}

			s.prRepo.(*mocks.PullRequestRepository).AssertExpectations(s.T())
		})
	}
}
