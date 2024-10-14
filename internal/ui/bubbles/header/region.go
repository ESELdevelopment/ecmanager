package header

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strconv"
)

type regions struct {
	currentRegion string
	index         int
}

var supportedRegions = []string{
	"us-east-1", "us-west-2", "us-west-1", "eu-west-1", "eu-central-1", "eu-north-1",
	"ap-northeast-1", "ap-northeast-2", "ap-southeast-1", "ap-southeast-2",
}

type RegionChanged struct {
	Value string
}

func (p regions) Init() tea.Cmd {
	p.index = 4
	return nil
}

var numbers = key.NewBinding(key.WithKeys("0", "1", "2", "3", "4", "5", "6", "7", "8", "9"))

func (p regions) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg.(tea.KeyMsg), numbers) {
			p.index, _ = strconv.Atoi(msg.(tea.KeyMsg).String())
			p.currentRegion = supportedRegions[p.index]
			return p, func() tea.Msg {
				return RegionChanged{Value: p.currentRegion}
			}
		}
	}
	return p, nil
}

func (p regions) View() string {
	enumeratorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
	itemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("212")).MarginRight(1)
	numCols := 2
	numRows := (len(supportedRegions) + numCols - 1) / numCols

	columns := make([]string, numCols)

	for i, item := range supportedRegions {
		col := i / numRows
		var formattedItem string
		if item == p.currentRegion {
			formattedItem = itemStyle.Foreground(lipgloss.Color("100")).Render(item)
		} else {
			formattedItem = itemStyle.Render(item)
		}
		columns[col] += enumeratorStyle.Render(fmt.Sprintf("%d) ", i)) + formattedItem + "\n"
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, columns...)
}
