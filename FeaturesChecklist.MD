
🟢 Phase 1: The Basics (Sequential)
[x] 1. Setup & Flags

[x] Create main.go.

[ ] Use the flag package to accept arguments (e.g., -host google.com or -start 1 -end 1024).

[x] 2. The Scanner Logic

[x] Write a function scanPort(protocol, hostname, port).

[x] Use net.DialTimeout to try connecting. If it connects, the port is OPEN. If it times out, it is CLOSED.

[x] 3. Sequential Execution

[x] Loop from start_port to end_port.

[x] Call scanPort one by one.

[x] Measure time taken (It will be slow).


🟡 Phase 2: The Speed (Concurrent)
[x] 4. Goroutines

[x] Wrap the scanPort call with the go keyword.

[x] The Problem: Observer the program exiting immediately (because main finishes before workers).

[x] 5. Synchronization (The Fix)

[x] Implement sync.WaitGroup.

[x] Add(1) before starting a thread.

[x] Done() when thread finishes.

[x] Wait() in main to block until finished.


🔴 Phase 3: The Safety (Thread-Safe Data)
[x] 6. Gathering Results

[x] Create a slice var openPorts []int.

[x] The Crash: Try appending to this slice from inside Goroutines. (Witness the Race Condition).

[x] 7. Mutex Locking

[x] Use sync.Mutex.

[x] Lock() before appending.

[x] Unlock() after appending.

[x] 8. Final Polish

[x] Sort the results (because threads finish in random order).

[x] Print "Open Ports: 80, 443, 8080".