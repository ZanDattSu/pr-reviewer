package user

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (s *SuiteService) TestGetUserStats() {
	stats := []model.UserStats{
		{UserID: "u1", TotalPR: 10},
		{UserID: "u2", TotalPR: 5},
	}

	tests := []struct {
		name        string
		top         int
		onlyActive  bool
		onlyOpen    bool
		setupMocks  func()
		want        []model.UserStats
		wantErr     bool
		errContains string
	}{
		{
			name:       "успех — без фильтров",
			top:        0,
			onlyActive: false,
			onlyOpen:   false,
			setupMocks: func() {
				s.userRepo.
					On("GetUserStats", s.ctx, 0, false, false).
					Return(stats, nil).Once()
			},
			want: stats,
		},

		{
			name:       "успех — with top limit",
			top:        1,
			onlyActive: false,
			onlyOpen:   false,
			setupMocks: func() {
				s.userRepo.
					On("GetUserStats", s.ctx, 1, false, false).
					Return([]model.UserStats{
						{UserID: "u1", TotalPR: 10},
					}, nil)
			},
			want: []model.UserStats{
				{UserID: "u1", TotalPR: 10},
			},
		},

		{
			name:       "успех — фильтр onlyActive",
			top:        0,
			onlyActive: true,
			onlyOpen:   false,
			setupMocks: func() {
				s.userRepo.
					On("GetUserStats", s.ctx, 0, true, false).
					Return(stats, nil).Once()
			},
			want: stats,
		},

		{
			name:       "успех — фильтр onlyOpen",
			top:        0,
			onlyActive: false,
			onlyOpen:   true,
			setupMocks: func() {
				s.userRepo.
					On("GetUserStats", s.ctx, 0, false, true).
					Return(stats, nil).Once()
			},
			want: stats,
		},

		{
			name:       "ошибка — NoDataError",
			top:        0,
			onlyActive: false,
			onlyOpen:   false,
			setupMocks: func() {
				s.userRepo.
					On("GetUserStats", mock.Anything, 0, false, false).
					Return([]model.UserStats(nil), apperror.NewNoDataError()).Once()
			},
			wantErr:     true,
			errContains: "no data",
		},

		{
			name:       "ошибка — репозиторий упал",
			top:        10,
			onlyActive: false,
			onlyOpen:   false,
			setupMocks: func() {
				s.userRepo.
					On("GetUserStats", s.ctx, 10, false, false).
					Return(nil, errors.New("db failure")).Once()
			},
			wantErr:     true,
			errContains: "db failure",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.setupMocks()

			got, err := s.service.GetUserStats(
				s.ctx,
				tt.top,
				tt.onlyActive,
				tt.onlyOpen,
			)

			if tt.wantErr {
				s.Require().Error(err)
				if tt.errContains != "" {
					s.Contains(err.Error(), tt.errContains)
				}
				return
			}

			s.Require().NoError(err)
			s.Equal(tt.want, got)
		})
	}
}
