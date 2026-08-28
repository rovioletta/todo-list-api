package task

import (
	"context"
	"fmt"

	"rovioletta/todo-list-api/internal/db"
)

func (s *Service) CreateTask(ctx context.Context, newTask Task) (taskID uint64, err error) {
	taskID, err = s.queries.CreateTask(ctx, &db.CreateTaskParams{
		Title:       *newTask.Title,
		Description: *newTask.Description,
		Status:      db.TaskStatusTodo,
	})
	if err != nil {
		return taskID, fmt.Errorf("insert task to db: %w", err)
	}

	return
}
