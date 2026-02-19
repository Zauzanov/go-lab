package main

import (
	"fmt"
	"net"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
}
