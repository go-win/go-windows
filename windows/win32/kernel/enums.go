// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package kernel implements the Windows.Win32.Kernel namespace.
package kernel

type EXCEPTION_DISPOSITION int32

const (
	ExceptionContinueExecution = 0
	ExceptionContinueSearch = 1
	ExceptionNestedException = 2
	ExceptionCollidedUnwind = 3
)

type EVENT_TYPE int32

const (
	NotificationEvent = 0
	SynchronizationEvent = 1
)

type TIMER_TYPE int32

const (
	NotificationTimer = 0
	SynchronizationTimer = 1
)

type WAIT_TYPE int32

const (
	WaitAll = 0
	WaitAny = 1
	WaitNotification = 2
	WaitDequeue = 3
)

type NT_PRODUCT_TYPE int32

const (
	NtProductWinNt = 1
	NtProductLanManNt = 2
	NtProductServer = 3
)

type SUITE_TYPE int32

const (
	SmallBusiness = 0
	Enterprise = 1
	BackOffice = 2
	CommunicationServer = 3
	TerminalServer = 4
	SmallBusinessRestricted = 5
	EmbeddedNT = 6
	DataCenter = 7
	SingleUserTS = 8
	Personal = 9
	Blade = 10
	EmbeddedRestricted = 11
	SecurityAppliance = 12
	StorageServer = 13
	ComputeServer = 14
	WHServer = 15
	PhoneNT = 16
	MultiUserTS = 17
	MaxSuiteType = 18
)

type COMPARTMENT_ID int32

const (
	UNSPECIFIED_COMPARTMENT_ID = 0
	DEFAULT_COMPARTMENT_ID = 1
)
