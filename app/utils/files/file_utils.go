package files

import (
	"io/ioutil"
	"k8s-management-go/app/utils/logger"
	"os"
	"strings"
)

type FileFilter struct {
	Prefix *string
	Suffix *string
}

// check if file exists
func FileOrDirectoryExists(fileNameWithPath string) bool {
	log := logger.Log()
	info, err := os.Stat(fileNameWithPath)
	if os.IsNotExist(err) {
		log.Error("Unable to find file [" + fileNameWithPath + "]")
		return false
	}
	return !info.IsDir()
}

// list files of a directory if it exists
func ListFilesOfDirectory(directory string) (files *[]string, err error) {
	files, err = ListFilesOfDirectoryWithFilter(directory, nil)
	return files, err
}

// list files of a directory if it exists with a filter
func ListFilesOfDirectoryWithFilter(directory string, filter *FileFilter) (files *[]string, err error) {
	log := logger.Log()
	// check if the directory exists before reading from directory
	directoryExists := FileOrDirectoryExists(directory)
	if directoryExists {
		fileList, err := ioutil.ReadDir(directory)
		if err != nil {
			log.Error(err)
			return files, err
		}

		var filesReturnValue []string
		for _, file := range fileList {
			if filterFilename(file.Name(), filter) {
				filesReturnValue = append(filesReturnValue, file.Name())
			}
		}
		return &filesReturnValue, err
	}
	return nil, err
}

// filter by filename and filter
func filterFilename(filename string, filter *FileFilter) bool {
	fileIsOk := true
	// no filter -> everything is ok
	if filter != nil {
		// filter prefix
		if filter.Prefix != nil {
			// filter prefix if exists
			if !strings.HasPrefix(filename, *filter.Prefix) {
				fileIsOk = false
			}
			// filter suffix if exists and file is still ok
			if fileIsOk && !strings.HasSuffix(filename, *filter.Suffix) {
				fileIsOk = false
			}
		}
	}
	return fileIsOk
}

// helper for adding new pathes
func AppendPath(originalPath string, pathExtension string) (extendedPath string) {
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
