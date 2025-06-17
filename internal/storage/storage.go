package storage

import (
	"github.com/Killazius/workmate/internal/domain"
)

type TaskStorage interface {
	Create() *domain.Task
	Get(id string) (*domain.Task, bool)
	Delete(id string)
}
