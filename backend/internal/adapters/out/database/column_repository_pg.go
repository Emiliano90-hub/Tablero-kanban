package database

import (
	"backend/internal/domain/models"
	"context"

	"gorm.io/gorm"
)

type ColumnRepositoryPG struct {
	db *gorm.DB
}

func NewColumnRepository(db *gorm.DB) *ColumnRepositoryPG {
	return &ColumnRepositoryPG{db: db}
}

func (r *ColumnRepositoryPG) Create(ctx context.Context, column *models.Column) error {
	return r.db.WithContext(ctx).Create(column).Error
}

func (r *ColumnRepositoryPG) Update(ctx context.Context, column *models.Column) error {
	return r.db.WithContext(ctx).Save(column).Error
}

func (r *ColumnRepositoryPG) Delete(ctx context.Context, columnID uint) error {
	return r.db.WithContext(ctx).Delete(&models.Column{}, columnID).Error
}

func (r *ColumnRepositoryPG) FindByID(ctx context.Context, id uint) (*models.Column, error) {
	var column models.Column
	err := r.db.WithContext(ctx).First(&column, id).Error
	return &column, err
}

func (r *ColumnRepositoryPG) FindAllColumns(ctx context.Context, boardID uint) ([]models.Column, error) {
	var columns []models.Column
	err := r.db.WithContext(ctx).Where("board_id = ?", boardID).Order("index").Find(&columns).Error
	return columns, err
}
