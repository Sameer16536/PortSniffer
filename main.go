package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. The Target
	// scanme.nmap.org is a service explicitly built for testing scanners.

	host := "scanme.nmap.org"

	fmt.Printf("Scanning %s (Sequential)...\n", host)
	start := time.Now()

	// 2. The Loop (Ports 75 to 85)
	// We only check a few because checking 1000 sequentially takes forever.

	for port := 75; port <= 85; port++ {
		scanPort(host, port)
	}

	elapsed := time.Since(start)
	fmt.Printf("Scanning completed in %s\n", elapsed)
}
