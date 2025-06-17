package taskstorage

import (
	"github.com/Killazius/workmate/internal/domain"
	"github.com/Killazius/workmate/internal/lib/id"
	"sync"
	"time"
)

type TaskStorage struct {
	mu    sync.RWMutex
	tasks map[int64]*domain.Task
	gen   *id.Generator
}

func New() *TaskStorage {
	return &TaskStorage{
		tasks: make(map[int64]*domain.Task),
		gen:   id.New(),
	}
}

func (s *TaskStorage) Create(time time.Time) *domain.Task {
	task := &domain.Task{
		ID:        s.gen.Next(),
		Status:    domain.Created,
		CreatedAt: time,
	}
	s.mu.Lock()
	s.tasks[task.ID] = task
	s.mu.Unlock()

	return task
}
func (s *TaskStorage) Update(task *domain.Task) (*domain.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[task.ID]; !exists {
		return nil, domain.ErrTaskNotFound
	}

	updatedTask := *task
	s.tasks[task.ID] = &updatedTask

	return &updatedTask, nil
}
func (s *TaskStorage) Get(id int64) (*domain.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	return task, exists
}

func (s *TaskStorage) Delete(id int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.tasks, id)
}
