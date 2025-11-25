package task_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"context"
)

type FindAllTasks struct {
	Repo repositories.TaskRepository
}

func NewFindAllTask(repo repositories.TaskRepository) *FindAllTasks {
	return &FindAllTasks{
		Repo: repo,
	}
}

func (uc *FindAllTasks) Execute(ctx context.Context, columnID uint) ([]models.Task, error) {
	if columnID == 0{
		return nil, domainErrors.ErrTaskInvalidColumnID
	}

	tasks, err := uc.Repo.FindAllTasks(ctx, columnID)
	if err != nil {
		return nil, domainErrors.ErrTaskColumnMismatch
	}	
	
	return tasks, nil
}