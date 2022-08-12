// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/db/db.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/COMP4050/square-team-5/api/internal/pkg/db/models"
	gomock "github.com/golang/mock/gomock"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// CreateAssignment mocks base method.
func (m *MockDatabase) CreateAssignment(name string, classID uint) (*models.Assignment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAssignment", name, classID)
	ret0, _ := ret[0].(*models.Assignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAssignment indicates an expected call of CreateAssignment.
func (mr *MockDatabaseMockRecorder) CreateAssignment(name, classID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAssignment", reflect.TypeOf((*MockDatabase)(nil).CreateAssignment), name, classID)
}

// CreateClass mocks base method.
func (m *MockDatabase) CreateClass(name string, unitID uint) (*models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClass", name, unitID)
	ret0, _ := ret[0].(*models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClass indicates an expected call of CreateClass.
func (mr *MockDatabaseMockRecorder) CreateClass(name, unitID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClass", reflect.TypeOf((*MockDatabase)(nil).CreateClass), name, unitID)
}

// CreateResult mocks base method.
func (m *MockDatabase) CreateResult(score float64, submissionID uint) (*models.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateResult", score, submissionID)
	ret0, _ := ret[0].(*models.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateResult indicates an expected call of CreateResult.
func (mr *MockDatabaseMockRecorder) CreateResult(score, submissionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResult", reflect.TypeOf((*MockDatabase)(nil).CreateResult), score, submissionID)
}

// CreateSubmission mocks base method.
func (m *MockDatabase) CreateSubmission(studentID string, assignmentID uint) (*models.Submission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubmission", studentID, assignmentID)
	ret0, _ := ret[0].(*models.Submission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubmission indicates an expected call of CreateSubmission.
func (mr *MockDatabaseMockRecorder) CreateSubmission(studentID, assignmentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubmission", reflect.TypeOf((*MockDatabase)(nil).CreateSubmission), studentID, assignmentID)
}

// CreateTest mocks base method.
func (m *MockDatabase) CreateTest(name string, assignmentID uint) (*models.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTest", name, assignmentID)
	ret0, _ := ret[0].(*models.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTest indicates an expected call of CreateTest.
func (mr *MockDatabaseMockRecorder) CreateTest(name, assignmentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTest", reflect.TypeOf((*MockDatabase)(nil).CreateTest), name, assignmentID)
}

// CreateUnit mocks base method.
func (m *MockDatabase) CreateUnit(name string) (*models.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUnit", name)
	ret0, _ := ret[0].(*models.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUnit indicates an expected call of CreateUnit.
func (mr *MockDatabaseMockRecorder) CreateUnit(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUnit", reflect.TypeOf((*MockDatabase)(nil).CreateUnit), name)
}

// CreateUser mocks base method.
func (m *MockDatabase) CreateUser(email, passwordHash string, role models.UserRole) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", email, passwordHash, role)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDatabaseMockRecorder) CreateUser(email, passwordHash, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDatabase)(nil).CreateUser), email, passwordHash, role)
}

// GetAllAssignments mocks base method.
func (m *MockDatabase) GetAllAssignments() ([]*models.Assignment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAssignments")
	ret0, _ := ret[0].([]*models.Assignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAssignments indicates an expected call of GetAllAssignments.
func (mr *MockDatabaseMockRecorder) GetAllAssignments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAssignments", reflect.TypeOf((*MockDatabase)(nil).GetAllAssignments))
}

// GetAllClasses mocks base method.
func (m *MockDatabase) GetAllClasses() ([]*models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllClasses")
	ret0, _ := ret[0].([]*models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllClasses indicates an expected call of GetAllClasses.
func (mr *MockDatabaseMockRecorder) GetAllClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllClasses", reflect.TypeOf((*MockDatabase)(nil).GetAllClasses))
}

// GetAllResults mocks base method.
func (m *MockDatabase) GetAllResults() ([]*models.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllResults")
	ret0, _ := ret[0].([]*models.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllResults indicates an expected call of GetAllResults.
func (mr *MockDatabaseMockRecorder) GetAllResults() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllResults", reflect.TypeOf((*MockDatabase)(nil).GetAllResults))
}

// GetAllSubmissions mocks base method.
func (m *MockDatabase) GetAllSubmissions() ([]*models.Submission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubmissions")
	ret0, _ := ret[0].([]*models.Submission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSubmissions indicates an expected call of GetAllSubmissions.
func (mr *MockDatabaseMockRecorder) GetAllSubmissions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubmissions", reflect.TypeOf((*MockDatabase)(nil).GetAllSubmissions))
}

// GetAllTests mocks base method.
func (m *MockDatabase) GetAllTests() ([]*models.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTests")
	ret0, _ := ret[0].([]*models.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTests indicates an expected call of GetAllTests.
func (mr *MockDatabaseMockRecorder) GetAllTests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTests", reflect.TypeOf((*MockDatabase)(nil).GetAllTests))
}

// GetAllUnits mocks base method.
func (m *MockDatabase) GetAllUnits() ([]*models.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUnits")
	ret0, _ := ret[0].([]*models.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUnits indicates an expected call of GetAllUnits.
func (mr *MockDatabaseMockRecorder) GetAllUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUnits", reflect.TypeOf((*MockDatabase)(nil).GetAllUnits))
}

// GetAssignment mocks base method.
func (m *MockDatabase) GetAssignment(id string) (*models.Assignment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssignment", id)
	ret0, _ := ret[0].(*models.Assignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssignment indicates an expected call of GetAssignment.
func (mr *MockDatabaseMockRecorder) GetAssignment(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssignment", reflect.TypeOf((*MockDatabase)(nil).GetAssignment), id)
}

// GetAssignmentsForClass mocks base method.
func (m *MockDatabase) GetAssignmentsForClass(classID string) ([]*models.Assignment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssignmentsForClass", classID)
	ret0, _ := ret[0].([]*models.Assignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssignmentsForClass indicates an expected call of GetAssignmentsForClass.
func (mr *MockDatabaseMockRecorder) GetAssignmentsForClass(classID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssignmentsForClass", reflect.TypeOf((*MockDatabase)(nil).GetAssignmentsForClass), classID)
}

// GetClass mocks base method.
func (m *MockDatabase) GetClass(id string) (*models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClass", id)
	ret0, _ := ret[0].(*models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClass indicates an expected call of GetClass.
func (mr *MockDatabaseMockRecorder) GetClass(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClass", reflect.TypeOf((*MockDatabase)(nil).GetClass), id)
}

// GetResult mocks base method.
func (m *MockDatabase) GetResult(id string) (*models.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResult", id)
	ret0, _ := ret[0].(*models.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResult indicates an expected call of GetResult.
func (mr *MockDatabaseMockRecorder) GetResult(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResult", reflect.TypeOf((*MockDatabase)(nil).GetResult), id)
}

// GetSubmission mocks base method.
func (m *MockDatabase) GetSubmission(id string) (*models.Submission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubmission", id)
	ret0, _ := ret[0].(*models.Submission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubmission indicates an expected call of GetSubmission.
func (mr *MockDatabaseMockRecorder) GetSubmission(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubmission", reflect.TypeOf((*MockDatabase)(nil).GetSubmission), id)
}

// GetSubmissionsForAssignment mocks base method.
func (m *MockDatabase) GetSubmissionsForAssignment(assignmentID string) ([]*models.Submission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubmissionsForAssignment", assignmentID)
	ret0, _ := ret[0].([]*models.Submission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubmissionsForAssignment indicates an expected call of GetSubmissionsForAssignment.
func (mr *MockDatabaseMockRecorder) GetSubmissionsForAssignment(assignmentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubmissionsForAssignment", reflect.TypeOf((*MockDatabase)(nil).GetSubmissionsForAssignment), assignmentID)
}

// GetTest mocks base method.
func (m *MockDatabase) GetTest(id string) (*models.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTest", id)
	ret0, _ := ret[0].(*models.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTest indicates an expected call of GetTest.
func (mr *MockDatabaseMockRecorder) GetTest(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTest", reflect.TypeOf((*MockDatabase)(nil).GetTest), id)
}

// GetTestsForAssignment mocks base method.
func (m *MockDatabase) GetTestsForAssignment(assignmentID string) ([]*models.Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestsForAssignment", assignmentID)
	ret0, _ := ret[0].([]*models.Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestsForAssignment indicates an expected call of GetTestsForAssignment.
func (mr *MockDatabaseMockRecorder) GetTestsForAssignment(assignmentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestsForAssignment", reflect.TypeOf((*MockDatabase)(nil).GetTestsForAssignment), assignmentID)
}

// GetUnitByID mocks base method.
func (m *MockDatabase) GetUnitByID(id string, fetchClasses bool) (*models.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitByID", id, fetchClasses)
	ret0, _ := ret[0].(*models.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitByID indicates an expected call of GetUnitByID.
func (mr *MockDatabaseMockRecorder) GetUnitByID(id, fetchClasses interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitByID", reflect.TypeOf((*MockDatabase)(nil).GetUnitByID), id, fetchClasses)
}

// GetUnitByName mocks base method.
func (m *MockDatabase) GetUnitByName(name string) (*models.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitByName", name)
	ret0, _ := ret[0].(*models.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitByName indicates an expected call of GetUnitByName.
func (mr *MockDatabaseMockRecorder) GetUnitByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitByName", reflect.TypeOf((*MockDatabase)(nil).GetUnitByName), name)
}

// GetUserByEmail mocks base method.
func (m *MockDatabase) GetUserByEmail(email string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockDatabaseMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockDatabase)(nil).GetUserByEmail), email)
}
