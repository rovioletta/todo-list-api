package user

import (
	"context"

	pb "rovioletta/todo-list-api/pkg/pb/user"
)

func (*UserService) Register(context.Context, *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, nil
}
