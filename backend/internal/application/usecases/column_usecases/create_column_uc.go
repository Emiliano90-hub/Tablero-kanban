package column_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"context"
)

type CreateColumn struct {
	Repo repositories.ColumnRepository
}

func NewCreateColumn(repo repositories.ColumnRepository) *CreateColumn {
	return &CreateColumn{
		Repo: repo,
	}
}

func (uc *CreateColumn) Execute(ctx context.Context, title string, boardID uint, index int) (*models.Column, error) {
	if title == "" {
		return nil, domainErrors.ErrColumnInvalidTitle
	}

	if boardID == 0 {
		return nil, domainErrors.ErrColumnInvalidBoardID
	}

	column := &models.Column{
		Title:   title,
		BoardID: boardID,
		Index:   index,
	}

	if err := uc.Repo.Create(ctx, column); err != nil {
		return nil, domainErrors.ErrColumnCreate
	}

	return column, nil
}
