package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync.Map version - safe for concurrent use
	var m sync.Map
	var wg sync.WaitGroup

	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func(g int) {
			for i := 0; i < 1000; i++ {
				key := g*1000 + i
				m.Store(key, i) // Concurrent write - SAFE!
			}
			wg.Done()
		}(g)
	}

	wg.Wait()

	// Count entries using Range
	count := 0
	m.Range(func(key, value interface{}) bool {
		count++
		return true
	})

	fmt.Println("sync.Map - Expected length: 50,000")
	fmt.Println("sync.Map - Actual length:", count)
}