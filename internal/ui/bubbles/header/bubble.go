package header

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type page struct {
	meta           metadata
	regions        regions
	selectedRegion string
}

var (
	modelStyle = lipgloss.NewStyle().
		Height(5).
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.HiddenBorder())
)

func New(currentRegion string) tea.Model {
	metaPage := metadata{}
	regionPage := regions{currentRegion: currentRegion}
	return page{metaPage, regionPage, currentRegion}
}

func (p page) Init() tea.Cmd {
	return nil
}

func (p page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	regionModel, regionCmd := p.regions.Update(msg)
	metadataModel, metadataCmd := p.meta.Update(msg)
	p.regions = regionModel.(regions)
	p.meta = metadataModel.(metadata)
	p, cmd := p.updateOnMessage(msg)
	return p, tea.Batch(regionCmd, metadataCmd, cmd)
}

func (p page) updateOnMessage(msg tea.Msg) (page, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		if msg.(tea.KeyMsg).String() == "q" {
			return p, tea.Quit
		}
		if msg.(tea.KeyMsg).String() == "a" {
			p.selectedRegion = "clicked"
			return p, nil
		}
	case RegionChanged:
		regionChanged := msg.(RegionChanged)
		p.selectedRegion = regionChanged.Value
		return p, nil
	default:
		return p, nil
	}
	return p, nil
}

func (p page) View() string {
	horizontal := lipgloss.JoinHorizontal(lipgloss.Top, p.meta.View(), p.regions.View())
	return modelStyle.Render(horizontal + "\n" + p.selectedRegion)
}
