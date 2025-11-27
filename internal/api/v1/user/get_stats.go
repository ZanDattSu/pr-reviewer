package user

import (
	"context"
	"errors"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/api/apierrors"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (h *userHandler) UsersStatsGet(ctx context.Context, params reviewerV1.UsersStatsGetParams) (reviewerV1.UsersStatsGetRes, error) {
	top := 0

	if v, ok := params.Top.Get(); ok {
		top = v
	}

	onlyActive := false

	if v, ok := params.OnlyActive.Get(); ok {
		onlyActive = v
	}

	onlyOpen := false

	if v, ok := params.OnlyOpen.Get(); ok {
		onlyOpen = v
	}

	stats, err := h.userService.GetUserStats(ctx, top, onlyActive, onlyOpen)
	if err != nil {

		var noDataErr *apperror.NoDataError

		if errors.As(err, &noDataErr) {
			return &reviewerV1.ErrorResponse{
				Error: apierrors.NewNoStatsError(),
			}, nil
		}

		return nil, err

	}

	resp := make([]reviewerV1.UsersStatsGetOKUsersItem, 0, len(stats))

	for _, s := range stats {
		resp = append(resp, reviewerV1.UsersStatsGetOKUsersItem{
			UserID: s.UserID,

			TotalPr: s.TotalPR,
		})
	}

	return &reviewerV1.UsersStatsGetOK{
		Users: resp,
	}, nil
}
