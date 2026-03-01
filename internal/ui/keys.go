package ui

import tea "github.com/charmbracelet/bubbletea"

type keyAction int

const (
	keyNone keyAction = iota
	keyQuit
	keyToggle
	keySkip
)

func handleKey(msg tea.KeyMsg) keyAction {
	switch msg.String() {
	case "q", "ctrl+c":
		return keyQuit
	case " ":
		return keyToggle
	case "s":
		return keySkip
	default:
		return keyNone
	}
}
