package setup

import (
	"flag"
	"fmt"
	"k8s-management-go/app/server"
	"k8s-management-go/app/utils/config"
	"k8s-management-go/app/utils/logger"
	"os"
	"strings"
)

// initial setup
// reads the configuration
func Setup() {
	// configure flags
	logFileFlag := flag.String("logfile", "", "Logging output file. If empty it logs to console.")
	logEncoding := flag.String("logencoding", "", "Logging output encoding. If empty it logs as json.")
	basePathFlag := flag.String("basepath", "", "base path to k8s-jcasc-management")
	serverStartFlag := flag.Bool("server", false, "start k8s-jcasc-management-go as a server")
	dryRunFlag := flag.Bool("dry-run", false, "execute helm charts with --dry-run --debug flags")
	helpFlag := flag.Bool("help", false, "show help")
	flag.Parse()

	// define main path
	logger.LogFilePath = *logFileFlag
	logger.LogEncoding = *logEncoding
	basePath := ""
	serverStart := *serverStartFlag
	dryRunDebug := *dryRunFlag
	if os.Getenv("K8S_MGMT_BASE_PATH") != "" {
		// base path from environment variables
		basePath = os.Getenv("K8S_MGMT_BASE_PATH")
	} else {
		basePath = *basePathFlag
	}

	if bool(*helpFlag) {
		showHelp()
		os.Exit(0)
	}

	// Logger
	log := logger.Log()

	// read configuration if base path was set, else go into panic mode
	if basePath != "" {
		// prepare basePath
		basePath = strings.Replace(basePath, "\"", "", -1)
		basePath = strings.TrimSuffix(basePath, "/")
	} else {
		currentPath, err := os.Getwd()
		if err != nil {
			log.Error(err)
			showHelp()
			log.Info("Can not find the base path! Please add it with -basepath flag.")
			os.Exit(1)
		}
		basePath = currentPath
	}

	// read configuration
	config.ReadConfiguration(basePath, dryRunDebug)
	config.ReadIpConfig()

	// start experimental server
	if serverStart {
		server.StartServer()
	}
}

func showHelp() {
	fmt.Println("k8s-jcasc-mgmt -basepath=<path> [-server] [-help]")
	fmt.Println()
	fmt.Println("  -logfile=<path/file.log>")
	fmt.Println("      * Optional *")
	fmt.Println("      File for logging output")
	fmt.Println("  -basepath=<path>")
	fmt.Println("      * Optional *")
	fmt.Println("      Add a base path to the k8s-jcasc-management (directory which contains version/configuration/templates) directory")
	fmt.Println("  -server")
	fmt.Println("      * Optional *")
	fmt.Println("      Start k8s-jcasc-mgmt as a server with a REST API (experimental; limited functionality)")
	fmt.Println("  -help")
	fmt.Println("       * Optional *")
	fmt.Println("       Show this help")
	fmt.Println()
}