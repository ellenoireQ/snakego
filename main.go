package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	lastKey string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.lastKey = "Up key pressed"
		case "down":
			m.lastKey = "Down key pressed"
		case "left":
			m.lastKey = "Left key pressed"
		case "right":
			m.lastKey = "Right key pressed"
		case "enter":
			m.lastKey = "Enter key pressed"
		case " ":
			m.lastKey = "Space key pressed"
		default:
			m.lastKey = fmt.Sprintf("%q key pressed", msg.String())
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.lastKey == "" {
		return "Press any key...\n\nPress q to quit.\n"
	}
	return fmt.Sprintf("%s\n\nPress q to quit.\n", m.lastKey)
}

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
