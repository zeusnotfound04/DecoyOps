package main

import (
	"flag"
	"fmt"

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
	
	fmt.Printf("%s%s", Green, Reset)
	fmt.Println(`
    [+] Tactical Misdirection & Counter-Intelligence Suite
    [*] Deployment Ready - Operational Security Enabled
`)
	fmt.Printf("%s", Reset)
}

func main() {
	displayBanner()

	processDiscovery := flag.Bool("process-discovery", false, "Simulate process discovery")
	fileEnum := flag.Bool("file-enum", false, "Simulate process enumeration")
	NetworkScan := flag.Bool("Network-scan", false, "Simulate network scanning")

	flag.Parse()

	logger.InitLogger()

	if *processDiscovery {
		executor.ProcessDiscovery()
	}

	if *fileEnum {
		executor.FileEnumeration()
	}

	if *NetworkScan {
		executor.NetworkScan()
	}

	if !*processDiscovery && !*fileEnum && !*NetworkScan {
		fmt.Printf("%s%s", Yellow, Bold)
		fmt.Println("No technique selected. Use --help to see available options.")
		fmt.Printf("%s", Reset)
	}
}