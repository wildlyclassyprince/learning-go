package sync

// Counter counts
type Counter struct {
	value int
}

// Inc increments value by 1
func (c *Counter) Inc() {
	c.value++
}

// Value returns final value
func (c *Counter) Value() int {
	return c.value
}
