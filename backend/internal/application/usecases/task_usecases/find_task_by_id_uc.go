package task_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"context"
)

type FindTaskByID struct {
	Repo repositories.TaskRepository
}

func NewFindTaskByID(repo repositories.TaskRepository) *FindTaskByID {
	return &FindTaskByID{
		Repo: repo,
	}
}

func (uc *FindTaskByID) Execute(ctx context.Context, taskID uint) (*models.Task, error) {
	if taskID == 0 {
		return nil, domainErrors.ErrTaskInvalidID
	}

	task, err := uc.Repo.FindByID(ctx, taskID)
	if err != nil {
		return nil, domainErrors.ErrTaskNotFound
	}

	return task, nil
}