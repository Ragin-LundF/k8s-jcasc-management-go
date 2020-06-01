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
	// read plain configuration
	readConfigurationFromFile(basePath + "/" + constants.DIR_CONFIG + "/" + constants.FILENAME_CONFIGURATION)
	// check if there is an custom configuration
	if FileExists(basePath + "/" + constants.DIR_CONFIG + "/" + constants.FILENAME_CONFIGURATION_CUSTOM) {
		readConfigurationFromFile(basePath + "/" + constants.DIR_CONFIG + "/" + constants.FILENAME_CONFIGURATION_CUSTOM)
	}
	// check if there is an alternative configuration path and try to read config from there
	configuration := *config.GetConfiguration()
	if configuration.AlternativeConfigFile != "" && FileExists(basePath+"/"+configuration.AlternativeConfigFile) {
		readConfigurationFromFile(basePath + "/" + configuration.AlternativeConfigFile)
	}
}

// Read configuration from k8s-management config file
func readConfigurationFromFile(configFile string) {
	// read configuration file. Replace unneeded double quotes if needed.
	data, err := os.Open(configFile)
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
				key, value := parseConfigurationLine(line)
				config.AssignToConfiguration(key, value)
			}
		}
	}
	// close file
	_ = data.Close()
}

// parse line of configuration and split it into key/value
func parseConfigurationLine(line string) (key string, value string) {
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
