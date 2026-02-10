package main

// This version adds concurremcy (goroutines).

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		go func(j int) { // Add a new goroutine: a concurrent execution thread. It takes a port number.
			// It runs this function concurrently for each i.
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address) // We wrap the Dial(network, address string) call in a goroutine, so many dials happen at once.
			if err != nil {                       // If dialing failed, the goroutine exits immediately via return.
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j) // Uses j (the goroutine’s copy), not i.
		}(i) // Immediately calls the anonymous function. It passes i as the argument, so inside the goroutine: j gets the value of i at that moment. So each goroutine gets its own port number snapshot.
	}
}
