package task

import (
	"log/slog"

	validate "buf.build/go/protovalidate"
	pb "rovioletta/todo-list-api/pkg/pb/task"
)

type taskService interface {
}

type Implementation struct {
	pb.UnimplementedTaskServiceServer

	v           validate.Validator
	logger      *slog.Logger
	taskService taskService
}

func NewImplementation(logger *slog.Logger) *Implementation {
	v, err := validate.New()
	if err != nil {
		panic(err)
	}

	return &Implementation{
		v:      v,
		logger: logger,
	}
}
