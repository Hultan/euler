package tools

import (
	"sync"
)

// SafeInt64Counter is safe to use concurrently.
type SafeFloat64Counter struct {
	value float64
	mux   sync.Mutex
}

// Increment increments the counter for the given key.
func (c *SafeFloat64Counter) Increment64(newValue float64) {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mux.Lock()
	defer c.mux.Unlock()
	c.value += newValue
}

// Value returns the current value of the counter for the given key.
func (c *SafeFloat64Counter) Value() float64 {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.value
}

// Value returns the current value of the counter for the given key.
func (c *SafeFloat64Counter) Reset() {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mux.Lock()
	defer c.mux.Unlock()
	c.value = 0
}
