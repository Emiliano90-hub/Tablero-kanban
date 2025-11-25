package task_usecases

import (
	domainErrors "backend/internal/domain/errors"
	"backend/internal/domain/repositories"
	"context"
	"fmt"
)

type DeleteTask struct {
	Repo repositories.TaskRepository
}

func NewDeleteTask(repo repositories.TaskRepository) *DeleteTask {
	return &DeleteTask{
		Repo: repo,
	}
}

func (uc *DeleteTask) Execute(ctx context.Context, taskID uint) error {
	if taskID == 0 {
		return domainErrors.ErrTaskInvalidID
	}

	task, err := uc.Repo.FindByID(ctx, taskID)
	if err != nil {
		return domainErrors.ErrTaskNotFound
	}

	oldIndex := task.Index
	columnID := task.ColumnID

	if err := uc.Repo.Delete(ctx, task.ID); err != nil {
		return fmt.Errorf("failed to delete task with ID %d: %w", columnID, err)
	}

	tasks, err := uc.Repo.FindAllTasks(ctx, columnID)
	if err != nil {
		return domainErrors.ErrColumnBoardMismatch
	}

	for i := range tasks {
		if tasks[i].Index > oldIndex {
			tasks[i].Index--
			_ = uc.Repo.Update(ctx, &tasks[i])
		}
	}

	return nil
}
