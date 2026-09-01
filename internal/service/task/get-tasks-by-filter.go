package task

import (
	"context"
	"fmt"

	"rovioletta/todo-list-api/internal/db"
)

func (s *Service) GetTasksByFilter(ctx context.Context, filter *TaskFilter) ([]Task, error) {
	params := &db.GetTasksByFilterParams{}

	if filter.Filter != nil {
		params.FilterSearchTitle = filter.Filter.SearchTitle
		params.FilterSearchStatus = (*string)(filter.Filter.SearchStatus)
		params.FilterCreatedFrom = filter.Filter.CreatedFrom
		params.FilterCreatedTo = filter.Filter.CreatedTo
	}

	if filter.Sort != nil {
		params.OrderByField = filter.Sort.Field
		params.OrderType = filter.Sort.Order
	}

	if filter.Pagination != nil {
		params.Limit = filter.Pagination.Limit
		params.Offset = filter.Pagination.Offset
	}

	res, err := s.queries.GetTasksByFilter(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("get tasks from db: %w", err)
	}

	tasks := make([]Task, 0, len(res))
	for _, t := range res {
		tasks = append(tasks, Task{
			ID:          t.ID,
			Title:       &t.Title,
			Description: &t.Description,
			Status:      repackStatus(t.Status),
			CreatedAt:   *t.CreatedAt,
		})
	}

	return tasks, nil
}
