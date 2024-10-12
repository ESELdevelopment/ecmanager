package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type examplePage struct {
}

func NewExamplePage() tea.Model {
	return examplePage{}
}

func (m examplePage) Init() tea.Cmd {
	return nil
}

func (m examplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		// any other key switches the screen
		return Router().Navigate(NewStartPage())
	}
	return m, nil
}

func (m examplePage) View() string {
	return lipgloss.NewStyle().Bold(true).Render("This is the example screen. Press any key to switch to the second screen.")
}
