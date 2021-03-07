package migration

import (
	"bufio"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/utils/files"
	"os"
	"strings"
)

// IP is the basic IP structure
type IP struct {
	Namespace string
	IP        string
}

// IPConfiguration contains a list of IPs
type IPConfiguration struct {
	IPs []IP
}

var ipConfig IPConfiguration

// AddIPAndNamespaceToConfiguration adds IP and namespace to the IP configuration
func AddIPAndNamespaceToConfiguration(namespace string, ip string) {
	ipConfig.IPs = append(ipConfig.IPs, IP{namespace, ip})
}

// ResetIPAndNamespaces will reset the configured IPs
func ResetIPAndNamespaces() {
	ipConfig.IPs = nil
}

// ReadIPConfig reads the IP configuration file
func ReadIPConfig() {
	// if IP config file does not exist, create it
	if !files.FileOrDirectoryExists(configuration.GetConfiguration().GetIPConfigurationFile()) {
		_, _ = os.Create(configuration.GetConfiguration().GetIPConfigurationFile())
	}

	// read configuration file. Replace unneeded double quotes if needed.
	data, err := os.Open(configuration.GetConfiguration().GetIPConfigurationFile())
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
			if line != "" && !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, configuration.GetConfiguration().K8SManagement.IPConfig.DummyPrefix) {
				namespace, ip := parseIPConfigurationLine(line)
				AddIPAndNamespaceToConfiguration(namespace, ip)
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
	} else if strings.Contains(line, " ") {
		lineArray = strings.Split(line, " ")
	} else {
		if strings.TrimSpace(line) != "" {
			return optimizeNamespaces(line), ""
		}
	}

	// assign to variables
	if len(lineArray) == 2 {
		namespace = lineArray[0]
		ip = lineArray[1]
		return optimizeNamespaces(namespace), strings.TrimSpace(ip)
	}
	return "", ""
}

func optimizeNamespaces(namespace string) string {
	// if value contains double quotes, replace them with empty string
	if strings.Contains(namespace, "\"") {
		namespace = strings.Replace(namespace, "\"", "", -1)
	}
	if strings.Contains(namespace, "'") {
		namespace = strings.Replace(namespace, "'", "", -1)
	}

	return strings.TrimSpace(namespace)
}
