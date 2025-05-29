package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zeusnotfound04/DecoyOps/internal/executor"
	"github.com/zeusnotfound04/DecoyOps/internal/logger"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
	Bold   = "\033[1m"
)

func displayBanner() {
	fmt.Printf("%s%s", Cyan, Bold)
	fmt.Println(`
 ____  ______ _____ ______     __  ____  _____   _____ 
|  _ \|  ____/ ____/ __ \ \   / / / __ \|  __ \ / ____|
| | | | |__| |   | |  | \ \_/ / | |  | | |__) | (___  
| | | |  __| |   | |  | |\   /  | |  | |  ___/ \___ \ 
| |_| | |__| |___| |__| | | |   | |__| | |     ____) |
|____/|______\____\____/  |_|    \____/|_|    |_____/ `)

	fmt.Printf("%s%s", Yellow, Bold)
	fmt.Println(`
           ◄═══ Advanced Deception Operations ═══►`)
	fmt.Printf("%s%s", Green, Bold)
	fmt.Println(`
    [+] Process Analysis & System Reconnaissance Tool
    [+] File System Enumeration Engine
    [+] Network Service Discovery Module
    [*] Version 1.0.0 - Ready for Deployment
`)
	fmt.Printf("%s", Reset)
}

func displaySummary(techniques map[string]bool) {
	fmt.Printf("\n%s%s═══════════════════ Execution Summary ═══════════════════%s\n", Purple, Bold, Reset)

	executedCount := 0
	for _, executed := range techniques {
		if executed {
			executedCount++
		}
	}

	fmt.Printf("%s%s[*] Total Techniques Executed: %d%s\n", Green, Bold, executedCount, Reset)

	if techniques["process"] {
		fmt.Printf("%s%s[+] Process Discovery - Completed%s\n", Cyan, Bold, Reset)
	}
	if techniques["file"] {
		fmt.Printf("%s%s[+] File Enumeration - Completed%s\n", Cyan, Bold, Reset)
	}
	if techniques["network"] {
		fmt.Printf("%s%s[+] Network Scanning - Completed%s\n", Cyan, Bold, Reset)
	}

	fmt.Printf("\n%s%s[*] Results saved in: output/output.json%s\n", Yellow, Bold, Reset)
	fmt.Printf("%s%s═════════════════════════════════════════════════════%s\n", Purple, Bold, Reset)
}

func displayHelp() {
	fmt.Printf("%s%s╔════════════════ DecoyOps Commands ════════════════╗%s\n", Cyan, Bold, Reset)
	fmt.Printf("%s%s║                                                   ║%s\n", Cyan, Bold, Reset)
	fmt.Printf("%s%s║  Available Commands:                             ║%s\n", Cyan, Bold, Reset)
	fmt.Printf("%s%s║                                                   ║%s\n", Cyan, Bold, Reset)
	fmt.Printf("%s%s║  1. Process Discovery:                           ║%s\n", Yellow, Bold, Reset)
	fmt.Printf("%s%s║     --process-discovery                          ║%s\n", White, Bold, Reset)
	fmt.Printf("%s%s║     Enumerates running processes and their info  ║%s\n", Green, Bold, Reset)
	fmt.Printf("%s%s║                                                   ║%s\n", Cyan, Bold, Reset)
	fmt.Printf("%s%s║  2. File Enumeration:                            ║%s\n", Yellow, Bold, Reset)
	fmt.Printf("%s%s║     --file-enum                                  ║%s\n", White, Bold, Reset)
	fmt.Printf("%s%s║     Lists files and directories in current path  ║%s\n", Green, Bold, Reset)
	fmt.Printf("%s%s║                                                   ║%s\n", Cyan, Bold, Reset)
	fmt.Printf("%s%s║  3. Network Scanning:                            ║%s\n", Yellow, Bold, Reset)
	fmt.Printf("%s%s║     --Network-scan                               ║%s\n", White, Bold, Reset)
	fmt.Printf("%s%s║     Discovers open ports and network services    ║%s\n", Green, Bold, Reset)
	fmt.Printf("%s%s║                                                   ║%s\n", Cyan, Bold, Reset)
	fmt.Printf("%s%s║  Usage Examples:                                 ║%s\n", Purple, Bold, Reset)
	fmt.Printf("%s%s║  DecoyOps.exe --process-discovery               ║%s\n", White, Bold, Reset)
	fmt.Printf("%s%s║  DecoyOps.exe --file-enum                       ║%s\n", White, Bold, Reset)
	fmt.Printf("%s%s║  DecoyOps.exe --Network-scan                    ║%s\n", White, Bold, Reset)
	fmt.Printf("%s%s║                                                   ║%s\n", Cyan, Bold, Reset)
	fmt.Printf("%s%s╚═══════════════════════════════════════════════════╝%s\n", Cyan, Bold, Reset)
}

func main() {
	displayBanner()

	helpFlag := flag.Bool("help", false, "Display detailed help information")
	processDiscovery := flag.Bool("process-discovery", false, "Enumerate running processes and their details")
	fileEnum := flag.Bool("file-enum", false, "List files and directories in the current path")
	NetworkScan := flag.Bool("Network-scan", false, "Discover open ports and network services")

	flag.Parse()

	if *helpFlag {
		displayHelp()
		os.Exit(0)
	}

	logger.InitLogger()

	techniques := make(map[string]bool)

	if *processDiscovery {
		executor.ProcessDiscovery()
		techniques["process"] = true
	}

	if *fileEnum {
		executor.FileEnumeration()
		techniques["file"] = true
	}

	if *NetworkScan {
		executor.NetworkScan()
		techniques["network"] = true
	}

	if !*processDiscovery && !*fileEnum && !*NetworkScan {
		fmt.Printf("%s%s", Yellow, Bold)
		fmt.Println("No technique selected. Use --help to see available options.")
		fmt.Printf("%s", Reset)
	} else {
		displaySummary(techniques)
	}
}
