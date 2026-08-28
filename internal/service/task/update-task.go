package task

import (
	"context"
	"fmt"

	"rovioletta/todo-list-api/internal/db"
)

func (s *Service) UpdateTask(ctx context.Context, updatedTask Task) (err error) {
	_, err = s.queries.UpdateTask(ctx, &db.UpdateTaskParams{
		ID:             updatedTask.ID,
		NewTitle:       updatedTask.Title,
		NewDescription: updatedTask.Description,
		Status:         (*db.TaskStatus)(updatedTask.Status),
	})
	if err != nil {
		return fmt.Errorf("insert task to db: %w", err)
	}

	return
}
