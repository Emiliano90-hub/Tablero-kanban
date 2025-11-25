package repositories

import (
	"backend/internal/domain/models"
	"context"
)

type TaskRepository interface {
	Create(ctx context.Context, task *models.Task) error
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*models.Task, error)
	FindAllTasks(ctx context.Context, columnID uint) ([]models.Task, error)
}
