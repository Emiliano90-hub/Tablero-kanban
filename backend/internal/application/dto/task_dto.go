package dto

type CreateTaskDTO struct {
	Title    string `json:"title" binding:"required"`
	ColumnID uint   `json:"column_id" binding:"required"`
	Index    int    `json:"index"`
}

type MoveTaskDTO struct {
	NewColumnID uint `json:"new_column_id" binding:"required"`
	NewIndex    int  `json:"new_index"`
}

type TaskResponseDTO struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	ColumnID uint   `json:"column_id"`
	Index    int    `json:"index"`
}
