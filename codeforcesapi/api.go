package codeforcesapi

import (
	"encoding/json"
	"net/http"

	"cf.practice.com/service"
	"cf.practice.com/utils"
)

var cache []byte = nil

func getAllProblems(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var ratingStart, ratingEnd int
	var tags []string

	// get these from the request
	ratingStart = 0
	ratingEnd = 10000
	tags = []string{"dp"}

	res := service.GetProblems(ratingStart, ratingEnd, tags)

	json.NewEncoder(w).Encode(res)
}

func CodeforcesRouter() *http.ServeMux {
	router := utils.NewRouter()

	router.HandleFunc("/problems", getAllProblems)

	return router.GetServeMux()
}
