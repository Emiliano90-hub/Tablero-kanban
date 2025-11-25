package errors

import "errors"

// Board
var (
	ErrBoardCreate        = errors.New("fail to create board")
	ErrBoardNotFound      = errors.New("board not found")
	ErrBoardAlreadyExists = errors.New("board already exists")
	ErrBoardInvalidID    = errors.New("invalid board ID")
	ErrBoardInvalidTitle  = errors.New("invalid board title")
)

// Column
var (
	ErrColumnCreate         = errors.New("fail to create column")
	ErrColumnNotFound       = errors.New("column not found")
	ErrColumnInvalidBoardID = errors.New("invalid board ID")
	ErrColumnAlreadyExists  = errors.New("column already exists")
	ErrColumnInvalidID      = errors.New("invalid column ID")
	ErrColumnInvalidTitle   = errors.New("invalid column title")
	ErrColumnBoardMismatch  = errors.New("column does not belong to the specified board")
)

// Task
var (
	ErrTaskCreate          = errors.New("fail to create task")
	ErrTaskNotFound        = errors.New("task not found")
	ErrTaskInvalidColumnID = errors.New("invalid column ID")
	ErrTaskAlreadyExists   = errors.New("task already exists")
	ErrTaskInvalidID       = errors.New("invalid task ID")
	ErrTaskInvalidTitle    = errors.New("invalid task title")
	ErrTaskInvalidIndex    = errors.New("invalid task index")
	ErrTaskColumnMismatch  = errors.New("task does not belong to the specified column")
)
