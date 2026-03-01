//go:build darwin

package notify

import (
	"fmt"
	"os/exec"
)

type darwinNotifier struct {
	noSound  bool
	noNotify bool
}

func newNotifier(noSound, noNotify bool) Notifier {
	return &darwinNotifier{noSound: noSound, noNotify: noNotify}
}

func (n *darwinNotifier) Notify(message string) {
	if !n.noNotify {
		script := fmt.Sprintf(
			`display notification %q with title "pomo" sound name "Glass"`,
			message,
		)
		exec.Command("osascript", "-e", script).Start()
	}

	if !n.noSound {
		exec.Command("afplay", "/System/Library/Sounds/Glass.aiff").Start()
	}
}
