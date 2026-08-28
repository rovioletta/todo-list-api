package task

import (
	"context"
	"errors"
	"log/slog"

	"rovioletta/todo-list-api/internal/apperrors"
	"rovioletta/todo-list-api/internal/service/task"
	pb "rovioletta/todo-list-api/pkg/pb/task"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*emptypb.Empty, error) {
	err := i.taskService.UpdateTask(ctx, task.Task{
		ID:          req.TaskId,
		Title:       req.Title,
		Description: req.Description,
		Status:      repackTaskStatusToModel(req.Status),
	})

	if err != nil {
		i.logger.Debug("i.taskService.UpdateTask", slog.String("error", err.Error()))

		if errors.Is(err, apperrors.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "task doesn't exist")
		}
		
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	return nil, nil
}

func repackTaskStatusToModel(status *pb.TaskStatus) *task.TaskStatus {
	if status == nil {
		return nil
	}

	switch *status {
	case pb.TaskStatus_TASK_STATUS_BACKLOG:
		repacked := task.TaskStatusBacklog
		return &repacked
	case pb.TaskStatus_TASK_STATUS_TODO:
		repacked := task.TaskStatusTodo
		return &repacked
	case pb.TaskStatus_TASK_STATUS_IN_PROGRESS:
		repacked := task.TaskStatusInProgress
		return &repacked
	case pb.TaskStatus_TASK_STATUS_IN_REVIEW:
		repacked := task.TaskStatusInReview
		return &repacked
	case pb.TaskStatus_TASK_STATUS_DONE:
		repacked := task.TaskStatusDone
		return &repacked
	case pb.TaskStatus_TASK_STATUS_CANCELED:
		repacked := task.TaskStatusCanceled
		return &repacked

	default:
		return nil
	}
}
