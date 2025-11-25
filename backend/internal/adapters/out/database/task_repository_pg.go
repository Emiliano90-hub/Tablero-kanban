package database

import (
	"context"
	"backend/internal/domain/models"
	"gorm.io/gorm"
)

type TaskRepositoryPG struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepositoryPG {
	return &TaskRepositoryPG{db: db}
}

func (r *TaskRepositoryPG) Create(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *TaskRepositoryPG) Update(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Save(task).Error
}

func (r *TaskRepositoryPG) Delete(ctx context.Context, taskID uint) error {
	return r.db.WithContext(ctx).Delete(&models.Task{}, taskID).Error
}

func (r *TaskRepositoryPG) FindByID(ctx context.Context, id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.WithContext(ctx).First(&task, id).Error
	return &task, err
}

func (r *TaskRepositoryPG) FindAllTasks(ctx context.Context, columnID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.WithContext(ctx).Where("column_id = ?", columnID).Order("index").Find(&tasks).Error
	return tasks, err
}
