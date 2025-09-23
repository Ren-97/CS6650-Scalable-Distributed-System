package main
// race condition
import (
	"fmt"
	"sync"
)

func main() {
	// Use regular integer (non-atomic)
	var ops int
	var wg sync.WaitGroup

	// Launch 50 goroutines
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				// Non-atomic operation
				ops++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Total operations ops:", ops)
}
