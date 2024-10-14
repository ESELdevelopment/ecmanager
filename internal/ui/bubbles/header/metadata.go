package header

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type metadata struct {
	msg string
}

func (p metadata) Init() tea.Cmd {
	return nil
}

func (p metadata) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case RegionChanged:
		regionChanged := msg.(RegionChanged)
		p.msg = regionChanged.Value
	}
	return p, nil
}

func (p metadata) View() string {
	return lipgloss.NewStyle().MarginRight(5).Render(fmt.Sprintf(`icon: 1.0.0
arn: %s
tbd
`, p.msg))
}
