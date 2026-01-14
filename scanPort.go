// scanPort.go: Contains the function to scan a single port on a host.
// Demonstrates TCP connection attempts with timeout and basic error handling.
package main

import (
	// For printing output
	"net"     // For network connections
	"strconv" // For converting int to string
	"time"    // For setting connection timeout
)

// scanPort tries to connect to a specific TCP port on the given host.
// If the connection is successful, the port is considered OPEN.
// Otherwise, it is assumed to be closed or filtered.
func scanPort(host string, port int) {
	// 1. Format the address as "host:port" (e.g., "scanme.nmap.org:80")
	address := host + ":" + strconv.Itoa(port)

	// 2. Attempt to establish a TCP connection with a 1-second timeout.
	// net.DialTimeout returns an error if the port is closed, filtered, or unreachable.
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		// Connection failed: port is closed or filtered, so do nothing.
		return
	}

	// 3. Connection succeeded: port is OPEN.
	// Always close the connection to avoid resource leaks.
	conn.Close()
}
