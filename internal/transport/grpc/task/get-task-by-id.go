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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) GetTaskByID(ctx context.Context, req *pb.GetTaskByIDRequest) (*pb.GetTaskByIDResponse, error) {
	task, err := i.taskService.GetTaskByID(ctx, req.TaskId)
	if err != nil {
		i.logger.Debug("i.taskService.GetTaskByID", slog.String("error", err.Error()))

		if errors.Is(err, apperrors.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "task doesn't exist")
		}

		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	return &pb.GetTaskByIDResponse{
		Task: repackTaskToPb(task),
	}, nil
}

func repackTaskToPb(task *task.Task) *pb.Task {
	return &pb.Task{
		Id:          task.ID,
		Title:       *task.Title,
		Description: *task.Description,
		Status:      repackTaskStatusToPb(task.Status),
		CreatedAt:   timestamppb.New(task.CreatedAt),
	}
}

func repackTaskStatusToPb(status *task.TaskStatus) pb.TaskStatus {
	if status == nil {
		return pb.TaskStatus_TASK_STATUS_UNKNOWN
	}

	switch *status {
	case task.TaskStatusBacklog:
		return pb.TaskStatus_TASK_STATUS_BACKLOG

	case task.TaskStatusTodo:
		return pb.TaskStatus_TASK_STATUS_TODO

	case task.TaskStatusInProgress:
		return pb.TaskStatus_TASK_STATUS_IN_PROGRESS

	case task.TaskStatusInReview:
		return pb.TaskStatus_TASK_STATUS_IN_REVIEW

	case task.TaskStatusDone:
		return pb.TaskStatus_TASK_STATUS_DONE

	case task.TaskStatusCanceled:
		return pb.TaskStatus_TASK_STATUS_CANCELED

	default:
		return pb.TaskStatus_TASK_STATUS_UNKNOWN
	}
}
