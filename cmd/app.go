package cmd

import (
	"fmt"
	"github.com/ESELDevelopment/ecmanager/internal/pages"
	"github.com/ESELDevelopment/ecmanager/internal/pages/start"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Start() {
	p := tea.NewProgram(
		start.New(pages.CreateRouter()),
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
		tea.WithFPS(120),
	)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
