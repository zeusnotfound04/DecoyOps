package executor

import (
	"fmt"
	"net"
	"time"

	"github.com/zeusnotfound04/DecoyOps/internal/logger"
)

func NetworkScan() {
	host := "127.0.0.1"
	openPorts := []int{}

	for port := 1; port <= 65535; port++ {	
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err == nil {
			openPorts = append(openPorts, port)
			conn.Close()
		}
	}

	logger.Log("T1046", "Network Service Scanning", fmt.Sprintf("Open Ports: %v", openPorts))
}
