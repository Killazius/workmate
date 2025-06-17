package domain

import "time"

type TaskStatus string

const (
	Created TaskStatus = "created"
	Running TaskStatus = "running"
	Failed  TaskStatus = "failed"
	Done    TaskStatus = "done"
)

type Task struct {
	ID          int64      `json:"id"`
	Status      TaskStatus `json:"status"`
	Result      string     `json:"result,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}
