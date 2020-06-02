package files

import (
	"os"
	"strings"
)

// check if file exists
func FileExists(fileNameWithPath string) bool {
	info, err := os.Stat(fileNameWithPath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// helper for adding new pathes
func AddFilePath(originalPath string, pathExtension string) (extendedPath string) {
	// path extension starts with "./" remove it
	if strings.HasPrefix(pathExtension, "./") {
		pathExtension = strings.TrimPrefix(pathExtension, "./")
	}

	// handle suffix and prefix to create proper path
	if strings.HasSuffix(originalPath, "/") {
		if strings.HasPrefix(pathExtension, "/") {
			// originalPath ends with "/" and path extension starts with "/"
			extendedPath = strings.TrimSuffix(originalPath, "/") + pathExtension
		} else {
			// original path ends with "/" and path extension does not start with "/"
			extendedPath = originalPath + pathExtension
		}
	} else if strings.HasPrefix(pathExtension, "/") {
		// original path does not end with "/" but pathExtension has "/" prefix
		extendedPath = originalPath + pathExtension
	} else {
		// original path does not end with "/" and path extension does not start with "/"
		extendedPath = originalPath + "/" + pathExtension
	}
	return extendedPath
}
