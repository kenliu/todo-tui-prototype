package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"todo-tui-prototype/tasklist"
)

type model struct {
	tl tasklist.Model
}

var (
	style = lipgloss.NewStyle().Border(lipgloss.ThickBorder(), true)
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		return m.tl.Update(msg)
	}
	return m, nil
}

func (m model) View() string {
	s := style.Render(m.tl.View())
	return s
}

func initialModel() model {
	return model{tl: tasklist.New()}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
