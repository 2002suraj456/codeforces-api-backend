package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"cf.practice.com/codeforcesapi"
	"cf.practice.com/models"
	"cf.practice.com/service"
	"cf.practice.com/utils"
)

func convertToInternalProblem(cfProblem codeforcesapi.CodeforcesProblem) models.Problem {
	return models.Problem{
		ContestID: cfProblem.ContestID,
		Index:     cfProblem.Index,
		Name:      cfProblem.Name,
		Rating:    cfProblem.Rating,
		Tags:      cfProblem.Tags,
	}
}

func Init() error {
	apiURL := "https://codeforces.com/api/problemset.problems"

	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var apiResponse codeforcesapi.CodeforcesProblemResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil
	}

	var internalProblems []models.Problem
	for _, cfProblem := range apiResponse.Result.Problem {
		internalProblems = append(internalProblems, convertToInternalProblem(cfProblem))
	}

	// Insert problems into database
	service.InsertProblems(internalProblems)

	return nil
}

func main() {

	if err := Init(); err != nil {
		panic("unable to start the app due to error : ")
	}

	router := utils.NewRouter()
	router.Handle("/api/codeforces/", codeforcesapi.CodeforcesRouter())

	fmt.Println("\nStarting server on :5729")
	log.Fatal(http.ListenAndServe(":5729", router.GetServeMux()))

}
