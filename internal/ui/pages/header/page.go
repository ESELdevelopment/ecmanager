package header

import (
	"github.com/ESELDevelopment/ecmanager/internal/ui/bubbles/header"
	tea "github.com/charmbracelet/bubbletea"
)

type page struct {
	header tea.Model
	main   tea.Model
}

func (p page) Init() tea.Cmd {
	return p.main.Init()
}

func (p page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	mainModel, mainCmd := p.main.Update(msg)
	headerModel, headerCmd := p.header.Update(msg)
	p.main = mainModel
	p.header = headerModel
	return p, tea.Batch(mainCmd, headerCmd)
}

func (p page) View() string {
	return p.header.View() + "\n" + p.main.View()
}

func New(mainModel tea.Model) tea.Model {
	return page{header.New("eu-central-1"), mainModel}
}
