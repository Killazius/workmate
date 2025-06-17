package service

import (
	"context"
	"github.com/Killazius/workmate/internal/domain"
	"time"
)

type Repository interface {
	Create(time time.Time) (*domain.Task, error)
	GetByID(id int64) (*domain.Task, error)
	Update(task *domain.Task) error
	Delete(id int64) error
}
type TaskService struct {
	repo Repository
}

func New(repo Repository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context) (*domain.Task, error) {
	task, err := s.repo.Create(time.Now())
	if err != nil {
		return nil, err
	}

	go s.processTask(ctx, task.ID)

	return task, nil
}

func (s *TaskService) GetTask(ctx context.Context, id int64) (*domain.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) DeleteTask(ctx context.Context, id int64) error {
	return s.repo.Delete(id)
}

func (s *TaskService) processTask(ctx context.Context, taskID int64) {
	task, err := s.repo.GetByID(taskID)
	if err != nil {
		return
	}
	t := time.Now()
	task.StartedAt = &t
	task.Status = domain.Running
	if err = s.repo.Update(task); err != nil {
		return
	}

	select {
	case <-time.After(3 * time.Minute):
	case <-ctx.Done():
		task.Status = domain.Failed
		task.Result = ctx.Err().Error()
		if err = s.repo.Update(task); err != nil {
			return
		}
		return
	}

	task.Status = domain.Done
	t2 := time.Now()
	task.CompletedAt = &t2
	task.Result = "task completed"
	if err = s.repo.Update(task); err != nil {
		return
	}

}
