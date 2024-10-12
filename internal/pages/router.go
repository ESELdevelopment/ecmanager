package pages

import tea "github.com/charmbracelet/bubbletea"

type RouterPage struct {
	model tea.Model // this will hold the current screen model
}

func Router() RouterPage {
	var defaultModel tea.Model = NewStartPage() // this is the default screen

	return RouterPage{
		model: defaultModel,
	}
}

func (m RouterPage) Init() tea.Cmd {
	return m.model.Init() // rest methods are just wrappers for the model's methods
}

func (m RouterPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.model.Update(msg)
}

func (m RouterPage) View() string {
	return m.model.View()
}

// Navigate this is the switcher which will switch between screens
func (m RouterPage) Navigate(model tea.Model) (tea.Model, tea.Cmd) {
	m.model = model
	return m.model, m.model.Init() // must return .Init() to initialize the screen (and here the magic happens)
}
