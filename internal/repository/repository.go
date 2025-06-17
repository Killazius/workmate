package repository

import (
	"github.com/Killazius/workmate/internal/domain"
	"github.com/Killazius/workmate/internal/storage"
)

type TaskRepo struct {
	storage storage.TaskStorage
}

func New(storage storage.TaskStorage) *TaskRepo {
	return &TaskRepo{
		storage: storage,
	}
}

func (r *TaskRepo) Create(task *domain.Task) error {
	if _, exists := r.storage.Get(task.ID); exists {
		return domain.ErrTaskAlreadyExists
	}

	createdTask := r.storage.Create()

	*createdTask = *task
	createdTask.ID = task.ID

	return nil
}

func (r *TaskRepo) GetByID(id string) (*domain.Task, error) {

	task, exists := r.storage.Get(id)
	if !exists {
		return nil, domain.ErrTaskNotFound
	}

	copyTask := *task
	return &copyTask, nil
}

func (r *TaskRepo) Update(task *domain.Task) error {

	if _, exists := r.storage.Get(task.ID); !exists {
		return domain.ErrTaskNotFound
	}

	r.storage.Delete(task.ID)
	newTask := r.storage.Create()
	*newTask = *task
	newTask.ID = task.ID
	return nil
}

func (r *TaskRepo) Delete(id string) error {

	if _, exists := r.storage.Get(id); !exists {
		return domain.ErrTaskNotFound
	}

	r.storage.Delete(id)
	return nil
}
