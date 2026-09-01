package task

import (
	"context"
	"fmt"

	"rovioletta/todo-list-api/internal/db"
)

func (s *Service) GetTaskByID(ctx context.Context, taskID uint64) (task *Task, err error) {
	row, err := s.queries.GetTaskByID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("get task by id from db: %w", err)
	}

	return &Task{
		ID:          row.ID,
		Title:       &row.Title,
		Description: &row.Description,
		Status:      repackStatus(row.Status),
		CreatedAt:   *row.CreatedAt,
	}, err
}

func repackStatus(status db.TaskStatus) *TaskStatus {
	switch status {
	case db.TaskStatusBacklog:
		repacked := TaskStatusBacklog
		return &repacked
	case db.TaskStatusTodo:
		repacked := TaskStatusTodo
		return &repacked
	case db.TaskStatusInProgress:
		repacked := TaskStatusInProgress
		return &repacked
	case db.TaskStatusInReview:
		repacked := TaskStatusInReview
		return &repacked
	case db.TaskStatusDone:
		repacked := TaskStatusDone
		return &repacked
	case db.TaskStatusCanceled:
		repacked := TaskStatusCanceled
		return &repacked

	default:
		return nil
	}
}
