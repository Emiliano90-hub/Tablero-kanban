package usecases

import (
	board_usecases "backend/internal/application/usecases/board_usecases"
)

type BoardUC struct {
	CreateBoardUC   board_usecases.CreateBoard
	UpdateBoardUC   board_usecases.UpdateBoard
	DeleteBoardUC   board_usecases.DeleteBoard
	FindBoardByIdUC board_usecases.FindBoardByID
	FindAllBoardsUC   board_usecases.FindAllBoards
}

func NewBoardUC(
	createBoard board_usecases.CreateBoard,
	updateBoard board_usecases.UpdateBoard,
	deleteBoard board_usecases.DeleteBoard,
	findBoardById board_usecases.FindBoardByID,
	findAllBoards board_usecases.FindAllBoards,

) *BoardUC {
	return &BoardUC{
		CreateBoardUC:   createBoard,
		UpdateBoardUC:   updateBoard,
		DeleteBoardUC:   deleteBoard,
		FindBoardByIdUC: findBoardById,
		FindAllBoardsUC:   findAllBoards,
	}
}
