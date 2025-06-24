package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/alexlangev/wheres-ur-timesheet/internal/clients"
	"github.com/alexlangev/wheres-ur-timesheet/internal/utils"
)

func main() {
	cli := flag.Bool("c", false, "Create worklog")
	dateString := flag.String("date", "", "Worklog date")
	startTimeString := flag.String("start", "", "Worklog start time")
	durationString := flag.String("duration", "", "Worklog duration")
	issueKey := flag.String("issue", "", "Issue Key")

	flag.Parse()

	switch {
	case *cli && *dateString != "" && *startTimeString != "" && *durationString != "" && *issueKey != "":
		fmt.Println("Both worked", *dateString, *startTimeString, *durationString, *issueKey)

	case *cli:
		fmt.Println("CLI here!")
		domain, email, jiraToken, _, err := utils.GetEnv()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		jiraClient := clients.NewJiraClient(10*time.Second, domain, email, jiraToken)

		// get user info like display name
		user, err := jiraClient.GetUser()
		if err != nil {
			fmt.Println("Error with authentication. Please review your env files", err)
			os.Exit(1)
		}

		fmt.Println("Welcome ", user.DisplayName)

		os.Exit(0)

	default:
		// TODO setup tui
		// tui.Run()
		fmt.Println("The TUI implementation is not ready.")
	}
}
