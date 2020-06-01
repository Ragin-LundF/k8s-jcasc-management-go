package utils

import (
	"os"
	"strings"
)

// initial setup
// reads the configuration
func Setup() {
	// define main path
	basePath := ""
	if os.Getenv("K8S_MGMT_BASE_PATH") != "" {
		// base path from environment variables
		basePath = os.Getenv("K8S_MGMT_BASE_PATH")
	} else if len(os.Args) == 2 {
		// base path as argument found
		basePath = os.Args[1]
	}

	// read configuration if base path was set, else go into panic mode
	if basePath != "" {
		// prepare basePath
		basePath = strings.Replace(basePath, "\"", "", -1)
		basePath = strings.TrimSuffix(basePath, "/")

		// read configuration
		ReadConfiguration(basePath)
		ReadIpConfig(basePath)
	} else {
		panic("No configuration path found...")
	}
}
