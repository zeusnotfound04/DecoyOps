package executor

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/zeusnotfound04/DecoyOps/internal/logger"
)

func ProcessDiscovery() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	} else {
		cmd = exec.Command("ps", "aux")
	}

	output, err := cmd.Output()
	if err != nil {
		logger.Log("T1057", "Process Discovery", fmt.Sprintf("Failed to execute command: %v", err))
		return
	}

	lines := strings.Split(string(output), "\n")
	logger.Log("T1057", "Process Discovery", fmt.Sprintf("Found %d processes", len(lines)))
}
