package board_usecases

import (
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"context"
	"fmt"
)

type FindAllBoards struct {
	Repo repositories.BoardRepository
}

func NewFindAllBoards(repo repositories.BoardRepository) *FindAllBoards {
	return &FindAllBoards{
		Repo: repo,
	}
}

func (uc *FindAllBoards) Execute(ctx context.Context) ([]models.Board, error) {
	boards, err := uc.Repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch boards: %w", err)
	}

	return  boards, nil
}