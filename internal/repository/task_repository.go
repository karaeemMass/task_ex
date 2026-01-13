package repository

import (
	"context"

	"gorm.io/gorm"
	"task_ex/internal/model"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, task *model.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *TaskRepository) List(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.WithContext(ctx).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) FindByID(ctx context.Context, id uint) (*model.Task, error) {
	var task model.Task
	err := r.db.WithContext(ctx).First(&task, id).Error
	return &task, err
}
