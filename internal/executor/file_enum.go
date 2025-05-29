package executor

import (
	"fmt"
	"os"
	"strings"

	"github.com/zeusnotfound04/DecoyOps/internal/logger"
)

func FileEnumeration() {
	files, err := os.ReadDir(".")

	if err != nil {
		logger.Log("T1083", "File and Directory Discovery", "Failed to enumerate files")
		return
	}

	var fileDetails []string
	for _, file := range files {
		info, err := file.Info()
		if err == nil {
			fileDetails = append(fileDetails, fmt.Sprintf("%s (Size: %d bytes, Modified: %s)",
				file.Name(),
				info.Size(),
				info.ModTime().Format("2006-01-02 15:04:05")))
		}
	}

	logger.Log("T1083", "File and Directory Discovery",
		fmt.Sprintf("Enumerated %d files in current directory:\n- %s",
			len(files),
			strings.Join(fileDetails, "\n- ")))
}
