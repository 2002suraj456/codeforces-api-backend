package codeforcesapi

type CodeforcesProblemResponse struct {
	Status string `json:"status"`
	Result struct {
		Problem []CodeforcesProblem `json:"problems"`
	} `json:"result"`
}

type CodeforcesProblem struct {
	ContestID      int      `json:"contestId"`
	ProblemsetName string   `json:"problemsetName"`
	Index          string   `json:"index"`
	Name           string   `json:"name"`
	Rating         int      `json:"rating,omitempty"`
	Tags           []string `json:"tags"`
}
