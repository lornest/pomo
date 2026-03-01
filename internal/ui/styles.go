package ui

import "github.com/charmbracelet/lipgloss"

// Gradient colors for work sessions (soft green → warm coral).
const (
	WorkGradientStart = "#86EFAC"
	WorkGradientEnd   = "#FCA5A5"
)

// Gradient colors for break sessions (sky blue → teal: calming).
const (
	BreakGradientStart = "#38BDF8"
	BreakGradientEnd   = "#2DD4BF"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			MarginBottom(1)

	timerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			MarginBottom(1)

	sessionStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A1A1AA")).
			MarginBottom(1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#71717A")).
			MarginTop(1)
)
