package executor

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/zeusnotfound04/DecoyOps/internal/logger"
)

func getFilePermissions(info os.FileInfo) string {
	mode := info.Mode()
	if runtime.GOOS == "windows" {
		attrs := make([]string, 0)
		if mode&1 != 0 {
			attrs = append(attrs, "ReadOnly")
		}
		if mode&2 != 0 {
			attrs = append(attrs, "Hidden")
		}
		if mode&4 != 0 {
			attrs = append(attrs, "System")
		}
		return strings.Join(attrs, ", ")
	} else {
		return mode.String()
	}
}

func FileEnumeration() {
	fileDetails := make([]string, 0)
	totalSize := int64(0)
	fileCount := 0
	dirCount := 0

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			dirCount++
		} else {
			fileCount++
			totalSize += info.Size()
		}

		perms := getFilePermissions(info)
		modTime := info.ModTime().Format("2006-01-02 15:04:05")
		size := fmt.Sprintf("%.2f MB", float64(info.Size())/1024/1024)

		fileType := "File"
		if info.IsDir() {
			fileType = "Directory"
		}

		detail := fmt.Sprintf("%s:\n  Path: %s\n  Size: %s\n  Modified: %s\n  Permissions: %s\n",
			fileType, path, size, modTime, perms)
		fileDetails = append(fileDetails, detail)

		return nil
	})

	if err != nil {
		logger.Log("T1083", "File and Directory Discovery", fmt.Sprintf("Error enumerating files: %v", err))
		return
	}

	summary := fmt.Sprintf("Summary:\n"+
		"Total Directories: %d\n"+
		"Total Files: %d\n"+
		"Total Size: %.2f MB\n\n"+
		"File Details:\n%s",
		dirCount,
		fileCount,
		float64(totalSize)/1024/1024,
		strings.Join(fileDetails, "\n"))

	logger.Log("T1083", "File and Directory Discovery", summary)
}
