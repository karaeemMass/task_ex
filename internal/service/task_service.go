package service

import (
	"context"
	"errors"

	"task_ex/internal/model"
	"task_ex/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(ctx context.Context, task *model.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	return s.repo.Create(ctx, task)
}

func (s *TaskService) List(ctx context.Context) ([]model.Task, error) {
	return s.repo.List(ctx)
}

func (s *TaskService) FindByID(ctx context.Context, id uint) (*model.Task, error) {
	return s.repo.FindByID(ctx, id)
}
