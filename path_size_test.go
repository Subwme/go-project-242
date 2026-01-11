package goproject242

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSize(t *testing.T) {
	resultDataCSV, err := GetPathSize("testdata/data.csv", false, false, false)
	assert.NoError(t, err)
	assert.Equal(t, "37B\ttestdata/data.csv", resultDataCSV)
}

func TestGetPathSizeDir(t *testing.T) {
	resultTestDir, err := GetPathSize("testdata", false, false, false)
	assert.NoError(t, err)
	assert.Equal(t, "198B\ttestdata", resultTestDir)
}