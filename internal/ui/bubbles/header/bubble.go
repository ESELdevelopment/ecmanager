package header

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type page struct {
	meta           metadata
	regions        regions
	selectedRegion string

	width int
}

func New(currentRegion string) tea.Model {
	metaPage := metadata{arn: "arn:aws:iam::123456789012:role/role-name", role: "role-name"}
	regionPage := regions{currentRegion: currentRegion}
	return page{metaPage, regionPage, currentRegion, 80}
}

func (p page) Init() tea.Cmd {
	regionCmd := p.regions.Init()
	metadataCmd := p.meta.Init()
	return tea.Batch(regionCmd, metadataCmd)
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
	case tea.WindowSizeMsg:
		p.width = msg.(tea.WindowSizeMsg).Width
		return p, nil
	case RegionChanged:
		regionChanged := msg.(RegionChanged)
		p.selectedRegion = regionChanged.Value
		return p, nil
	default:
		return p, nil
	}
}

func (p page) View() string {
	regionView := p.regions.View()
	metaView := p.meta.View()
	regionWith := lipgloss.Width(regionView)
	metaSize := p.width - regionWith
	horizontal := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Width(metaSize).Render(metaView),
		lipgloss.NewStyle().Width(regionWith).Render(regionView),
	)
	return horizontal
}
