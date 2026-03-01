package notify

// Notifier sends a notification when a session completes.
type Notifier interface {
	Notify(message string)
}

// New returns a platform-appropriate Notifier.
func New(noSound, noNotify bool) Notifier {
	return newNotifier(noSound, noNotify)
}
