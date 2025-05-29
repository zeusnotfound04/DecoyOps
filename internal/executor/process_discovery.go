package executor

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/zeusnotfound04/DecoyOps/internal/logger"
)

type ProcessInfo struct {
	Name      string
	PID       string
	Memory    string
	CPU       string
	Status    string
	StartTime string
}

type ProcessStats struct {
	TotalProcesses   int
	TotalMemoryUsage float64
	RunningProcesses int
	SystemProcesses  int
	UserProcesses    int
	TopMemoryProcess ProcessInfo
}

func calculateProcessStats(processes []ProcessInfo) ProcessStats {
	stats := ProcessStats{}
	stats.TotalProcesses = len(processes)

	for _, proc := range processes {
		if runtime.GOOS == "windows" {
			status := strings.ToLower(strings.TrimSpace(proc.Status))
			if status == "" || status == "unknown" || !strings.Contains(status, "stopped") &&
				!strings.Contains(status, "terminated") && !strings.Contains(status, "suspended") {
				stats.RunningProcesses++
			}
		} else {
			if strings.Contains(strings.ToLower(proc.Status), "running") ||
				strings.Contains(strings.ToLower(proc.Status), "sleep") ||
				strings.Contains(strings.ToLower(proc.Status), "wait") {
				stats.RunningProcesses++
			}
		}

		if strings.HasSuffix(proc.Memory, "MB") {
			memStr := strings.TrimSuffix(proc.Memory, " MB")
			if mem, err := strconv.ParseFloat(memStr, 64); err == nil {
				stats.TotalMemoryUsage += mem

				currentTopMem := 0.0
				if strings.HasSuffix(stats.TopMemoryProcess.Memory, "MB") {
					currentTopMem, _ = strconv.ParseFloat(strings.TrimSuffix(stats.TopMemoryProcess.Memory, " MB"), 64)
				}
				if mem > currentTopMem {
					stats.TopMemoryProcess = proc
				}
			}
		}

		processName := strings.ToLower(proc.Name)
		if strings.HasPrefix(processName, "system") ||
			strings.HasPrefix(processName, "svc") ||
			strings.HasPrefix(processName, "win") ||
			strings.Contains(processName, "service") {
			stats.SystemProcesses++
		} else {
			stats.UserProcesses++
		}
	}

	return stats
}

