package code

import (
	"os"
	"path/filepath"
	"strings"
)

func calculateSize(path string, recursive, all bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var totalSize int64
	for _, file := range files {
		if !all && strings.HasPrefix(file.Name(), ".") {
			continue
		}

		fullPath := filepath.Join(path, file.Name())

		if file.IsDir() && recursive {
			subSize, err := calculateSize(fullPath, recursive, all)
			if err != nil {
				return 0, err
			}
			totalSize += subSize
		} else {
			fileInfo, err := file.Info()
			if err != nil {
				return 0, err
			}
			totalSize += fileInfo.Size()
		}
	}

	return totalSize, nil
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := calculateSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	result := FormatSize(size, human)
	return result, nil
}
