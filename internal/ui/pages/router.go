package pages

import tea "github.com/charmbracelet/bubbletea"

type Router interface {
	tea.Model
	Navigate(model tea.Model) (tea.Model, tea.Cmd)
}

// CreateRouter Don't use this page as tea model
func CreateRouter() Router {
	var defaultModel tea.Model // this is the default screen

	return routerImpl{
		model: defaultModel,
	}
}

type routerImpl struct {
	model tea.Model // this will hold the current screen model
}

func (m routerImpl) Init() tea.Cmd {
	return m.model.Init() // rest methods are just wrappers for the model's methods
}

func (m routerImpl) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.model.Update(msg)
}

func (m routerImpl) View() string {
	return m.model.View()
}

// Navigate this is the switcher which will switch between screens
func (m routerImpl) Navigate(model tea.Model) (tea.Model, tea.Cmd) {
	m.model = model
	return m.model, m.model.Init() // must return .Init() to initialize the screen (and here the magic happens)
}
