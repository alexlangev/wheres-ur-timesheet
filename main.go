package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ticket struct {
	Key       string    `json:"key"`
	Summary   string    `json:"summary"`
	Status    string    `json:"status"`
	Project   string    `json:"project"`
	Assignee  string    `json:"assignee"`
	IssueType string    `json:"issueType"`
	Updated   time.Time `json:"updated"`
}

type ticketList map[int]ticket

type trackedTickets []ticket

// will be replaced with http GET
// for time being, read mock file
func getTickets() (map[int]ticket, error) {
	data, err := os.ReadFile("mock.json")

	if err != nil {
		return nil, err
	}

	var tickets []ticket

	if err := json.Unmarshal(data, &tickets); err != nil {
		return nil, err
	}

	m := make(map[int]ticket)

	for i, t := range tickets {
		m[i] = t
	}

	return m, nil
}

func main() {
	fmt.Println("I don't know what I'm doing...")
	tickets, err := getTickets()

	if err != nil {
		fmt.Printf("Failed to read from mock file: %v", err)
	}

	for _, v := range tickets {
		fmt.Println("ticket: ", v.Key)
	}
}
