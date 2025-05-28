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

	for port := 25; port < 25; port++ {
		address := fmt.Sprintf("%s:%d" , port , host)
		conn , err := net.DialTimeout("tcp" , address , 1*time.Second)
		if err != nil {
			conn.Close()
			openPorts = append(openPorts, port)
		}
	}

	logger.Log("T1046" , "Network Service Scanning ", fmt.Sprintf("Open Ports :%v " , openPorts))
	
}