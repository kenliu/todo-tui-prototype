package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	"os"
	"todo-tui-prototype/tasklist"
)

type mainModel struct {
	tasklist tasklist.Model
}

var (
	style  = lipgloss.NewStyle().Border(lipgloss.ThickBorder(), true)
	logger log.Logger
)

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	//logger.Println(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	m.tasklist, cmd = m.tasklist.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
	//return m, cmd
}

func (m mainModel) View() string {
	s := style.Render(m.tasklist.View())
	return s
}

func initialModel() mainModel {
	return mainModel{tasklist: tasklist.New()}
}

func main() {
	f, err := tea.LogToFile("logs/debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	//logger := log.New(f, "", log.LstdFlags)
	//logger.Println("starting program")
	defer f.Close()

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
