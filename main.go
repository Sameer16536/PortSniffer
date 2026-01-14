// PortSniffer: A simple concurrent port scanner in Go
// This program scans a range of ports on a given host to check which ports are open.
// It demonstrates the use of goroutines, WaitGroup, and TCP connections with timeouts.
package main

import (
	"fmt" // For printing output
	"sort"
	"sync" // For synchronizing goroutines
	"time" // For measuring elapsed time and setting timeouts
)

func main() {
	// 1. Define the target host to scan.
	// scanme.nmap.org is a public host for testing port scanners.
	host := "scanme.nmap.org"

	// 2. Print a message and record the start time for performance measurement.
	fmt.Printf("Scanning %s (Concurrent)...\n", host)
	start := time.Now()

	// 3. Create a WaitGroup to wait for all port scan goroutines to finish.
	wg := sync.WaitGroup{}

	// Shared state
	var openPorts []int
	var mutex sync.Mutex
	// 4. Loop through the desired port range (75 to 85 inclusive).
	// For each port, start a goroutine to scan it concurrently.
	for port := 75; port <= 85; port++ {
		wg.Add(1) // Increment WaitGroup counter for each goroutine

		go func(p int) {
			defer wg.Done()   // Decrement counter when goroutine completes
			scanPort(host, p) // Scan the port
			mutex.Lock()
			openPorts = append(openPorts, p)
			mutex.Unlock()
		}(port) // Pass the current port as an argument to avoid closure issues
	}

	// 5. Wait for all port scan goroutines to finish before continuing.
	wg.Wait()

	sort.Ints(openPorts)

	// 6. Print the total time taken for the scan.
	elapsed := time.Since(start)
	fmt.Printf("Scanning completed in %s\n", elapsed)
	fmt.Printf("Open Ports: %v\n", openPorts)
}
