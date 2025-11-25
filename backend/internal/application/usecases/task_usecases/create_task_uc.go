package task_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"context"
)

type CreateTask struct {
	Repo repositories.TaskRepository
}

func NewCreateTask(repo repositories.TaskRepository) *CreateTask {
	return &CreateTask{
		Repo: repo,
	}
}

func (uc *CreateTask) Execute(ctx context.Context, title string, columnID uint, index int) (*models.Task, error) {
	if title == "" {
		return nil, domainErrors.ErrTaskInvalidTitle
	}

	if columnID == 0 {
		return nil, domainErrors.ErrTaskInvalidColumnID
	}

	task := &models.Task{
		Title: title,
		ColumnID: columnID,
		Index: index,
	}

	if err := uc.Repo.Create(ctx, task); err != nil {
		return nil, domainErrors.ErrTaskCreate
	}

	return task, nil

}