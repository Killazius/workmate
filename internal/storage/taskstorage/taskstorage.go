package taskstorage

import (
	"github.com/Killazius/workmate/internal/domain"
	"github.com/google/uuid"
	"sync"
	"time"
)

type TaskStorage struct {
	mu    sync.RWMutex
	tasks map[string]*domain.Task
}

func New() *TaskStorage {
	return &TaskStorage{
		tasks: make(map[string]*domain.Task),
	}
}

func (s *TaskStorage) Create() *domain.Task {
	task := &domain.Task{
		ID:        uuid.New().String(),
		Status:    "",
		CreatedAt: time.Now(),
	}

	s.mu.Lock()
	s.tasks[task.ID] = task
	s.mu.Unlock()

	return task
}

func (s *TaskStorage) Get(id string) (*domain.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	return task, exists
}

func (s *TaskStorage) Delete(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.tasks, id)
}
