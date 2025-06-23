package main

import (
	"flag"
	"fmt"
)

func main() {
	direct := flag.Bool("c", false, "Create worklog")

	flag.Parse()

	switch {
	case *direct:
		fmt.Println("CLI here!")
		// env incomplete?
		// validate Jira auth with currentUser()
		// print valid auth or not

	default:
		// TODO setup tui
		// tui.Run()
		fmt.Println("The TUI implementation is not ready.")
	}
}
