package code

import "fmt"

func formatSize(size int64, human bool) string {
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
