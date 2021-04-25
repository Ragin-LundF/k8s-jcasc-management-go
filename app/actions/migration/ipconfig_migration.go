package migration

import (
	"bufio"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"k8s-management-go/app/configuration"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/loggingstate"
	"os"
	"strings"
)

// MigrateDeploymentIPConfigurationV3 starts deployment IP configuration migration
func MigrateDeploymentIPConfigurationV3() string {
	return migrateIpConfigFromCnfToYaml()
}

//NOSONAR
// readIPConfig reads the IP configuration file
func readIPConfig() *configuration.DeploymentYAMLConfig {
	var ipDeploymentCfgFile = configuration.GetConfiguration().GetIPConfigurationFile()
	if strings.HasSuffix(ipDeploymentCfgFile, ".yaml") {
		ipDeploymentCfgFile = strings.Replace(ipDeploymentCfgFile, ".yaml", ".cnf", -1)
	}
	// if IP config file does not exist, create it
	if files.FileOrDirectoryExists(ipDeploymentCfgFile) {
		// read configuration file. Replace unneeded double quotes if needed.
		data, err := os.Open(ipDeploymentCfgFile)

		// check for error
		if err != nil {
			panic(err)
		} else {
			// everything seems to be ok. Read data with line scanner
			scanner := bufio.NewScanner(data)
			scanner.Split(bufio.ScanLines)

			var deploymentYaml = configuration.DeploymentYAMLConfig{}

			// iterate over every line
			for scanner.Scan() {
				var deploymentIpConfig = configuration.DeploymentStruct{}
				// trim the line to avoid problems
				line := strings.TrimSpace(scanner.Text())
				// if line is not a comment (marker: "#") parse the configuration and assign it to the config
				if line != "" && !strings.HasPrefix(line, "#") {
					namespace, ip := parseIPConfigurationLine(line)
					if !strings.HasPrefix(line, configuration.GetConfiguration().K8SManagement.IPConfig.DummyPrefix) {
						deploymentIpConfig.Dummy = ""
					} else {
						deploymentIpConfig.Dummy = "true"
					}
					deploymentIpConfig.IPAddress = ip
					deploymentIpConfig.Namespace = namespace

					deploymentYaml.K8SManagement.IPConfig.Deployments = append(deploymentYaml.K8SManagement.IPConfig.Deployments, deploymentIpConfig)
				}
			}
			_ = data.Close()
			return &deploymentYaml
		}
	}
	return nil
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

func migrateIpConfigFromCnfToYaml() string {
	var deploymentYaml = readIPConfig()

	if deploymentYaml != nil {
		var deployYamlOutByte []byte
		deployYamlOutByte, err := yaml.Marshal(deploymentYaml)
		if err != nil {
			loggingstate.AddErrorEntryAndDetails("Unable to create IP deployment YAML file", err.Error())
			return "FAILED - Unable to create IP deployment YAML file"
		}
		var yamlConfig = string(deployYamlOutByte)
		loggingstate.AddInfoEntryAndDetails("New custom IP deployment configuration", yamlConfig)

		var configFile = configuration.GetConfiguration().GetIPConfigurationFile()
		if files.FileOrDirectoryExists(configFile) {
			return "FAILED - Config already exists"
		}
		err = ioutil.WriteFile(configFile, deployYamlOutByte, 0644)
		if err != nil {
			return "FAILED - Unable to write new IP config"
		}
	}
	return "SUCCESS"
}
