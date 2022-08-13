package models

import (
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	Name         string
	StoragePath  string
	AssignmentID uint // foreign key
}
