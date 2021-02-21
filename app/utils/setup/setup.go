package setup

import (
	"flag"
	"fmt"
	"k8s-management-go/app/actions/migration"
	"k8s-management-go/app/models"
	"k8s-management-go/app/server"
	"k8s-management-go/app/utils/cmdexecutor"
	"k8s-management-go/app/utils/config"
	"k8s-management-go/app/utils/files"
	"k8s-management-go/app/utils/logger"
	"os"
	"strings"
)

// Setup is the initial setup and reads the configuration
func Setup() {
	// configure flags
	logFileFlag := flag.String("logfile", "", "Logging output file. If empty it logs to console.")
	logEncoding := flag.String("logencoding", "", "Logging output encoding. If empty it logs as json.")
	basePathFlag := flag.String("basepath", "", "base path to k8s-jcasc-management")
	serverStartFlag := flag.Bool("server", false, "start k8s-jcasc-management-go as a server")
	dryRunFlag := flag.Bool("dry-run", false, "execute helm charts with --dry-run --debug flags")
	cliOnly := flag.Bool("cli", false, "Start in CLI mode")
	migrateTemplates := flag.Bool("migrate-templates-v2", false, "Migrate templates from v2 to v3")
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

	// show help?
	if *helpFlag {
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
			log.Errorf(err.Error())
			showHelp()
			log.Infof("Cannot find the base path! Please add it with -basepath flag.")
			fmt.Print("Cannot find the base path! Please add it with -basepath flag.")
			os.Exit(1)
		}
		basePath = currentPath
	}

	// configure (read configuration and do additional configuration)
	configure(basePath, dryRunDebug, *cliOnly)

	if *migrateTemplates {
		println("Starting migration...")
		var migrationStatus = migration.MigrateTemplatesToV3()
		println(migrationStatus)
		os.Exit(0)
	}

	// start experimental server
	if serverStart {
		server.StartServer()
	}

	// Set the OS executor for exec.Command().CombinedOutput() execution
	cmdexecutor.Executor = cmdexecutor.OsCommandExec{}
}

func configure(basePath string, dryRunDebug bool, cliOnly bool) {
	// read configuration
	config.ReadConfiguration(basePath, dryRunDebug, cliOnly)
	config.ReadIPConfig()

	// overwrite logging
	if logger.LogEncoding == "" && models.GetConfiguration().K8sManagement.Logging.LogEncoding != "" {
		logger.LogEncoding = models.GetConfiguration().K8sManagement.Logging.LogEncoding
	}
	if logger.LogFilePath == "" && models.GetConfiguration().K8sManagement.Logging.LogFile != "" {
		logger.LogFilePath = models.GetConfiguration().K8sManagement.Logging.LogFile
	}
	if logger.LogFilePath != "" && models.GetConfiguration().K8sManagement.Logging.LogOverwriteOnStart {
		if files.FileOrDirectoryExists(logger.LogFilePath) {
			_ = os.Rename(logger.LogFilePath, logger.LogFilePath+".1")
		}
	}
}

func showHelp() {
	fmt.Println("k8s-jcasc-mgmt -basepath=<path> [-server] [-help]")
	fmt.Println()
	fmt.Println("  -logfile=<path/file.log>")
	fmt.Println("      * Optional *")
	fmt.Println("      File for logging output")
	fmt.Println("  -logencoding=<console | json>")
	fmt.Println("      * Optional *")
	fmt.Println("      Defines logging output format (console or json)")
	fmt.Println("  -basepath=<path>")
	fmt.Println("      * Optional *")
	fmt.Println("      Add a base path to the k8s-jcasc-management (directory which contains version/configuration/templates) directory")
	fmt.Println("  -server")
	fmt.Println("      * Optional *")
	fmt.Println("      Start k8s-jcasc-mgmt as a server with a REST API (experimental; limited functionality)")
	fmt.Println("  -migrate-templates-v2")
	fmt.Println("      * Optional *")
	fmt.Println("      Migrate the templates from v2 to v3")
	fmt.Println("  -help")
	fmt.Println("       * Optional *")
	fmt.Println("       Show this help")
	fmt.Println()
}
