package main

import (
	"flag"
	"fmt"

	"github.com/zeusnotfound04/DecoyOps/internal/executor"
	"github.com/zeusnotfound04/DecoyOps/internal/logger"
)

func main() {
	processDiscovery := flag.Bool("process-discovery" , false , "Simulate process discovery")
	fileEnum := flag.Bool("file-enum" , false , "Simulate process enumeration")
	NetworkScan := flag.Bool("Network-scan" , false , "Simulate network scanning")

	flag.Parse()



	logger.InitLogger()

	if *processDiscovery {
		executor.ProcessDiscovery()
	}

	if *fileEnum {
		executor.ProcessDiscovery()
	}

	if *NetworkScan {
		executor.NetworkScan()
	}

	if !*processDiscovery && !*fileEnum && !*NetworkScan {
        fmt.Println("No technique selected. Use --help to see available options.")
    }
}