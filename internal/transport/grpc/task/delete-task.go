package task

import (
	"context"
	"errors"
	"log/slog"

	"rovioletta/todo-list-api/internal/apperrors"
	pb "rovioletta/todo-list-api/pkg/pb/task"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*emptypb.Empty, error) {
	err := i.taskService.DeleteTask(ctx, req.TaskId)

	if err != nil {
		i.logger.Debug("i.taskService.DeleteTask", slog.String("error", err.Error()))

		if errors.Is(err, apperrors.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "task doesn't exist")
		}
		
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	return nil, nil
}
