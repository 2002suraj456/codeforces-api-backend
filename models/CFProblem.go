package models

type Problem struct {
	ContestID      int      `json:"contestId"`
	ProblemsetName string   `json:"problemsetName"`
	Index          string   `json:"index"`
	Name           string   `json:"name"`
	ProblemType    string   `json:"type"`
	Points         string   `json:"points"`
	Rating         int      `json:"rating,omitempty"`
	Tags           []string `json:"tags"`
}

type CodeforcesProblemsAPIResponse struct {
	Status string `json:"status"`
	Result struct {
		Problem []Problem `json:"problems"`
	} `json:"result"`
}
