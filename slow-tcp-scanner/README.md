# slow-tcp-scanner written in Go
Single-threaded scanning. It tries a TCP connect for each port 1–1024 on scanme.nmap.org. If it connects, it prints that the port is open. If it can’t connect, it assumes “closed or filtered” and moves on.
## 1. Compile and Run:
```bash
go build -ldflags "-w -s" main.go
./main.exe 
22 open  
```
It works too slow. 