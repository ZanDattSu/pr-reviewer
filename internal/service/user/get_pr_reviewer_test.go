package user

import (
	"errors"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (s *SuiteService) TestGetPRReviewer() {
	tests := []struct {
		name           string
		userID         string
		setupMocks     func()
		expectedPRs    []model.UserAssignedPR
		expectedErr    error
		expectedErrTyp error
	}{
		{
			name:   "успех — пользователь существует, есть PR'ы",
			userID: "u2",
			setupMocks: func() {
				s.userRepo.
					On("CheckUserExists", s.ctx, "u2").
					Return(true, nil).
					Once()

				s.userRepo.
					On("GetPRReviewer", s.ctx, "u2").
					Return([]model.UserAssignedPR{
						{
							PullRequestID:   "pr-1001",
							PullRequestName: "Add search",
							AuthorID:        "u1",
							Status:          model.StatusOpen,
						},
					}, nil).
					Once()
			},
			expectedPRs: []model.UserAssignedPR{
				{
					PullRequestID:   "pr-1001",
					PullRequestName: "Add search",
					AuthorID:        "u1",
					Status:          model.StatusOpen,
				},
			},
		},

		{
			name:   "успех — пользователь существует, но PR'ов нет",
			userID: "u5",
			setupMocks: func() {
				s.userRepo.
					On("CheckUserExists", s.ctx, "u5").
					Return(true, nil).
					Once()

				s.userRepo.
					On("GetPRReviewer", s.ctx, "u5").
					Return([]model.UserAssignedPR{}, nil).
					Once()
			},
			expectedPRs: []model.UserAssignedPR{},
		},

		{
			name:   "ошибка — пользователь не найден",
			userID: "ghost",
			setupMocks: func() {
				s.userRepo.
					On("CheckUserExists", s.ctx, "ghost").
					Return(false, nil).
					Once()
			},
			expectedErrTyp: apperror.NewUserNotFoundError("ghost"),
		},

		{
			name:   "ошибка — ошибка в CheckUserExists (DB error)",
			userID: "u1",
			setupMocks: func() {
				s.userRepo.
					On("CheckUserExists", s.ctx, "u1").
					Return(false, errors.New("db failure")).
					Once()
			},
			expectedErr: errors.New("db failure"),
		},

		{
			name:   "ошибка — DB error во время GetPRReviewer",
			userID: "u2",
			setupMocks: func() {
				s.userRepo.
					On("CheckUserExists", s.ctx, "u2").
					Return(true, nil).
					Once()

				s.userRepo.
					On("GetPRReviewer", s.ctx, "u2").
					Return(nil, errors.New("timeout")).
					Once()
			},
			expectedErr: errors.New("timeout"),
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.setupMocks()

			prList, err := s.service.UserGetPRReviewer(s.ctx, tt.userID)

			switch {
			case tt.expectedErr != nil:
				s.Error(err)
				s.Equal(tt.expectedErr.Error(), err.Error())

			case tt.expectedErrTyp != nil:
				s.Error(err)
				s.IsType(tt.expectedErrTyp, err)

			default:
				s.NoError(err)
				s.Equal(tt.expectedPRs, prList)
			}

			s.userRepo.AssertExpectations(s.T())
		})
	}
}
