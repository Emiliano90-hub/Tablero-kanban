package column_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/repositories"
	"context"
	"fmt"
)

type DeleteColumn struct {
	Repo repositories.ColumnRepository
}

func NewDeleteColumn(repo repositories.ColumnRepository) *DeleteColumn {
	return &DeleteColumn{
		Repo: repo,
	}
}

func (uc *DeleteColumn) Execute(ctx context.Context, columnID uint) error {
	if columnID == 0 {
		return domainErrors.ErrColumnInvalidID
	}
	
	column, err := uc.Repo.FindByID(ctx, columnID)
	if err != nil {
		return domainErrors.ErrColumnNotFound
	}

	oldIndex := column.Index
	boardID := column.BoardID

	if err := uc.Repo.Delete(ctx, column.ID); err != nil {
		return fmt.Errorf("failed to delete column with ID %d: %w", column.ID, err)
	}

	columns, err := uc.Repo.FindAllColumns(ctx, boardID)
	if err != nil {
		return domainErrors.ErrColumnBoardMismatch
	}

	for i := range columns {
		if columns[i].Index > oldIndex {
			columns[i].Index--
			_ = uc.Repo.Update(ctx, &columns[i])
		}
	}

	return nil
}
