package board_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"context"
)

type FindBoardByID struct {
	Repo repositories.BoardRepository
}

func NewFindBoardByID(repo repositories.BoardRepository) *FindBoardByID {
	return &FindBoardByID{
		Repo: repo,
	}
}

func (uc *FindBoardByID) Execute(ctx context.Context, boardID uint) (*models.Board, error) {
	if boardID == 0 {
		return nil, domainErrors.ErrBoardInvalidID
	}
	
	board, err := uc.Repo.FindByID(ctx, boardID)
	if err != nil {
		return nil, domainErrors.ErrBoardNotFound
	}

	return  board, nil
}