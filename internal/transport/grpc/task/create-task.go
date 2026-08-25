package task

import (
	"context"

	pb "rovioletta/todo-list-api/pkg/pb/task"
)

func (i *Implementation) CreateTask(context.Context, *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	return nil, nil
}
