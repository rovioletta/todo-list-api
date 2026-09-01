package task

import (
	"context"
	"log/slog"

	"rovioletta/todo-list-api/internal/service/task"
	pb "rovioletta/todo-list-api/pkg/pb/task"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetTasksByFilter(ctx context.Context, req *pb.GetTasksByFilterRequest) (*pb.GetTasksByFilterResponse, error) {
	raw, err := i.taskService.GetTasksByFilter(ctx, repackFilterToModel(req))
	if err != nil {
		i.logger.Debug("i.taskService.GetTasksByFilter", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	tasks := make([]*pb.Task, 0, len(raw))
	for _, t := range raw {
		tasks = append(tasks, repackTaskToPb(&t))
	}

	return &pb.GetTasksByFilterResponse{
		Tasks: tasks,
	}, nil
}

func repackFilterToModel(raw *pb.GetTasksByFilterRequest) *task.TaskFilter {
	filter := &task.TaskFilter{}

	if raw.Filter != nil {
		filter.Filter = &task.Filter{}
		filter.Filter.SearchTitle = raw.Filter.SearchTitle

		status := repackTaskStatusToModel(raw.Filter.SearchStatus)
		filter.Filter.SearchStatus = status

		if raw.Filter.CreatedFrom != nil {
			createdFrom := raw.Filter.CreatedFrom.AsTime()
			filter.Filter.CreatedFrom = &createdFrom
		}

		if raw.Filter.CreatedTo != nil {
			createdTo := raw.Filter.CreatedTo.AsTime()
			filter.Filter.CreatedTo = &createdTo
		}
	}

	if raw.Pagination != nil {
		filter.Pagination = &task.Pagination{}
		filter.Pagination.Limit = raw.Pagination.Limit
		filter.Pagination.Offset = raw.Pagination.Offset
	}

	if raw.Sort != nil {
		filter.Sort = &task.Sort{}
		filter.Sort.Field = raw.Sort.Field
		filter.Sort.Order = raw.Sort.Order
	}

	return filter
}
