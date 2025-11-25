package usecases

import (
	task_usecases "backend/internal/application/usecases/task_usecases"
)

type TaskUC struct {
	CreateTaskUC   task_usecases.CreateTask
	UpdateTaskUC   task_usecases.UpdateTask
	DeleteTaskUC   task_usecases.DeleteTask
	FindTaskByIdUC task_usecases.FindTaskByID
	FindAllTasksUC task_usecases.FindAllTasks
}

func NewTaskUC(
	createTask task_usecases.CreateTask,
	updateTask task_usecases.UpdateTask,
	deleteTask task_usecases.DeleteTask,
	findTaskById task_usecases.FindTaskByID,
	findAllTasks task_usecases.FindAllTasks,

) *TaskUC {
	return &TaskUC{
		CreateTaskUC:   createTask,
		UpdateTaskUC:   updateTask,
		DeleteTaskUC:   deleteTask,
		FindTaskByIdUC: findTaskById,
		FindAllTasksUC: findAllTasks,
	}
}
