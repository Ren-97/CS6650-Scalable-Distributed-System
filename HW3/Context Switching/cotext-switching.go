package main

import (
	"fmt"
	"runtime"
	"time"
)

func pingPongSingleChannel(max int) time.Duration {
	ch := make(chan struct{})
	done := make(chan struct{})

	start := time.Now()

	go func() { // Goroutine A
		for i := 0; i < max; i++ {
			ch <- struct{}{}  // A sends signal to B (blocks until B receives)
			<-ch              // A receives signal from B 
		}
		close(done)
	}()

	go func() { // Goroutine B
		for i := 0; i < max; i++ {
			<-ch              // B receives signal from A 
			ch <- struct{}{}  // B sends signal to A (blocks until A receives)
		}
	}()

	<-done
	return time.Since(start)
}

func main() {
	const iterations = 1000000

	// Test 1: Single thread
	runtime.GOMAXPROCS(1)
	fmt.Println("=== Single Thread Mode ===")
	duration1 := pingPongSingleChannel(iterations)
	avgSwitch1 := duration1 / (2 * time.Duration(iterations))
	fmt.Printf("Total time: %v\n", duration1)
	fmt.Printf("Average switch time: %v\n", avgSwitch1)
	fmt.Printf("Switches per second: %.0f\n", float64(2*iterations)/duration1.Seconds())

	// Test 2: Multi-thread
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("\n=== Multi-Thread Mode ===")
	duration2 := pingPongSingleChannel(iterations)
	avgSwitch2 := duration2 / (2 * time.Duration(iterations))
	fmt.Printf("Total time: %v\n", duration2)
	fmt.Printf("Average switch time: %v\n", avgSwitch2)
	fmt.Printf("Switches per second: %.0f\n", float64(2*iterations)/duration2.Seconds())

	// Result comparison
	fmt.Println("\n=== Result Comparison ===")
	if avgSwitch1 < avgSwitch2 {
		fmt.Printf("Single thread is faster by %.2fx\n", float64(avgSwitch2)/float64(avgSwitch1))
	} else {
		fmt.Printf("Multi-thread is faster by %.2fx\n", float64(avgSwitch1)/float64(avgSwitch2))
	}
}