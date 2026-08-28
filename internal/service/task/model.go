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
