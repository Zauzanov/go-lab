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
	ports := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)

}
