package timer

import (
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	tm := New(25 * time.Minute)
	if tm.State() != Idle {
		t.Errorf("expected Idle, got %d", tm.State())
	}
	if tm.TotalDuration() != 25*time.Minute {
		t.Errorf("expected 25m, got %v", tm.TotalDuration())
	}
	if tm.Remaining() != 25*time.Minute {
		t.Errorf("expected 25m remaining, got %v", tm.Remaining())
	}
}

func TestStartTransition(t *testing.T) {
	tm := New(5 * time.Second)
	tm.Start()
	if tm.State() != Running {
		t.Errorf("expected Running, got %d", tm.State())
	}
}

func TestPauseResume(t *testing.T) {
	tm := New(5 * time.Second)
	tm.Start()
	time.Sleep(10 * time.Millisecond)

	tm.Pause()
	if tm.State() != Paused {
		t.Errorf("expected Paused, got %d", tm.State())
	}
	elapsed := tm.Elapsed()
	if elapsed == 0 {
		t.Error("expected non-zero elapsed after pause")
	}

	// Elapsed should not change while paused.
	time.Sleep(10 * time.Millisecond)
	if tm.Elapsed() != elapsed {
		t.Error("elapsed should not change while paused")
	}

	tm.Resume()
	if tm.State() != Running {
		t.Errorf("expected Running after resume, got %d", tm.State())
	}
}

func TestToggle(t *testing.T) {
	tm := New(5 * time.Second)
	tm.Start()
	tm.Toggle()
	if tm.State() != Paused {
		t.Errorf("expected Paused after toggle, got %d", tm.State())
	}
	tm.Toggle()
	if tm.State() != Running {
		t.Errorf("expected Running after second toggle, got %d", tm.State())
	}
}

func TestCompletion(t *testing.T) {
	tm := New(10 * time.Millisecond)
	tm.Start()
	time.Sleep(20 * time.Millisecond)

	completed := tm.Tick()
	if !completed {
		t.Error("expected timer to complete")
	}
	if tm.State() != Completed {
		t.Errorf("expected Completed, got %d", tm.State())
	}
	if tm.Remaining() != 0 {
		t.Errorf("expected 0 remaining, got %v", tm.Remaining())
	}
}

func TestProgressBounds(t *testing.T) {
	tm := New(10 * time.Millisecond)
	if tm.Progress() != 0 {
		t.Errorf("expected 0 progress in idle, got %f", tm.Progress())
	}

	tm.Start()
	time.Sleep(20 * time.Millisecond)
	tm.Tick()

	if tm.Progress() != 1.0 {
		t.Errorf("expected 1.0 progress when completed, got %f", tm.Progress())
	}
}

func TestReset(t *testing.T) {
	tm := New(5 * time.Second)
	tm.Start()
	time.Sleep(10 * time.Millisecond)
	tm.Reset(10 * time.Second)

	if tm.State() != Idle {
		t.Errorf("expected Idle after reset, got %d", tm.State())
	}
	if tm.TotalDuration() != 10*time.Second {
		t.Errorf("expected 10s, got %v", tm.TotalDuration())
	}
	if tm.Elapsed() != 0 {
		t.Error("expected 0 elapsed after reset")
	}
}

func TestZeroDuration(t *testing.T) {
	tm := New(0)
	if tm.Progress() != 1.0 {
		t.Errorf("expected 1.0 for zero duration, got %f", tm.Progress())
	}
}
