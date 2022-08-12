package models

import (
	"gorm.io/gorm"
)

type Submission struct {
	gorm.Model
	StudentID    string
	Result       Result
	AssignmentID uint // foreign key
}
