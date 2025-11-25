package user

import (
	"errors"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (s *SuiteService) TestUpdateUserStatus() {
	tests := []struct {
		name           string
		userID         string
		isActive       bool
		setupMocks     func()
		expectedUser   model.User
		expectedErr    error
		expectedErrTyp error
	}{
		{
			name:     "успех — обновление статуса пользователя",
			userID:   "u1",
			isActive: false,
			setupMocks: func() {
				s.userRepo.
					On("UpdateUserStatus", s.ctx, "u1", false).
					Return(model.User{
						UserID:   "u1",
						Username: "alice",
						TeamName: "backend",
						IsActive: false,
					}, nil).
					Once()
			},
			expectedUser: model.User{
				UserID:   "u1",
				Username: "alice",
				TeamName: "backend",
				IsActive: false,
			},
		},

		{
			name:     "успех — включение активности",
			userID:   "u2",
			isActive: true,
			setupMocks: func() {
				s.userRepo.
					On("UpdateUserStatus", s.ctx, "u2", true).
					Return(model.User{
						UserID:   "u2",
						Username: "bob",
						TeamName: "mobile",
						IsActive: true,
					}, nil).
					Once()
			},
			expectedUser: model.User{
				UserID:   "u2",
				Username: "bob",
				TeamName: "mobile",
				IsActive: true,
			},
		},

		{
			name:     "ошибка — пользователь не найден",
			userID:   "ghost",
			isActive: false,
			setupMocks: func() {
				s.userRepo.
					On("UpdateUserStatus", s.ctx, "ghost", false).
					Return(model.User{}, apperror.NewUserNotFoundError("ghost")).
					Once()
			},
			expectedErrTyp: apperror.NewUserNotFoundError("ghost"),
		},

		{
			name:     "ошибка базы данных",
			userID:   "u1",
			isActive: true,
			setupMocks: func() {
				s.userRepo.
					On("UpdateUserStatus", s.ctx, "u1", true).
					Return(model.User{}, errors.New("db timeout")).
					Once()
			},
			expectedErr: errors.New("db timeout"),
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.setupMocks()

			actualUser, err := s.service.UpdateUserStatus(s.ctx, tt.userID, tt.isActive)

			//nolint:gocritic
			if tt.expectedErr != nil {
				s.Error(err)
				s.Equal(tt.expectedErr.Error(), err.Error())
			} else if tt.expectedErrTyp != nil {
				s.Error(err)
				s.IsType(tt.expectedErrTyp, err)
			} else {
				s.NoError(err)
				s.Equal(tt.expectedUser, actualUser)
			}

			s.userRepo.AssertExpectations(s.T())
		})
	}
}
