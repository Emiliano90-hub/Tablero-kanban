package http

import (
	"net/http"
	"strconv"

	"backend/internal/application/dto"
	"backend/internal/application/usecases"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, boardUC *usecases.BoardUC, columnUC *usecases.ColumnUC, taskUC *usecases.TaskUC) {
	boards := r.Group("/boards")
	{
		// Crear un nuevo board
		boards.POST("", func(c *gin.Context) {
			var boardDTO dto.CreateBoardDTO
			if err := c.ShouldBindJSON(&boardDTO); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			board, err := boardUC.CreateBoardUC.Execute(c.Request.Context(), boardDTO.Title)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, dto.BoardResponseDTO{
				ID:    board.ID,
				Title: board.Title,
			})
		})

		boards.GET("/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
				return
			}

			board, err := boardUC.FindBoardByIdUC.Execute(c.Request.Context(), uint(id))
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, dto.BoardResponseDTO{
				ID:    board.ID,
				Title: board.Title,
			})
		})

		// Listar todos los boards
		boards.GET("", func(c *gin.Context) {
			list, err := boardUC.FindAllBoardsUC.Execute(c.Request.Context())
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			boardsResponse := make([]dto.BoardResponseDTO, len(list))
			for i, b := range list {
				boardsResponse[i] = dto.BoardResponseDTO{
					ID:    b.ID,
					Title: b.Title,
				}
			}

			c.JSON(http.StatusOK, boardsResponse)
		})
	}

	columns := r.Group("/columns")
	{
		// Crear columna
		columns.POST("", func(c *gin.Context) {
			var columnDTO dto.CreateColumnDTO
			if err := c.ShouldBindJSON(&columnDTO); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			column, err := columnUC.CreateColumnUC.Execute(c.Request.Context(), columnDTO.Title, columnDTO.BoardID, columnDTO.Index)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, gin.H{
				"id":      column.ID,
				"title":   column.Title,
				"boardId": column.BoardID,
			})
		})

		// Obtener columnas por board
		columns.GET("/:board_id", func(c *gin.Context) {
			boardID, _ := strconv.ParseUint(c.Param("board_id"), 10, 64)

			list, err := columnUC.FindAllColumnsUC.Execute(c.Request.Context(), uint(boardID))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, list)
		})
	}
	tasks := r.Group("/tasks")
	{
		// Crear task
		tasks.POST("", func(c *gin.Context) {
			var taskDTO dto.CreateTaskDTO
			if err := c.ShouldBindJSON(&taskDTO); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			task, err := taskUC.CreateTaskUC.Execute(c.Request.Context(), taskDTO.Title, taskDTO.ColumnID, taskDTO.Index)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, task)
		})

		// Obtener tasks por columna
		tasks.GET("/:column_id", func(c *gin.Context) {
			columnID, _ := strconv.ParseUint(c.Param("column_id"), 10, 64)

			list, err := taskUC.FindAllTasksUC.Execute(c.Request.Context(), uint(columnID))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, list)
		})

		// Eliminar task
		tasks.DELETE("/:id", func(c *gin.Context) {
			id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

			if err := taskUC.DeleteTaskUC.Execute(c.Request.Context(), uint(id)); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"status": "deleted"})
		})
	}
}
