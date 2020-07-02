package config

import (
	"bufio"
	"fmt"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"k8s-management-go/app/utils/loggingstate"
	"os"
	"strings"
)

// ReadIPConfig reads the IP configuration file
func ReadIPConfig() {
	configuration := models.GetConfiguration()

	// if IP config file does not exist, create it
	if !files.FileOrDirectoryExists(models.GetIPConfigurationFile()) {
		os.Create(models.GetIPConfigurationFile())
	}

	// read configuration file. Replace unneeded double quotes if needed.
	data, err := os.Open(models.GetIPConfigurationFile())
	defer data.Close()

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
			if line != "" && !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, configuration.IPConfig.IPConfigFileDummyPrefix) {
				namespace, ip := parseIPConfigurationLine(line)
				models.AddIPAndNamespaceToConfiguration(namespace, ip)
			}
		}
	}
}

// parse line of configuration and split it into key/value
func parseIPConfigurationLine(line string) (namespace string, ip string) {
	// split line on "="
	var lineArray []string
	if strings.Contains(line, "=") {
		lineArray = strings.Split(line, "=")
	} else {
		lineArray = strings.Split(line, " ")
	}

	// assign to variables
	if cap(lineArray) == 2 {
		namespace = lineArray[0]
		ip = lineArray[1]
		// if value contains double quotes, replace them with empty string
		if strings.Contains(namespace, "\"") {
			namespace = strings.Replace(namespace, "\"", "", -1)
		}
		if strings.Contains(namespace, "'") {
			namespace = strings.Replace(namespace, "'", "", -1)
		}
		return strings.TrimSpace(namespace), strings.TrimSpace(ip)
	}
	return "", ""
}

// AddToIPConfigFile adds an IP to the IP config file
func AddToIPConfigFile(namespace string, ip string) (success bool, err error) {
	log := logger.Log()
	ipconfigFile, err := os.OpenFile(models.GetIPConfigurationFile(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to open IP config file [%s]", models.GetIPConfigurationFile()), err.Error())
		log.Errorf("[AddToIPConfigFile] Unable to open IP config file [%s]. \n%s", models.GetIPConfigurationFile(), err.Error())
		return false, err
	}
	defer ipconfigFile.Close()

	if _, err := ipconfigFile.WriteString(namespace + " " + ip + "\n"); err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to add new IP and namespace to file [%s]", models.GetIPConfigurationFile()), err.Error())
		log.Errorf("[AddToIPConfigFile] Unable to add new IP and namespace to file [%s]. \n%s", models.GetIPConfigurationFile(), err.Error())
		return false, err
	}
	return true, err
}
