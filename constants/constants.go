package constants

// common config
const DirConfig = "config"
const FilenameConfiguration = "k8s_jcasc_mgmt.cnf"
const FilenameConfigurationCustom = "k8s_jcasc_custom.cnf"
const FilenamePvcClaim = "pvc_claim.yaml"
const SecretsFileEncodedEnding = ".gpg"

// commands
const CommandMenu = "menu"
const CommandInstall = "install"
const CommandUninstall = "uninstall"
const CommandUpgrade = "upgrade"
const CommandEncryptSecrets = "encryptSecrets"
const CommandDecryptSecrets = "decryptSecrets"
const CommandApplySecrets = "applySecrets"
const CommandApplySecretsToAll = "applySecretsToAll"
const CommandCreateProject = "createProject"
const CommandCreateDeploymentOnlyProject = "createDeploymentOnlyProject"
const CommandCreateJenkinsUserPassword = "createJenkinsUserPassword"
const CommandQuit = "quit"

// error
const ErrorPromptFailed = "prompt failed"

// colors
const ColorNormal = "\033[0m"
const ColorInfo = "\033[1;34m"
const ColorNotice = "\033[1;36m"
const ColorWarning = "\033[1;33m"
const ColorError = "\033[1;31m"
const ColorDebug = "\033[0;36m"
