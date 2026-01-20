package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		human    bool
		expected string
	}{
		{"1 EB human", 1152921504606846976, true, "1.0EB"},
		{"1 PB human", 1125899906842624, true, "1.0PB"},
		{"1 TB human", 1099511627776, true, "1.0TB"},
		{"1 GB human", 1073741824, true, "1.0GB"},
		{"1 MB human", 1048576, true, "1.0MB"},
		{"2 KB human", 2048, true, "2.0KB"},
		{"512 B human", 512, true, "512B"},
		{"5.1 MB human", 5109122, true, "4.9MB"},

		{"1 EB bytes", 1152921504606846976, false, "1152921504606846976B"},
		{"1 PB bytes", 1125899906842624, false, "1125899906842624B"},
		{"1 TB bytes", 1099511627776, false, "1099511627776B"},
		{"1 GB bytes", 1073741824, false, "1073741824B"},
		{"1 MB bytes", 1048576, false, "1048576B"},
		{"2 KB bytes", 2048, false, "2048B"},
		{"512 B bytes", 512, false, "512B"},
		{"5.1 MB bytes", 5109122, false, "5109122B"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatSize(tt.size, tt.human)
			assert.Equal(t, tt.expected, result)
		})
	}
}
