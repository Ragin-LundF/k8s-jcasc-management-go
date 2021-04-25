package constants

// CLI Workflow
const LogWizardStartProjectWizardDialogs = "Starting Project Wizard: Dialogs..."
const LogWizardStartProjectWizardDialogsDone = "Starting Project Wizard: Dialogs...done"
const LogWizardStartProcessingTemplates = "Starting Project Wizard: Template processing..."
const LogWizardStartProcessingTemplatesFailed = "Starting Project Wizard: Template processing...failed"
const LogWizardStartProcessingTemplatesDone = "Starting Project Wizard: Template processing...done"

// Secret file handling
const LogAskForSecretsFile = "  -> Ask for secrets file to apply..."
const LogAskForSecretsFileFailed = "Ask for secrets file to apply...failed"
const LogAskForSecretsFileDone = "Ask for secrets file to apply...done"

// Password of secrets file
const LogAskForPasswordOfSecretsFile = "  -> Ask for the password for secret file..."                          // NOSONAR
const LogAskForPasswordOfSecretsFileFailed = "  -> Ask for the password for secret file...failed"              // NOSONAR
const LogAskForPasswordOfSecretsFileDone = "  -> Ask for the password for secret file...done"                  // NOSONAR
const LogAskForConfirmationPasswordOfSecretsFile = "  -> Ask for the confirmation password for secret file..." // NOSONAR
const LogErrPasswordDidNotMatch = "  -> Passwords did not match!"                                              // NOSONAR
const LogInfoPasswordDidMatchStartEncrypting = "  -> Passwords did match! Starting encryption...."             // NOSONAR

// Namespace to apply secrets
const LogAskForNamespace = "-> Ask for namespace..."
const LogAskForNamespaceDone = "-> Ask for namespace...done"
const LogAskForNamespaceForSecretApply = "  -> Ask for namespace to apply secrets..."
const LogAskForNamespaceForSecretApplyFailed = "  -> Ask for namespace to apply secrets......failed"
const LogAskForNamespaceForSecretApplyDone = "  -> Ask for namespace to apply secrets......done"

const LogApplySecretsToNamespace = "  -> Apply secrets to namespace..."
const LogApplySecretsToNamespaceFailed = "  -> Apply secrets to namespace......failed"
const LogApplySecretsToNamespaceDone = "  -> Apply secrets to namespace......done"

const LogUnableToGetNameOfNewNamespace = "  -> Unable to get name of new namespace!"
const LogUnableToGetNameOfAdditionalNamespace = "  -> Unable to get name of additional namespace!"
const LogUnableToGetStoreConfigOnly = "  -> Unable to get store config only prompt!"

// IP address
const LogAskForIPAddress = "-> Ask for IP address..."
const LogAskForIPAddressDone = "-> Ask for IP address...done"
const LogErrUnableToGetIPAddress = "  -> Unable to get the IP address."

// domain name
const LogAskForJenkinsUrl = "-> Ask for Jenkins domain..."
const LogAskForJenkinsUrlDone = "-> Ask for Jenkins domain...done"

// Cloud templates
const LogAskForCloudTemplates = "-> Ask for cloud templates..."
const LogAskForCloudTemplatesDone = "-> Ask for cloud templates...done"

// Jenkins system message
const LogAskForJenkinsSystemMessage = "-> Ask for Jenkins system message..."
const LogAskForJenkinsSystemMessageDone = "-> Ask for Jenkins system message...done"
const LogUnableToGetJenkinsSystemMessage = "  -> Unable to get the Jenkins system message."

// Jobs Configuration Repository
const LogAskForJobsConfigurationRepository = "-> Ask for jobs configuration repository..."
const LogAskForJobsConfigurationRepositoryDone = "-> Ask for jobs configuration repository...done"
const LogUnableToGetJobsConfigurationRepository = "  -> Unable to get the jobs configuration repository."

// PVC
const LogAskForPvc = "-> Ask for persistent volume claim..."
const LogAskForPvcDone = "-> Ask for persistent volume claim...done"
const LogUnableToGetPvc = "  -> Unable to get persistent volume claim."
