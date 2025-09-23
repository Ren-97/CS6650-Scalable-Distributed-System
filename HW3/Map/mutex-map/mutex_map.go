package main

import (
	"fmt"
	"sync"
	"time"
)

// Container holds a map of counters with a mutex for synchronization
type Container struct {
	mu       sync.Mutex     // Mutex to protect concurrent access
	counters map[string]int // Map storing counter values
}

// inc safely increments a named counter
func (c *Container) inc(name string) {
	c.mu.Lock()         // Acquire lock before accessing shared data
	defer c.mu.Unlock() // Ensure lock is released when function exits
	c.counters[name]++  // Critical section: increment the counter
}

// length safely returns the number of keys in the map
func (c *Container) length() int {
	c.mu.Lock()         // Need to lock for reading map structure
	defer c.mu.Unlock()
	return len(c.counters)
}

func main() {
	// Initialize container with initial counter values
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish

	// Record start time
	startTime := time.Now()

	// Helper function to increment a counter multiple times
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ { // Loop n times
			c.inc(name) // Safely increment the counter
		}
	}

	// Launch first goroutine: increment counter "a" 10,000 times
	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("a", 10000)
	}()

	// Launch second goroutine: also increment counter "a" 10,000 times
	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("a", 10000)
	}()

	// Launch third goroutine: increment counter "b" 10,000 times
	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("b", 10000)
	}()

	// Wait for all goroutines to complete their work
	wg.Wait()

	// Calculate total time taken
	totalTime := time.Since(startTime)

	// Print the required results
	fmt.Println("=== Mutex Protected HashMap Results ===")
	fmt.Printf("Length of hashmap: %d\n", c.length())
	fmt.Printf("Total time taken: %v\n", totalTime)
	fmt.Printf("Final counter values: %v\n", c.counters)
	fmt.Printf("Expected values: map[a:20000 b:10000]\n")
}
