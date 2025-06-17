package storage

import (
	"github.com/Killazius/workmate/internal/domain"
	"time"
)

type TaskStorage interface {
	Create(time time.Time) *domain.Task
	Get(id int64) (*domain.Task, bool)
	Delete(id int64)
	Update(task *domain.Task) (*domain.Task, error)
}
