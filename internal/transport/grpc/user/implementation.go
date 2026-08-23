package user

import (
	"context"
	"log/slog"

	validate "buf.build/go/protovalidate"
	pb "rovioletta/todo-list-api/pkg/pb/user"
)

type userService interface {
	CreateUser(ctx context.Context, login, password string) (userID uint64, err error)
}

type Implementation struct {
	pb.UnimplementedUserServiceServer

	v           validate.Validator
	logger      *slog.Logger
	userService userService
}

func NewImplementation(logger *slog.Logger, userService userService) *Implementation {
	v, err := validate.New()
	if err != nil {
		panic(err)
	}

	return &Implementation{
		v:           v,
		logger:      logger,
		userService: userService,
	}
}
