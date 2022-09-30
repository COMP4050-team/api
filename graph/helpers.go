package graph

import (
	"fmt"

	"github.com/COMP4050/square-team-5/api/internal/pkg/db"
	"github.com/COMP4050/square-team-5/api/internal/pkg/db/models"
)

func getOffset(from *int) int {
	if from == nil {
		return 1
	}

	if *from < 1 {
		return 1
	}

	return *from
}

func getAssignment(dbClient db.Database, id string) (*models.Assignment, error) {
	assignment, err := dbClient.GetAssignment(id)
	if err != nil {
		return nil, err
	}
	if assignment == nil {
		return nil, fmt.Errorf("assignment not found")
	}

	return assignment, nil
}

func getClass(dbClient db.Database, id string) (*models.Class, error) {
	class, err := dbClient.GetClass(id)
	if err != nil {
		return nil, err
	}
	if class == nil {
		return nil, fmt.Errorf("class not found")
	}

	return class, nil
}

func getUnit(dbClient db.Database, id string) (*models.Unit, error) {
	unit, err := dbClient.GetUnitByID(id, false)
	if err != nil {
		return nil, err
	}
	if unit == nil {
		return nil, fmt.Errorf("unit not found")
	}

	return unit, nil
}

func getSubmission(dbClient db.Database, id string) (*models.Submission, error) {
	submission, err := dbClient.GetSubmission(id)
	if err != nil {
		return nil, err
	}
	if submission == nil {
		return nil, fmt.Errorf("submission not found")
	}

	return submission, nil
}

func getTest(dbClient db.Database, id string) (*models.Test, error) {
	test, err := dbClient.GetTest(id)
	if err != nil {
		return nil, err
	}
	if test == nil {
		return nil, fmt.Errorf("test not found")
	}

	return test, nil
}
