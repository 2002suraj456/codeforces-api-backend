package service_test

import (
	"testing"

	"cf.practice.com/models"
	"cf.practice.com/service"
)

func getStubProblems() []models.Problem {
	return []models.Problem{
		{
			ContestID: 1234,
			Index:     "A",
			Name:      "Problem A",
			Rating:    1200,
			Tags:      []string{"dp", "greedy"},
		},
		{
			ContestID: 1234,
			Index:     "B",
			Name:      "Problem B",
			Rating:    1300,
			Tags:      []string{"greedy"},
		},
		{
			ContestID: 1234,
			Index:     "C",
			Name:      "Problem C",
			Rating:    1400,
			Tags:      []string{"dp", "greedy"},
		},
	}
}

func TestGetProblems(t *testing.T) {

	var cfProblems []models.Problem

	cfProblems = getStubProblems()

	service.InsertProblems(cfProblems)

	problems := service.GetProblems(1200, 1400, []string{"dp", "greedy"})

	if len(problems) != 2 {
		t.Errorf("Expected 2 problems, got %d", len(problems))
	}
}
