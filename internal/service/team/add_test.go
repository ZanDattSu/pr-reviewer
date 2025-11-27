package team

import (
	"context"
	"errors"

	"github.com/stretchr/testify/mock"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (s *SuiteService) mockTM() {
	s.trm.
		On("Do", mock.Anything, mock.AnythingOfType("func(context.Context) error")).
		Return(func(ctx context.Context, fn func(context.Context) error) error {
			return fn(ctx)
		})
}

func (s *SuiteService) TestAddTeam() {
	tests := []struct {
		name           string
		inputTeam      model.Team
		setupMocks     func()
		expectedTeam   model.Team
		expectedErr    error
		expectedErrTyp error
	}{
		{
			name: "успешное добавление команды с двумя участниками",
			inputTeam: model.Team{
				TeamName: "backend",
				Members: []model.TeamMember{
					{UserID: "u1", Username: "alice", IsActive: true},
					{UserID: "u2", Username: "bob", IsActive: true},
				},
			},
			setupMocks: func() {
				s.mockTM()

				s.teamRepo.
					On("AddTeam", s.ctx, mock.MatchedBy(func(t model.Team) bool {
						return t.TeamName == "backend" &&
							len(t.Members) == 2 &&
							t.Members[0].UserID == "u1"
					})).
					Return(model.Team{
						TeamName: "backend",
						Members: []model.TeamMember{
							{UserID: "u1", Username: "alice", IsActive: true},
							{UserID: "u2", Username: "bob", IsActive: true},
						},
					}, nil).
					Once()
			},
			expectedTeam: model.Team{
				TeamName: "backend",
				Members: []model.TeamMember{
					{UserID: "u1", Username: "alice", IsActive: true},
					{UserID: "u2", Username: "bob", IsActive: true},
				},
			},
		},

		{
			name: "успех: команда без участников",
			inputTeam: model.Team{
				TeamName: "empty",
				Members:  []model.TeamMember{},
			},
			setupMocks: func() {
				s.mockTM()

				s.teamRepo.
					On("AddTeam", s.ctx, mock.Anything).
					Return(model.Team{TeamName: "empty", Members: []model.TeamMember{}}, nil).
					Once()
			},
			expectedTeam: model.Team{TeamName: "empty", Members: []model.TeamMember{}},
		},

		{
			name: "ошибка: команда уже существует",
			inputTeam: model.Team{
				TeamName: "backend",
				Members:  []model.TeamMember{},
			},
			setupMocks: func() {
				s.mockTM()

				s.teamRepo.
					On("AddTeam", s.ctx, mock.Anything).
					Return(model.Team{}, apperror.NewTeamExistsError("backend")).
					Once()
			},
			expectedErrTyp: apperror.NewTeamExistsError("backend"),
		},

		{
			name: "ошибка: один из участников в другой команде",
			inputTeam: model.Team{
				TeamName: "backend",
				Members: []model.TeamMember{
					{UserID: "u777", Username: "heh", IsActive: true},
				},
			},
			setupMocks: func() {
				s.mockTM()

				s.teamRepo.
					On("AddTeam", s.ctx, mock.Anything).
					Return(model.Team{}, apperror.NewUserInAnotherTeamError("u777")).
					Once()
			},
			expectedErrTyp: apperror.NewUserInAnotherTeamError("u777"),
		},

		{
			name: "ошибка базы данных",
			inputTeam: model.Team{
				TeamName: "backend",
				Members:  []model.TeamMember{},
			},
			setupMocks: func() {
				s.mockTM()

				s.teamRepo.
					On("AddTeam", s.ctx, mock.Anything).
					Return(model.Team{}, errors.New("db timeout")).
					Once()
			},
			expectedErr: errors.New("db timeout"),
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.setupMocks()

			actualTeam, err := s.service.AddTeam(s.ctx, tt.inputTeam)

			switch {
			case tt.expectedErr != nil:
				s.Error(err)
				s.Equal(tt.expectedErr.Error(), err.Error())

			case tt.expectedErrTyp != nil:
				s.Error(err)
				s.IsType(tt.expectedErrTyp, err)

			default:
				s.NoError(err)
				s.Equal(tt.expectedTeam, actualTeam)
			}

			s.teamRepo.AssertExpectations(s.T())
			s.trm.AssertExpectations(s.T())
		})
	}
}
