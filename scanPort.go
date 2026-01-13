package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func scanPort(host string, port int) {
	// 3. Format address
	// (e.g., "scanme.nmap.org:80")

	address := host + ":" + strconv.Itoa(port)

	// 4. Try to Connect (TCP)
	// We give it 1 second timeout.
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		// Port is closed or filtered.
		return
	}

	// 5. Success!
	// If we get here, the port is OPEN.
	// We must close the connection, otherwise we leak resources.
	conn.Close()
	fmt.Printf("✅ Port %d is OPEN\n", port)
}
