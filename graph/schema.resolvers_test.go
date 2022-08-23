package graph

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/COMP4050/square-team-5/api/fixtures/mocks"
	"github.com/COMP4050/square-team-5/api/graph/generated"
	"github.com/COMP4050/square-team-5/api/graph/model"
	"github.com/COMP4050/square-team-5/api/internal/pkg/db"
	"github.com/COMP4050/square-team-5/api/internal/pkg/db/models"
)

func TestRootResolver(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mocks.NewMockDatabase(ctrl)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: db}})))

	t.Run("introspection", func(t *testing.T) {
		// Make sure we can run the graphiql introspection query without errors
		var resp interface{}
		c.MustPost(introspection.Query, &resp)
	})
}

func TestUnitResolver(t *testing.T) {
	t.Parallel()

	t.Run("Get Unit", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUnitByID("1", false).Return(&models.Unit{Model: gorm.Model{ID: 1}, Name: "COMP1000"}, nil)
		mockDB.EXPECT().GetUnitByID("1", true).Return(&models.Unit{Model: gorm.Model{ID: 1}, Name: "COMP1000",
			Classes: []models.Class{{Model: gorm.Model{ID: 1}, Name: "Class 1"}},
		}, nil)

		var resp struct {
			Unit struct {
				ID, Name string
				Classes  []model.Class
			}
		}
		c.MustPost(`{ unit(id:"1") { id name classes { id name } } }`, &resp)

		assert.Equal(t, "1", resp.Unit.ID)
		assert.Equal(t, "COMP1000", resp.Unit.Name)
		assert.Equal(t, []model.Class{{ID: "1", Name: "Class 1"}}, resp.Unit.Classes)
	})

	t.Run("Get All Units", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetAllUnits(1).Return([]*models.Unit{
			{Model: gorm.Model{ID: 1}, Name: "COMP1000"},
			{Model: gorm.Model{ID: 2}, Name: "COMP1010"},
		}, nil)

		var resp struct {
			Units []struct{ ID, Name string }
		}
		c.MustPost(`{ units() { id name } }`, &resp)

		assert.Equal(t, "1", resp.Units[0].ID)
		assert.Equal(t, "2", resp.Units[1].ID)
		assert.Equal(t, "COMP1000", resp.Units[0].Name)
		assert.Equal(t, "COMP1010", resp.Units[1].Name)
	})

	t.Run("Get All Units - Error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		customErr := errors.New("custom error")
		mockDB.EXPECT().GetAllUnits(1).Return([]*models.Unit{
			{Model: gorm.Model{ID: 1}, Name: "COMP1000"},
			{Model: gorm.Model{ID: 2}, Name: "COMP1010"},
		}, customErr)

		var resp struct {
			Units []struct{ ID, Name string }
		}
		err := c.Post(`{ units() { id name } }`, &resp)

		assert.ErrorContains(t, err, customErr.Error())

		assert.Nil(t, resp.Units)
	})

	t.Run("Get Unit Not Found", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUnitByID("1", false).Return(nil, db.ErrRecordNotFound)

		var resp struct {
			Unit struct{ ID, Name string }
		}
		err := c.Post(`{ unit(id:"1") { id name } }`, &resp)

		assert.ErrorContains(t, err, "record not found")
		assert.NotEqual(t, "1", resp.Unit.ID)
	})

	t.Run("Create Unit", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUnitByName("COMP1000").Return(nil, db.ErrRecordNotFound)
		mockDB.EXPECT().CreateUnit("COMP1000").Return(&models.Unit{Model: gorm.Model{ID: 1}, Name: "COMP1000"}, nil)

		var resp struct {
			CreateUnit struct{ ID, Name string }
		}
		c.MustPost(`mutation { createUnit(input: {name: "COMP1000"}) { id name } }`, &resp)

		assert.Equal(t, "1", resp.CreateUnit.ID)
		assert.Equal(t, "COMP1000", resp.CreateUnit.Name)
	})

	t.Run("Create Unit - Already Exists", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUnitByName("COMP1000").Return(&models.Unit{Model: gorm.Model{ID: 1}, Name: "COMP1000"}, nil)

		var resp struct {
			CreateUnit struct{ ID, Name string }
		}
		err := c.Post(`mutation { createUnit(input: {name: "COMP1000"}) { id name } }`, &resp)

		assert.ErrorContains(t, err, "unit already exists")
		assert.NotEqual(t, "1", resp.CreateUnit.ID)
	})

	t.Run("Create Unit - Error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUnitByName("COMP1000").Return(&models.Unit{Model: gorm.Model{ID: 1}, Name: "COMP1000"}, fmt.Errorf("error"))
		// mockDB.EXPECT().CreateUnit("COMP1000").Return(&models.Unit{Model: gorm.Model{ID: 1}, Name: "COMP1000"}, nil)

		var resp struct {
			CreateUnit struct{ ID, Name string }
		}
		err := c.Post(`mutation { createUnit(input: {name: "COMP1000"}) { id name } }`, &resp)

		assert.ErrorContains(t, err, "error getting unit: error")
		assert.NotEqual(t, "1", resp.CreateUnit.ID)
	})
}

