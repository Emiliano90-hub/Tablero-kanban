package column_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"context"
	"fmt"
)

type UpdateColumn struct {
	Repo repositories.ColumnRepository
}

func NewUpdateColumn(repo repositories.ColumnRepository) *UpdateColumn {
	return &UpdateColumn{
		Repo: repo,
	}
}

func (uc *UpdateColumn) Execute(ctx context.Context, columnID uint, newTitle string) (*models.Column, error) {
	
	if columnID == 0 {
		return nil, domainErrors.ErrColumnInvalidID
	}

	if newTitle == "" {
		return nil, domainErrors.ErrColumnInvalidTitle
	}

	column, err := uc.Repo.FindByID(ctx, columnID)
	if err != nil {
		return nil, domainErrors.ErrColumnNotFound
	}

	column.Title = newTitle

	if err := uc.Repo.Update(ctx, column); err != nil {
		return nil, fmt.Errorf("failed to update column with ID %d: %w", column.ID, err)
	}

	return column, nil
}