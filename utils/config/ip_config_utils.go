package config

import (
	"bufio"
	"k8s-management-go/models/config"
	"os"
	"strings"
)

func ReadIpConfig(basePath string) {
	configuration := *config.GetConfiguration()

	// read configuration file. Replace unneeded double quotes if needed.
	data, err := os.Open(basePath + "/" + configuration.IpConfig.IpConfigFile)

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
				config.AddIpAndNamespaceToConfiguration(namespace, ip)
			}
		}
	}
	// close file
	_ = data.Close()
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
