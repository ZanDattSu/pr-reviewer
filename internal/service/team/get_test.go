package team

import (
	"errors"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (s *SuiteService) TestGetTeam() {
	tests := []struct {
		name string

		teamName string

		setupMocks func()

		expectedTeam model.Team

		expectedErr error

		expectedErrTyp error
	}{
		{
			name: "успех: команда найдена, есть участники",

			teamName: "backend",

			setupMocks: func() {
				s.teamRepo.
					On("GetTeam", s.ctx, "backend").
					Return(model.Team{
						TeamName: "backend",

						Members: []model.TeamMember{
							{UserID: "u1", Username: "alice", IsActive: true},

							{UserID: "u2", Username: "bob", IsActive: false},
						},
					}, nil).
					Once()
			},

			expectedTeam: model.Team{
				TeamName: "backend",

				Members: []model.TeamMember{
					{UserID: "u1", Username: "alice", IsActive: true},

					{UserID: "u2", Username: "bob", IsActive: false},
				},
			},
		},

		{
			name: "успех: команда найдена, но участников нет",

			teamName: "empty",

			setupMocks: func() {
				s.teamRepo.
					On("GetTeam", s.ctx, "empty").
					Return(model.Team{
						TeamName: "empty",

						Members: []model.TeamMember{},
					}, nil).
					Once()
			},

			expectedTeam: model.Team{
				TeamName: "empty",

				Members: []model.TeamMember{},
			},
		},

		{
			name: "ошибка: команда не найдена",

			teamName: "ghost",

			setupMocks: func() {
				s.teamRepo.
					On("GetTeam", s.ctx, "ghost").
					Return(model.Team{}, apperror.NewTeamNotFoundError("ghost")).
					Once()
			},

			expectedErrTyp: apperror.NewTeamNotFoundError("ghost"),
		},

		{
			name: "ошибка репозитория (DB error)",

			teamName: "backend",

			setupMocks: func() {
				s.teamRepo.
					On("GetTeam", s.ctx, "backend").
					Return(model.Team{}, errors.New("db timeout")).
					Once()
			},

			expectedErr: errors.New("db timeout"),
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.setupMocks()

			team, err := s.service.GetTeam(s.ctx, tt.teamName)

			switch {
			case tt.expectedErr != nil:
				s.Error(err)
				s.Equal(tt.expectedErr.Error(), err.Error())

			case tt.expectedErrTyp != nil:
				s.Error(err)
				s.IsType(tt.expectedErrTyp, err)

			default:
				s.NoError(err)
				s.Equal(tt.expectedTeam, team)
			}

			s.teamRepo.AssertExpectations(s.T())
		})
	}
}
