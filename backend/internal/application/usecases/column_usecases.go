package usecases

import (
	column_usecases "backend/internal/application/usecases/column_usecases"
)

type ColumnUC struct {
	CreateColumnUC   column_usecases.CreateColumn
	UpdateColumnUC   column_usecases.UpdateColumn
	DeleteColumnUC   column_usecases.DeleteColumn
	FindColumnByIdUC column_usecases.FindColumnByID
	FindAllColumnsUC column_usecases.FindAllColumns
}

func NewColumnUC(
	createColumn column_usecases.CreateColumn,
	updateColumn column_usecases.UpdateColumn,
	deleteColumn column_usecases.DeleteColumn,
	findColumnById column_usecases.FindColumnByID,
	findAllColumns column_usecases.FindAllColumns,

) *ColumnUC {
	return &ColumnUC{
		CreateColumnUC:   createColumn,
		UpdateColumnUC:   updateColumn,
		DeleteColumnUC:   deleteColumn,
		FindColumnByIdUC: findColumnById,
		FindAllColumnsUC: findAllColumns,
	}
}
