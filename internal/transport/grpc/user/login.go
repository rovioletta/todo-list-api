package user

import (
	"context"

	pb "rovioletta/todo-list-api/pkg/pb/user"
)

func (*Implementation) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}
