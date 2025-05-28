package executor

import (
	"fmt"
	"io/ioutil"

	"github.com/zeusnotfound04/DecoyOps/internal/logger"
)

func FileEnumeration() {
	files, err := ioutil.ReadDir(".")


	if err != nil {
		logger.Log("T1083" , "File and Directory Discovery" , "Failed to enumerate files")
		return
	}

	logger.Log("T1083" , "File and Directory Discovery" , fmt.Sprintf("Enumerated %d files in current directory", len(files)))
}

