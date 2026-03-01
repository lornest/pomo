package session

import (
	"testing"
	"time"
)

func TestNewManager(t *testing.T) {
	m := NewManager(25*time.Minute, 5*time.Minute, 15*time.Minute, 4)
	if m.CurrentType() != Work {
		t.Errorf("expected Work, got %v", m.CurrentType())
	}
	if m.CurrentWork() != 1 {
		t.Errorf("expected work 1, got %d", m.CurrentWork())
	}
	if m.Duration() != 25*time.Minute {
		t.Errorf("expected 25m, got %v", m.Duration())
	}
}

func TestLabel(t *testing.T) {
	m := NewManager(25*time.Minute, 5*time.Minute, 15*time.Minute, 4)
	if m.Label() != "Pomodoro 1/4" {
		t.Errorf("expected 'Pomodoro 1/4', got %q", m.Label())
	}
	m.Advance()
	if m.Label() != "Short Break" {
		t.Errorf("expected 'Short Break', got %q", m.Label())
	}
}

func TestFullCycle(t *testing.T) {
	m := NewManager(25*time.Minute, 5*time.Minute, 15*time.Minute, 4)

	// Work 1 → Short Break
	assertSession(t, m, Work, 1)
	m.Advance()
	assertSession(t, m, ShortBreak, 1)

	// Short Break → Work 2
	m.Advance()
	assertSession(t, m, Work, 2)

	// Work 2 → Short Break
	m.Advance()
	assertSession(t, m, ShortBreak, 2)

	// Short Break → Work 3
	m.Advance()
	assertSession(t, m, Work, 3)

	// Work 3 → Short Break
	m.Advance()
	assertSession(t, m, ShortBreak, 3)

	// Short Break → Work 4
	m.Advance()
	assertSession(t, m, Work, 4)

	// Work 4 → Long Break (4th pomodoro triggers long break)
	m.Advance()
	assertSession(t, m, LongBreak, 4)

	// Long Break → Work 1 (cycle resets)
	m.Advance()
	assertSession(t, m, Work, 1)
}

func TestDurations(t *testing.T) {
	m := NewManager(25*time.Minute, 5*time.Minute, 15*time.Minute, 4)

	if m.Duration() != 25*time.Minute {
		t.Errorf("work duration: expected 25m, got %v", m.Duration())
	}
	m.Advance() // → short break
	if m.Duration() != 5*time.Minute {
		t.Errorf("short break duration: expected 5m, got %v", m.Duration())
	}

	// Advance through to long break
	m.Advance() // → work 2
	m.Advance() // → short break
	m.Advance() // → work 3
	m.Advance() // → short break
	m.Advance() // → work 4
	m.Advance() // → long break
	if m.Duration() != 15*time.Minute {
		t.Errorf("long break duration: expected 15m, got %v", m.Duration())
	}
}

func TestSessionTypeString(t *testing.T) {
	tests := []struct {
		st   Type
		want string
	}{
		{Work, "Work"},
		{ShortBreak, "Short Break"},
		{LongBreak, "Long Break"},
	}
	for _, tc := range tests {
		if tc.st.String() != tc.want {
			t.Errorf("expected %q, got %q", tc.want, tc.st.String())
		}
	}
}

func assertSession(t *testing.T, m *Manager, expectedType Type, expectedWork int) {
	t.Helper()
	if m.CurrentType() != expectedType {
		t.Errorf("expected type %v, got %v", expectedType, m.CurrentType())
	}
	if m.CurrentWork() != expectedWork {
		t.Errorf("expected work %d, got %d", expectedWork, m.CurrentWork())
	}
}
