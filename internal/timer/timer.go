package timer

import "time"

type State int

const (
	Idle State = iota
	Running
	Paused
	Completed
)

type Timer struct {
	totalDuration time.Duration
	state         State
	startedAt     time.Time
	pausedElapsed time.Duration // elapsed time accumulated before pause
}

func New(d time.Duration) *Timer {
	return &Timer{
		totalDuration: d,
		state:         Idle,
	}
}

func (t *Timer) Start() {
	if t.state == Idle {
		t.startedAt = time.Now()
		t.pausedElapsed = 0
		t.state = Running
	}
}

func (t *Timer) Pause() {
	if t.state == Running {
		t.pausedElapsed += time.Since(t.startedAt)
		t.state = Paused
	}
}

func (t *Timer) Resume() {
	if t.state == Paused {
		t.startedAt = time.Now()
		t.state = Running
	}
}

func (t *Timer) Toggle() {
	switch t.state {
	case Running:
		t.Pause()
	case Paused:
		t.Resume()
	}
}

func (t *Timer) Reset(d time.Duration) {
	t.totalDuration = d
	t.state = Idle
	t.pausedElapsed = 0
}

func (t *Timer) State() State {
	return t.state
}

func (t *Timer) TotalDuration() time.Duration {
	return t.totalDuration
}

// Elapsed returns total elapsed time, accounting for pauses.
func (t *Timer) Elapsed() time.Duration {
	switch t.state {
	case Running:
		return t.pausedElapsed + time.Since(t.startedAt)
	case Paused, Completed:
		return t.pausedElapsed
	default:
		return 0
	}
}

// Remaining returns time left. Returns 0 if completed.
func (t *Timer) Remaining() time.Duration {
	r := t.totalDuration - t.Elapsed()
	if r < 0 {
		return 0
	}
	return r
}

// Progress returns a value between 0.0 and 1.0.
func (t *Timer) Progress() float64 {
	if t.totalDuration == 0 {
		return 1.0
	}
	p := float64(t.Elapsed()) / float64(t.totalDuration)
	if p > 1.0 {
		return 1.0
	}
	return p
}

// Tick checks whether the timer has completed and transitions state.
// Returns true if the timer just completed on this tick.
func (t *Timer) Tick() bool {
	if t.state == Running && t.Elapsed() >= t.totalDuration {
		t.pausedElapsed = t.totalDuration
		t.state = Completed
		return true
	}
	return false
}
