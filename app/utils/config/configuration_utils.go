package config

import (
	"bufio"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os"
	"strconv"
	"strings"
)

// Read configuration from k8s-management
func ReadConfiguration(basePath string, dryRunDebug bool) {
	// read plain configuration
	readConfigurationFromFile(files.AppendPath(files.AppendPath(basePath, constants.DirConfig), constants.FilenameConfiguration))
	// check if there is an custom configuration
	if files.FileOrDirectoryExists(files.AppendPath(files.AppendPath(basePath, constants.DirConfig), constants.FilenameConfigurationCustom)) {
		readConfigurationFromFile(files.AppendPath(files.AppendPath(basePath, constants.DirConfig), constants.FilenameConfigurationCustom))
	}
	// check if there is an alternative configuration path and try to read config from there
	configuration := models.GetConfiguration()
	models.AssignToConfiguration("K8S_MGMT_BASE_PATH", basePath)
	models.AssignToConfiguration("K8S_MGMT_DRY_RUN_DEBUG", strconv.FormatBool(dryRunDebug))

	if configuration.AlternativeConfigFile != "" && files.FileOrDirectoryExists(files.AppendPath(basePath, configuration.AlternativeConfigFile)) {
		readConfigurationFromFile(files.AppendPath(basePath, configuration.AlternativeConfigFile))
	}
}

// Read configuration from k8s-management config file
func readConfigurationFromFile(configFile string) {
	log := logger.Log()
	// read configuration file. Replace unneeded double quotes if needed.
	data, err := os.Open(configFile)
	// check for error
	if err != nil {
		log.Errorf(err.Error())
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
				models.AssignToConfiguration(key, value)
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
