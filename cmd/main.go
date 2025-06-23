package main

import (
	"flag"
	"fmt"

	"github.com/alexlangev/wheres-ur-timesheet/internal/utils"
)

func main() {
	cli := flag.Bool("c", false, "Create worklog")

	flag.Parse()

	switch {
	case *cli:
		fmt.Println("CLI here!")
		// get env
		// print incomplete
		_, _, _, _, err := utils.GetEnv()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("is the env working?")

	default:
		// TODO setup tui
		// tui.Run()
		fmt.Println("The TUI implementation is not ready.")
	}
}
