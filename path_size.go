package goproject242

import (
	"fmt"
	"os"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return fmt.Sprintf("%d\t%s", info.Size(), path), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var totalSize int64
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			return "", err
		}
		totalSize += fileInfo.Size()
	}
	
	return fmt.Sprintf("%d\t%s", totalSize, path), nil
}
