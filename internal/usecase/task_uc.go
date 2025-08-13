package usecase

import (
	"context"
	"errors"

	"github.com/soulstalker/task-api/internal/domain"
)

var (
	ErrNotFound = errors.New("task not found")
	ErrInvalid  = errors.New("invalid input")
)

type TaskUC struct {
	repo domain.TaskRepo
}

func NewTaskUC(repo domain.TaskRepo) *TaskUC {
	return &TaskUC{repo: repo}
}

func (s *TaskUC) All(ctx context.Context, status *domain.Status) ([]domain.Task, error) {
	return s.repo.All(ctx, status)
}

func (s *TaskUC) GetById(ctx context.Context, id uint) (domain.Task, error) {
	return s.repo.GetById(ctx, id)
}

func (s *TaskUC) Create(ctx context.Context, task domain.Task) (domain.Task, error) {
	if task.Title == "" {
		return domain.Task{}, ErrInvalid
	}
	return s.repo.Create(ctx, task)
}
