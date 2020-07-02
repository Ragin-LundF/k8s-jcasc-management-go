package files

import (
	"testing"
)

func TestFilterFilenamePrefix(t *testing.T) {
	var filename = "my_file.txt"
	var prefix = "my_"
	var filter = FileFilter{
		Prefix: &prefix,
	}
	isValid := filterFilename(filename, &filter)

	if isValid {
		t.Log("Success. Found file with correct prefix.")
	} else {
		t.Errorf("Failed. File [%s] with prefix filter [%s] should return true", filename, prefix)
	}
}

func TestFilterFilenameInvalidPrefix(t *testing.T) {
	var filename = "my_file.txt"
	var prefix = "d_"
	var filter = FileFilter{
		Prefix: &prefix,
	}
	isValid := filterFilename(filename, &filter)

	if !isValid {
		t.Log("Success. No file found.")
	} else {
		t.Errorf("Failed. File [%s] with prefix filter [%s] should return false", filename, prefix)
	}
}

func TestFilterFilenameSuffix(t *testing.T) {
	var filename = "my_file.txt"
	var suffix = ".txt"
	var filter = FileFilter{
		Suffix: &suffix,
	}
	isValid := filterFilename(filename, &filter)

	if isValid {
		t.Log("Success. Found file with correct suffix.")
	} else {
		t.Errorf("Failed. File [%s] with suffix filter [%s] should return true", filename, suffix)
	}
}

func TestFilterFilenameInvalidSuffix(t *testing.T) {
	var filename = "my_file.txt"
	var suffix = ".log"
	var filter = FileFilter{
		Suffix: &suffix,
	}
	isValid := filterFilename(filename, &filter)

	if !isValid {
		t.Log("Success. No file found with invalid suffix.")
	} else {
		t.Errorf("Failed. File [%s] with suffix filter [%s] should return false", filename, suffix)
	}
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

	if isValid {
		t.Log("Success. File found with prefix and suffix.")
	} else {
		t.Errorf("Failed. File [%s] with prefix [%s] and suffix filter [%s] should return false", filename, prefix, suffix)
	}
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

	if !isValid {
		t.Log("Success. No file found with invalid prefix and valid suffix.")
	} else {
		t.Errorf("Failed. File [%s] with prefix [%s] and suffix filter [%s] should return false", filename, *filter.Prefix, *filter.Suffix)
	}
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

	if !isValid {
		t.Log("Success. No file found with invalid prefix and valid suffix.")
	} else {
		t.Errorf("Failed. File [%s] with prefix [%s] and suffix filter [%s] should return false", filename, *filter.Prefix, *filter.Suffix)
	}
}

func TestAppendPathTrailingSlashes(t *testing.T) {
	var basepath = "/mypath"
	var secondpath = "/yourpath"
	var expectedPath = "/mypath/yourpath"

	path := AppendPath(basepath, secondpath)
	if path == expectedPath {
		t.Logf("Success. Path is: [%s]", path)
	} else {
		t.Errorf("Failed. Wrong path is [%s]", path)
	}
}

func TestAppendPathTrailingAndLeadingSlashes(t *testing.T) {
	var basepath = "/mypath/"
	var secondpath = "/yourpath"
	var expectedPath = "/mypath/yourpath"

	path := AppendPath(basepath, secondpath)
	if path == expectedPath {
		t.Logf("Success. Path is: [%s]", path)
	} else {
		t.Errorf("Failed. Wrong path is [%s]", path)
	}
}

func TestAppendPathSecondWithoutSlashes(t *testing.T) {
	var basepath = "/mypath/"
	var secondpath = "yourpath"
	var expectedPath = "/mypath/yourpath"

	path := AppendPath(basepath, secondpath)
	if path == expectedPath {
		t.Logf("Success. Path is: [%s]", path)
	} else {
		t.Errorf("Failed. Wrong path is [%s]", path)
	}
}

func TestAppendPathWithoutSlashes(t *testing.T) {
	var basepath = "mypath"
	var secondpath = "yourpath"
	var expectedPath = "mypath/yourpath"

	path := AppendPath(basepath, secondpath)
	if path == expectedPath {
		t.Logf("Success. Path is: [%s]", path)
	} else {
		t.Errorf("Failed. Wrong path is [%s]", path)
	}
}
