package task

import (
	"context"
	"fmt"
)

func (s *Service) DeleteTask(ctx context.Context, taskID uint64) (err error) {
	_, err = s.queries.DeleteTask(ctx, taskID)
	if err != nil {
		return fmt.Errorf("delete task from db: %w", err)
	}

	return
}
