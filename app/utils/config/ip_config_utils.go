package config

import (
	"bufio"
	"fmt"
	"k8s-management-go/app/cli/loggingstate"
	"k8s-management-go/app/models"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os"
	"strings"
)

func ReadIpConfig() {
	configuration := models.GetConfiguration()

	// if IP config file does not exist, create it
	if !files.FileOrDirectoryExists(models.GetIpConfigurationFile()) {
		os.Create(models.GetIpConfigurationFile())
	}

	// read configuration file. Replace unneeded double quotes if needed.
	data, err := os.Open(models.GetIpConfigurationFile())
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
			if line != "" && !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, configuration.IpConfig.IpConfigFileDummyPrefix) {
				namespace, ip := parseIpConfigurationLine(line)
				models.AddIpAndNamespaceToConfiguration(namespace, ip)
			}
		}
	}
}

// parse line of configuration and split it into key/value
func parseIpConfigurationLine(line string) (namespace string, ip string) {
	// split line on "="
	lineArray := strings.Split(line, " ")
	// assign to variables
	namespace = lineArray[0]
	ip = lineArray[1]
	// if value contains double quotes, replace them with empty string
	if strings.Contains(namespace, "\"") {
		namespace = strings.Replace(namespace, "\"", "", -1)
	}
	return namespace, ip
}

// Add IP to IP config file
func AddToIpConfigFile(namespace string, ip string) (success bool, err error) {
	log := logger.Log()
	ipconfigFile, err := os.OpenFile(models.GetIpConfigurationFile(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to open IP config file [%s]", models.GetIpConfigurationFile()), err.Error())
		log.Errorf("[AddToIpConfigFile] Unable to open IP config file [%s]. \n%s", models.GetIpConfigurationFile(), err.Error())
		return false, err
	}
	defer ipconfigFile.Close()

	if _, err := ipconfigFile.WriteString(namespace + " " + ip + "\n"); err != nil {
		loggingstate.AddErrorEntryAndDetails(fmt.Sprintf("  -> Unable to add new IP and namespace to file [%s]", models.GetIpConfigurationFile()), err.Error())
		log.Errorf("[AddToIpConfigFile] Unable to add new IP and namespace to file [%s]. \n%s", models.GetIpConfigurationFile(), err.Error())
		return false, err
	}
	return true, err
}
