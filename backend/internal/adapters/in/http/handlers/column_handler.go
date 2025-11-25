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

type ColumnHandler struct {
	UC *usecases.ColumnUC
}

func NewColumnHandler(uc *usecases.ColumnUC) *ColumnHandler {
	return &ColumnHandler{UC: uc}
}

// POST /columns
func (h *ColumnHandler) CreateColumn(c *gin.Context) {
	var body dto.CreateColumnDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	column, err := h.UC.CreateColumnUC.Execute(c.Request.Context(), body.Title, body.BoardID, body.Index)
	if err != nil {
		if errors.Is(err, domainErrors.ErrColumnInvalidTitle) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create column"})
		return
	}

	c.JSON(http.StatusCreated, dto.ColumnResponseDTO{
		ID:   column.ID,
		Title: column.Title,
		Index: column.Index,
	})
}

// GET /columns/:id
func (h *ColumnHandler) GetColumnByID(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	column, err := h.UC.FindColumnByIdUC.Execute(c.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, domainErrors.ErrColumnNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, dto.ColumnResponseDTO{
		ID:   column.ID,
		Title: column.Title,
		Index: column.Index,
	})
}

// DELETE /columns/:id
func (h *ColumnHandler) DeleteColumn(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.UC.DeleteColumnUC.Execute(c.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, domainErrors.ErrColumnNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
