package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/COMP4050/square-team-5/api/graph/generated"
	"github.com/COMP4050/square-team-5/api/graph/model"
	"github.com/COMP4050/square-team-5/api/internal/pkg/db"
	"github.com/COMP4050/square-team-5/api/internal/pkg/db/models"
	"github.com/golang-jwt/jwt"
)

// Tests is the resolver for the tests field.
func (r *assignmentResolver) Tests(ctx context.Context, obj *model.Assignment) ([]*model.Test, error) {
	tests, err := r.DB.GetTestsForAssignment(obj.ID)
	if err != nil {
		return nil, err
	}
	if tests == nil {
		return nil, nil
	}

	var gqlTests []*model.Test

	for _, test := range tests {
		gqlTests = append(gqlTests, &model.Test{
			ID:   fmt.Sprintf("%d", test.ID),
			Name: test.Name,
		})
	}

	return gqlTests, nil
}

// Submissions is the resolver for the submissions field.
func (r *assignmentResolver) Submissions(ctx context.Context, obj *model.Assignment) ([]*model.Submission, error) {
	submissions, err := r.DB.GetSubmissionsForAssignment(obj.ID)
	if err != nil {
		return nil, err
	}
	if submissions == nil {
		return nil, nil
	}

	var gqlSubmissions []*model.Submission

	for _, submission := range submissions {
		gqlSubmissions = append(gqlSubmissions, &model.Submission{
			ID:        fmt.Sprintf("%d", submission.ID),
			StudentID: submission.StudentID,
			Result: &model.Result{
				ID:           fmt.Sprintf("%d", submission.Result.ID),
				Score:        submission.Result.Score,
				Date:         submission.Result.CreatedAt.Format("02/01/2006"),
				SubmissionID: fmt.Sprintf("%d", submission.Result.SubmissionID),
			},
		})
	}

	return gqlSubmissions, nil
}

