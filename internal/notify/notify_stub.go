//go:build !darwin

package notify

type stubNotifier struct{}

func newNotifier(_, _ bool) Notifier {
	return &stubNotifier{}
}

func (n *stubNotifier) Notify(string) {}
