package cmd

import (
	"fmt"
	"github.com/ESELDevelopment/ecmanager/internal/pages"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Start() {
	p := tea.NewProgram(pages.NewStartPage(), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
