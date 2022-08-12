package models

import (
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	Score        float64
	SubmissionID uint // foreign key
}
