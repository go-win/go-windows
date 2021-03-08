// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package restartmanager implements the Windows.Win32.RestartManager namespace.
package restartmanager

type RM_APP_TYPE int32

const (
	RmUnknownApp = 0
	RmMainWindow = 1
	RmOtherWindow = 2
	RmService = 3
	RmExplorer = 4
	RmConsole = 5
	RmCritical = 1000
)

type RM_SHUTDOWN_TYPE int32

const (
	RmForceShutdown = 1
	RmShutdownOnlyRegistered = 16
)

type RM_APP_STATUS int32

const (
	RmStatusUnknown = 0
	RmStatusRunning = 1
	RmStatusStopped = 2
	RmStatusStoppedOther = 4
	RmStatusRestarted = 8
	RmStatusErrorOnStop = 16
	RmStatusErrorOnRestart = 32
	RmStatusShutdownMasked = 64
	RmStatusRestartMasked = 128
)

type RM_REBOOT_REASON int32

const (
	RmRebootReasonNone = 0
	RmRebootReasonPermissionDenied = 1
	RmRebootReasonSessionMismatch = 2
	RmRebootReasonCriticalProcess = 4
	RmRebootReasonCriticalService = 8
	RmRebootReasonDetectedSelf = 16
)

type RM_FILTER_TRIGGER int32

const (
	RmFilterTriggerInvalid = 0
	RmFilterTriggerFile = 1
	RmFilterTriggerProcess = 2
	RmFilterTriggerService = 3
)

type RM_FILTER_ACTION int32

const (
	RmInvalidFilterAction = 0
	RmNoRestart = 1
	RmNoShutdown = 2
)
