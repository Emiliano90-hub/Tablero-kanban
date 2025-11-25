package board_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/repositories"
	"context"
	"fmt"
)

type DeleteBoard struct {
	Repo repositories.BoardRepository
}

func NewDeleteBoard(repo repositories.BoardRepository) *DeleteBoard {
	return &DeleteBoard{
		Repo: repo,
	}
}

func (uc *DeleteBoard) Execute(ctx context.Context, boardID uint) error {
	if boardID == 0 {
		return domainErrors.ErrBoardInvalidID
	}
	
	board, err := uc.Repo.FindByID(ctx, boardID)
	if err != nil {
		return domainErrors.ErrBoardNotFound
	}

	if err := uc.Repo.Delete(ctx, board.ID); err != nil {
		return fmt.Errorf("failed to delete board with ID %d: %w", board.ID, err)
	}

	return nil

}
