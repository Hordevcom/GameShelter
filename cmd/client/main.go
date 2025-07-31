package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	apiURL      = "http://localhost:8080/api/users/a31lulus234ha/games"
	httpTimeout = 10 * time.Second
)

var keyMap = struct {
	Quit   string
	Up     string
	Down   string
	Select string
}{
	Quit:   "q",
	Up:     "up",
	Down:   "down",
	Select: "enter",
}

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	status   int
	err      error
}

func checkGames() tea.Msg {
	client := &http.Client{Timeout: httpTimeout}
	res, err := client.Get(apiURL)
	if err != nil {
		return err
	}
	return res.Body
}

func initialModel() model {
	return model{
		choices:  []string{"Equip gear", "Find dolly", "Do some stuff"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", keyMap.Quit:
			return m, tea.Quit
		case keyMap.Up:
			if m.cursor > 0 {
				m.cursor--
			}
		case keyMap.Down:
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case keyMap.Select:
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "What quest we do?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += fmt.Sprintf("\nPress '%s' to quit.\n", keyMap.Quit)

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
