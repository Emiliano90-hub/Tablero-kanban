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

type BoardHandler struct {
	UC *usecases.BoardUC
}

func NewBoardHandler(uc *usecases.BoardUC) *BoardHandler {
	return &BoardHandler{UC: uc}
}

// POST /boards
func (h *BoardHandler) CreateBoard(c *gin.Context) {
	var body dto.CreateBoardDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	board, err := h.UC.CreateBoardUC.Execute(c.Request.Context(), body.Title)
	if err != nil {
		if errors.Is(err, domainErrors.ErrBoardInvalidTitle) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create board"})
		return
	}

	c.JSON(http.StatusCreated, dto.BoardResponseDTO{
		ID:   board.ID,
		Title: board.Title,
	})
}

// GET /boards/:id
func (h *BoardHandler) GetBoardByID(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	board, err := h.UC.FindBoardByIdUC.Execute(c.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, domainErrors.ErrBoardNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, dto.BoardResponseDTO{
		ID:   board.ID,
		Title: board.Title,
	})
}

// DELETE /boards/:id
func (h *BoardHandler) DeleteBoard(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.UC.DeleteBoardUC.Execute(c.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, domainErrors.ErrBoardNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
