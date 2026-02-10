package main

// Synchronized Scanning using WaitGroup involves launching multiple goroutines:
// wg.Add(int): Increments the counter by the number of goroutines to wait for.
// wg.Done(): Decrements the counter by one, called when a goroutine finishes.
// wg.Wait(): Blocks execution until the counter reaches zero.

// It involves launching multiple goroutines, tracking them with a counter,
// and blocking the main program until all workers report they are finished.
// It ensures the program doesn't exit before all tasks are complete

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
}
