package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/ZanDattSu/pr-reviewer/internal/model"
)

func (s *SuiteService) mockTM() {
	s.trm.
		On("Do", mock.Anything, mock.Anything).
		Return(func(ctx context.Context, fn func(context.Context) error) error {
			return fn(ctx)
		})
}

func (s *SuiteService) TestDeactivateUsersAndReassignPR() {
	now := time.Now()

	type fields struct {
		userIDs []string
	}

	tests := []struct {
		name        string
		fields      fields
		setupMocks  func()
		wantErr     bool
		errContains string
		wantPRs     []model.ReassignedPR
		checkResult func(got []model.ReassignedPR)
	}{
		{
			name:   "успех — 1 деактивирован, 2 PR перераспределены",
			fields: fields{userIDs: []string{"u1"}},
			setupMocks: func() {
				s.mockTM()

				s.userRepo.
					On("DeactivateUsers", s.ctx, []string{"u1"}).
					Return([]string{"u1"}, nil).
					Once()

				s.prRepo.
					On("FindOpenPRsWithReviewers", s.ctx, []string{"u1"}).
					Return([]model.OpenPR{
						{PRID: "pr1", OldReviewer: "u1"},
						{PRID: "pr2", OldReviewer: "u1"},
					}, nil).
					Once()

				s.prService.
					On("ReassignPullRequest", s.ctx, "pr1", "u1").
					Return(model.PullRequest{PullRequestID: "pr1", CreatedAt: &now}, "new1", nil).
					Once()

				s.prService.
					On("ReassignPullRequest", s.ctx, "pr2", "u1").
					Return(model.PullRequest{PullRequestID: "pr2", CreatedAt: &now}, "new2", nil).
					Once()
			},
			wantErr: false,
			wantPRs: []model.ReassignedPR{
				{
					PullRequestID: model.PullRequest{PullRequestID: "pr1", CreatedAt: &now},
					ReplacedBy:    "new1",
				},
				{
					PullRequestID: model.PullRequest{PullRequestID: "pr2", CreatedAt: &now},
					ReplacedBy:    "new2",
				},
			},
		},

		{
			name:   "успех — открытых PR нет → возвращаем пустой список",
			fields: fields{userIDs: []string{"u1"}},
			setupMocks: func() {
				s.mockTM()

				s.userRepo.
					On("DeactivateUsers", s.ctx, []string{"u1"}).
					Return([]string{"u1"}, nil).
					Once()

				s.prRepo.
					On("FindOpenPRsWithReviewers", s.ctx, []string{"u1"}).
					Return([]model.OpenPR{}, nil).
					Once()
			},
			wantErr: false,
			wantPRs: []model.ReassignedPR{},
		},
		{
			name:   "успех — деактивация 3 пользователей с различными PR",
			fields: fields{userIDs: []string{"u1", "u2", "u3"}},
			setupMocks: func() {
				s.mockTM()

				s.userRepo.
					On("DeactivateUsers", s.ctx, []string{"u1", "u2", "u3"}).
					Return([]string{"u1", "u2", "u3"}, nil).
					Once()

				s.prRepo.
					On("FindOpenPRsWithReviewers", s.ctx, []string{"u1", "u2", "u3"}).
					Return([]model.OpenPR{
						{PRID: "pr1", OldReviewer: "u1"},
						{PRID: "pr2", OldReviewer: "u2"},
						{PRID: "pr3", OldReviewer: "u1"},
						{PRID: "pr4", OldReviewer: "u3"},
					}, nil).
					Once()

				s.prService.
					On("ReassignPullRequest", s.ctx, "pr1", "u1").
					Return(model.PullRequest{PullRequestID: "pr1", CreatedAt: &now}, "reviewer1", nil).
					Once()

				s.prService.
					On("ReassignPullRequest", s.ctx, "pr2", "u2").
					Return(model.PullRequest{PullRequestID: "pr2", CreatedAt: &now}, "reviewer2", nil).
					Once()

				s.prService.
					On("ReassignPullRequest", s.ctx, "pr3", "u1").
					Return(model.PullRequest{PullRequestID: "pr3", CreatedAt: &now}, "reviewer3", nil).
					Once()

				s.prService.
					On("ReassignPullRequest", s.ctx, "pr4", "u3").
					Return(model.PullRequest{PullRequestID: "pr4", CreatedAt: &now}, "reviewer4", nil).
					Once()
			},
			wantErr: false,
			checkResult: func(got []model.ReassignedPR) {
				s.Require().Len(got, 4, "должно быть 4 переназначенных PR")
				s.Equal("pr1", got[0].PullRequestID.PullRequestID)
				s.Equal("pr2", got[1].PullRequestID.PullRequestID)
				s.Equal("pr3", got[2].PullRequestID.PullRequestID)
				s.Equal("pr4", got[3].PullRequestID.PullRequestID)
			},
		},
		{
			name:   "успех — частичная деактивация (запрошено 3, деактивировано 2)",
			fields: fields{userIDs: []string{"u1", "u2", "u999"}},
			setupMocks: func() {
				s.mockTM()

				s.userRepo.
					On("DeactivateUsers", s.ctx, []string{"u1", "u2", "u999"}).
					Return([]string{"u1", "u2"}, nil).
					Once()

				s.prRepo.
					On("FindOpenPRsWithReviewers", s.ctx, []string{"u1", "u2"}).
					Return([]model.OpenPR{
						{PRID: "pr1", OldReviewer: "u1"},
					}, nil).
					Once()

				s.prService.
					On("ReassignPullRequest", s.ctx, "pr1", "u1").
					Return(model.PullRequest{PullRequestID: "pr1", CreatedAt: &now}, "new1", nil).
					Once()
			},
			wantErr: false,
			wantPRs: []model.ReassignedPR{
				{
					PullRequestID: model.PullRequest{PullRequestID: "pr1", CreatedAt: &now},
					ReplacedBy:    "new1",
				},
			},
		},
		{
			name:   "успех — один пользователь с 10 открытыми PR",
			fields: fields{userIDs: []string{"u1"}},
			setupMocks: func() {
				s.mockTM()

				s.userRepo.
					On("DeactivateUsers", s.ctx, []string{"u1"}).
					Return([]string{"u1"}, nil).
					Once()

				openPRs := make([]model.OpenPR, 10)
				for i := 0; i < 10; i++ {
					prID := fmt.Sprintf("pr%d", i+1)
					openPRs[i] = model.OpenPR{PRID: prID, OldReviewer: "u1"}

					s.prService.
						On("ReassignPullRequest", s.ctx, prID, "u1").
						Return(model.PullRequest{PullRequestID: prID, CreatedAt: &now}, fmt.Sprintf("reviewer%d", i+1), nil).
						Once()
				}

				s.prRepo.
					On("FindOpenPRsWithReviewers", s.ctx, []string{"u1"}).
					Return(openPRs, nil).
					Once()
			},
			wantErr: false,
			checkResult: func(got []model.ReassignedPR) {
				s.Require().Len(got, 10, "должно быть 10 переназначенных PR")
				for i := 0; i < 10; i++ {
					s.Equal(fmt.Sprintf("pr%d", i+1), got[i].PullRequestID.PullRequestID)
					s.Equal(fmt.Sprintf("reviewer%d", i+1), got[i].ReplacedBy)
				}
			},
		},

		{
			name:   "успех — массовая деактивация 5 пользователей без PR",
			fields: fields{userIDs: []string{"u1", "u2", "u3", "u4", "u5"}},
			setupMocks: func() {
				s.mockTM()

				s.userRepo.
					On("DeactivateUsers", s.ctx, []string{"u1", "u2", "u3", "u4", "u5"}).
					Return([]string{"u1", "u2", "u3", "u4", "u5"}, nil).
					Once()

				s.prRepo.
					On("FindOpenPRsWithReviewers", s.ctx, []string{"u1", "u2", "u3", "u4", "u5"}).
					Return([]model.OpenPR{}, nil).
					Once()
			},
			wantErr: false,
			wantPRs: []model.ReassignedPR{},
		},
		{
			name:   "ошибка — пустой список userIDs",
			fields: fields{userIDs: []string{}},
			setupMocks: func() {
				s.mockTM()

				s.userRepo.
					On("DeactivateUsers", s.ctx, []string{}).
					Return([]string{}, nil).
					Once()
			},
			wantErr:     true,
			errContains: "no users were deactivated",
		},
		{
			name:   "ошибка — nil список userIDs",
			fields: fields{userIDs: nil},
			setupMocks: func() {
				s.mockTM()

				s.userRepo.
					On("DeactivateUsers", s.ctx, []string(nil)).
					Return([]string{}, nil).
					Once()
			},
			wantErr:     true,
			errContains: "no users were deactivated",
		},
		{
			name:   "ошибка — DeactivateUsers падает с ошибкой БД",
			fields: fields{userIDs: []string{"u1", "u2"}},
			setupMocks: func() {
				s.mockTM()

				s.userRepo.
					On("DeactivateUsers", s.ctx, []string{"u1", "u2"}).
					Return(nil, errors.New("database connection lost")).
					Once()
			},
			wantErr:     true,
			errContains: "failed to deactivate users",
		},
		{
			name:   "ошибка — все запрошенные пользователи не найдены",
			fields: fields{userIDs: []string{"u999", "u998"}},
			setupMocks: func() {
				s.mockTM()

				s.userRepo.
					On("DeactivateUsers", s.ctx, []string{"u999", "u998"}).
					Return([]string{}, nil).
					Once()
			},
			wantErr:     true,
			errContains: "no users were deactivated",
		},
		{
			name:   "ошибка — FindOpenPRsWithReviewers падает",
			fields: fields{userIDs: []string{"u1"}},
			setupMocks: func() {
				s.mockTM()

				s.userRepo.
					On("DeactivateUsers", s.ctx, []string{"u1"}).
					Return([]string{"u1"}, nil).
					Once()

				s.prRepo.
					On("FindOpenPRsWithReviewers", s.ctx, []string{"u1"}).
					Return(nil, errors.New("query timeout")).
					Once()
			},
			wantErr:     true,
			errContains: "failed to find open PRs",
		},
		{
			name:   "ошибка — ReassignPullRequest падает на первом PR",
			fields: fields{userIDs: []string{"u1"}},
			setupMocks: func() {
				s.mockTM()

				s.userRepo.
					On("DeactivateUsers", s.ctx, []string{"u1"}).
					Return([]string{"u1"}, nil).
					Once()

				s.prRepo.
					On("FindOpenPRsWithReviewers", s.ctx, []string{"u1"}).
					Return([]model.OpenPR{
						{PRID: "pr1", OldReviewer: "u1"},
					}, nil).
					Once()

				s.prService.
					On("ReassignPullRequest", s.ctx, "pr1", "u1").
					Return(model.PullRequest{}, "", errors.New("no available reviewers")).
					Once()
			},
			wantErr:     true,
			errContains: "failed to reassign PR pr1",
		},
		{
			name:   "ошибка — ReassignPullRequest падает на втором PR из трёх",
			fields: fields{userIDs: []string{"u1"}},
			setupMocks: func() {
				s.mockTM()

				s.userRepo.
					On("DeactivateUsers", s.ctx, []string{"u1"}).
					Return([]string{"u1"}, nil).
					Once()

				s.prRepo.
					On("FindOpenPRsWithReviewers", s.ctx, []string{"u1"}).
					Return([]model.OpenPR{
						{PRID: "pr1", OldReviewer: "u1"},
						{PRID: "pr2", OldReviewer: "u1"},
						{PRID: "pr3", OldReviewer: "u1"},
					}, nil).
					Once()

				s.prService.
					On("ReassignPullRequest", s.ctx, "pr1", "u1").
					Return(model.PullRequest{PullRequestID: "pr1", CreatedAt: &now}, "new1", nil).
					Once()

				s.prService.
					On("ReassignPullRequest", s.ctx, "pr2", "u1").
					Return(model.PullRequest{}, "", errors.New("reviewer overloaded")).
					Once()
			},
			wantErr:     true,
			errContains: "failed to reassign PR pr2",
		},
		{
			name:   "ошибка — TransactionManager.Do возвращает ошибку (rollback)",
			fields: fields{userIDs: []string{"u1"}},
			setupMocks: func() {
				s.trm.
					On("Do", mock.Anything, mock.Anything).
					Return(errors.New("transaction deadlock")).
					Once()
			},
			wantErr:     true,
			errContains: "transaction deadlock",
		},
		{
			name:   "увольнение команды (5 человек, 15 PR)",
			fields: fields{userIDs: []string{"team-lead", "dev1", "dev2", "dev3", "dev4"}},
			setupMocks: func() {
				s.mockTM()

				team := []string{"team-lead", "dev1", "dev2", "dev3", "dev4"}

				s.userRepo.
					On("DeactivateUsers", s.ctx, team).
					Return(team, nil).
					Once()

				// Разные члены команды имеют разное количество PR
				openPRs := []model.OpenPR{
					{PRID: "pr1", OldReviewer: "team-lead"},
					{PRID: "pr2", OldReviewer: "team-lead"},
					{PRID: "pr3", OldReviewer: "dev1"},
					{PRID: "pr4", OldReviewer: "dev1"},
					{PRID: "pr5", OldReviewer: "dev1"},
					{PRID: "pr6", OldReviewer: "dev2"},
					{PRID: "pr7", OldReviewer: "dev2"},
					{PRID: "pr8", OldReviewer: "dev2"},
					{PRID: "pr9", OldReviewer: "dev2"},
					{PRID: "pr10", OldReviewer: "dev3"},
					{PRID: "pr11", OldReviewer: "dev3"},
					{PRID: "pr12", OldReviewer: "dev3"},
					{PRID: "pr13", OldReviewer: "dev4"},
					{PRID: "pr14", OldReviewer: "dev4"},
					{PRID: "pr15", OldReviewer: "dev4"},
				}

				s.prRepo.
					On("FindOpenPRsWithReviewers", s.ctx, team).
					Return(openPRs, nil).
					Once()

				for i, pr := range openPRs {
					s.prService.
						On("ReassignPullRequest", s.ctx, pr.PRID, pr.OldReviewer).
						Return(model.PullRequest{PullRequestID: pr.PRID, CreatedAt: &now}, fmt.Sprintf("backup-reviewer-%d", i), nil).
						Once()
				}
			},
			wantErr: false,
			checkResult: func(got []model.ReassignedPR) {
				s.Require().Len(got, 15, "все 15 PR должны быть переназначены")
				// Проверяем, что все PR уникальны
				prIDs := make(map[string]bool)
				for _, pr := range got {
					s.False(prIDs[pr.PullRequestID.PullRequestID], "PR ID должен быть уникальным")
					prIDs[pr.PullRequestID.PullRequestID] = true
				}
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Пересоздаём моки для каждого подтеста
			s.SetupTest()
			tt.setupMocks()

			got, err := s.service.DeactivateUsersAndReassignPR(s.ctx, tt.fields.userIDs)

			if tt.wantErr {
				s.Require().Error(err)
				if tt.errContains != "" {
					s.Contains(err.Error(), tt.errContains)
				}
				return
			}

			s.Require().NoError(err)

			if tt.checkResult != nil {
				tt.checkResult(got)
			} else {
				s.Equal(tt.wantPRs, got)
			}
		})
	}
}
