package repository

import (
	"github.com/Killazius/workmate/internal/domain"
	"github.com/Killazius/workmate/internal/storage"
	"time"
)

type TaskRepo struct {
	storage storage.TaskStorage
}

func New(storage storage.TaskStorage) *TaskRepo {
	return &TaskRepo{
		storage: storage,
	}
}

func (r *TaskRepo) Create(time time.Time) (*domain.Task, error) {
	createdTask := r.storage.Create(time)
	return createdTask, nil
}

func (r *TaskRepo) GetByID(id int64) (*domain.Task, error) {

	task, exists := r.storage.Get(id)
	if !exists {
		return nil, domain.ErrTaskNotFound
	}

	return task, nil
}

func (r *TaskRepo) Update(task *domain.Task) error {
	_, err := r.storage.Update(task)
	return err
}

func (r *TaskRepo) Delete(id int64) error {

	if _, exists := r.storage.Get(id); !exists {
		return domain.ErrTaskNotFound
	}

	r.storage.Delete(id)
	return nil
}
