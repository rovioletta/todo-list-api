package user

import (
	"context"
	"errors"
	"log/slog"

	"rovioletta/todo-list-api/internal/apperrors"
	pb "rovioletta/todo-list-api/pkg/pb/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := i.userService.Login(ctx, req.GetLogin(), req.GetPassword())
	if err != nil {
		i.logger.Debug("i.userService.Login", slog.String("error", err.Error()))

		if errors.Is(err, apperrors.ErrWrongPassword) {
			return nil, status.Errorf(codes.PermissionDenied, "wrong password")
		}

		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}
