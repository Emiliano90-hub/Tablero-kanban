package task_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"context"
)

type UpdateTask struct {
	Repo repositories.TaskRepository
}

func NewUpdateTask(repo repositories.TaskRepository) *UpdateTask {
	return &UpdateTask{
		Repo: repo,
	}
}

func (uc *UpdateTask) Execute(ctx context.Context, taskID uint, newTitle string) (*models.Task, error) {
	if taskID == 0 {
		return nil, domainErrors.ErrTaskInvalidID
	}

	if newTitle == "" {
		return nil, domainErrors.ErrTaskInvalidTitle
	}

	task, err := uc.Repo.FindByID(ctx, taskID)
	if err != nil {
		return nil, domainErrors.ErrTaskNotFound
	}

	return task, nil
}
