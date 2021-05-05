package files

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileOrDirectoryExists(t *testing.T) {
	var exists = FileOrDirectoryExists("../../utils")
	assert.True(t, exists)
}

func TestFileOrDirectoryExistsErr(t *testing.T) {
	var exists = FileOrDirectoryExists("../../abcdefg")
	assert.False(t, exists)
}

func TestListFilesOfDirectory(t *testing.T) {
	directoryContent, err := ListFilesOfDirectory("./")
	assert.Nil(t, err)
	assert.NotNil(t, directoryContent)
	assert.True(t, len(*directoryContent) == 2)
}

func TestListFilesOfDirectoryWithFilter(t *testing.T) {
	var fileUtils = "file_utils.go"
	var fileFilter = FileFilter{
		Prefix: &fileUtils,
	}
	directoryContent, err := ListFilesOfDirectoryWithFilter("./", &fileFilter)

	assert.Nil(t, err)
	assert.NotNil(t, directoryContent)
	assert.True(t, len(*directoryContent) == 1)

	var directoryArray = *directoryContent
	assert.Equal(t, directoryArray[0], fileUtils)
}

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

func TestLoadTemplateFiles(t *testing.T) {
	var expectedFiles = []string{
		"../../../templates/jcasc_config.yaml",
		"../../../templates/jenkins_helm_values.yaml",
		"../../../templates/nginx_ingress_helm_values.yaml",
		"../../../templates/pvc_claim.yaml",
	}

	templates, err := LoadTemplateFilesOfDirectory("../../../templates/")
	assert.Nil(t, err)
	assert.NotNil(t, templates)

	for i := 0; i < len(expectedFiles); i++ {
		assert.Equal(t, expectedFiles[i], templates[i])
	}
}
