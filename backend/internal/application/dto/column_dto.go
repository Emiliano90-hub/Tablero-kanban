package dto

type CreateColumnDTO struct {
	Title   string `json:"title" binding:"required"`
	BoardID uint   `json:"board_id" binding:"required"`
	Index   int    `json:"index"`
}

type ColumnResponseDTO struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	BoardID uint   `json:"board_id"`
	Index   int    `json:"index"`
}
