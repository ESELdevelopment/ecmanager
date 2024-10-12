package pages

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
)

type tickMsg time.Time

type StartPage struct {
	progress progress.Model

	quitting bool
	err      error
}

var quitKeys = key.NewBinding(
	key.WithKeys("q", "ctrl+c"),
	key.WithHelp("", "press q/crtl+c to quit"),
)

var navKeys = key.NewBinding(
	key.WithKeys("n", "ctrl+n"),
	key.WithHelp("", "press n/crtl+n to navigate"),
)

var style = lipgloss.NewStyle().
	Padding(2, 4).
	Align(lipgloss.Center, lipgloss.Center)

func NewStartPage() StartPage {
	p := progress.New(progress.WithDefaultGradient())

	return StartPage{progress: p}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m StartPage) Init() tea.Cmd {
	cmd := []tea.Cmd{
		tea.SetWindowTitle("ECManager"),
		tickCmd(),
	}
	return tea.Batch(cmd...)
}

func (m StartPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, quitKeys) {
			m.quitting = true
			return m, tea.Quit
		}
		if key.Matches(msg, navKeys) {
			return Router().Navigate(NewExamplePage())
		}
		return m, nil
	case tea.WindowSizeMsg:
		style = style.Width(msg.Width).Height(msg.Height)
		return m, nil

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			return m, nil
		}

		cmd := m.progress.IncrPercent(0.04)
		return m, tea.Batch(tickCmd(), cmd)

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m StartPage) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	if m.quitting {
		return "Goodbye!"
	}
	msg := "Loading... \n" + m.progress.View()
	if m.progress.Percent() == 1.0 {
		msg = "Hello, World! \n" + navKeys.Help().Desc
	}

	return style.Render(msg + "\n" + quitKeys.Help().Desc)
}
