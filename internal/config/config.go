package config

import (
	"flag"
	"time"
)

type Config struct {
	WorkDuration   time.Duration
	ShortBreak     time.Duration
	LongBreak      time.Duration
	Intervals      int
	NoSound        bool
	NoNotify       bool
}

func Default() Config {
	return Config{
		WorkDuration: 25 * time.Minute,
		ShortBreak:   5 * time.Minute,
		LongBreak:    15 * time.Minute,
		Intervals:    4,
	}
}

func ParseFlags() Config {
	cfg := Default()

	flag.DurationVar(&cfg.WorkDuration, "work", cfg.WorkDuration, "Work session duration")
	flag.DurationVar(&cfg.ShortBreak, "short", cfg.ShortBreak, "Short break duration")
	flag.DurationVar(&cfg.LongBreak, "long", cfg.LongBreak, "Long break duration")
	flag.IntVar(&cfg.Intervals, "intervals", cfg.Intervals, "Pomodoros before long break")
	flag.BoolVar(&cfg.NoSound, "no-sound", cfg.NoSound, "Disable sound")
	flag.BoolVar(&cfg.NoNotify, "no-notify", cfg.NoNotify, "Disable notifications")

	flag.Parse()
	return cfg
}
