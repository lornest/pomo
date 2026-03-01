package session

import (
	"fmt"
	"time"
)

type Type int

const (
	Work Type = iota
	ShortBreak
	LongBreak
)

func (t Type) String() string {
	switch t {
	case Work:
		return "Work"
	case ShortBreak:
		return "Short Break"
	case LongBreak:
		return "Long Break"
	default:
		return "Unknown"
	}
}

type Manager struct {
	workDuration   time.Duration
	shortBreak     time.Duration
	longBreak      time.Duration
	intervals      int // pomodoros before long break
	currentWork    int // 1-based work session counter
	currentType    Type
}

func NewManager(work, short, long time.Duration, intervals int) *Manager {
	return &Manager{
		workDuration: work,
		shortBreak:   short,
		longBreak:    long,
		intervals:    intervals,
		currentWork:  1,
		currentType:  Work,
	}
}

func (m *Manager) CurrentType() Type {
	return m.currentType
}

func (m *Manager) CurrentWork() int {
	return m.currentWork
}

func (m *Manager) Intervals() int {
	return m.intervals
}

// Duration returns the duration for the current session type.
func (m *Manager) Duration() time.Duration {
	switch m.currentType {
	case Work:
		return m.workDuration
	case ShortBreak:
		return m.shortBreak
	case LongBreak:
		return m.longBreak
	default:
		return m.workDuration
	}
}

// Advance moves to the next session in the cycle.
// Work → ShortBreak (or LongBreak after `intervals` work sessions)
// Break → Work (incrementing work counter)
func (m *Manager) Advance() {
	switch m.currentType {
	case Work:
		if m.currentWork >= m.intervals {
			m.currentType = LongBreak
		} else {
			m.currentType = ShortBreak
		}
	case ShortBreak, LongBreak:
		if m.currentType == LongBreak {
			m.currentWork = 1
		} else {
			m.currentWork++
		}
		m.currentType = Work
	}
}

// Label returns a display string like "Work 2/4" or "Short Break".
func (m *Manager) Label() string {
	switch m.currentType {
	case Work:
		return fmt.Sprintf("Pomodoro %d/%d", m.currentWork, m.intervals)
	default:
		return m.currentType.String()
	}
}
