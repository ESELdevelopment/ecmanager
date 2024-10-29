package header

import (
	"fmt"
	"github.com/ESELDevelopment/ecmanager/cmd/info"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type metadata struct {
	arn    string
	userId string
	width  int
}

func (p metadata) Init() tea.Cmd {
	return nil
}

func (p metadata) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.width = msg.Width / 3
	}
	return p, nil
}

func (p metadata) View() string {
	return lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).Width(p.width).MarginRight(1).
		Render(fmt.Sprintf(`Rev: %s
arn: %s
userId: %s`, info.GetVersion(), p.arn, p.userId))
}
