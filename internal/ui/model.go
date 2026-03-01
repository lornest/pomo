package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"

	"pomo/internal/config"
	"pomo/internal/notify"
	"pomo/internal/session"
	"pomo/internal/timer"
)

const tickInterval = 100 * time.Millisecond

type tickMsg time.Time

type Model struct {
	timer    *timer.Timer
	session  *session.Manager
	progress progress.Model
	notifier notify.Notifier
	config   config.Config
	width    int
}

func NewModel(cfg config.Config) Model {
	sess := session.NewManager(cfg.WorkDuration, cfg.ShortBreak, cfg.LongBreak, cfg.Intervals)
	p := progress.New(
		progress.WithGradient(WorkGradientStart, WorkGradientEnd),
		progress.WithoutPercentage(),
	)

	return Model{
		timer:    timer.New(sess.Duration()),
		session:  sess,
		progress: p,
		notifier: notify.New(cfg.NoSound, cfg.NoNotify),
		config:   cfg,
		width:    40,
	}
}

func (m Model) Init() tea.Cmd {
	m.timer.Start()
	return tickCmd()
}

func tickCmd() tea.Cmd {
	return tea.Tick(tickInterval, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch handleKey(msg) {
		case keyQuit:
			return m, tea.Quit
		case keyToggle:
			m.timer.Toggle()
		case keySkip:
			m.advanceSession()
		}

	case tickMsg:
		if m.timer.Tick() {
			m.notifier.Notify(m.session.CurrentType().String() + " complete!")
			m.advanceSession()
		}
		return m, tickCmd()

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - 4
		if m.progress.Width > 80 {
			m.progress.Width = 80
		}
		m.width = msg.Width
	}

	return m, nil
}

func (m *Model) advanceSession() {
	m.session.Advance()
	m.timer.Reset(m.session.Duration())
	m.timer.Start()
	m.updateGradient()
}

func (m *Model) updateGradient() {
	switch m.session.CurrentType() {
	case session.Work:
		m.progress = progress.New(
			progress.WithGradient(WorkGradientStart, WorkGradientEnd),
			progress.WithoutPercentage(),
		)
	default:
		m.progress = progress.New(
			progress.WithGradient(BreakGradientStart, BreakGradientEnd),
			progress.WithoutPercentage(),
		)
	}
	m.progress.Width = m.width - 4
	if m.progress.Width > 80 {
		m.progress.Width = 80
	}
}
