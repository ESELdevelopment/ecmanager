package testdata

import tea "github.com/charmbracelet/bubbletea"

type MockRouter struct {
}

func (m MockRouter) Init() tea.Cmd {
	//TODO implement me
	return nil
}

func (m MockRouter) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//TODO implement me
	return m, nil
}

func (m MockRouter) View() string {
	return "Mock Navigated"
}

func (m MockRouter) Navigate(_ tea.Model) (tea.Model, tea.Cmd) {
	return m, nil
}
