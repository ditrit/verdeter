package verdeter

import (
	"os"
)

// Return true if the file exists
func FileExist(path string) bool {
	if stats, err := os.Stat(path); !os.IsNotExist(err) {
		return !stats.IsDir()
	}
	return false
}
