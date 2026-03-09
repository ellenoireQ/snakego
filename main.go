package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	width  = 20
	height = 10
)

type Coord struct {
	X int
	Y int
}

type model struct {
	snake []Coord
	dir   Coord
}

func initialModel() model {
	return model{
		snake: []Coord{
			{10, 5},
			{9, 5},
			{8, 5},
		},
		dir: Coord{1, 0},
	}
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
			m.dir = Coord{0, -1}

		case "down":
			m.dir = Coord{0, 1}

		case "left":
			m.dir = Coord{-1, 0}

		case "right":
			m.dir = Coord{1, 0}
		}

		m.moveSnake()
	}

	return m, nil
}

func (m *model) moveSnake() {
	head := m.snake[0]

	newHead := Coord{
		X: head.X + m.dir.X,
		Y: head.Y + m.dir.Y,
	}

	m.snake = append([]Coord{newHead}, m.snake...)
	m.snake = m.snake[:len(m.snake)-1]
}

func (m model) isSnake(x, y int) bool {
	for _, s := range m.snake {
		if s.X == x && s.Y == y {
			return true
		}
	}
	return false
}

func (m model) View() string {
	out := ""

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if m.isSnake(x, y) {
				out += "█"
			} else {
				out += "-"
			}
		}
		out += "\n"
	}

	return out + "\nPress arrows to move • q to quit\n"
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
