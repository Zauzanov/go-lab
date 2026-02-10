package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ { // keep going while i is ≤ 1024. Increment i by 1 after each iteration.
		address := fmt.Sprintf("scanme.nmap.org:%d", i) // Returns a formatted string (it does not print).
		// %d means “insert an integer here (base 10)”.
		// i is the value used for %d.
		conn, err := net.Dial("tcp", address) // Tries to open a TCP connection to that address.
		// conn (net.Conn): the connection object if it succeeds
		//
		if err != nil { // means the dial failed.
			// port is closed or filtered
			continue // Skips the rest of the loop body and jumps to the next port
		}
		conn.Close() // Closes the TCP connection immediately.
		fmt.Printf("%d open\n", i)
	}
}