// Assignments is the resolver for the assignments field.
func (r *classResolver) Assignments(ctx context.Context, obj *model.Class) ([]*model.Assignment, error) {
	classID, err := strconv.ParseUint(obj.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	assignments, err := r.DB.GetAssignmentsForClass(uint(classID))
	if err != nil {
		return nil, err
	}
	if assignments == nil {
		return nil, nil
	}

	var gqlAssignments []*model.Assignment

	for _, assignment := range assignments {
		gqlAssignments = append(gqlAssignments, &model.Assignment{
			ID:          fmt.Sprintf("%d", assignment.ID),
			Name:        assignment.Name,
			DueDate:     int(assignment.DueDate.Unix()),
			Tests:       []*model.Test{},
			Submissions: []*model.Submission{},
		})
	}

	return gqlAssignments, nil
}

// CreateUnit is the resolver for the createUnit field.
func (r *mutationResolver) CreateUnit(ctx context.Context, input model.NewUnit) (*model.Unit, error) {
	// Check if a unit with the same name already exists
	existingUnit, err := r.DB.GetUnitByName(input.Name)
	if err != nil && !errors.Is(err, db.ErrRecordNotFound) {
		return nil, fmt.Errorf("error getting unit: %w", err)
	}
	if existingUnit != nil {
		return nil, fmt.Errorf("unit already exists")
	}

	unit, err := r.DB.CreateUnit(input.Name)
	if err != nil {
		return nil, fmt.Errorf("error creating unit: %w", err)
	}
	if unit == nil {
		return nil, fmt.Errorf("error creating unit")
	}

	gqlUnit := &model.Unit{ID: fmt.Sprintf("%d", unit.ID), Name: unit.Name, Classes: []*model.Class{}}

	return gqlUnit, nil
}

// CreateClass is the resolver for the createClass field.
func (r *mutationResolver) CreateClass(ctx context.Context, input model.NewClass) (*model.Class, error) {
	unitID, err := strconv.ParseUint(input.UnitID, 10, 64)
	if err != nil {
		return nil, err
	}

	// TODO: Don't allow creating a class with the same name for the same unit

	unit, err := r.DB.GetUnitByID(input.UnitID, false)
	if err != nil {
		return nil, fmt.Errorf("error creating class: %w", err)
	}
	if unit == nil {
		return nil, fmt.Errorf("unit with id: %d does not exist", unitID)
	}

	class, err := r.DB.CreateClass(input.Name, uint(unitID))
	if err != nil {
		return nil, fmt.Errorf("error creating class: %w", err)
	}
	if class == nil {
		return nil, nil
	}

	gqlClass := &model.Class{ID: fmt.Sprintf("%d", class.ID), Name: class.Name}

	return gqlClass, nil
}

// CreateAssignment is the resolver for the createAssignment field.
func (r *mutationResolver) CreateAssignment(ctx context.Context, input model.NewAssignment) (*model.Assignment, error) {
	id, err := strconv.ParseUint(input.ClassID, 10, 64)
	if err != nil {
		return nil, err
	}

	assignment, err := r.DB.CreateAssignment(input.Name, input.DueDate, uint(id))
	if err != nil {
		return nil, fmt.Errorf("error creating assignment: %w", err)
	}
	if assignment == nil {
		return nil, nil
	}

	gqlAssignment := &model.Assignment{ID: fmt.Sprintf("%d", assignment.ID), Name: assignment.Name}

	return gqlAssignment, nil
}

// CreateTest is the resolver for the createTest field.
func (r *mutationResolver) CreateTest(ctx context.Context, input model.NewTest) (*model.Test, error) {
	id, err := strconv.ParseUint(input.AssignmentID, 10, 64)
	if err != nil {
		return nil, err
	}

	if input.Name == "" || input.StoragePath == "" {
		return nil, fmt.Errorf("name and storage path are required")
	}

	test, err := r.DB.CreateTest(input.Name, input.StoragePath, uint(id))
	if err != nil {
		return nil, fmt.Errorf("error creating test: %w", err)
	}
	if test == nil {
		return nil, nil
	}

	gqlTest := &model.Test{ID: fmt.Sprintf("%d", test.ID), Name: test.Name}

	return gqlTest, nil
}

// RunTest is the resolver for the runTest field.
func (r *mutationResolver) RunTest(ctx context.Context, testID string) (bool, error) {
	body := map[string]string{
		"s3Key": fmt.Sprintf("tests/test_%s.java", testID),
	}

	json, err := json.Marshal(body)
	if err != nil {
		return false, err
	}

	testExecutorEndpoint := r.Config.TestExecutorEndpoint

	res, err := http.Post(testExecutorEndpoint, "application/json", bytes.NewBuffer(json))

	if err != nil {
		return false, fmt.Errorf("error running test: %w", err)
	}

	if res.StatusCode != 200 {
		return false, fmt.Errorf("error running test: %w", err)
	}

	return true, nil
}

// CreateSubmission is the resolver for the createSubmission field.
func (r *mutationResolver) CreateSubmission(ctx context.Context, input model.NewSubmission) (*model.Submission, error) {
	assignmentID, err := strconv.ParseUint(input.AssignmentID, 10, 64)
	if err != nil {
		return nil, err
	}

	submission, err := r.DB.CreateSubmission(input.StudentID, uint(assignmentID))
	if err != nil {
		return nil, fmt.Errorf("error creating submission: %w", err)
	}
	if submission == nil {
		return nil, nil
	}

	gqlSubmission := &model.Submission{ID: fmt.Sprintf("%d", submission.ID), StudentID: submission.StudentID}

	return gqlSubmission, nil
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, email string, password string) (string, error) {
	if email == "" || password == "" {
		return "", fmt.Errorf("email or password must not be empty")
	}

	user, err := r.DB.GetUserByEmail(email)
	if err != nil && !errors.Is(err, db.ErrRecordNotFound) {
		return "", fmt.Errorf("error getting user: %w", err)
	}
	if user != nil {
		return "", fmt.Errorf("user already exists")
	}

	passwordHash, err := models.HashPassword(password)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	user, err = r.DB.CreateUser(email, passwordHash, models.UserRoleAdmin)
	if err != nil {
		return "", fmt.Errorf("error creating user: %w", err)
	}
	if user == nil {
		return "", fmt.Errorf("error creating user")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
	})

	tokenString, err := token.SignedString([]byte(r.Config.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (string, error) {
	if email == "" || password == "" {
		return "", fmt.Errorf("email or password must not be empty")
	}

	user, err := r.DB.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("error getting user: %w", err)
	}
	if user == nil {
		return "", fmt.Errorf("user with email: %s does not exist", email)
	}

	err = user.CheckPassword(password)
	if err != nil {
		return "", fmt.Errorf("incorrect username or password: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
	})

	tokenString, err := token.SignedString([]byte(r.Config.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}

// Units is the resolver for the units field.
func (r *queryResolver) Units(ctx context.Context) ([]*model.Unit, error) {
	units, err := r.DB.GetAllUnits()
	if err != nil {
		return nil, fmt.Errorf("error getting units: %w", err)
	}
	if units == nil {
		return nil, fmt.Errorf("error getting units")
	}

	gqlUnits := []*model.Unit{}
	for _, unit := range units {
		gqlUnit := &model.Unit{ID: fmt.Sprintf("%d", unit.ID), Name: unit.Name, Classes: []*model.Class{}}
		gqlUnits = append(gqlUnits, gqlUnit)
	}

	return gqlUnits, nil
}

// Unit is the resolver for the unit field.
func (r *queryResolver) Unit(ctx context.Context, id string) (*model.Unit, error) {
	unit, err := r.DB.GetUnitByID(id, false)
	if err != nil {
		return nil, err
	}
	if unit == nil {
		return nil, nil
	}

	gqlUnit := &model.Unit{ID: id, Name: unit.Name, Classes: []*model.Class{}}

	return gqlUnit, nil
}

// Classes is the resolver for the classes field.
func (r *queryResolver) Classes(ctx context.Context) ([]*model.Class, error) {
	classes, err := r.DB.GetAllClasses()
	if err != nil {
		return nil, fmt.Errorf("error getting classes: %w", err)
	}
	if classes == nil {
		return nil, nil
	}

	gqlClasses := []*model.Class{}
	for _, class := range classes {
		gqlClass := &model.Class{ID: fmt.Sprintf("%d", class.ID), Name: class.Name, UnitID: fmt.Sprintf("%d", class.UnitID)}
		gqlClasses = append(gqlClasses, gqlClass)
	}

	return gqlClasses, nil
}

// Class is the resolver for the class field.
func (r *queryResolver) Class(ctx context.Context, id string) (*model.Class, error) {
	class, err := r.DB.GetClass(id)
	if err != nil {
		return nil, fmt.Errorf("error getting class: %w", err)
	}
	if class == nil {
		return nil, nil
	}

	return &model.Class{ID: id, Name: class.Name}, nil
}

// Assignments is the resolver for the assignments field.
func (r *queryResolver) Assignments(ctx context.Context) ([]*model.Assignment, error) {
	assignments, err := r.DB.GetAllAssignments()
	if err != nil {
		return nil, fmt.Errorf("error getting assignments: %w", err)
	}
	if assignments == nil {
		return nil, nil
	}

	gqlAssignments := []*model.Assignment{}
	for _, assignment := range assignments {
		gqlAssignments = append(gqlAssignments, &model.Assignment{
			ID:          fmt.Sprintf("%d", assignment.ID),
			Name:        assignment.Name,
			DueDate:     int(assignment.DueDate.Unix()),
			Tests:       []*model.Test{},
			Submissions: []*model.Submission{},
		})
	}

	return gqlAssignments, nil
}

// Assignment is the resolver for the assignment field.
func (r *queryResolver) Assignment(ctx context.Context, id string) (*model.Assignment, error) {
	assignment, err := r.DB.GetAssignment(id)
	if err != nil {
		return nil, fmt.Errorf("error getting assignment: %w", err)
	}
	if assignment == nil {
		return nil, nil
	}

	return &model.Assignment{ID: id, Name: assignment.Name}, nil
}

// Tests is the resolver for the tests field.
func (r *queryResolver) Tests(ctx context.Context) ([]*model.Test, error) {
	tests, err := r.DB.GetAllTests()
	if err != nil {
		return nil, fmt.Errorf("error getting tests: %w", err)
	}
	if tests == nil {
		return nil, nil
	}

	gqlTests := []*model.Test{}
	for _, test := range tests {
		gqlTest := &model.Test{ID: fmt.Sprintf("%d", test.ID), Name: test.Name, AssignmentID: fmt.Sprintf("%d", test.AssignmentID)}
		gqlTests = append(gqlTests, gqlTest)
	}

	return gqlTests, nil
}

// Test is the resolver for the test field.
func (r *queryResolver) Test(ctx context.Context, id string) (*model.Test, error) {
	test, err := r.DB.GetTest(id)
	if err != nil {
		return nil, fmt.Errorf("error getting test: %w", err)
	}
	if test == nil {
		return nil, nil
	}

	return &model.Test{ID: id, Name: test.Name}, nil
}

// Submissions is the resolver for the submissions field.
func (r *queryResolver) Submissions(ctx context.Context) ([]*model.Submission, error) {
	submissions, err := r.DB.GetAllSubmissions()
	if err != nil {
		return nil, fmt.Errorf("error getting submissions: %w", err)
	}
	if submissions == nil {
		return nil, nil
	}

	gqlSubmissions := []*model.Submission{}
	for _, submission := range submissions {
		gqlSubmission := &model.Submission{ID: fmt.Sprintf("%d", submission.ID), StudentID: submission.StudentID}
		gqlSubmissions = append(gqlSubmissions, gqlSubmission)
	}

	return gqlSubmissions, nil
}

// Submission is the resolver for the submission field.
func (r *queryResolver) Submission(ctx context.Context, id string) (*model.Submission, error) {
	submission, err := r.DB.GetSubmission(id)
	if err != nil {
		return nil, fmt.Errorf("error getting submission: %w", err)
	}
	if submission == nil {
		return nil, nil
	}

	return &model.Submission{ID: id, StudentID: submission.StudentID}, nil
}

// Results is the resolver for the results field.
func (r *queryResolver) Results(ctx context.Context) ([]*model.Result, error) {
	results, err := r.DB.GetAllResults()
	if err != nil {
		return nil, fmt.Errorf("error getting results: %w", err)
	}
	if results == nil {
		return nil, nil
	}

	gqlResults := []*model.Result{}
	for _, result := range results {
		gqlResult := &model.Result{
			ID:           fmt.Sprintf("%d", result.ID),
			Score:        result.Score,
			Date:         result.CreatedAt.Format("02/01/2006"),
			SubmissionID: fmt.Sprintf("%d", result.SubmissionID),
		}
		gqlResults = append(gqlResults, gqlResult)
	}

	return gqlResults, nil
}

// Result is the resolver for the result field.
func (r *queryResolver) Result(ctx context.Context, id string) (*model.Result, error) {
	result, err := r.DB.GetResult(id)
	if err != nil {
		return nil, fmt.Errorf("error getting result: %w", err)
	}
	if result == nil {
		return nil, nil
	}

	return &model.Result{ID: fmt.Sprintf("%d", result.ID), Score: result.Score, SubmissionID: fmt.Sprintf("%d", result.SubmissionID), Date: result.CreatedAt.Format("02/01/2006")}, nil
}

// Result is the resolver for the result field.
func (r *submissionResolver) Result(ctx context.Context, obj *model.Submission) (*model.Result, error) {
	submission, err := r.DB.GetSubmission(obj.ID)
	if err != nil {
		return nil, err
	}
	if submission == nil {
		return nil, nil
	}

	return &model.Result{ID: fmt.Sprintf("%d", submission.Result.ID), Score: submission.Result.Score}, nil
}

// Classes is the resolver for the classes field.
func (r *unitResolver) Classes(ctx context.Context, obj *model.Unit) ([]*model.Class, error) {
	unit, err := r.DB.GetUnitByID(obj.ID, true)
	if err != nil {
		return nil, err
	}
	if unit == nil {
		return nil, nil
	}

	var gqlClasses []*model.Class

	for _, class := range unit.Classes {
		gqlClasses = append(gqlClasses, &model.Class{ID: fmt.Sprintf("%d", class.ID), Name: class.Name, UnitID: fmt.Sprintf("%d", class.UnitID)})
	}

	return gqlClasses, nil
}

// Assignment returns generated.AssignmentResolver implementation.
func (r *Resolver) Assignment() generated.AssignmentResolver { return &assignmentResolver{r} }

// Class returns generated.ClassResolver implementation.
func (r *Resolver) Class() generated.ClassResolver { return &classResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Submission returns generated.SubmissionResolver implementation.
func (r *Resolver) Submission() generated.SubmissionResolver { return &submissionResolver{r} }

// Unit returns generated.UnitResolver implementation.
func (r *Resolver) Unit() generated.UnitResolver { return &unitResolver{r} }

type assignmentResolver struct{ *Resolver }
type classResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type submissionResolver struct{ *Resolver }
type unitResolver struct{ *Resolver }
