package example

import (
	"github.com/ESELDevelopment/ecmanager/internal/pages"
	"github.com/ESELDevelopment/ecmanager/internal/pages/quit"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type examplePage struct {
	router pages.Router
}

func New(router pages.Router) tea.Model {
	return examplePage{router: router}
}

func (m examplePage) Init() tea.Cmd {
	return nil
}

func (m examplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		// any other key switches the screen
		return m.router.Navigate(quit.New())
	}
	return m, nil
}

func (m examplePage) View() string {
	return lipgloss.NewStyle().Bold(true).Render("This is the example screen. Press any key to switch to the second screen.")
}
