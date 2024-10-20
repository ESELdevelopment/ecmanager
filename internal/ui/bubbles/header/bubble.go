package header

import (
	"context"
	"github.com/ESELDevelopment/ecmanager/internal/aws/sts"
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
	account, userId := getAwsMetadata()
	metaPage := metadata{arn: *account, userId: *userId}
	regionPage := regions{currentRegion: currentRegion}
	return page{metaPage, regionPage, currentRegion, 80}
}

func getAwsMetadata() (*string, *string) {
	ctx := context.Background()
	stsService := sts.GetService(ctx)
	identity, _ := stsService.GetCallerIdentity(ctx)
	return identity.Arn, identity.UserId
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
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.width = msg.Width
		return p, nil
	case RegionChanged:
		p.selectedRegion = msg.Value
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
