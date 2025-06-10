package main

import (
	"fmt"

	jira "github.com/alexlangev/wheres-ur-timesheet/internal/client"
)

func main() {
	fmt.Println("I don't know what I'm doing...")
	tickets, err := jira.GetTickets()

	if err != nil {
		fmt.Printf("Failed to read from mock file: %v", err)
	}

	for _, v := range tickets {
		fmt.Println("ticket: ", v.Key)
	}
}