func TestClassResolver(t *testing.T) {
	t.Parallel()

	t.Run("Get Class", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetClass("1").Return(&models.Class{Model: gorm.Model{ID: 1}, Name: "Class 1"}, nil)
		mockDB.EXPECT().GetAssignmentsForClass(uint(1)).Return([]*models.Assignment{{Model: gorm.Model{ID: 1}, Name: "Assignment 1"}}, nil)

		var resp struct {
			Class struct {
				ID, Name    string
				Assignments []model.Assignment
			}
		}
		c.MustPost(`{ class(id:"1") { id name assignments { id name } } }`, &resp)

		assert.Equal(t, "1", resp.Class.ID)
		assert.Equal(t, "Class 1", resp.Class.Name)
	})

	t.Run("Get All Classes", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetAllClasses(1).Return([]*models.Class{
			{Model: gorm.Model{ID: 1}, Name: "Class 1", UnitID: 1},
			{Model: gorm.Model{ID: 2}, Name: "Class 2", UnitID: 1},
		}, nil)
		mockDB.EXPECT().GetAssignmentsForClass(uint(1)).Return([]*models.Assignment{{Model: gorm.Model{ID: 1}, Name: "Assignment 1"}}, nil)
		mockDB.EXPECT().GetAssignmentsForClass(uint(2)).Return([]*models.Assignment{{Model: gorm.Model{ID: 2}, Name: "Assignment 2"}}, nil)

		var resp struct {
			Classes []struct {
				ID, Name, UnitID string
				Assignments      []model.Assignment
			}
		}
		c.MustPost(`{ classes() { id name assignments { id name } unitID } }`, &resp)

		assert.Equal(t, "1", resp.Classes[0].ID)
		assert.Equal(t, "2", resp.Classes[1].ID)
		assert.Equal(t, "Class 1", resp.Classes[0].Name)
		assert.Equal(t, "Class 2", resp.Classes[1].Name)
		assert.Equal(t, []model.Assignment{{ID: "1", Name: "Assignment 1"}}, resp.Classes[0].Assignments)
		assert.Equal(t, []model.Assignment{{ID: "2", Name: "Assignment 2"}}, resp.Classes[1].Assignments)
		assert.Equal(t, "1", resp.Classes[0].UnitID)
		assert.Equal(t, "1", resp.Classes[1].UnitID)
	})

	t.Run("Get Class Not Found", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetClass("1").Return(nil, db.ErrRecordNotFound)

		var resp struct {
			Class struct{ ID, Name string }
		}
		err := c.Post(`{ class(id:"1") { id name } }`, &resp)

		assert.ErrorContains(t, err, "record not found")
		assert.NotEqual(t, "1", resp.Class.ID)
	})

	t.Run("Create Class", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUnitByID("1", false).Return(&models.Unit{Model: gorm.Model{ID: 1}, Name: "COMP1000"}, nil)
		mockDB.EXPECT().CreateClass("Class 1", uint(1)).Return(&models.Class{Model: gorm.Model{ID: 1}, Name: "Class 1"}, nil)

		var resp struct {
			CreateClass struct{ ID, Name string }
		}
		c.MustPost(`mutation { createClass(input: {name: "Class 1", unitID: "1"}) { id name } }`, &resp)

		assert.Equal(t, "1", resp.CreateClass.ID)
		assert.Equal(t, "Class 1", resp.CreateClass.Name)
	})
}

