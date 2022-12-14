package db

import (
	"fmt"
	"time"

	"github.com/COMP4050/square-team-5/api/internal/pkg/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database interface {
	ResetDB() (Database, error)

	CreateUser(email, passwordHash string, role models.UserRole) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)

	CreateUnit(name string) (*models.Unit, error)
	GetAllUnits(from int) ([]*models.Unit, error)
	GetUnitByID(id string, fetchClasses bool) (*models.Unit, error)
	GetUnitByName(name string) (*models.Unit, error)

	CreateClass(name string, unitID uint) (*models.Class, error)
	GetAllClasses(from int) ([]*models.Class, error)
	GetClass(id string) (*models.Class, error)

	CreateAssignment(name string, dueDate int, classID uint) (*models.Assignment, error)
	GetAllAssignments(from int) ([]*models.Assignment, error)
	GetAssignment(id string) (*models.Assignment, error)
	GetAssignmentsForClass(classID uint) ([]*models.Assignment, error)

	CreateTest(name string, assignmentID uint) (*models.Test, error)
	GetAllTests(from int) ([]*models.Test, error)
	GetTest(id string) (*models.Test, error)
	GetTestsForAssignment(assignmentID string) ([]*models.Test, error)

	CreateSubmission(studentID string, assignmentID uint) (*models.Submission, error)
	GetAllSubmissions(from int) ([]*models.Submission, error)
	GetSubmission(id string) (*models.Submission, error)
	GetSubmissionsForAssignment(assignmentID string) ([]*models.Submission, error)

	CreateResult(score float64, submissionID uint) (*models.Result, error)
	GetAllResults(from int) ([]*models.Result, error)
	GetResult(id string) (*models.Result, error)
}

type database struct {
	client   *gorm.DB
	filePath string
}

const PAGE_SIZE = 50

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound

	allModels = []interface{}{
		&models.Unit{},
		&models.Class{},
		&models.Assignment{},
		&models.Test{},
		&models.Submission{},
		&models.User{},
	}
)

func NewDB(dbFilePath string) Database {
	db, err := gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	for _, model := range allModels {
		err = db.AutoMigrate(model)
		if err != nil {
			panic("failed to migrate database")
		}
	}

	return &database{client: db, filePath: dbFilePath}
}

func (db *database) ResetDB() (Database, error) {
	for _, model := range allModels {
		err := db.client.Migrator().DropTable(model)
		if err != nil {
			return db, fmt.Errorf("error dropping table: %w", err)
		}
	}

	return NewDB(db.filePath), nil
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

func (db *database) GetAllUnits(from int) ([]*models.Unit, error) {
	var units []*models.Unit
	tx := db.client.Where("id >= ? AND id < ?", from, from+PAGE_SIZE).Find(&units)
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

func (db *database) GetAllClasses(from int) ([]*models.Class, error) {
	var classes []*models.Class
	tx := db.client.Where("id >= ? AND id < ?", from, from+PAGE_SIZE).Find(&classes)
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

func (db *database) CreateAssignment(name string, dueDate int, classID uint) (*models.Assignment, error) {
	assignment := models.Assignment{Name: name, DueDate: time.Unix(int64(dueDate), 0), ClassID: classID}
	tx := db.client.Create(&assignment)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &assignment, nil
}

func (db *database) GetAllAssignments(from int) ([]*models.Assignment, error) {
	var assignments []*models.Assignment
	tx := db.client.Where("id >= ? AND id < ?", from, from+PAGE_SIZE).Find(&assignments)
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

func (db *database) GetAssignmentsForClass(classID uint) ([]*models.Assignment, error) {
	var assignments []*models.Assignment
	tx := db.client.Where("class_id = ?", classID).Find(&assignments)
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

func (db *database) GetAllTests(from int) ([]*models.Test, error) {
	var tests []*models.Test
	tx := db.client.Where("id >= ? AND id < ?", from, from+PAGE_SIZE).Find(&tests)
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
	tx := db.client.Where("assignment_id = ?", assignmentID).Find(&tests)
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

func (db *database) GetAllSubmissions(from int) ([]*models.Submission, error) {
	var submissions []*models.Submission
	tx := db.client.Where("id >= ? AND id < ?", from, from+PAGE_SIZE).Find(&submissions)
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
	tx := db.client.Where("assignment_id = ?", assignmentID).Find(&submissions)
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

func (db *database) GetAllResults(from int) ([]*models.Result, error) {
	var results []*models.Result
	tx := db.client.Where("id >= ? AND id < ?", from, from+PAGE_SIZE).Find(&results)
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
