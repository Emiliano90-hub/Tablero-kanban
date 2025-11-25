package http

import (
	"backend/internal/application/dto"
	"backend/internal/application/usecases"
	domainErrors "backend/internal/domain/errors"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	UC *usecases.TaskUC
}

func NewTaskHandler(uc *usecases.TaskUC) *TaskHandler {
	return &TaskHandler{UC: uc}
}

// POST /columns
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var body dto.CreateTaskDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	task, err := h.UC.CreateTaskUC.Execute(c.Request.Context(), body.Title, body.ColumnID, body.Index)
	if err != nil {
		if errors.Is(err, domainErrors.ErrTaskInvalidTitle) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, dto.TaskResponseDTO{
		ID:   task.ID,
		Title: task.Title,
		Index: task.Index,
	})
}

// GET /columns/:id
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	task, err := h.UC.FindTaskByIdUC.Execute(c.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, domainErrors.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, dto.TaskResponseDTO{
		ID:   task.ID,
		Title: task.Title,
		Index: task.Index,
	})
}

// DELETE /columns/:id
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.UC.DeleteTaskUC.Execute(c.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, domainErrors.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
