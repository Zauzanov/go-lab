package main

// Synchronized Scanning using WaitGroup involves launching multiple goroutines:
// wg.Add(int): Increments the counter by the number of goroutines to wait for.
// wg.Done(): Decrements the counter by one, called when a goroutine finishes.
// wg.Wait(): Blocks execution until the counter reaches zero.

// It involves launching multiple goroutines, tracking them with a counter,
// and blocking the main program until all workers report they are finished.
// It ensures the program doesn't exit before all tasks are complete

// Without goroutines: check port 1, wait, check port 2, wait… (slow)
// With goroutines: kick off checks for many ports at once (fast). It doesn't block, it runs in the background while we continue.

import (
	"fmt"
	"net"
	"sync" // synchronization primitives (WaitGroup)
)

func main() {
	var wg sync.WaitGroup // Declares a variable named wg, where type is sync.WaitGroup.
	for i := 1; i <= 1024; i++ {
		wg.Add(1) // Increments the WaitGroup’s internal counter by 1.
		go func(j int) {
			defer wg.Done() // defer schedules a call to run when the surrounding function returns.
			// wg.Done() subtracts 1 from the WaitGroup counter.
			// No matter how this goroutine exits(ok, error...), it marks it finished.
			// This is important because we do return on error later.
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return // Because we used defer wg.Done(), the WaitGroup counter still gets decremented even on this early return.
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait() // Don’t exit until every goroutine has called Done().
	// Blocks main() here until the WaitGroup counter returns to 0.
	// main waits for them all to finish via wg.Wait().
	// If main exits, goroutines die. That’s why we used a WaitGroup (wg.Wait()), to keep the program alive until all scans finish.
}
