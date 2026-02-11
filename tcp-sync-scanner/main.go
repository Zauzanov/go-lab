package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) { // a bidirectional channel(can receive and send; inside worker only receives) that carries int values (port numbers).
	// Plus a pointer to a WaitGroup — all workers share and update the same WaitGroup.
	for p := range ports { // Reads from the channel in a loop. range ports keeps receiving values until the channel is closed and drained. Each received value is assigned to p.
		// It keeps working as long as there are port numbers coming in.
		fmt.Println(p) // Prints the port number p.
		wg.Done()      // Decrements the WaitGroup counter by 1, signaling that it finished processing one job.
	}
}

func main() {
	ports := make(chan int, 100)      // Creates a buffered channel of int, buffer size: 100.
	var wg sync.WaitGroup             // Declares a WaitGroup named wg (zero value ready to use).
	for i := 0; i < cap(ports); i++ { // Run the loop 100 times/Start 100 worker goroutines: loop to start workers. cap(ports) returns the channel’s buffer capacity (here, 100). So this spawns 100 worker goroutines.
		// capacity 100 doesn't mean only 100 ports total. Capacity only controls how many jobs can be queued up waiting at once, not how many jobs can be processed overall.
		go worker(ports, &wg) // Starts a goroutine running worker. Passes: ports channel (shared job queue); &wg address of wg (pointer) so all workers call Done() on the same WaitGroup.
	}
	for i := 1; i <= 1024; i++ { // Loops over port numbers 1..1024.
		wg.Add(1)  // Increments the WaitGroup counter.
		ports <- i // Buffer size matters: sending (ports <- i) can proceed without blocking until the buffer is full.
		// Sends the integer i into the ports channel. This is enqueueing a job. If the channel buffer is full, this send blocks until a worker receives a value.
	}
	wg.Wait()    // Blocks until the WaitGroup counter reaches 0: until workers have called Done() exactly 1024 times.
	close(ports) // Closes the channel.

}
