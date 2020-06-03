package cli

// workflow for Jenkins installation
func InstallJenkins() (info string, err error) {
	// ask for namespace
	namespace, err := DialogAskForNamespace()
	if err != nil {
		return info, err
	}
	// check if namespace is available or create a new one if not
	infoLog, err := CheckAndCreateNamespace(namespace)
	info = info + infoLog
	if err != nil {
		return info, err
	}
	// install secrets
	infoLog, err = ApplySecretsToNamespace(namespace)
	info = info + infoLog
	if err != nil {
		return info, err
	}
	// check if PVC was specified and install it if needed
	infoLog, err = InstallPersistenceVolumeClaim(namespace)
	info = info + infoLog
	if err != nil {
		return info, err
	}

	return info, err
}