func TestAssignmentResolver(t *testing.T) {
	t.Parallel()

	t.Run("Get Assignment", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetAssignment("1").Return(&models.Assignment{Model: gorm.Model{ID: 1}, Name: "Assignment 1"}, nil)

		var resp struct {
			Assignment struct{ ID, Name string }
		}
		c.MustPost(`{ assignment(id:"1") { id name } }`, &resp)

		assert.Equal(t, "1", resp.Assignment.ID)
		assert.Equal(t, "Assignment 1", resp.Assignment.Name)
	})

	t.Run("Get All Assignments", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		dueDate := time.Now().Add(time.Hour * 24 * 7)

		mockDB.EXPECT().GetAllAssignments(1).Return([]*models.Assignment{
			{Model: gorm.Model{ID: 1}, Name: "Assignment 1", DueDate: dueDate, Tests: nil, Submissions: nil, ClassID: 1},
			{Model: gorm.Model{ID: 2}, Name: "Assignment 2", DueDate: dueDate, Tests: nil, Submissions: nil, ClassID: 1},
		}, nil)
		mockDB.EXPECT().GetSubmissionsForAssignment("1").Return([]*models.Submission{{Model: gorm.Model{ID: 1}}}, nil)
		mockDB.EXPECT().GetSubmissionsForAssignment("2").Return([]*models.Submission{{Model: gorm.Model{ID: 2}}}, nil)
		mockDB.EXPECT().GetTestsForAssignment("1").Return([]*models.Test{{Model: gorm.Model{ID: 1}}}, nil)
		mockDB.EXPECT().GetTestsForAssignment("2").Return([]*models.Test{{Model: gorm.Model{ID: 2}}}, nil)

		var resp struct {
			Assignments []struct {
				ID, Name    string
				DueDate     int
				Tests       []model.Test
				Submissions []model.Submission
			}
		}
		c.MustPost(`{ assignments() { id name dueDate tests { id } submissions { id } } }`, &resp)

		assert.Equal(t, "1", resp.Assignments[0].ID)
		assert.Equal(t, "2", resp.Assignments[1].ID)
		assert.Equal(t, "Assignment 1", resp.Assignments[0].Name)
		assert.Equal(t, "Assignment 2", resp.Assignments[1].Name)
		assert.Equal(t, int(dueDate.Unix()), resp.Assignments[0].DueDate)
		assert.Equal(t, int(dueDate.Unix()), resp.Assignments[1].DueDate)
		assert.Equal(t, []model.Test{{ID: "1"}}, resp.Assignments[0].Tests)
		assert.Equal(t, []model.Test{{ID: "2"}}, resp.Assignments[1].Tests)
		assert.Equal(t, []model.Submission{{ID: "1"}}, resp.Assignments[0].Submissions)
		assert.Equal(t, []model.Submission{{ID: "2"}}, resp.Assignments[1].Submissions)
	})

	t.Run("Get Assignment Not Found", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetAssignment("1").Return(nil, db.ErrRecordNotFound)

		var resp struct {
			Assignment struct{ ID, Name string }
		}
		err := c.Post(`{ assignment(id:"1") { id name } }`, &resp)

		assert.ErrorContains(t, err, "record not found")
		assert.NotEqual(t, "1", resp.Assignment.ID)
	})

	t.Run("Create Assignment", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().CreateAssignment("Assignment 1", 1660657596, uint(1)).Return(&models.Assignment{Model: gorm.Model{ID: 1}, Name: "Assignment 1"}, nil)

		var resp struct {
			CreateAssignment struct{ ID, Name string }
		}
		c.MustPost(`mutation { createAssignment(input: {name: "Assignment 1", dueDate: 1660657596, classID: "1"}) { id name } }`, &resp)

		assert.Equal(t, "1", resp.CreateAssignment.ID)
		assert.Equal(t, "Assignment 1", resp.CreateAssignment.Name)
	})
}

