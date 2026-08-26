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

func (i *Implementation) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	userID, err := i.userService.CreateUser(ctx, req.GetLogin(), req.GetPassword())
	if err != nil {
		i.logger.Debug("i.userService.CreateUser", slog.String("error", err.Error()))

		if errors.Is(err, apperrors.ErrAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "user already exists")
		}

		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	return &pb.RegisterResponse{
		UserId: userID,
	}, nil
}
