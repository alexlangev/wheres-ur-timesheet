package main

import (
	"fmt"
	"log"
	"time"

	"github.com/alexlangev/wheres-ur-timesheet/internal/clients"
)

func main() {
	jira := clients.NewClient(10 * time.Second)

	fmt.Println("Welcome to my app!")
	fmt.Println()

	user, err := jira.GetUser()
	if err != nil {
		log.Fatalf("Failed to get Jira user: %v", err)
	}

	fmt.Println("Where's your timesheet " + user.DisplayName + "?")
}