func TestTestResolver(t *testing.T) {
	t.Parallel()

	t.Run("Get Test", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetTest("1").Return(&models.Test{Model: gorm.Model{ID: 1}, Name: "Test 1"}, nil)

		var resp struct {
			Test struct{ ID, Name string }
		}
		c.MustPost(`{ test(id:"1") { id name } }`, &resp)

		assert.Equal(t, "1", resp.Test.ID)
		assert.Equal(t, "Test 1", resp.Test.Name)
	})

	t.Run("Get All Tests", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetAllTests(1).Return([]*models.Test{
			{Model: gorm.Model{ID: 1}, Name: "Test 1", AssignmentID: 1},
			{Model: gorm.Model{ID: 1}, Name: "Test 2", AssignmentID: 1},
		}, nil)

		var resp struct {
			Tests []struct {
				ID, Name, AssignmentID string
			}
		}
		c.MustPost(`{ tests() { id name assignmentID } }`, &resp)

		assert.Equal(t, "1", resp.Tests[0].ID)
		assert.Equal(t, "1", resp.Tests[1].ID)
		assert.Equal(t, "Test 1", resp.Tests[0].Name)
		assert.Equal(t, "Test 2", resp.Tests[1].Name)
		assert.Equal(t, "1", resp.Tests[0].AssignmentID)
		assert.Equal(t, "1", resp.Tests[1].AssignmentID)
	})

	t.Run("Get Test Not Found", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetTest("1").Return(nil, db.ErrRecordNotFound)

		var resp struct {
			Test struct{ ID, Name string }
		}
		err := c.Post(`{ test(id:"1") { id name } }`, &resp)

		assert.ErrorContains(t, err, "record not found")
		assert.NotEqual(t, "1", resp.Test.ID)
	})

	t.Run("Create Test", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		// Storage path here is tests/{assignmentID}/test_{testID}.java
		mockDB.EXPECT().CreateTest("Test 1", "tests/1/test_1.java", uint(1)).Return(&models.Test{Model: gorm.Model{ID: 1}, Name: "Test 1"}, nil)

		var resp struct {
			CreateTest struct{ ID, Name string }
		}
		c.MustPost(`mutation { createTest(input: {name: "Test 1", storagePath: "tests/1/test_1.java",assignmentID: "1"}) { id name } }`, &resp)

		assert.Equal(t, "1", resp.CreateTest.ID)
		assert.Equal(t, "Test 1", resp.CreateTest.Name)
	})
}

func TestSubmissionResolver(t *testing.T) {
	t.Parallel()

	t.Run("Get Submission", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetSubmission("1").Return(&models.Submission{Model: gorm.Model{ID: 1}, StudentID: "44444444"}, nil)

		var resp struct {
			Submission struct{ ID, StudentID string }
		}
		c.MustPost(`{ submission(id:"1") { id studentID} }`, &resp)

		assert.Equal(t, "1", resp.Submission.ID)
		assert.Equal(t, "44444444", resp.Submission.StudentID)
	})

	t.Run("Get All Submissions", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetAllSubmissions(1).Return([]*models.Submission{
			{Model: gorm.Model{ID: 1}, StudentID: "44444444", Result: models.Result{Model: gorm.Model{ID: 1}, Score: 99}, AssignmentID: 1},
			{Model: gorm.Model{ID: 2}, StudentID: "44444445", Result: models.Result{Model: gorm.Model{ID: 2}, Score: 51}, AssignmentID: 1},
		}, nil)
		mockDB.EXPECT().GetSubmission("1").Return(
			&models.Submission{
				Model:     gorm.Model{ID: 1},
				StudentID: "44444444",
				Result: models.Result{
					Model: gorm.Model{ID: 1},
					Score: 99,
				},
				AssignmentID: 1}, nil)
		mockDB.EXPECT().GetSubmission("2").Return(
			&models.Submission{
				Model:     gorm.Model{ID: 2},
				StudentID: "44444445",
				Result: models.Result{
					Model: gorm.Model{ID: 2},
					Score: 51,
				},
				AssignmentID: 1}, nil)

		var resp struct {
			Submissions []struct {
				ID, StudentID string
				Result        model.Result
			}
		}
		c.MustPost(`{ submissions() { id studentID result { id score } } }`, &resp)

		assert.Equal(t, "1", resp.Submissions[0].ID)
		assert.Equal(t, "2", resp.Submissions[1].ID)
		assert.Equal(t, "44444444", resp.Submissions[0].StudentID)
		assert.Equal(t, "44444445", resp.Submissions[1].StudentID)
		assert.Equal(t, model.Result{ID: "1", Score: 99}, resp.Submissions[0].Result)
		assert.Equal(t, model.Result{ID: "2", Score: 51}, resp.Submissions[1].Result)
	})

	t.Run("Get Submission Not Found", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetSubmission("1").Return(nil, db.ErrRecordNotFound)

		var resp struct {
			Submission struct{ ID, StudentID string }
		}
		err := c.Post(`{ submission(id:"1") { id studentID } }`, &resp)

		assert.ErrorContains(t, err, "record not found")
		assert.NotEqual(t, "1", resp.Submission.ID)
	})

	t.Run("Create Submission", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().CreateSubmission("44444444", uint(1)).Return(&models.Submission{Model: gorm.Model{ID: 1}, StudentID: "44444444"}, nil)

		var resp struct {
			CreateSubmission struct{ ID, StudentID string }
		}
		c.MustPost(`mutation { createSubmission(input: {studentID: "44444444", assignmentID: "1"}) { id studentID} }`, &resp)

		assert.Equal(t, "1", resp.CreateSubmission.ID)
		assert.Equal(t, "44444444", resp.CreateSubmission.StudentID)
	})
}

