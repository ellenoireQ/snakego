package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Coord struct {
	x int
	y int
}

func (c *Coord) updatex() {
	c.x++
}

func (c *Coord) updatey() {
	c.y++
}

type model struct {
	lastKey string
	coord   Coord
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
			m.coord.updatex()
		case "down":
			m.lastKey = "Down key pressed"
		case "left":
			m.lastKey = "Left key pressed"
		case "right":
			m.coord.updatey()
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
	return fmt.Sprintf(
		"X: %d\nY: %d\n\nPress q to quit.\n",
		m.coord.x,
		m.coord.y,
	)
}

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
