package goproject242

import (
	"fmt"
	"os"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		result := FormatSize(info.Size(), human)
		return fmt.Sprintf("%s\t%s", result, path), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var totalSize int64
	for _, file := range files {
		if !all && strings.HasPrefix(file.Name(), ".") {
			continue
		}
		
		fileInfo, err := file.Info()
		if err != nil {
			return "", err
		}
		totalSize += fileInfo.Size()
	}
	
	result := FormatSize(totalSize, human)
	return fmt.Sprintf("%s\t%s", result, path), nil
}

func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	const (
		KB = 1 << (10 * 1)
		MB = 1 << (10 * 2)
		GB = 1 << (10 * 3)
		TB = 1 << (10 * 4)
		PB = 1 << (10 * 5)
		EB = 1 << (10 * 6)
	)

	switch {
	case size >= EB:
		return fmt.Sprintf("%.1fEB", float64(size)/float64(EB))
	case size >= PB:
		return fmt.Sprintf("%.1fPB", float64(size)/float64(PB))
	case size >= TB:
		return fmt.Sprintf("%.1fTB", float64(size)/float64(TB))
	case size >= GB:
		return fmt.Sprintf("%.1fGB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.1fMB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.1fKB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%dB", size)
	}
}
