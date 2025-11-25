package dto

type CreateBoardDTO struct {
	Title string `json:"title" binding:"required"`
}

type BoardResponseDTO struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
