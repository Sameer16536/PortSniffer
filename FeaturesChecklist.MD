🟢 Phase 1: The Basics (Sequential)
[ ] 1. Setup & Flags

[ ] Create main.go.

[ ] Use the flag package to accept arguments (e.g., -host google.com or -start 1 -end 1024).

[ ] 2. The Scanner Logic

[ ] Write a function scanPort(protocol, hostname, port).

[ ] Use net.DialTimeout to try connecting. If it connects, the port is OPEN. If it times out, it is CLOSED.

[ ] 3. Sequential Execution

[ ] Loop from start_port to end_port.

[ ] Call scanPort one by one.

[ ] Measure time taken (It will be slow).

🟡 Phase 2: The Speed (Concurrent)
[ ] 4. Goroutines

[ ] Wrap the scanPort call with the go keyword.

[ ] The Problem: Observer the program exiting immediately (because main finishes before workers).

[ ] 5. Synchronization (The Fix)

[ ] Implement sync.WaitGroup.

[ ] Add(1) before starting a thread.

[ ] Done() when thread finishes.

[ ] Wait() in main to block until finished.

🔴 Phase 3: The Safety (Thread-Safe Data)
[ ] 6. Gathering Results

[ ] Create a slice var openPorts []int.

[ ] The Crash: Try appending to this slice from inside Goroutines. (Witness the Race Condition).

[ ] 7. Mutex Locking

[ ] Use sync.Mutex.

[ ] Lock() before appending.

[ ] Unlock() after appending.

[ ] 8. Final Polish

[ ] Sort the results (because threads finish in random order).

[ ] Print "Open Ports: 80, 443, 8080".