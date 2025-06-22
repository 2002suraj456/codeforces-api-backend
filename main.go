package main

import (
	"fmt"
	"log"
	"net/http"

	"cf.practice.com/codeforcesapi"
	"cf.practice.com/utils"
)

func main() {
	router := utils.NewRouter()
	router.Handle("/api/codeforces/", codeforcesapi.CodeforcesRouter())

	fmt.Println("\nStarting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", router.GetServeMux()))

}
