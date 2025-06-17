package domain

import "time"

type Task struct {
	ID          string     `json:"id"`
	Status      string     `json:"status"`
	Result      string     `json:"result,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}
