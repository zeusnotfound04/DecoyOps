package executor

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/zeusnotfound04/DecoyOps/internal/logger"
)

func portStatus(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err == nil {
		conn.Close()
		return true
	}
	return false
}

func getServiceName(port int) string {
	switch port {
	case 21:
		return "FTP"
	case 22:
		return "SSH"
	case 23:
		return "Telnet"
	case 25:
		return "SMTP"
	case 53:
		return "DNS"
	case 80:
		return "HTTP"
	case 443:
		return "HTTPS"
	case 3306:
		return "MySQL"
	case 5432:
		return "PostgreSQL"
	default:
		return "Unknown"
	}
}

func NetworkScan() {
	host := "127.0.0.1"
	var openPorts []string
	commonPorts := []int{21, 22, 23, 25, 53, 80, 443, 3306, 5432, 8080}

	logger.Log("T1046", "Network Service Scanning", "Starting network scan...")

	for _, port := range commonPorts {
		if portStatus(host, port) {
			service := getServiceName(port)
			portInfo := fmt.Sprintf("Port %d (%s)", port, service)
			openPorts = append(openPorts, portInfo)
		}
	}

	if len(openPorts) > 0 {
		logger.Log("T1046", "Network Service Scanning", fmt.Sprintf("Open Ports:\n%s", strings.Join(openPorts, "\n")))
	} else {
		logger.Log("T1046", "Network Service Scanning", "No open ports found in common port range")
	}

	if runtime.GOOS == "windows" {
		cmd := exec.Command("ipconfig", "/all")
		output, err := cmd.Output()
		if err == nil {
			logger.Log("T1046", "Network Configuration", fmt.Sprintf("Network Configuration:\n%s", string(output)))
		}
	} else {
		cmd := exec.Command("ifconfig", "-a")
		output, err := cmd.Output()
		if err == nil {
			logger.Log("T1046", "Network Configuration", fmt.Sprintf("Network Configuration:\n%s", string(output)))
		}
	}
}
