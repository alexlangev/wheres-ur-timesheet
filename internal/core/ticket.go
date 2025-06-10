package core

import (
	"time"
)

// The atomic datastructure of the app
// everything revolves aroung tickets

type Ticket struct {
	Key       string    `json:"key"`
	Summary   string    `json:"summary"`
	Status    string    `json:"status"`
	Project   string    `json:"project"`
	Assignee  string    `json:"assignee"`
	IssueType string    `json:"issueType"`
	Updated   time.Time `json:"updated"`
}
