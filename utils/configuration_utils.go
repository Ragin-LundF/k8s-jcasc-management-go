package utils

import (
	"bufio"
	"k8s-management-go/constants"
	"k8s-management-go/models/config"
	"os"
	"strings"
)

// Read configuration from k8s-management
func ReadConfiguration(basePath string) {
	// read configuration file. Replace unneeded double quotes if needed.
	data, err := os.Open(basePath + "/" + constants.DIR_CONFIG + "/" + constants.FILENAME_CONFIGURATION)
	// check for error
	if err != nil {
		panic(err)
	} else {
		// everything seems to be ok. Read data with line scanner
		scanner := bufio.NewScanner(data)
		scanner.Split(bufio.ScanLines)

		// iterate over every line
		for scanner.Scan() {
			// trim the line to avoid problems
			line := strings.TrimSpace(scanner.Text())
			// if line is not a comment (marker: "#") parse the configuration and assign it to the config
			if line != "" && !strings.HasPrefix(line, "#") {
				key, value := ParseConfigurationLine(line)
				config.AssignToConfiguration(key, value)
			}
		}
	}
}

// parse line of configuration and split it into key/value
func ParseConfigurationLine(line string) (key string, value string) {
	// split line on "="
	lineArray := strings.Split(line, "=")
	// assign to variables
	key = lineArray[0]
	value = lineArray[1]
	// if value contains double quotes, replace them with empty string
	if strings.Contains(value, "\"") {
		value = strings.Replace(value, "\"", "", -1)
	}
	return key, value
}