func TestResultResolver(t *testing.T) {
	t.Parallel()

	var resp struct {
		Result struct {
			ID, Date, SubmissionID string
			Score                  float64
		}
	}

	t.Run("Get Result", func(t *testing.T) {
		t.Parallel()

		now := time.Now()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetResult("1").Return(&models.Result{Model: gorm.Model{ID: 1, CreatedAt: now}, Score: 99, SubmissionID: 1}, nil)

		c.MustPost(`{ result(id:"1") { id date score submissionID } }`, &resp)

		assert.Equal(t, "1", resp.Result.ID)
		year, month, day := now.Date()
		assert.Equal(t, fmt.Sprintf("%02d/%02d/%d", day, month, year), resp.Result.Date)
		assert.Equal(t, "1", resp.Result.SubmissionID)
		assert.Equal(t, float64(99), resp.Result.Score)
	})

	t.Run("Get All Results", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		now := time.Now()

		mockDB.EXPECT().GetAllResults(1).Return([]*models.Result{
			{Model: gorm.Model{ID: 1, CreatedAt: now}, Score: 99, SubmissionID: 1},
			{Model: gorm.Model{ID: 2, CreatedAt: now}, Score: 51, SubmissionID: 2},
		}, nil)

		var resp struct {
			Results []struct {
				ID, Date, SubmissionID string
				Score                  float64
			}
		}
		c.MustPost(`{ results() { id score date submissionID } }`, &resp)

		assert.Equal(t, "1", resp.Results[0].ID)
		assert.Equal(t, "2", resp.Results[1].ID)
		assert.Equal(t, now.Format("02/01/2006"), resp.Results[0].Date)
		assert.Equal(t, now.Format("02/01/2006"), resp.Results[1].Date)
		assert.Equal(t, float64(99), resp.Results[0].Score)
		assert.Equal(t, float64(51), resp.Results[1].Score)
		assert.Equal(t, "1", resp.Results[0].SubmissionID)
		assert.Equal(t, "2", resp.Results[1].SubmissionID)
	})

	t.Run("Get Result Not Found", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetResult("1").Return(nil, db.ErrRecordNotFound)

		err := c.Post(`{ result(id:"1") { id date score submissionID } }`, &resp)

		assert.ErrorContains(t, err, "record not found")
		assert.NotEqual(t, "1", resp.Result.ID)
	})
}

