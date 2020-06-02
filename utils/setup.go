package utils

import (
	"flag"
	"fmt"
	"k8s-management-go/server"
	"k8s-management-go/utils/config"
	"os"
	"strings"
)

// initial setup
// reads the configuration
func Setup() {
	// configure flags
	basePathFlag := flag.String("basepath", "", "base path to k8s-jcasc-management")
	serverStartFlag := flag.Bool("server", false, "start k8s-jcasc-management-go as a server")
	helpFlag := flag.Bool("help", false, "show help")
	flag.Parse()

	// define main path
	basePath := ""
	serverStart := false
	if os.Getenv("K8S_MGMT_BASE_PATH") != "" {
		// base path from environment variables
		basePath = os.Getenv("K8S_MGMT_BASE_PATH")
	} else {
		basePath = *basePathFlag
		serverStart = *serverStartFlag
	}

	if bool(*helpFlag) {
		showHelp()
		os.Exit(0)
	}

	// read configuration if base path was set, else go into panic mode
	if basePath != "" {
		// prepare basePath
		basePath = strings.Replace(basePath, "\"", "", -1)
		basePath = strings.TrimSuffix(basePath, "/")

		// read configuration
		config.ReadConfiguration(basePath)
		config.ReadIpConfig(basePath)

		if serverStart {
			server.StartServer()
		}
	} else {
		showHelp()
		fmt.Errorf("No configuration path found...")
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Println("k8s-jcasc-mgmt -basepath=<path> [-server] [-help]")
	fmt.Println()
	fmt.Println("  -basepath=<path>")
	fmt.Println("      * Required *")
	fmt.Println("      Add a basepath to the k8s-jcasc-management (bash version/configuration/templates) directory")
	fmt.Println("  -server")
	fmt.Println("      * Optional *")
	fmt.Println("      Start k8s-jcasc-mgmt as a server with a REST API (experimental; limited functionality)")
	fmt.Println("  -help")
	fmt.Println("       * Optional *")
	fmt.Println("       Show this help")
	fmt.Println()
}
