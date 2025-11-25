package repositories

import (
	"context"
	"backend/internal/domain/models"
)

type BoardRepository interface {
	Create(ctx context.Context, board *models.Board) error
	Update(ctx context.Context, board *models.Board) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*models.Board, error)
	FindAll(ctx context.Context) ([]models.Board, error)
}
