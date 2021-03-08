// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsupdateagent implements the Windows.Win32.WindowsUpdateAgent namespace.
package windowsupdateagent

type AutomaticUpdatesNotificationLevel int32

const (
	aunlNotConfigured = 0
	aunlDisabled = 1
	aunlNotifyBeforeDownload = 2
	aunlNotifyBeforeInstallation = 3
	aunlScheduledInstallation = 4
)

type AutomaticUpdatesScheduledInstallationDay int32

const (
	ausidEveryDay = 0
	ausidEverySunday = 1
	ausidEveryMonday = 2
	ausidEveryTuesday = 3
	ausidEveryWednesday = 4
	ausidEveryThursday = 5
	ausidEveryFriday = 6
	ausidEverySaturday = 7
)

type DownloadPhase int32

const (
	dphInitializing = 1
	dphDownloading = 2
	dphVerifying = 3
)

type DownloadPriority int32

const (
	dpLow = 1
	dpNormal = 2
	dpHigh = 3
	dpExtraHigh = 4
)

type AutoSelectionMode int32

const (
	asLetWindowsUpdateDecide = 0
	asAutoSelectIfDownloaded = 1
	asNeverAutoSelect = 2
	asAlwaysAutoSelect = 3
)

type AutoDownloadMode int32

const (
	adLetWindowsUpdateDecide = 0
	adNeverAutoDownload = 1
	adAlwaysAutoDownload = 2
)

type InstallationImpact int32

const (
	iiNormal = 0
	iiMinor = 1
	iiRequiresExclusiveHandling = 2
)

type InstallationRebootBehavior int32

const (
	irbNeverReboots = 0
	irbAlwaysRequiresReboot = 1
	irbCanRequestReboot = 2
)

type OperationResultCode int32

const (
	orcNotStarted = 0
	orcInProgress = 1
	orcSucceeded = 2
	orcSucceededWithErrors = 3
	orcFailed = 4
	orcAborted = 5
)

type ServerSelection int32

const (
	ssDefault = 0
	ssManagedServer = 1
	ssWindowsUpdate = 2
	ssOthers = 3
)

type UpdateType int32

const (
	utSoftware = 1
	utDriver = 2
)

type UpdateOperation int32

const (
	uoInstallation = 1
	uoUninstallation = 2
)

type DeploymentAction int32

const (
	daNone = 0
	daInstallation = 1
	daUninstallation = 2
	daDetection = 3
	daOptionalInstallation = 4
)

type UpdateExceptionContext int32

const (
	uecGeneral = 1
	uecWindowsDriver = 2
	uecWindowsInstaller = 3
	uecSearchIncomplete = 4
)

type AutomaticUpdatesUserType int32

const (
	auutCurrentUser = 1
	auutLocalAdministrator = 2
)

type AutomaticUpdatesPermissionType int32

const (
	auptSetNotificationLevel = 1
	auptDisableAutomaticUpdates = 2
	auptSetIncludeRecommendedUpdates = 3
	auptSetFeaturedUpdatesEnabled = 4
	auptSetNonAdministratorsElevated = 5
)

type UpdateServiceRegistrationState int32

const (
	usrsNotRegistered = 1
	usrsRegistrationPending = 2
	usrsRegistered = 3
)

type SearchScope int32

const (
	searchScopeDefault = 0
	searchScopeMachineOnly = 1
	searchScopeCurrentUserOnly = 2
	searchScopeMachineAndCurrentUser = 3
	searchScopeMachineAndAllUsers = 4
	searchScopeAllUsers = 5
)

type UpdateLockdownOption int32

const (
	uloForWebsiteAccess = 1
)

type AddServiceFlag int32

const (
	asfAllowPendingRegistration = 1
	asfAllowOnlineRegistration = 2
	asfRegisterServiceWithAU = 4
)

type UpdateServiceOption int32

const (
	usoNonVolatileService = 1
)
