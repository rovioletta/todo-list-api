package task

import "time"

type TaskStatus string

const (
	TaskStatusBacklog    TaskStatus = "backlog"
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusInReview   TaskStatus = "in_review"
	TaskStatusDone       TaskStatus = "done"
	TaskStatusCanceled   TaskStatus = "canceled"
)

type Task struct {
	ID          uint64
	Title       *string
	Description *string
	Status      *TaskStatus
	CreatedAt   time.Time
}

type Filter struct {
	SearchTitle  *string
	SearchStatus *TaskStatus
	CreatedFrom  *time.Time
	CreatedTo    *time.Time
}

type Pagination struct {
	Limit  uint64
	Offset uint64
}

type Sort struct {
	Field *string
	Order *string
}

type TaskFilter struct {
	Filter     *Filter
	Pagination *Pagination
	Sort       *Sort
}
