package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "rovioletta/todo-list-api/pkg/pb/user"
)

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if err := s.v.Validate(req); err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "%s", err.Error())
	}

	return nil, nil
}
