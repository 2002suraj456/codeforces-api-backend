package codeforcesapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"cf.practice.com/models"
	"cf.practice.com/service"
	"cf.practice.com/utils"
)

var cache []byte = nil

func getAllProblems(w http.ResponseWriter, r *http.Request) {

	if cache != nil {
		// fmt.Fprintln(w, "successfully got & written the data")
		return
	}

	apiURL := "https://codeforces.com/api/problemset.problems"

	resp, err := http.Get(apiURL)

	if err != nil {
		// write a error response to the client that unalbe to fetch problems
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// write to a the client the error that it got from the api
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		// write to the client error reading the response
	}

	var apiResponse models.CodeforcesProblemsAPIResponse
	err = json.Unmarshal(body, &apiResponse)

	if err != nil {
		// write to the client about not able to process api correctly
	}

	// service.InsertProblems(apiResponse.Result.Problem)
	// val s = apiResponse
	var s = apiResponse.Result.Problem
	service.InsertProblems(s)

	data, err := json.MarshalIndent(apiResponse, "", " ")
	if err != nil {
		panic(err)
	}

	cache = data

	err = os.WriteFile("problems.json", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "successfully got & written the data new")

}

func CodeforcesRouter() *http.ServeMux {
	router := utils.NewRouter()

	router.HandleFunc("/problems", getAllProblems)

	return router.GetServeMux()
}
