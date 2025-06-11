package main

import (
	"fmt"
	"os"

	jira "github.com/alexlangev/wheres-ur-timesheet/internal/client"
	core "github.com/alexlangev/wheres-ur-timesheet/internal/core"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list map[int]core.Ticket // good for now...
	err  error
}

type listMsg map[int]core.Ticket

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func (m model) Init() tea.Cmd {
	// abstract into a func?
	return func() tea.Msg {
		tickets, err := jira.GetTickets()
		if err != nil {
			return errMsg{err}
		}

		return listMsg(tickets)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
		return m, tea.Quit

	case listMsg:
		m.list = msg
		return m, tea.Quit
	}

	return m, nil
}

func (m model) View() string {

	if m.err != nil {
		return fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err)
	}

	s := ""
	for _, v := range m.list {
		s += " "
		s += v.Key
		s += "\n"
	}

	return "\n" + s + "\n\n"
}

func main() {
	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v", err)
		os.Exit(1)
	}
}
