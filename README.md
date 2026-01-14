# PortSniffer

PortSniffer is a simple, concurrent port scanner written in Go. It scans a range of ports on a specified host to determine which ports are open. The project demonstrates key Go concurrency concepts such as goroutines, WaitGroups, and mutexes, and is structured to be easy to understand and extend.

## What does it do?
- Scans a target host (e.g., scanme.nmap.org) for open TCP ports in a specified range.
- Uses concurrency to speed up the scanning process.
- Collects and displays a sorted list of open ports.
- Demonstrates safe concurrent access to shared data.

## How does it work?
1. **Target Selection**: The host to scan is specified in the code (default: scanme.nmap.org).
2. **Port Range**: The scanner loops through a range of ports (75 to 85 by default).
3. **Concurrency**: Each port scan runs in its own goroutine, allowing multiple ports to be checked simultaneously.
4. **Synchronization**: A `sync.WaitGroup` ensures the main function waits for all scans to complete.
5. **Thread-Safe Data**: A `sync.Mutex` protects the shared slice of open ports from race conditions.
6. **Result Collection**: Open ports are collected, sorted, and printed at the end.

## Key Go Concepts Used

### Goroutines
Goroutines are lightweight threads managed by the Go runtime. By prefixing a function call with `go`, it runs concurrently with other goroutines. In PortSniffer, each port scan runs in its own goroutine, enabling fast, parallel scanning.

### WaitGroup
A `sync.WaitGroup` is used to wait for a collection of goroutines to finish. We call `Add(1)` before starting a goroutine, `Done()` when it finishes, and `Wait()` in main to block until all are done.

### Mutex
A `sync.Mutex` is used to ensure that only one goroutine at a time can append to the shared `openPorts` slice. This prevents race conditions and data corruption.

### TCP Connections with Timeout
The scanner uses `net.DialTimeout` to attempt a TCP connection to each port, with a 1-second timeout. If the connection succeeds, the port is open; otherwise, it is closed or filtered.

## Features Checklist

🟢 Phase 1: The Basics (Sequential)
- [x] 1. Setup & Flags
- [x] Create main.go.
- [ ] Use the flag package to accept arguments (e.g., -host google.com or -start 1 -end 1024).
- [x] 2. The Scanner Logic
- [x] Write a function scanPort(protocol, hostname, port).
- [x] Use net.DialTimeout to try connecting. If it connects, the port is OPEN. If it times out, it is CLOSED.
- [x] 3. Sequential Execution
- [x] Loop from start_port to end_port.
- [x] Call scanPort one by one.
- [x] Measure time taken (It will be slow).

🟡 Phase 2: The Speed (Concurrent)
- [x] 4. Goroutines
- [x] Wrap the scanPort call with the go keyword.
- [x] The Problem: Observer the program exiting immediately (because main finishes before workers).
- [x] 5. Synchronization (The Fix)
- [x] Implement sync.WaitGroup.
- [x] Add(1) before starting a thread.
- [x] Done() when thread finishes.
- [x] Wait() in main to block until finished.

🔴 Phase 3: The Safety (Thread-Safe Data)
- [x] 6. Gathering Results
- [x] Create a slice var openPorts []int.
- [x] The Crash: Try appending to this slice from inside Goroutines. (Witness the Race Condition).
- [x] 7. Mutex Locking
- [x] Use sync.Mutex.
- [x] Lock() before appending.
- [x] Unlock() after appending.
- [x] 8. Final Polish
- [x] Sort the results (because threads finish in random order).
- [x] Print "Open Ports: 80, 443, 8080".

## Usage

1. Run the program:
   ```sh
   go run .
   ```
2. The scanner will print open ports and the time taken.

## Future Improvements
- Add command-line flags for host and port range selection.
- Support for scanning UDP ports.
- Improved output formatting.
