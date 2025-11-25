package board_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"context"
)

type CreateBoard struct {
	Repo repositories.BoardRepository
}

func NewCreateBoard(repo repositories.BoardRepository) *CreateBoard {
	return &CreateBoard{
		Repo: repo,
	}
}

func (uc *CreateBoard) Execute(ctx context.Context, title string) (*models.Board, error) {
	if title == "" {
		return nil, domainErrors.ErrBoardInvalidTitle
	}

	board := &models.Board{
		Title: title,
	}

	if err := uc.Repo.Create(ctx, board); err != nil {
		return nil, domainErrors.ErrBoardCreate
	}

	return board, nil
}
