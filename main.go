package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"pomo/internal/config"
	"pomo/internal/ui"
)

func main() {
	cfg := config.ParseFlags()

	p := tea.NewProgram(
		ui.NewModel(cfg),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
