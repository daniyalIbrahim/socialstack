package util

import "os"

var (
	logger = GetLogger()
)

func CreateDir(path string) {
	logger.Errorf("Creating Directory: %v", path)
	// Create dir if not exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			logger.Errorf("Error: %v", err)
		}
	}
}
