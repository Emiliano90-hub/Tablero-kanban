package board_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"context"
	"fmt"
)

type UpdateBoard struct {
	Repo repositories.BoardRepository
}

func NewUpdateBoard(repo repositories.BoardRepository) *UpdateBoard {
	return &UpdateBoard{
		Repo: repo,
	}
}

func (uc *UpdateBoard) Execute(ctx context.Context, boardID uint, newTitle string) (*models.Board, error) {
	
	if boardID == 0 {
		return nil, domainErrors.ErrBoardInvalidID
	}

	if newTitle == "" {
		return nil, domainErrors.ErrBoardInvalidTitle
	}

	board, err := uc.Repo.FindByID(ctx, boardID)
	if err != nil {
		return nil, domainErrors.ErrBoardNotFound
	}

	board.Title = newTitle

	if err := uc.Repo.Update(ctx, board); err != nil {
		return nil, fmt.Errorf("failed to update board with ID %d: %w", board.ID, err)
	}

	return board, nil
}