func TestRegisterResolver(t *testing.T) {
	t.Parallel()

	var resp struct {
		Register string
	}

	user := &models.User{Model: gorm.Model{ID: 1}, Email: "a@b.com", PasswordHash: "password"}

	t.Run("New User", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUserByEmail("a@b.com").Return(nil, db.ErrRecordNotFound)
		mockDB.EXPECT().CreateUser("a@b.com", gomock.Any(), models.UserRoleAdmin).Return(user, nil)

		c.MustPost(`mutation { register(email:"a@b.com", password: "password") }`, &resp)

		assert.NotEmpty(t, resp.Register)
	})

	t.Run("User Exists", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUserByEmail("a@b.com").Return(nil, db.ErrRecordNotFound)
		mockDB.EXPECT().CreateUser("a@b.com", gomock.Any(), models.UserRoleAdmin).Return(user, nil)
		mockDB.EXPECT().GetUserByEmail("a@b.com").Return(user, db.ErrRecordNotFound)

		c.MustPost(`mutation { register(email:"a@b.com", password: "password") }`, &resp)
		err := c.Post(`mutation { register(email:"a@b.com", password: "password") }`, &resp)
		assert.Error(t, err)
	})

	t.Run("No Email", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		err := c.Post(`mutation { register(email:"", password: "password") }`, &resp)
		assert.Error(t, err)
	})

	t.Run("No Password", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		err := c.Post(`mutation { register(email:"a@b.com", password: "") }`, &resp)
		assert.Error(t, err)
	})

	t.Run("Error Getting User", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		customErr := errors.New("my cool error")
		mockDB.EXPECT().GetUserByEmail("a@b.com").Return(nil, customErr)

		err := c.Post(`mutation { register(email:"a@b.com", password: "password") }`, &resp)

		assert.ErrorContains(t, err, customErr.Error())
	})

	t.Run("Error Creating User - error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		customErr := errors.New("my cool error")
		mockDB.EXPECT().GetUserByEmail("a@b.com").Return(nil, db.ErrRecordNotFound)
		mockDB.EXPECT().CreateUser("a@b.com", gomock.Any(), models.UserRoleAdmin).Return(nil, customErr)

		err := c.Post(`mutation { register(email:"a@b.com", password: "password") }`, &resp)

		assert.ErrorContains(t, err, customErr.Error())
	})

	t.Run("Error Creating User - nil user", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUserByEmail("a@b.com").Return(nil, db.ErrRecordNotFound)
		mockDB.EXPECT().CreateUser("a@b.com", gomock.Any(), models.UserRoleAdmin).Return(nil, nil)

		err := c.Post(`mutation { register(email:"a@b.com", password: "password") }`, &resp)

		assert.ErrorContains(t, err, "error creating user")
	})
}

func TestLoginResolver(t *testing.T) {
	t.Parallel()

	var resp struct {
		Login string
	}

	user := &models.User{Model: gorm.Model{ID: 1}, Email: "a@b.com", PasswordHash: "$2y$10$iVyaKJWb4LzkbCMNKl6biuNQNdBG1WSsn3/cMkg3VHg5RSpQTJW0K"}

	t.Run("User", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUserByEmail("a@b.com").Return(user, nil)

		c.MustPost(`mutation { login(email:"a@b.com", password: "password") }`, &resp)

		assert.NotEmpty(t, resp.Login)
	})

	t.Run("No Email", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		err := c.Post(`mutation { login(email:"", password: "password") }`, &resp)
		assert.Error(t, err)
	})

	t.Run("No Password", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		err := c.Post(`mutation { login(email:"a@b.com", password: "") }`, &resp)
		assert.Error(t, err)
	})

	t.Run("User Does Not Exist - not found error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUserByEmail("a@b.com").Return(nil, db.ErrRecordNotFound)

		err := c.Post(`mutation { login(email:"a@b.com", password: "password") }`, &resp)
		assert.Error(t, err)
	})

	t.Run("User Does Not Exist - nil user", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUserByEmail("a@b.com").Return(nil, nil)

		err := c.Post(`mutation { login(email:"a@b.com", password: "password") }`, &resp)
		assert.ErrorContains(t, err, "user with email: a@b.com does not exist")
	})

	t.Run("User Incorrect Password", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUserByEmail("a@b.com").Return(user, nil)

		err := c.Post(`mutation { login(email:"a@b.com", password: "wrong_password") }`, &resp)
		assert.Error(t, err)
	})
}
