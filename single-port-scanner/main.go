package main

/*
This is a single-port reachability check.
Can I establish a TCP handshake to port 80?
It’s not an HTTP request. It just checks the TCP connect.
*/

import (
	"fmt" // Standard formatting/printing package.
	"net" // Network package.
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80") // network + address:host+port.
	// := declares and assigns variables in one step.
	// We use _ bc we only want to know whether the connection succeeds. We ignore the connection object.
	if err == nil {
		fmt.Println("Connection successful") // In Go, nil is the “no value / no error” value. So it checks whether the error is nil.
	}
}
