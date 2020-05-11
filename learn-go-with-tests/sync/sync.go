package sync

import "sync"

// Counter counts
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc increments value by 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns final value
func (c *Counter) Value() int {
	return c.value
}
