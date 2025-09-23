package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// Use atomic integer type as counter
	var ops atomic.Uint64
	
	// Use WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup
	
	// Launch 50 goroutines
	for i := 0; i < 50; i++ {
		wg.Add(1) // Add count for each goroutine
		
		go func() {
			// Each goroutine increments the counter 1000 times
			for j := 0; j < 1000; j++ {
				ops.Add(1) // Atomic increment
			}
			wg.Done() // Notify completion of one goroutine
		}()
	}
	
	// Wait for all goroutines to complete
	wg.Wait()
	
	// Safely read the final value
	fmt.Println("Total operations ops:", ops.Load())
}