package codeforcesapi

import (
	"fmt"
	"net/http"

	"cf.practice.com/utils"
)

func getAllProblems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "problem from codeforces api")
}

func CodeforcesRouter() *http.ServeMux {
	router := utils.NewRouter()

	router.HandleFunc("/problems", getAllProblems)

	return router.GetServeMux()
}
