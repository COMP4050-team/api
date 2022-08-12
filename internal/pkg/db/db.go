package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/COMP4050/square-team-5/api/internal/pkg/db/models"
)

type Database interface {
	CreateUser(email, passwordHash string, role models.UserRole) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)

	CreateUnit(name string) (*models.Unit, error)
	GetAllUnits() ([]*models.Unit, error)
	GetUnitByID(id string, fetchClasses bool) (*models.Unit, error)
	GetUnitByName(name string) (*models.Unit, error)

	CreateClass(name string, unitID uint) (*models.Class, error)
	GetAllClasses() ([]*models.Class, error)
	GetClass(id string) (*models.Class, error)

	CreateAssignment(name string, classID uint) (*models.Assignment, error)
	GetAllAssignments() ([]*models.Assignment, error)
	GetAssignment(id string) (*models.Assignment, error)
	GetAssignmentsForClass(classID string) ([]*models.Assignment, error)

	CreateTest(name string, assignmentID uint) (*models.Test, error)
	GetAllTests() ([]*models.Test, error)
	GetTest(id string) (*models.Test, error)
	GetTestsForAssignment(assignmentID string) ([]*models.Test, error)

	CreateSubmission(studentID string, assignmentID uint) (*models.Submission, error)
	GetAllSubmissions() ([]*models.Submission, error)
	GetSubmission(id string) (*models.Submission, error)
	GetSubmissionsForAssignment(assignmentID string) ([]*models.Submission, error)

	CreateResult(score float64, submissionID uint) (*models.Result, error)
	GetAllResults() ([]*models.Result, error)
	GetResult(id string) (*models.Result, error)
}

type database struct {
	client *gorm.DB
}

var ErrRecordNotFound = gorm.ErrRecordNotFound

func NewDB() Database {
	db, err := gorm.Open(sqlite.Open("/app/data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(
		&models.Unit{},
		&models.Class{},
		&models.Assignment{},
		&models.Test{},
		&models.Submission{},
		&models.User{},
	)
	if err != nil {
		panic("failed to migrate database")
	}

	return &database{client: db}
}

func (db *database) CreateUser(email, passwordHash string, role models.UserRole) (*models.User, error) {
	user := models.User{Email: email, PasswordHash: passwordHash, Role: role}
	tx := db.client.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (db *database) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	tx := db.client.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (db *database) CreateUnit(name string) (*models.Unit, error) {
	unit := models.Unit{Name: name}
	tx := db.client.Create(&unit)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &unit, nil
}

func (db *database) GetAllUnits() ([]*models.Unit, error) {
	var units []*models.Unit
	tx := db.client.Find(&units)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return units, nil
}

func (db *database) GetUnitByID(id string, fetchClasses bool) (*models.Unit, error) {
	var unit models.Unit

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

func (db *database) GetUnitByName(name string) (*models.Unit, error) {
	var unit models.Unit

	tx := db.client.Where("name = ?", name).First(&unit)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &unit, nil
}

func (db *database) CreateClass(name string, unitID uint) (*models.Class, error) {
	class := models.Class{Name: name, UnitID: unitID}
	tx := db.client.Create(&class)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &class, nil
}

func (db *database) GetAllClasses() ([]*models.Class, error) {
	var classes []*models.Class
	tx := db.client.Find(&classes)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return classes, nil
}

func (db *database) GetClass(id string) (*models.Class, error) {
	var class models.Class
	tx := db.client.First(&class, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &class, nil
}

func (db *database) CreateAssignment(name string, classID uint) (*models.Assignment, error) {
	assignment := models.Assignment{Name: name, ClassID: classID}
	tx := db.client.Create(&assignment)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &assignment, nil
}

func (db *database) GetAllAssignments() ([]*models.Assignment, error) {
	var assignments []*models.Assignment
	tx := db.client.Find(&assignments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return assignments, nil
}

func (db *database) GetAssignment(id string) (*models.Assignment, error) {
	var assignment models.Assignment
	tx := db.client.First(&assignment, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &assignment, nil
}

func (db *database) GetAssignmentsForClass(classID string) ([]*models.Assignment, error) {
	var assignments []*models.Assignment
	tx := db.client.Find(&assignments).Where("class_id = ?", classID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return assignments, nil
}

func (db *database) CreateTest(name string, assignmentID uint) (*models.Test, error) {
	test := models.Test{Name: name, AssignmentID: assignmentID}
	tx := db.client.Create(&test)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &test, nil
}

func (db *database) GetAllTests() ([]*models.Test, error) {
	var tests []*models.Test
	tx := db.client.Find(&tests)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tests, nil
}

func (db *database) GetTest(id string) (*models.Test, error) {
	var test models.Test
	tx := db.client.First(&test, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &test, nil
}

func (db *database) GetTestsForAssignment(assignmentID string) ([]*models.Test, error) {
	var tests []*models.Test
	tx := db.client.Find(&tests).Where("assignment_id = ?", assignmentID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tests, nil
}

func (db *database) CreateSubmission(studentID string, assignmentID uint) (*models.Submission, error) {
	submission := models.Submission{StudentID: studentID, AssignmentID: assignmentID}
	tx := db.client.Create(&submission)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &submission, nil
}

func (db *database) GetAllSubmissions() ([]*models.Submission, error) {
	var submissions []*models.Submission
	tx := db.client.Find(&submissions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return submissions, nil
}

func (db *database) GetSubmission(id string) (*models.Submission, error) {
	var submission models.Submission
	tx := db.client.First(&submission, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &submission, nil
}

func (db *database) GetSubmissionsForAssignment(assignmentID string) ([]*models.Submission, error) {
	var submissions []*models.Submission
	tx := db.client.Find(&submissions).Where("assignment_id = ?", assignmentID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return submissions, nil
}

func (db *database) CreateResult(score float64, submissionID uint) (*models.Result, error) {
	result := models.Result{Score: score, SubmissionID: submissionID}
	tx := db.client.Create(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &result, nil
}

func (db *database) GetAllResults() ([]*models.Result, error) {
	var results []*models.Result
	tx := db.client.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return results, nil
}

func (db *database) GetResult(id string) (*models.Result, error) {
	var result models.Result
	tx := db.client.First(&result, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &result, nil
}
