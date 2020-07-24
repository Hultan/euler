package tools

import (
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	value float64
	mux   sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(newValue int) {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mux.Lock()
	defer c.mux.Unlock()
	//fmt.Printf("Adding %v to sum : %f\n", newValue, c.value)
	c.value += float64(newValue)
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value() float64 {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.value
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Reset() {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mux.Lock()
	defer c.mux.Unlock()
	c.value = 0
}
