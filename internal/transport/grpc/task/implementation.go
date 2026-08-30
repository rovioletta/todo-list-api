package task

import (
	"context"
	"log/slog"

	"rovioletta/todo-list-api/internal/service/task"
	pb "rovioletta/todo-list-api/pkg/pb/task"
)

type taskService interface {
	CreateTask(ctx context.Context, newTask task.Task) (taskID uint64, err error)
	UpdateTask(ctx context.Context, updatedTask task.Task) (err error)
	GetTaskByID(ctx context.Context, taskID uint64) (task *task.Task, err error)
}

type Implementation struct {
	pb.UnimplementedTaskServiceServer

	logger      *slog.Logger
	taskService taskService
}

func NewImplementation(logger *slog.Logger, taskService taskService) *Implementation {
	return &Implementation{
		logger:      logger,
		taskService: taskService,
	}
}
