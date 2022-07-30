package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database interface {
	CreateUnit(name string) (*Unit, error)
	GetAllUnits() ([]*Unit, error)
	GetUnitByID(id string, fetchClasses bool) (*Unit, error)
	GetUnitByName(name string) (*Unit, error)

	CreateClass(name string, unitID uint) (*Class, error)
	GetAllClasses() ([]*Class, error)
	GetClass(id string) (*Class, error)

	CreateAssignment(name string, classID uint) (*Assignment, error)
	GetAllAssignments() ([]*Assignment, error)
	GetAssignment(id string) (*Assignment, error)

	CreateTest(name string, assignmentID uint) (*Test, error)
	GetAllTests() ([]*Test, error)
	GetTest(id string) (*Test, error)

	CreateSubmission(studentID string, assignmentID uint) (*Submission, error)
	GetAllSubmissions() ([]*Submission, error)
	GetSubmission(id string) (*Submission, error)

	CreateResult(score float64, submissionID uint) (*Result, error)
	GetAllResults() ([]*Result, error)
	GetResult(id string) (*Result, error)
}

type database struct {
	client *gorm.DB
}

var ErrRecordNotFound = gorm.ErrRecordNotFound

func NewDB() Database {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(
		&Unit{},
		&Class{},
		&Assignment{},
		&Test{},
		&Submission{},
	)
	if err != nil {
		panic("failed to migrate database")
	}

	return &database{client: db}

	// // Create
	// test := models.Test{Name: "Tests 1"}
	// db.Create(&test)

	// assignment := models.Assignment{
	// 	Name:    "Assignment 1",
	// 	DueDate: time.Now().Add(time.Hour * 24 * 7),
	// 	Tests:   []models.Test{test},
	// }
	// db.Create(&assignment)

	// class := models.Class{
	// 	Name:        "Tutorial 1",
	// 	Assignments: []models.Assignment{assignment},
	// }
	// db.Create(&class)

	// db.Create(&models.Unit{
	// 	Name:    "COMP1000",
	// 	Classes: []models.Class{class},
	// })

	// // Read
	// var unit models.Unit
	// db.Preload("Classes").First(&unit)

	// fmt.Printf("%+v", unit)

	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(models.Unit{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
}

func (db *database) CreateUnit(name string) (*Unit, error) {
	unit := Unit{Name: name}
	tx := db.client.Create(&unit)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &unit, nil
}

func (db *database) GetAllUnits() ([]*Unit, error) {
	var units []*Unit
	tx := db.client.Find(&units)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return units, nil
}

func (db *database) GetUnitByID(id string, fetchClasses bool) (*Unit, error) {
	var unit Unit

	if fetchClasses {
		tx := db.client.Preload("Classes").First(&unit, id)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx := db.client.First(&unit, id)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return &unit, nil
}

func (db *database) GetUnitByName(name string) (*Unit, error) {
	var unit Unit

	tx := db.client.Where("name = ?", name).First(&unit)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &unit, nil
}

func (db *database) CreateClass(name string, unitID uint) (*Class, error) {
	class := Class{Name: name, UnitID: unitID}
	tx := db.client.Create(&class)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &class, nil
}

func (db *database) GetAllClasses() ([]*Class, error) {
	var classes []*Class
	tx := db.client.Find(&classes)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return classes, nil
}

func (db *database) GetClass(id string) (*Class, error) {
	var class Class
	tx := db.client.First(&class, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &class, nil
}

func (db *database) CreateAssignment(name string, classID uint) (*Assignment, error) {
	assignment := Assignment{Name: name, ClassID: classID}
	tx := db.client.Create(&assignment)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &assignment, nil
}

func (db *database) GetAllAssignments() ([]*Assignment, error) {
	var assignments []*Assignment
	tx := db.client.Find(&assignments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return assignments, nil
}

func (db *database) GetAssignment(id string) (*Assignment, error) {
	var assignment Assignment
	tx := db.client.First(&assignment, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &assignment, nil
}

func (db *database) CreateTest(name string, assignmentID uint) (*Test, error) {
	test := Test{Name: name, AssignmentID: assignmentID}
	tx := db.client.Create(&test)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &test, nil
}

func (db *database) GetAllTests() ([]*Test, error) {
	var tests []*Test
	tx := db.client.Find(&tests)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tests, nil
}

func (db *database) GetTest(id string) (*Test, error) {
	var test Test
	tx := db.client.First(&test, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &test, nil
}

func (db *database) CreateSubmission(studentID string, assignmentID uint) (*Submission, error) {
	submission := Submission{StudentID: studentID, AssignmentID: assignmentID}
	tx := db.client.Create(&submission)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &submission, nil
}

func (db *database) GetAllSubmissions() ([]*Submission, error) {
	var submissions []*Submission
	tx := db.client.Find(&submissions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return submissions, nil
}

func (db *database) GetSubmission(id string) (*Submission, error) {
	var submission Submission
	tx := db.client.First(&submission, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &submission, nil
}

func (db *database) CreateResult(score float64, submissionID uint) (*Result, error) {
	result := Result{Score: score, SubmissionID: submissionID}
	tx := db.client.Create(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &result, nil
}

func (db *database) GetAllResults() ([]*Result, error) {
	var results []*Result
	tx := db.client.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return results, nil
}

func (db *database) GetResult(id string) (*Result, error) {
	var result Result
	tx := db.client.First(&result, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &result, nil
}
