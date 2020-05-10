package sync

// Counter counts
type Counter struct{}

// Inc increments value by 1
func (c *Counter) Inc() {
}

// Value returns final value
func (c *Counter) Value() int {
	return 0
}
