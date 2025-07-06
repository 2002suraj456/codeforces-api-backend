package db_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"

	"cf.practice.com/codeforcesapi"
	"cf.practice.com/models"
	"cf.practice.com/utils"
)

var problems []models.Problem

func TestMain(m *testing.M) {
	apiURL := "https://codeforces.com/api/problemset.problems"

	resp, err := http.Get(apiURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return

	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return

	}

	var apiResponse codeforcesapi.CodeforcesProblemResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return
	}

	for _, cfProblem := range apiResponse.Result.Problem {
		problems = append(problems, models.Problem{
			ContestID: cfProblem.ContestID,
			Index:     cfProblem.Index,
			Name:      cfProblem.Name,
			Rating:    cfProblem.Rating,
			Tags:      cfProblem.Tags,
		})
	}

	code := m.Run()

	os.Exit(code)
}

func getPopulatedDBInstance() (models.ProblemDBInterface, error) {
	newdb := models.ProblemDB{}
	newdb.Init()

	for _, problem := range problems {
		if problem.Rating != 0 && len(problem.Tags) > 0 {
			newdb.InsertProblem(problem.ContestID, problem.Index, problem.Name, problem.Rating, problem.Tags)
		}
	}
	return &newdb, nil
}

func queryFromCFDirect(ratingstart int, ratingend int, tags []string) ([]models.Problem, error) {

	result := utils.Filter(problems, func(problem models.Problem) bool {
		if problem.Rating == 0 {
			return false
		}
		if problem.Rating > ratingend || problem.Rating < ratingstart {
			return false
		}
		for _, tag := range tags {
			if utils.Find(problem.Tags, tag) == false {
				return false
			}
		}

		return true
	})

	return result, nil
}

func TestWithSingleTag(t *testing.T) {
	db, err := getPopulatedDBInstance()

	if err != nil {
		t.Fatal("error getting populated db")
	}

	tags := []models.Tag{"dp"}
	ratingstart := 0
	ratingend := 15000

	resSeqIDs := db.Query(models.Rating(ratingstart), models.Rating(ratingend), tags)

	var resProblems []models.Problem

	for _, val := range resSeqIDs {
		resProblems = append(resProblems, db.GetProblem(val))
	}

	directRes, err := queryFromCFDirect(ratingstart, ratingend, []string{"dp"})

	if err != nil {
		t.Fatal("error query from cfdirectly")
	}

	var match map[string]int = make(map[string]int)

	for _, problem := range resProblems {
		match[strconv.Itoa(problem.ContestID)+problem.Index]++
	}

	for _, problem := range directRes {
		match[strconv.Itoa(problem.ContestID)+problem.Index]--
	}

	var notmatch map[string]int = make(map[string]int)

	for key, val := range match {
		if val != 0 {
			notmatch[key] = val
		}
	}

	if len(notmatch) > 0 {
		fmt.Println(notmatch)
		t.Fatal("result not matching")
	}

}

func TestWithMultipleTag(t *testing.T) {
	db, err := getPopulatedDBInstance()

	if err != nil {
		t.Fatal("error getting populated db")
	}

	tags := []models.Tag{"dp", "greedy"}
	ratingstart := 0
	ratingend := 15000

	resSeqIDs := db.Query(models.Rating(ratingstart), models.Rating(ratingend), tags)

	var resProblems []models.Problem

	for _, val := range resSeqIDs {
		resProblems = append(resProblems, db.GetProblem(val))
	}

	directRes, err := queryFromCFDirect(ratingstart, ratingend, []string{"dp", "greedy"})

	if err != nil {
		t.Fatal("error query from cfdirectly")
	}

	var match map[string]int = make(map[string]int)

	for _, problem := range resProblems {
		match[strconv.Itoa(problem.ContestID)+problem.Index]++
	}

	for _, problem := range directRes {
		match[strconv.Itoa(problem.ContestID)+problem.Index]--
	}

	var notmatch map[string]int = make(map[string]int)

	for key, val := range match {
		if val != 0 {
			notmatch[key] = val
		}
	}

	if len(notmatch) > 0 {
		fmt.Println(notmatch)
		t.Fatal("result not matching")
	}
}
