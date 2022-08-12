package models

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	Name        string
	DueDate     time.Time
	Tests       []Test
	Submissions []Submission
	ClassID     uint // foreign key
}
