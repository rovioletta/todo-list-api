package task

import (
	"context"
	"log/slog"

	"rovioletta/todo-list-api/internal/service/task"
	pb "rovioletta/todo-list-api/pkg/pb/task"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	taskID, err := i.taskService.CreateTask(ctx, task.Task{
		Title:       &req.Title,
		Description: &req.Description,
	})

	if err != nil {
		i.logger.Debug("i.taskService.CreateTask", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	return &pb.CreateTaskResponse{
		TaskId: taskID,
	}, nil
}
