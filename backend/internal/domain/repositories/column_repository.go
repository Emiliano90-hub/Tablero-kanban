package repositories

import (
	"backend/internal/domain/models"
	"context"
)

type ColumnRepository interface {
	Create(ctx context.Context, column *models.Column) error
	Update(ctx context.Context, column *models.Column) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*models.Column, error)
	FindAllColumns(ctx context.Context, boardID uint) ([]models.Column, error)
}
