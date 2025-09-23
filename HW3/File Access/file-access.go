package main

/*
System calls: Unbuffered triggers syscall for every write, Buffered batches calls
Performance: Buffered is typically tens of times faster
Data safety: Unbuffered is safer, Buffered may lose unflushed data
Memory usage: Buffered requires additional memory space
Simple choice: Use Buffered for most cases, Unbuffered for critical data
*/

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	const n = 100000
	const file = "test.txt"
	
	// Unbuffered write
	start1 := time.Now()
	f1, _ := os.Create("unbuffered_" + file)
	for i := 0; i < n; i++ {
		f1.Write([]byte("Hello World\n"))
	}
	f1.Close()
	t1 := time.Since(start1)
	
	// Buffered write
	start2 := time.Now()
	f2, _ := os.Create("buffered_" + file)
	w := bufio.NewWriter(f2)
	for i := 0; i < n; i++ {
		w.WriteString("Hello World\n")
	}
	w.Flush()
	f2.Close()
	t2 := time.Since(start2)
	
	fmt.Printf("Unbuffered: %v\n", t1)
	fmt.Printf("Buffered: %v\n", t2)
	fmt.Printf("Buffered is %.1f times faster\n", float64(t1)/float64(t2))
}