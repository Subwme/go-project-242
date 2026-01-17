package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatSize(t *testing.T) {
	//human
	assert.Equal(t, "1.0EB", FormatSize(1152921504606846976, true))
	assert.Equal(t, "1.0PB", FormatSize(1125899906842624, true))
	assert.Equal(t, "1.0TB", FormatSize(1099511627776, true))
	assert.Equal(t, "1.0GB", FormatSize(1073741824, true))
	assert.Equal(t, "1.0MB", FormatSize(1048576, true))
	assert.Equal(t, "2.0KB", FormatSize(2048, true))
	assert.Equal(t, "512B", FormatSize(512, true))

	//reality size
	assert.Equal(t, "5109122B", FormatSize(5109122, false))
	assert.Equal(t, "4.9MB", FormatSize(5109122, true))

	//human=false
	assert.Equal(t, "1152921504606846976B", FormatSize(1152921504606846976, false))
	assert.Equal(t, "1125899906842624B", FormatSize(1125899906842624, false))
	assert.Equal(t, "1099511627776B", FormatSize(1099511627776, false))
	assert.Equal(t, "1073741824B", FormatSize(1073741824, false))
	assert.Equal(t, "1048576B", FormatSize(1048576, false))
	assert.Equal(t, "2048B", FormatSize(2048, false))
	assert.Equal(t, "512B", FormatSize(512, false))
}
