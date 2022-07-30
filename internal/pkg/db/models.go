package db

import (
	"time"

	"gorm.io/gorm"
)

type Unit struct {
	gorm.Model
	Name    string
	Classes []Class
}

type Class struct {
	gorm.Model
	Name        string
	Assignments []Assignment
	UnitID      uint // foreign key
}

type Assignment struct {
	gorm.Model
	Name        string
	DueDate     time.Time
	Tests       []Test
	Submissions []Submission
	ClassID     uint // foreign key
}

type Test struct {
	gorm.Model
	Name         string
	AssignmentID uint // foreign key
}

type Submission struct {
	gorm.Model
	StudentID    string
	Result       Result
	AssignmentID uint // foreign key
}

type Result struct {
	gorm.Model
	Score        float64
	SubmissionID uint // foreign key
}
