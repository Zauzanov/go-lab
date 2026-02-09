# single-port-scanner written in Go
This is a single-port reachability check. Can I establish a TCP handshake to port 80? It’s not an HTTP request. It just checks the TCP connect. 
## 1. Run the source code:
```bash
go run main.go
Connection successful
```
## 2. Compile an executable binary file:
```bash
go build -ldflags "-w -s" main.go
./main.exe
```