package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. Set up a plain map[int]int
	m := make(map[int]int)
	var wg sync.WaitGroup

	// 2. Spawn 50 goroutines
	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func(g int) {
			// 3. Each goroutine writes 1000 key/value pairs
			for i := 0; i < 1000; i++ {
				key := g*1000 + i
				m[key] = i // Concurrent write to map!
			}
			wg.Done()
		}(g)
	}

	// 4. Wait for all goroutines to finish
	wg.Wait()

	// 5. Print the length of the map
	fmt.Println("Expected length: 50,000")
	fmt.Println("Actual length:", len(m))
}


/* The underlying operations of a map are not atomic. 
For example, m[key] = value actually involves multiple steps:

1. Calculate the hash value of the key
2. Locate the corresponding hash bucket based on the hash
3. Handle hash collisions if multiple keys map to the same bucket
4. Update or insert the key-value pair into the data structure

When two goroutines execute these steps simultaneously, they interfere with each other, 
leading to corruption of the map's internal data structure.

It is not possible to safely operate on the same "bucket" simultaneously. Concurrent operations can lead to:

1. Overwriting each other's modifications 
- When multiple goroutines write to the same bucket, their changes may overwrite one another, 
resulting in lost updates.

2. Corruption of the linked list structure 
- The internal linked list within each bucket can become damaged when multiple goroutines simultaneously 
modify pointer connections between entries.

3. Data loss or corruption 
- These conflicts can cause entries to disappear, become inaccessible, or contain inconsistent values, 
ultimately leading to map data integrity failure.
*/