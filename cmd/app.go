package cmd

import (
  "fmt"
  "github.com/ESELDevelopment/ecmanager/internal/pages"
  "os"

  tea "github.com/charmbracelet/bubbletea"
)

func Start() {
  p := tea.NewProgram(
    pages.Router(),
    tea.WithAltScreen(),
    tea.WithMouseAllMotion(),
    tea.WithFPS(120),
  )
  if _, err := p.Run(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
