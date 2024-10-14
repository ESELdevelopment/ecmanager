package quit

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type examplePage struct {
}

func New() tea.Model {
	return examplePage{}
}

func (m examplePage) Init() tea.Cmd {
	return nil
}

func (m examplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		// any other key switches the screen
		return m, tea.Quit
	}
	return m, nil
}

func (m examplePage) View() string {
	return lipgloss.NewStyle().Bold(true).Render("This is the end screen. Press any key to quit.")
}
