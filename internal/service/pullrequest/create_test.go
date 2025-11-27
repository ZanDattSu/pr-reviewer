package pullrequest

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (s *SuiteService) TestCreatePullRequest() {
	dbErr := errors.New("database error")

	tests := []struct {
		name              string
		pullRequestID     string
		pullRequestName   string
		authorID          string
		setupMocks        func()
		expectedPR        model.PullRequest
		expectedError     error
		expectedErrorType error
	}{
		{
			name:            "успешное создание PR с двумя ревьюверами",
			pullRequestID:   "pr-1001",
			pullRequestName: "Add search",
			authorID:        "u1",
			setupMocks: func() {
				s.prRepo.
					On("CheckPRExists", s.ctx, "pr-1001").
					Return(false, nil).
					Once()

				s.userRepo.
					On("CheckUserExists", s.ctx, "u1").
					Return(true, nil).
					Once()

				s.userRepo.
					On("GetTeamActiveMembers", s.ctx, "u1").
					Return([]string{"u2", "u3", "u4"}, nil).
					Once()

				s.prRepo.
					On("InsertPR", s.ctx, mock.MatchedBy(func(pr model.PullRequest) bool {
						return pr.PullRequestID == "pr-1001" &&
							pr.PullRequestName == "Add search" &&
							pr.AuthorID == "u1" &&
							pr.Status == model.StatusOpen &&
							len(pr.AssignedReviewers) == 2
					})).
					Return(model.PullRequest{
						PullRequestID:     "pr-1001",
						PullRequestName:   "Add search",
						AuthorID:          "u1",
						Status:            model.StatusOpen,
						AssignedReviewers: []string{"u2", "u3"},
					}, nil).
					Once()
			},
			expectedPR: model.PullRequest{
				PullRequestID:     "pr-1001",
				PullRequestName:   "Add search",
				AuthorID:          "u1",
				Status:            model.StatusOpen,
				AssignedReviewers: []string{"u2", "u3"},
			},
			expectedError: nil,
		},
		{
			name:            "успешное создание PR с одним ревьювером (доступен только один)",
			pullRequestID:   "pr-1002",
			pullRequestName: "Fix bug",
			authorID:        "u1",
			setupMocks: func() {
				s.prRepo.
					On("CheckPRExists", s.ctx, "pr-1002").
					Return(false, nil).
					Once()

				s.userRepo.
					On("CheckUserExists", s.ctx, "u1").
					Return(true, nil).
					Once()

				s.userRepo.
					On("GetTeamActiveMembers", s.ctx, "u1").
					Return([]string{"u2"}, nil).
					Once()

				s.prRepo.
					On("InsertPR", s.ctx, mock.MatchedBy(func(pr model.PullRequest) bool {
						return pr.PullRequestID == "pr-1002" &&
							len(pr.AssignedReviewers) == 1 &&
							pr.AssignedReviewers[0] == "u2"
					})).
					Return(model.PullRequest{
						PullRequestID:     "pr-1002",
						PullRequestName:   "Fix bug",
						AuthorID:          "u1",
						Status:            model.StatusOpen,
						AssignedReviewers: []string{"u2"},
					}, nil).
					Once()
			},
			expectedPR: model.PullRequest{
				PullRequestID:     "pr-1002",
				PullRequestName:   "Fix bug",
				AuthorID:          "u1",
				Status:            model.StatusOpen,
				AssignedReviewers: []string{"u2"},
			},
			expectedError: nil,
		},
		{
			name:            "успешное создание PR без ревьюверов (нет доступных членов команды)",
			pullRequestID:   "pr-1003",
			pullRequestName: "Update docs",
			authorID:        "u1",
			setupMocks: func() {
				s.prRepo.
					On("CheckPRExists", s.ctx, "pr-1003").
					Return(false, nil).
					Once()

				s.userRepo.
					On("CheckUserExists", s.ctx, "u1").
					Return(true, nil).
					Once()

				s.userRepo.
					On("GetTeamActiveMembers", s.ctx, "u1").
					Return([]string{}, nil).
					Once()

				s.prRepo.
					On("InsertPR", s.ctx, mock.MatchedBy(func(pr model.PullRequest) bool {
						return pr.PullRequestID == "pr-1003" &&
							len(pr.AssignedReviewers) == 0
					})).
					Return(model.PullRequest{
						PullRequestID:     "pr-1003",
						PullRequestName:   "Update docs",
						AuthorID:          "u1",
						Status:            model.StatusOpen,
						AssignedReviewers: []string{},
					}, nil).
					Once()
			},
			expectedPR: model.PullRequest{
				PullRequestID:     "pr-1003",
				PullRequestName:   "Update docs",
				AuthorID:          "u1",
				Status:            model.StatusOpen,
				AssignedReviewers: []string{},
			},
			expectedError: nil,
		},
		{
			name:            "ошибка: PR уже существует",
			pullRequestID:   "pr-1001",
			pullRequestName: "Duplicate PR",
			authorID:        "u1",
			setupMocks: func() {
				s.prRepo.
					On("CheckPRExists", s.ctx, "pr-1001").
					Return(true, nil).
					Once()
			},
			expectedPR:        model.PullRequest{},
			expectedErrorType: apperror.NewPRExistsError("pr-1001"),
		},
		{
			name:            "ошибка: автор не найден",
			pullRequestID:   "pr-1004",
			pullRequestName: "New feature",
			authorID:        "non-existent-user",
			setupMocks: func() {
				s.prRepo.
					On("CheckPRExists", s.ctx, "pr-1004").
					Return(false, nil).
					Once()

				s.userRepo.
					On("CheckUserExists", s.ctx, "non-existent-user").
					Return(false, nil).
					Once()
			},
			expectedPR:        model.PullRequest{},
			expectedErrorType: apperror.NewUserNotFoundError("non-existent-user"),
		},
		{
			name:            "ошибка при проверке существования PR",
			pullRequestID:   "pr-1005",
			pullRequestName: "Test PR",
			authorID:        "u1",
			setupMocks: func() {
				s.prRepo.
					On("CheckPRExists", s.ctx, "pr-1005").
					Return(false, dbErr).
					Once()
			},
			expectedPR:    model.PullRequest{},
			expectedError: dbErr,
		},
		{
			name:            "ошибка при проверке существования пользователя",
			pullRequestID:   "pr-1006",
			pullRequestName: "Test PR",
			authorID:        "u1",
			setupMocks: func() {
				s.prRepo.
					On("CheckPRExists", s.ctx, "pr-1006").
					Return(false, nil).
					Once()

				s.userRepo.
					On("CheckUserExists", s.ctx, "u1").
					Return(false, dbErr).
					Once()
			},
			expectedPR:    model.PullRequest{},
			expectedError: dbErr,
		},
		{
			name:            "ошибка при получении членов команды",
			pullRequestID:   "pr-1007",
			pullRequestName: "Test PR",
			authorID:        "u1",
			setupMocks: func() {
				s.prRepo.
					On("CheckPRExists", s.ctx, "pr-1007").
					Return(false, nil).
					Once()

				s.userRepo.
					On("CheckUserExists", s.ctx, "u1").
					Return(true, nil).
					Once()

				s.userRepo.
					On("GetTeamActiveMembers", s.ctx, "u1").
					Return(nil, dbErr).
					Once()
			},
			expectedPR:    model.PullRequest{},
			expectedError: dbErr,
		},
		{
			name:            "ошибка при вставке PR в базу",
			pullRequestID:   "pr-1008",
			pullRequestName: "Test PR",
			authorID:        "u1",
			setupMocks: func() {
				s.prRepo.
					On("CheckPRExists", s.ctx, "pr-1008").
					Return(false, nil).
					Once()

				s.userRepo.
					On("CheckUserExists", s.ctx, "u1").
					Return(true, nil).
					Once()

				s.userRepo.
					On("GetTeamActiveMembers", s.ctx, "u1").
					Return([]string{"u2", "u3"}, nil).
					Once()

				s.prRepo.
					On("InsertPR", s.ctx, mock.Anything).
					Return(model.PullRequest{}, dbErr).
					Once()
			},
			expectedPR:    model.PullRequest{},
			expectedError: dbErr,
		},
		{
			name:            "пустой pullRequestID",
			pullRequestID:   "",
			pullRequestName: "Test PR",
			authorID:        "u1",
			setupMocks: func() {
				s.prRepo.
					On("CheckPRExists", s.ctx, "").
					Return(false, nil).
					Once()

				s.userRepo.
					On("CheckUserExists", s.ctx, "u1").
					Return(true, nil).
					Once()

				s.userRepo.
					On("GetTeamActiveMembers", s.ctx, "u1").
					Return([]string{"u2"}, nil).
					Once()

				s.prRepo.
					On("InsertPR", s.ctx, mock.MatchedBy(func(pr model.PullRequest) bool {
						return pr.PullRequestID == ""
					})).
					Return(model.PullRequest{
						PullRequestID:     "",
						PullRequestName:   "Test PR",
						AuthorID:          "u1",
						Status:            model.StatusOpen,
						AssignedReviewers: []string{"u2"},
					}, nil).
					Once()
			},
			expectedPR: model.PullRequest{
				PullRequestID:     "",
				PullRequestName:   "Test PR",
				AuthorID:          "u1",
				Status:            model.StatusOpen,
				AssignedReviewers: []string{"u2"},
			},
			expectedError: nil,
		},
		{
			name:            "пустой pullRequestName",
			pullRequestID:   "pr-2001",
			pullRequestName: "",
			authorID:        "u1",
			setupMocks: func() {
				s.prRepo.
					On("CheckPRExists", s.ctx, "pr-2001").
					Return(false, nil).
					Once()

				s.userRepo.
					On("CheckUserExists", s.ctx, "u1").
					Return(true, nil).
					Once()

				s.userRepo.
					On("GetTeamActiveMembers", s.ctx, "u1").
					Return([]string{"u2"}, nil).
					Once()

				s.prRepo.
					On("InsertPR", s.ctx, mock.MatchedBy(func(pr model.PullRequest) bool {
						return pr.PullRequestName == ""
					})).
					Return(model.PullRequest{
						PullRequestID:     "pr-2001",
						PullRequestName:   "",
						AuthorID:          "u1",
						Status:            model.StatusOpen,
						AssignedReviewers: []string{"u2"},
					}, nil).
					Once()
			},
			expectedPR: model.PullRequest{
				PullRequestID:     "pr-2001",
				PullRequestName:   "",
				AuthorID:          "u1",
				Status:            model.StatusOpen,
				AssignedReviewers: []string{"u2"},
			},
			expectedError: nil,
		},
		{
			name:            "пустой authorID",
			pullRequestID:   "pr-2002",
			pullRequestName: "Test PR",
			authorID:        "",
			setupMocks: func() {
				s.prRepo.
					On("CheckPRExists", s.ctx, "pr-2002").
					Return(false, nil).
					Once()

				s.userRepo.
					On("CheckUserExists", s.ctx, "").
					Return(false, nil).
					Once()
			},
			expectedPR:        model.PullRequest{},
			expectedErrorType: apperror.NewUserNotFoundError(""),
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.setupMocks()

			actualPR, err := s.service.CreatePullRequest(
				s.ctx,
				tt.pullRequestID,
				tt.pullRequestName,
				tt.authorID,
			)

			switch {
			case tt.expectedError != nil:
				s.Error(err)
				s.Equal(tt.expectedError.Error(), err.Error())

			case tt.expectedErrorType != nil:
				s.Error(err)
				s.IsType(tt.expectedErrorType, err)

			default:
				s.NoError(err)
				s.Equal(tt.expectedPR.PullRequestID, actualPR.PullRequestID)
				s.Equal(tt.expectedPR.PullRequestName, actualPR.PullRequestName)
				s.Equal(tt.expectedPR.AuthorID, actualPR.AuthorID)
				s.Equal(tt.expectedPR.Status, actualPR.Status)
				s.Equal(len(tt.expectedPR.AssignedReviewers), len(actualPR.AssignedReviewers))
			}

			s.prRepo.AssertExpectations(s.T())
			s.userRepo.AssertExpectations(s.T())
		})
	}
}
