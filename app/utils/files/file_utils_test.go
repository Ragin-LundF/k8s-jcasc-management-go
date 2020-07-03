package files

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterFilenamePrefix(t *testing.T) {
	var filename = "my_file.txt"
	var prefix = "my_"
	var filter = FileFilter{
		Prefix: &prefix,
	}
	isValid := filterFilename(filename, &filter)

	assert.True(t, isValid)
}

func TestFilterFilenameInvalidPrefix(t *testing.T) {
	var filename = "my_file.txt"
	var prefix = "d_"
	var filter = FileFilter{
		Prefix: &prefix,
	}
	isValid := filterFilename(filename, &filter)

	assert.False(t, isValid)
}

func TestFilterFilenameSuffix(t *testing.T) {
	var filename = "my_file.txt"
	var suffix = ".txt"
	var filter = FileFilter{
		Suffix: &suffix,
	}
	isValid := filterFilename(filename, &filter)

	assert.True(t, isValid)
}

func TestFilterFilenameInvalidSuffix(t *testing.T) {
	var filename = "my_file.txt"
	var suffix = ".log"
	var filter = FileFilter{
		Suffix: &suffix,
	}
	isValid := filterFilename(filename, &filter)

	assert.False(t, isValid)
}

func TestFilterFilenamePrefixSuffix(t *testing.T) {
	var filename = "my_file.txt"
	var prefix = "my_"
	var suffix = ".txt"
	var filter = FileFilter{
		Suffix: &suffix,
		Prefix: &prefix,
	}
	isValid := filterFilename(filename, &filter)

	assert.True(t, isValid)
}

func TestFilterFilenameInvalidPrefixSuffix(t *testing.T) {
	var filename = "my_file.txt"
	var prefix = "d_"
	var suffix = ".txt"
	var filter = FileFilter{
		Suffix: &suffix,
		Prefix: &prefix,
	}
	isValid := filterFilename(filename, &filter)

	assert.False(t, isValid)
}

func TestFilterFilenamePrefixInvalidSuffix(t *testing.T) {
	var filename = "my_file.txt"
	var prefix = "my_"
	var suffix = ".log"
	var filter = FileFilter{
		Suffix: &suffix,
		Prefix: &prefix,
	}
	isValid := filterFilename(filename, &filter)

	assert.False(t, isValid)
}

func TestAppendPathTrailingSlashes(t *testing.T) {
	var basepath = "/mypath"
	var secondpath = "/yourpath"
	var expectedPath = "/mypath/yourpath"

	path := AppendPath(basepath, secondpath)

	assert.Equal(t, expectedPath, path)
}

func TestAppendPathTrailingAndLeadingSlashes(t *testing.T) {
	var basepath = "/mypath/"
	var secondpath = "/yourpath"
	var expectedPath = "/mypath/yourpath"

	path := AppendPath(basepath, secondpath)

	assert.Equal(t, expectedPath, path)
}

func TestAppendPathSecondWithoutSlashes(t *testing.T) {
	var basepath = "/mypath/"
	var secondpath = "yourpath"
	var expectedPath = "/mypath/yourpath"

	path := AppendPath(basepath, secondpath)

	assert.Equal(t, expectedPath, path)
}

func TestAppendPathWithoutSlashes(t *testing.T) {
	var basepath = "mypath"
	var secondpath = "yourpath"
	var expectedPath = "mypath/yourpath"

	path := AppendPath(basepath, secondpath)

	assert.Equal(t, expectedPath, path)
}
