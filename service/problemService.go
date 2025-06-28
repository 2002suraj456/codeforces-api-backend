package service

import (
	dbapi "cf.practice.com/db"
	"cf.practice.com/models"
)

func InsertProblems(cfProblems []models.Problem) {
	db := dbapi.GetInstance()

	for _, problem := range cfProblems {
		if problem.Rating != 0 && len(problem.Tags) > 0 {
			db.InsertProblem(problem.ContestID, problem.Index, problem.Name, problem.Rating, problem.Tags)
		}
	}
}

func GetProblems(ratingStart, ratingEnd int, tags []string) []models.Problem {
	var result []models.Problem

	db := dbapi.GetInstance()

	var dbTags []models.Tag
	for _, t := range tags {
		dbTags = append(dbTags, models.Tag(t))
	}

	ids := db.Query(models.Rating(ratingStart), models.Rating(ratingEnd), dbTags)

	for _, id := range ids {
		result = append(result, db.GetProblem(id))
	}

	return result
}
