package tools

import (
	"sync"
)

// SafeInt64Counter is safe to use concurrently.
type SafeInt64Counter struct {
	value int64
	mux   sync.Mutex
}

// Increment increments the counter for the given key.
func (c *SafeInt64Counter) Increment(newValue int) {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mux.Lock()
	defer c.mux.Unlock()
	c.value += int64(newValue)
}

// Increment increments the counter for the given key.
func (c *SafeInt64Counter) Increment64(newValue int64) {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mux.Lock()
	defer c.mux.Unlock()
	c.value += newValue
}

// Value returns the current value of the counter for the given key.
func (c *SafeInt64Counter) Value() int64 {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.value
}

// Value returns the current value of the counter for the given key.
func (c *SafeInt64Counter) Reset() {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mux.Lock()
	defer c.mux.Unlock()
	c.value = 0
}
