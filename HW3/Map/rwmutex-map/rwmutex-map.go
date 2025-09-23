package main

import (
	"fmt"
	"sync"
	"time"
)

// Container holds a map of counters with a RWMutex for synchronization
/* 
	Multiple readers can access simultaneously
	Only one writer can access at a time
	No reads during writes
	No writes during reads
*/
type Container struct {
	mu       sync.RWMutex     // RWMutex instead of regular Mutex
	counters map[string]int   // Map storing counter values
}

// inc safely increments a named counter using RWMutex
func (c *Container) inc(name string) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Ensure lock is released
	c.counters[name]++  // Critical section: increment the counter
}

// get safely reads a counter value using RWMutex
func (c *Container) get(name string) int {
	c.mu.RLock()         // Acquire read lock
	defer c.mu.RUnlock() // Ensure read lock is released
	return c.counters[name]
}

// length safely returns the map length
func (c *Container) length() int {
	c.mu.RLock()         // Read lock for reading length
	defer c.mu.RUnlock()
	return len(c.counters)
}

func main() {
	// Initialize container with initial counter values
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup
	start := time.Now()

	// Helper function to increment a counter multiple times
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name) // Safely increment the counter
		}
	}

	// Launch goroutines
	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("a", 10000)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("a", 10000)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("b", 10000)
	}()

	// Wait for all goroutines to complete
	wg.Wait()

	elapsed := time.Since(start)

	// Print results
	fmt.Println("=== RWMutex Results ===")
	fmt.Println("Final counters:", c.counters)
	fmt.Println("Map length:", c.length())
	fmt.Println("Total time taken:", elapsed)
}