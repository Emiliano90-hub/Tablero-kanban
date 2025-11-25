package column_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"context"
)

type FindColumnByID struct {
	Repo repositories.ColumnRepository
}

func NewFindColumnByID(repo repositories.ColumnRepository) *FindColumnByID {
	return &FindColumnByID{
		Repo: repo,
	}
}

func (uc *FindColumnByID) Execute(ctx context.Context, columnID uint) (*models.Column, error) {
	if columnID == 0 {
		return nil, domainErrors.ErrColumnInvalidID
	}
	
	column, err := uc.Repo.FindByID(ctx, columnID)
	if err != nil {
		return nil, domainErrors.ErrColumnNotFound
	}

	return  column, nil
}