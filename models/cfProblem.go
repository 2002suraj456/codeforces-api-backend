package models

// Problem represents a Codeforces problem
type Problem struct {
	ContestID int      `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Rating    int      `json:"rating,omitempty"`
	Tags      []string `json:"tags"`
	SeqID     int      `json:"__seq_id"`
}
