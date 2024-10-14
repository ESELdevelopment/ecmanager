package start

import (
	"fmt"
	"github.com/ESELDevelopment/ecmanager/internal/ui/pages"
	"github.com/ESELDevelopment/ecmanager/internal/ui/pages/example"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type page struct {
	spinner spinner.Model

	router pages.Router
	err    error
}

var quitKeys = key.NewBinding(
	key.WithKeys("q", "ctrl+c"),
	key.WithHelp("", "press q/crtl+c to quit"),
)

var stopKeys = key.NewBinding(
	key.WithKeys("s"),
	key.WithHelp("", "press s to stop spinner"),
)

var style = lipgloss.NewStyle().
	Align(lipgloss.Center, lipgloss.Center)

func New(router pages.Router) tea.Model {
	s := spinner.New()
	s.Spinner = spinner.Monkey
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return page{spinner: s, router: router}
}

func (m page) Init() tea.Cmd {
	cmd := []tea.Cmd{
		tea.SetWindowTitle("ECManager"),
		m.spinner.Tick,
	}
	return tea.Batch(cmd...)
}

func (m page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, quitKeys) {
			return m, tea.Quit
		}
		if key.Matches(msg, stopKeys) {
			return m.router.Navigate(example.New(m.router))
		}
		return m, nil
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.WindowSizeMsg:
		style = style.Width(msg.Width).Height(msg.Height)
		return m, nil
	default:
		return m, nil
	}
}

func (m page) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	msg := fmt.Sprintf("%s Loading... \n", m.spinner.View())

	return style.Render(
		msg,
		stopKeys.Help().Desc+"\n",
		quitKeys.Help().Desc,
	)
}
