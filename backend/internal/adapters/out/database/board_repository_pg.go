package database

import (
	"context"
	"backend/internal/domain/models"
	"gorm.io/gorm"
)

type BoardRepositoryPG struct {
	db *gorm.DB
}

func NewBoardRepository(db *gorm.DB) *BoardRepositoryPG {
	return &BoardRepositoryPG{db: db}
}

func (r *BoardRepositoryPG) Create(ctx context.Context, board *models.Board) error {
	return r.db.WithContext(ctx).Create(board).Error
}

func (r *BoardRepositoryPG) Update(ctx context.Context, board *models.Board) error {
	return r.db.WithContext(ctx).Save(board).Error
}

func (r *BoardRepositoryPG) Delete(ctx context.Context, boardID uint) error {
	return r.db.WithContext(ctx).Delete(&models.Board{}, boardID).Error
}

func (r *BoardRepositoryPG) FindByID(ctx context.Context, id uint) (*models.Board, error) {
	var board models.Board
	err := r.db.WithContext(ctx).First(&board, id).Error
	return &board, err
}

func (r *BoardRepositoryPG) FindAll(ctx context.Context) ([]models.Board, error) {
	var boards []models.Board
	err := r.db.WithContext(ctx).Order("created_at").Find(&boards).Error
	return boards, err
}