func ProcessDiscovery() {
	var cmd *exec.Cmd
	var processInfos []ProcessInfo

	if runtime.GOOS == "windows" {
		cmd = exec.Command("wmic", "process", "get", "ProcessId,Name,WorkingSetSize,Status,CreationDate,ThreadCount", "/format:csv")
	} else {
		cmd = exec.Command("ps", "-eo", "pid,comm,pcpu,pmem,stat,start,time")
	}

	output, err := cmd.Output()
	if err != nil {
		logger.Log("T1057", "Process Discovery", fmt.Sprintf("Failed to execute command: %v", err))
		return
	}

	lines := strings.Split(string(output), "\n")

	if runtime.GOOS == "windows" {
		for i, line := range lines {
			if i < 2 || len(line) == 0 {
				continue
			}
			parts := strings.Split(line, ",")
			if len(parts) >= 6 {
				memoryMB := "N/A"
				if memBytes, err := strconv.ParseInt(strings.TrimSpace(parts[3]), 10, 64); err == nil {
					memoryMB = fmt.Sprintf("%.2f MB", float64(memBytes)/1024/1024)
				}

				startTime := "N/A"
				if len(parts[5]) > 14 {
					if t, err := time.Parse("20060102150405.000000-0700", strings.TrimSpace(parts[5])); err == nil {
						startTime = t.Format("2006-01-02 15:04:05")
					}
				}

				processInfos = append(processInfos, ProcessInfo{
					Name:      strings.TrimSpace(parts[2]),
					PID:       strings.TrimSpace(parts[1]),
					Memory:    memoryMB,
					Status:    strings.TrimSpace(parts[4]),
					StartTime: startTime,
					CPU:       fmt.Sprintf("Threads: %s", strings.TrimSpace(parts[6])),
				})
			}
		}
	} else {
		for i, line := range lines {
			if i == 0 || len(line) == 0 { 
				continue
			}
			fields := strings.Fields(line)
			if len(fields) >= 7 {
				pid := fields[0]
				name := fields[1]
				cpu := fields[2]
				memory := fields[3]
				status := fields[4]
				startTime := fields[5]
				runTime := fields[6]

				processInfos = append(processInfos, ProcessInfo{
					Name:      name,
					PID:       pid,
					Memory:    fmt.Sprintf("%.2f %%", parseFloat(memory)),
					Status:    status,
					StartTime: startTime,
					CPU:       fmt.Sprintf("CPU: %s%%, Runtime: %s", cpu, runTime),
				})
			}
		}
	}

	stats := calculateProcessStats(processInfos)

	var details strings.Builder
	details.WriteString("\n🔍 Process Discovery Report\n")
	details.WriteString("════════════════════════════════\n")

	details.WriteString("📊 Summary Statistics:\n")
	details.WriteString(fmt.Sprintf("  ├─ Total Processes: %d\n", stats.TotalProcesses))
	details.WriteString(fmt.Sprintf("  ├─ Running Processes: %d\n", stats.RunningProcesses))
	details.WriteString(fmt.Sprintf("  ├─ System Processes: %d\n", stats.SystemProcesses))
	details.WriteString(fmt.Sprintf("  ├─ User Processes: %d\n", stats.UserProcesses))
	details.WriteString(fmt.Sprintf("  ├─ Total Memory Usage: %.2f MB\n", stats.TotalMemoryUsage))
	details.WriteString(fmt.Sprintf("  └─ Average Memory Per Process: %.2f MB\n", stats.TotalMemoryUsage/float64(stats.TotalProcesses)))
	details.WriteString("\n")

	details.WriteString("💾 Highest Memory Consumer:\n")
	details.WriteString(fmt.Sprintf("  ├─ Process: %s\n", stats.TopMemoryProcess.Name))
	details.WriteString(fmt.Sprintf("  ├─ PID: %s\n", stats.TopMemoryProcess.PID))
	details.WriteString(fmt.Sprintf("  └─ Memory: %s\n", stats.TopMemoryProcess.Memory))
	details.WriteString("\n")

	details.WriteString("📦 Process Details:\n")
	details.WriteString("────────────────────────────────\n")

	for _, proc := range processInfos {
		details.WriteString(fmt.Sprintf("Process: %s\n", proc.Name))
		details.WriteString(fmt.Sprintf("  ├─ 🔢 PID: %s\n", proc.PID))
		details.WriteString(fmt.Sprintf("  ├─ 💾 Memory: %s\n", proc.Memory))
		details.WriteString(fmt.Sprintf("  ├─ ⚡ Status: %s\n", proc.Status))
		details.WriteString(fmt.Sprintf("  ├─ 🔄 %s\n", proc.CPU))
		details.WriteString(fmt.Sprintf("  └─ 🕒 Start Time: %s\n", proc.StartTime))
		details.WriteString("────────────────────────────────\n")
	}

	details.WriteString("\n📋 Final Summary:\n")
	details.WriteString("════════════════════════════════\n")
	details.WriteString(fmt.Sprintf("✅ Scanned %d processes successfully\n", stats.TotalProcesses))
	details.WriteString(fmt.Sprintf("📈 System Load: %.2f%% processes running\n",
		float64(stats.RunningProcesses)/float64(stats.TotalProcesses)*100))
	details.WriteString(fmt.Sprintf("💻 System/User Process Ratio: %.2f\n",
		float64(stats.SystemProcesses)/float64(stats.UserProcesses)))
	details.WriteString("════════════════════════════════")

	logger.Log("T1057", "Process Discovery", details.String())
}

func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return f
}
