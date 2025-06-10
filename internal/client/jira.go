package jira

import (
	"encoding/json"
	"os"

	core "github.com/alexlangev/wheres-ur-timesheet/internal/core"
)

// will be replaced with http GET
// for time being, read mock file
func GetTickets() (map[int]core.Ticket, error) {
	data, err := os.ReadFile("mock.json")

	if err != nil {
		return nil, err
	}

	// will rename better
	var tickets []core.Ticket

	if err := json.Unmarshal(data, &tickets); err != nil {
		return nil, err
	}

	m := make(map[int]core.Ticket)

	for i, t := range tickets {
		m[i] = t
	}

	return m, nil
}
