package config

import (
	"bufio"
	"k8s-management-go/app/constants"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os"
	"strings"
)

// Read configuration from k8s-management
func ReadConfiguration(basePath string, dryRunDebug bool, cliOnly bool) {
	// read plain configuration
	readConfigurationFromFile(files.AppendPath(files.AppendPath(basePath, constants.DirConfig), constants.FilenameConfiguration))
	// check if there is an custom configuration
	if files.FileOrDirectoryExists(files.AppendPath(files.AppendPath(basePath, constants.DirConfig), constants.FilenameConfigurationCustom)) {
		readConfigurationFromFile(files.AppendPath(files.AppendPath(basePath, constants.DirConfig), constants.FilenameConfigurationCustom))
		if models.GetConfiguration().BasePath != "" {
			basePath = models.GetConfiguration().BasePath
		}
	}
	// check if there is an alternative configuration path and try to read config from there
	models.AssignToConfiguration("K8S_MGMT_BASE_PATH", basePath)
	models.AssignDryRun(dryRunDebug)
	models.AssignCliOnlyMode(cliOnly)

	configuration := models.GetConfiguration()
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
		// close file after this method was finished
		defer data.Close()

		// everything seems to be ok. Read data with line scanner
		scanner := bufio.NewScanner(data)
		scanner.Split(bufio.ScanLines)

		// iterate over every line
		for scanner.Scan() {
			// trim the line to avoid problems
			line := strings.TrimSpace(scanner.Text())
			processLine(line)
		}
	}
}

// function to check if line is empty / a comment or a valid configuration.
// if line looks valid, try to append it to internal configuration structure
func processLine(line string) {
	// if line is not a comment (marker: "#") parse the configuration and assign it to the config
	if line != "" && !strings.HasPrefix(line, "#") {
		key, value := parseConfigurationLine(line)
		models.AssignToConfiguration(key, value)
	}
}

// parse line of configuration and split it into key/value
func parseConfigurationLine(line string) (key string, value string) {
	// split line on "="
	lineArray := strings.Split(line, "=")
	if cap(lineArray) == 2 {
		// assign to variables
		key = lineArray[0]
		value = lineArray[1]
		// if value contains double quotes, replace them with empty string
		if strings.Contains(value, "\"") {
			value = strings.Replace(value, "\"", "", -1)
		}
		if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
			value = strings.Replace(value, "'", "", -1)
		}
		return strings.TrimSpace(key), strings.TrimSpace(value)
	}
	return "", ""
}
