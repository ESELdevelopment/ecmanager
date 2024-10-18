package header

import (
	"fmt"
	"github.com/ESELDevelopment/ecmanager/cmd/info"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type metadata struct {
	arn    string
	role   string
	region string
	width  int
}

func (p metadata) Init() tea.Cmd {
	return nil
}

func (p metadata) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case RegionChanged:
		regionChanged := msg.(RegionChanged)
		p.region = regionChanged.Value
	case tea.WindowSizeMsg:
		p.width = msg.(tea.WindowSizeMsg).Width / 3
	}
	return p, nil
}

func (p metadata) View() string {
	return lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).BorderForeground(lipgloss.Color("123")).Width(p.width).MarginRight(1).Render(fmt.Sprintf(`Rev: %s
region: %s
account: %s
role: %s`, info.GetVersion(), p.region, p.arn, p.role))
}
