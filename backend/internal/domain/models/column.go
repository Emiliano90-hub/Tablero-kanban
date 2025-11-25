package models

import "gorm.io/gorm"

type Column struct {
	gorm.Model
	Title   string
	Index   int
	BoardID uint
	Task    []Task
}
