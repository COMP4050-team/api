package models

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name        string
	Assignments []Assignment
	UnitID      uint // foreign key
}
