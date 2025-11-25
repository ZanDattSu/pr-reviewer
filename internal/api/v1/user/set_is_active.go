package user

import (
	"context"
	"errors"

	reviewerV1 "github.com/ZanDattSu/pr-reviewer/api/pkg/reviewer/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/api/apierrors"
	"github.com/ZanDattSu/pr-reviewer/internal/converter"
	"github.com/ZanDattSu/pr-reviewer/internal/model/apperror"
)

func (h *userHandler) UsersSetIsActivePost(ctx context.Context, req *reviewerV1.UsersSetIsActivePostReq) (reviewerV1.UsersSetIsActivePostRes, error) {
	activeUser, err := h.userService.UpdateUserStatus(ctx, req.UserID, req.IsActive)
	if err != nil {

		var userNF *apperror.UserNotFoundError

		if errors.As(err, &userNF) {
			return &reviewerV1.ErrorResponse{
				Error: apierrors.NewUserNotFoundErr(userNF.UserID()),
			}, nil
		}

		return nil, err

	}

	return &reviewerV1.UsersSetIsActivePostOK{
		User: reviewerV1.NewOptUser(
			converter.ServiceUserToAPI(activeUser),
		),
	}, nil
}
