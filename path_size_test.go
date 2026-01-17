package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSize(t *testing.T) {
	resultDataCSV, err := GetPathSize("testdata/data.csv", false, false, false)
	assert.NoError(t, err)
	assert.Equal(t, "37B", resultDataCSV)
}

func TestGetPathSizeDir(t *testing.T) {
	resultTestDir, err := GetPathSize("testdata", false, false, false)
	assert.NoError(t, err)
	assert.Contains(t, resultTestDir, "B")
}

func TestGetPathSizeWithoutAll(t *testing.T) {
	result, err := GetPathSize("testdata", false, false, false)
	assert.NoError(t, err)
	assert.Contains(t, result, "B")
}

func TestGetPathSizeWithAll(t *testing.T) {
	result, err := GetPathSize("testdata", false, false, true)
	assert.NoError(t, err)
	assert.Contains(t, result, "B")
}

func TestGetPathSizeDiff(t *testing.T) {
	resultWithoutAll, err := GetPathSize("testdata", false, false, false)
	assert.NoError(t, err)

	resultWithAll, err := GetPathSize("testdata", false, false, true)
	assert.NoError(t, err)

	assert.NotEqual(t, resultWithoutAll, resultWithAll)
}

func TestGetPathSizeHiddenFile(t *testing.T) {
	result, err := GetPathSize("testdata/.hidden_file.txt", false, false, false)
	assert.NoError(t, err)
	assert.Contains(t, result, "B")
}

func TestGetPathSizeRecursive(t *testing.T) {
	resultWithoutRecursive, err := GetPathSize("testdata", false, false, true)
	assert.NoError(t, err)

	resultWithRecursive, err := GetPathSize("testdata", true, false, true)
	assert.NoError(t, err)

	assert.NotEqual(t, resultWithoutRecursive, resultWithRecursive)
}
