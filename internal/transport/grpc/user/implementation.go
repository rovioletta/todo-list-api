package user

import (
	"context"
	"log/slog"

	pb "rovioletta/todo-list-api/pkg/pb/user"
)

type userService interface {
	CreateUser(ctx context.Context, login, password string) (userID uint64, err error)
	Login(ctx context.Context, login, password string) (token string, err error)
}

type Implementation struct {
	pb.UnimplementedUserServiceServer

	logger      *slog.Logger
	userService userService
}

func NewImplementation(logger *slog.Logger, userService userService) *Implementation {
	return &Implementation{
		logger:      logger,
		userService: userService,
	}
}
