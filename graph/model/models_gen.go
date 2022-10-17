// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Assignment struct {
	ID          string        `json:"id"`
	Class       *Class        `json:"class"`
	Unit        *Unit         `json:"unit"`
	Name        string        `json:"name"`
	DueDate     int           `json:"dueDate"`
	Tests       []*Test       `json:"tests"`
	Submissions []*Submission `json:"submissions"`
}

type Class struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Unit        *Unit         `json:"unit"`
	Assignments []*Assignment `json:"assignments"`
}

type NewAssignment struct {
	Name    string `json:"name"`
	DueDate int    `json:"dueDate"`
	ClassID string `json:"classID"`
}

type NewClass struct {
	Name   string `json:"name"`
	UnitID string `json:"unitID"`
}

type NewSubmission struct {
	StudentID    string `json:"studentID"`
	AssignmentID string `json:"assignmentID"`
}

type NewTest struct {
	Name         string `json:"name"`
	AssignmentID string `json:"assignmentID"`
}

type NewUnit struct {
	Name string `json:"name"`
}

type Result struct {
	ID           string  `json:"id"`
	Score        float64 `json:"score"`
	Date         string  `json:"date"`
	SubmissionID string  `json:"submissionID"`
}

type Submission struct {
	ID         string      `json:"id"`
	StudentID  string      `json:"studentID"`
	Result     *Result     `json:"result"`
	Unit       *Unit       `json:"unit"`
	Class      *Class      `json:"class"`
	Assignment *Assignment `json:"assignment"`
}

type Test struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Unit       *Unit       `json:"unit"`
	Class      *Class      `json:"class"`
	Assignment *Assignment `json:"assignment"`
}

type Unit struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Classes []*Class `json:"classes"`
}
