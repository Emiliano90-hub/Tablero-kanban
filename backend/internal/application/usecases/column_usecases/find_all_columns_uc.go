package column_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"context"
)

type FindAllColumns struct {
	Repo repositories.ColumnRepository
}

func NewFindAllColumns(repo repositories.ColumnRepository) *FindAllColumns {
	return &FindAllColumns{
		Repo: repo,
	}
}

func (uc *FindAllColumns) Execute(ctx context.Context, boardID uint) ([]models.Column, error) {
	if boardID == 0 {
		return nil, domainErrors.ErrColumnInvalidBoardID
	}

	columns, err := uc.Repo.FindAllColumns(ctx, boardID)
	if err != nil {
		return nil, domainErrors.ErrColumnBoardMismatch
	}

	return columns, nil
}
