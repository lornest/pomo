package ui

import (
	"fmt"

	"github.com/lornest/pomo/internal/timer"
)

func (m Model) View() string {
	remaining := m.timer.Remaining()
	minutes := int(remaining.Minutes())
	seconds := int(remaining.Seconds()) % 60

	timeStr := fmt.Sprintf("%02d:%02d", minutes, seconds)

	label := m.session.Label()

	var stateHint string
	switch m.timer.State() {
	case timer.Paused:
		stateHint = " (paused)"
	case timer.Completed:
		stateHint = " (completed)"
	}

	s := titleStyle.Render("🍅 pomo") + "\n\n"
	s += sessionStyle.Render(label+stateHint) + "\n"
	s += timerStyle.Render(timeStr) + "\n\n"
	s += m.progress.ViewAs(m.timer.Progress()) + "\n"
	s += helpStyle.Render("space: pause • s: skip • q: quit")

	return s
}
