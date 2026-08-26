package task

import (
	"log/slog"

	pb "rovioletta/todo-list-api/pkg/pb/task"
)

type taskService interface {
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
