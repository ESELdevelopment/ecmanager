package cmd

import (
	"fmt"
	"github.com/ESELDevelopment/ecmanager/internal/ui/bubbles/header"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Start() {
	p := tea.NewProgram(
		header.New("eu-central-1"),
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
		tea.WithFPS(120),
	)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
