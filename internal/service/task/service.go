package task

import (
	"context"

	"rovioletta/todo-list-api/internal/db"
)

type queries interface {
	CreateTask(ctx context.Context, arg *db.CreateTaskParams) (uint64, error)
	UpdateTask(ctx context.Context, arg *db.UpdateTaskParams) (uint64, error)
	GetTaskByID(ctx context.Context, id uint64) (db.Task, error)
	DeleteTask(ctx context.Context, id uint64) (uint64, error)
}

type Service struct {
	queries queries
}

func NewService(queries queries) *Service {
	return &Service{
		queries: queries,
	}
}
