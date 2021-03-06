// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsprogramming implements the Windows.Win32.WindowsProgramming namespace.
package windowsprogramming

type PROCESS_CREATION_FLAGS uint32

const (
	DEBUG_PROCESS                    PROCESS_CREATION_FLAGS = 1
	DEBUG_ONLY_THIS_PROCESS          PROCESS_CREATION_FLAGS = 2
	CREATE_SUSPENDED                 PROCESS_CREATION_FLAGS = 4
	DETACHED_PROCESS                 PROCESS_CREATION_FLAGS = 8
	CREATE_NEW_CONSOLE               PROCESS_CREATION_FLAGS = 16
	NORMAL_PRIORITY_CLASS            PROCESS_CREATION_FLAGS = 32
	IDLE_PRIORITY_CLASS              PROCESS_CREATION_FLAGS = 64
	HIGH_PRIORITY_CLASS              PROCESS_CREATION_FLAGS = 128
	REALTIME_PRIORITY_CLASS          PROCESS_CREATION_FLAGS = 256
	CREATE_NEW_PROCESS_GROUP         PROCESS_CREATION_FLAGS = 512
	CREATE_UNICODE_ENVIRONMENT       PROCESS_CREATION_FLAGS = 1024
	CREATE_SEPARATE_WOW_VDM          PROCESS_CREATION_FLAGS = 2048
	CREATE_SHARED_WOW_VDM            PROCESS_CREATION_FLAGS = 4096
	CREATE_FORCEDOS                  PROCESS_CREATION_FLAGS = 8192
	BELOW_NORMAL_PRIORITY_CLASS      PROCESS_CREATION_FLAGS = 16384
	ABOVE_NORMAL_PRIORITY_CLASS      PROCESS_CREATION_FLAGS = 32768
	INHERIT_PARENT_AFFINITY          PROCESS_CREATION_FLAGS = 65536
	INHERIT_CALLER_PRIORITY          PROCESS_CREATION_FLAGS = 131072
	CREATE_PROTECTED_PROCESS         PROCESS_CREATION_FLAGS = 262144
	EXTENDED_STARTUPINFO_PRESENT     PROCESS_CREATION_FLAGS = 524288
	PROCESS_MODE_BACKGROUND_BEGIN    PROCESS_CREATION_FLAGS = 1048576
	PROCESS_MODE_BACKGROUND_END      PROCESS_CREATION_FLAGS = 2097152
	CREATE_SECURE_PROCESS            PROCESS_CREATION_FLAGS = 4194304
	CREATE_BREAKAWAY_FROM_JOB        PROCESS_CREATION_FLAGS = 16777216
	CREATE_PRESERVE_CODE_AUTHZ_LEVEL PROCESS_CREATION_FLAGS = 33554432
	CREATE_DEFAULT_ERROR_MODE        PROCESS_CREATION_FLAGS = 67108864
	CREATE_NO_WINDOW                 PROCESS_CREATION_FLAGS = 134217728
	PROFILE_USER                     PROCESS_CREATION_FLAGS = 268435456
	PROFILE_KERNEL                   PROCESS_CREATION_FLAGS = 536870912
	PROFILE_SERVER                   PROCESS_CREATION_FLAGS = 1073741824
	CREATE_IGNORE_SYSTEM_DEFAULT     PROCESS_CREATION_FLAGS = 2147483648
)

type HANDLE_FLAG_OPTIONS int32

const (
	HANDLE_FLAG_INHERIT            HANDLE_FLAG_OPTIONS = 1
	HANDLE_FLAG_PROTECT_FROM_CLOSE HANDLE_FLAG_OPTIONS = 2
)

type DUPLICATE_HANDLE_OPTIONS int32

const (
	DUPLICATE_CLOSE_SOURCE DUPLICATE_HANDLE_OPTIONS = 1
	DUPLICATE_SAME_ACCESS  DUPLICATE_HANDLE_OPTIONS = 2
)

type STD_HANDLE_TYPE uint32

const (
	STD_INPUT_HANDLE  STD_HANDLE_TYPE = 4294967286
	STD_OUTPUT_HANDLE STD_HANDLE_TYPE = 4294967285
	STD_ERROR_HANDLE  STD_HANDLE_TYPE = 4294967284
)

type VER_FLAGS uint32

const (
	VER_MINORVERSION     VER_FLAGS = 1
	VER_MAJORVERSION     VER_FLAGS = 2
	VER_BUILDNUMBER      VER_FLAGS = 4
	VER_PLATFORMID       VER_FLAGS = 8
	VER_SERVICEPACKMINOR VER_FLAGS = 16
	VER_SERVICEPACKMAJOR VER_FLAGS = 32
	VER_SUITENAME        VER_FLAGS = 64
	VER_PRODUCT_TYPE     VER_FLAGS = 128
)

type FORMAT_MESSAGE_OPTIONS int32

const (
	FORMAT_MESSAGE_ALLOCATE_BUFFER FORMAT_MESSAGE_OPTIONS = 256
	FORMAT_MESSAGE_IGNORE_INSERTS  FORMAT_MESSAGE_OPTIONS = 512
	FORMAT_MESSAGE_FROM_STRING     FORMAT_MESSAGE_OPTIONS = 1024
	FORMAT_MESSAGE_FROM_HMODULE    FORMAT_MESSAGE_OPTIONS = 2048
	FORMAT_MESSAGE_FROM_SYSTEM     FORMAT_MESSAGE_OPTIONS = 4096
	FORMAT_MESSAGE_ARGUMENT_ARRAY  FORMAT_MESSAGE_OPTIONS = 8192
	FORMAT_MESSAGE_MAX_WIDTH_MASK  FORMAT_MESSAGE_OPTIONS = 255
)

type FIRMWARE_TYPE int32

const (
	FirmwareTypeUnknown FIRMWARE_TYPE = 0
	FirmwareTypeBios    FIRMWARE_TYPE = 1
	FirmwareTypeUefi    FIRMWARE_TYPE = 2
	FirmwareTypeMax     FIRMWARE_TYPE = 3
)

type UpdateImpactLevel int32

const (
	UpdateImpactLevel_None   UpdateImpactLevel = 0
	UpdateImpactLevel_Low    UpdateImpactLevel = 1
	UpdateImpactLevel_Medium UpdateImpactLevel = 2
	UpdateImpactLevel_High   UpdateImpactLevel = 3
)

type UpdateAssessmentStatus int32

const (
	UpdateAssessmentStatus_Latest                   UpdateAssessmentStatus = 0
	UpdateAssessmentStatus_NotLatestSoftRestriction UpdateAssessmentStatus = 1
	UpdateAssessmentStatus_NotLatestHardRestriction UpdateAssessmentStatus = 2
	UpdateAssessmentStatus_NotLatestEndOfSupport    UpdateAssessmentStatus = 3
	UpdateAssessmentStatus_NotLatestServicingTrain  UpdateAssessmentStatus = 4
	UpdateAssessmentStatus_NotLatestDeferredFeature UpdateAssessmentStatus = 5
	UpdateAssessmentStatus_NotLatestDeferredQuality UpdateAssessmentStatus = 6
	UpdateAssessmentStatus_NotLatestPausedFeature   UpdateAssessmentStatus = 7
	UpdateAssessmentStatus_NotLatestPausedQuality   UpdateAssessmentStatus = 8
	UpdateAssessmentStatus_NotLatestManaged         UpdateAssessmentStatus = 9
	UpdateAssessmentStatus_NotLatestUnknown         UpdateAssessmentStatus = 10
	UpdateAssessmentStatus_NotLatestTargetedVersion UpdateAssessmentStatus = 11
)

type THREAD_INFORMATION_CLASS int32

const (
	ThreadMemoryPriority      THREAD_INFORMATION_CLASS = 0
	ThreadAbsoluteCpuPriority THREAD_INFORMATION_CLASS = 1
	ThreadDynamicCodePolicy   THREAD_INFORMATION_CLASS = 2
	ThreadPowerThrottling     THREAD_INFORMATION_CLASS = 3
	ThreadInformationClassMax THREAD_INFORMATION_CLASS = 4
)

type COMPUTER_NAME_FORMAT int32

const (
	ComputerNameNetBIOS                   COMPUTER_NAME_FORMAT = 0
	ComputerNameDnsHostname               COMPUTER_NAME_FORMAT = 1
	ComputerNameDnsDomain                 COMPUTER_NAME_FORMAT = 2
	ComputerNameDnsFullyQualified         COMPUTER_NAME_FORMAT = 3
	ComputerNamePhysicalNetBIOS           COMPUTER_NAME_FORMAT = 4
	ComputerNamePhysicalDnsHostname       COMPUTER_NAME_FORMAT = 5
	ComputerNamePhysicalDnsDomain         COMPUTER_NAME_FORMAT = 6
	ComputerNamePhysicalDnsFullyQualified COMPUTER_NAME_FORMAT = 7
	ComputerNameMax                       COMPUTER_NAME_FORMAT = 8
)

type DEP_SYSTEM_POLICY_TYPE int32

const (
	DEPPolicyAlwaysOff  DEP_SYSTEM_POLICY_TYPE = 0
	DEPPolicyAlwaysOn   DEP_SYSTEM_POLICY_TYPE = 1
	DEPPolicyOptIn      DEP_SYSTEM_POLICY_TYPE = 2
	DEPPolicyOptOut     DEP_SYSTEM_POLICY_TYPE = 3
	DEPTotalPolicyCount DEP_SYSTEM_POLICY_TYPE = 4
)

type PROC_THREAD_ATTRIBUTE_NUM int32

const (
	ProcThreadAttributeParentProcess                PROC_THREAD_ATTRIBUTE_NUM = 0
	ProcThreadAttributeHandleList                   PROC_THREAD_ATTRIBUTE_NUM = 2
	ProcThreadAttributeGroupAffinity                PROC_THREAD_ATTRIBUTE_NUM = 3
	ProcThreadAttributePreferredNode                PROC_THREAD_ATTRIBUTE_NUM = 4
	ProcThreadAttributeIdealProcessor               PROC_THREAD_ATTRIBUTE_NUM = 5
	ProcThreadAttributeUmsThread                    PROC_THREAD_ATTRIBUTE_NUM = 6
	ProcThreadAttributeMitigationPolicy             PROC_THREAD_ATTRIBUTE_NUM = 7
	ProcThreadAttributeSecurityCapabilities         PROC_THREAD_ATTRIBUTE_NUM = 9
	ProcThreadAttributeProtectionLevel              PROC_THREAD_ATTRIBUTE_NUM = 11
	ProcThreadAttributeJobList                      PROC_THREAD_ATTRIBUTE_NUM = 13
	ProcThreadAttributeChildProcessPolicy           PROC_THREAD_ATTRIBUTE_NUM = 14
	ProcThreadAttributeAllApplicationPackagesPolicy PROC_THREAD_ATTRIBUTE_NUM = 15
	ProcThreadAttributeWin32kFilter                 PROC_THREAD_ATTRIBUTE_NUM = 16
	ProcThreadAttributeSafeOpenPromptOriginClaim    PROC_THREAD_ATTRIBUTE_NUM = 17
	ProcThreadAttributeDesktopAppPolicy             PROC_THREAD_ATTRIBUTE_NUM = 18
	ProcThreadAttributePseudoConsole                PROC_THREAD_ATTRIBUTE_NUM = 22
)

type DOMNodeType int32

const (
	NODE_INVALID                DOMNodeType = 0
	NODE_ELEMENT                DOMNodeType = 1
	NODE_ATTRIBUTE              DOMNodeType = 2
	NODE_TEXT                   DOMNodeType = 3
	NODE_CDATA_SECTION          DOMNodeType = 4
	NODE_ENTITY_REFERENCE       DOMNodeType = 5
	NODE_ENTITY                 DOMNodeType = 6
	NODE_PROCESSING_INSTRUCTION DOMNodeType = 7
	NODE_COMMENT                DOMNodeType = 8
	NODE_DOCUMENT               DOMNodeType = 9
	NODE_DOCUMENT_TYPE          DOMNodeType = 10
	NODE_DOCUMENT_FRAGMENT      DOMNodeType = 11
	NODE_NOTATION               DOMNodeType = 12
)

type XMLEMEM_TYPE int32

const (
	XMLELEMTYPE_ELEMENT  XMLEMEM_TYPE = 0
	XMLELEMTYPE_TEXT     XMLEMEM_TYPE = 1
	XMLELEMTYPE_COMMENT  XMLEMEM_TYPE = 2
	XMLELEMTYPE_DOCUMENT XMLEMEM_TYPE = 3
	XMLELEMTYPE_DTD      XMLEMEM_TYPE = 4
	XMLELEMTYPE_PI       XMLEMEM_TYPE = 5
	XMLELEMTYPE_OTHER    XMLEMEM_TYPE = 6
)

type FILE_INFORMATION_CLASS int32

const (
	FileDirectoryInformation FILE_INFORMATION_CLASS = 1
)

type PROCESSINFOCLASS int32

const (
	ProcessBasicInformation   PROCESSINFOCLASS = 0
	ProcessDebugPort          PROCESSINFOCLASS = 7
	ProcessWow64Information   PROCESSINFOCLASS = 26
	ProcessImageFileName      PROCESSINFOCLASS = 27
	ProcessBreakOnTermination PROCESSINFOCLASS = 29
)

type THREADINFOCLASS int32

const (
	ThreadIsIoPending THREADINFOCLASS = 16
)

type SYSTEM_INFORMATION_CLASS int32

const (
	SystemBasicInformation                SYSTEM_INFORMATION_CLASS = 0
	SystemPerformanceInformation          SYSTEM_INFORMATION_CLASS = 2
	SystemTimeOfDayInformation            SYSTEM_INFORMATION_CLASS = 3
	SystemProcessInformation              SYSTEM_INFORMATION_CLASS = 5
	SystemProcessorPerformanceInformation SYSTEM_INFORMATION_CLASS = 8
	SystemInterruptInformation            SYSTEM_INFORMATION_CLASS = 23
	SystemExceptionInformation            SYSTEM_INFORMATION_CLASS = 33
	SystemRegistryQuotaInformation        SYSTEM_INFORMATION_CLASS = 37
	SystemLookasideInformation            SYSTEM_INFORMATION_CLASS = 45
	SystemCodeIntegrityInformation        SYSTEM_INFORMATION_CLASS = 103
	SystemPolicyInformation               SYSTEM_INFORMATION_CLASS = 134
)

type OBJECT_INFORMATION_CLASS int32

const (
	ObjectBasicInformation OBJECT_INFORMATION_CLASS = 0
	ObjectTypeInformation  OBJECT_INFORMATION_CLASS = 2
)

type KEY_SET_INFORMATION_CLASS int32

const (
	KeyWriteTimeInformation         KEY_SET_INFORMATION_CLASS = 0
	KeyWow64FlagsInformation        KEY_SET_INFORMATION_CLASS = 1
	KeyControlFlagsInformation      KEY_SET_INFORMATION_CLASS = 2
	KeySetVirtualizationInformation KEY_SET_INFORMATION_CLASS = 3
	KeySetDebugInformation          KEY_SET_INFORMATION_CLASS = 4
	KeySetHandleTagsInformation     KEY_SET_INFORMATION_CLASS = 5
	MaxKeySetInfoClass              KEY_SET_INFORMATION_CLASS = 6
)

type WINSTATIONINFOCLASS int32

const (
	WinStationInformation WINSTATIONINFOCLASS = 8
)

type EUSERALLOCATIONSTATE int32

const (
	AllocationStateUnknown EUSERALLOCATIONSTATE = 0
	AllocationStateBusy    EUSERALLOCATIONSTATE = 1
	AllocationStateFree    EUSERALLOCATIONSTATE = 2
)

type EHEAPALLOCATIONSTATE int32

const (
	HeapFullPageHeap EHEAPALLOCATIONSTATE = 1073741824
	HeapMetadata     EHEAPALLOCATIONSTATE = -2147483648
	HeapStateMask    EHEAPALLOCATIONSTATE = -65536
)

type EHEAPENUMERATIONLEVEL int32

const (
	HeapEnumerationEverything EHEAPENUMERATIONLEVEL = 0
	HeapEnumerationStop       EHEAPENUMERATIONLEVEL = -1
)

type EHANDLE_TRACE_OPERATIONS int32

const (
	OperationDbUnused EHANDLE_TRACE_OPERATIONS = 0
	OperationDbOPEN   EHANDLE_TRACE_OPERATIONS = 1
	OperationDbCLOSE  EHANDLE_TRACE_OPERATIONS = 2
	OperationDbBADREF EHANDLE_TRACE_OPERATIONS = 3
)

type EAVRFRESOURCETYPES int32

const (
	AvrfResourceHeapAllocation EAVRFRESOURCETYPES = 0
	AvrfResourceHandleTrace    EAVRFRESOURCETYPES = 1
	AvrfResourceMax            EAVRFRESOURCETYPES = 2
)

type CameraUIControlMode int32

const (
	Browse CameraUIControlMode = 0
	Linear CameraUIControlMode = 1
)

type CameraUIControlLinearSelectionMode int32

const (
	Single   CameraUIControlLinearSelectionMode = 0
	Multiple CameraUIControlLinearSelectionMode = 1
)

type CameraUIControlCaptureMode int32

const (
	PhotoOrVideo CameraUIControlCaptureMode = 0
	Photo        CameraUIControlCaptureMode = 1
	Video        CameraUIControlCaptureMode = 2
)

type CameraUIControlPhotoFormat int32

const (
	Jpeg   CameraUIControlPhotoFormat = 0
	Png    CameraUIControlPhotoFormat = 1
	JpegXR CameraUIControlPhotoFormat = 2
)

type CameraUIControlVideoFormat int32

const (
	Mp4 CameraUIControlVideoFormat = 0
	Wmv CameraUIControlVideoFormat = 1
)

type CameraUIControlViewType int32

const (
	SingleItem CameraUIControlViewType = 0
	ItemList   CameraUIControlViewType = 1
)

type FCIERROR int32

const (
	FCIERR_NONE             FCIERROR = 0
	FCIERR_OPEN_SRC         FCIERROR = 1
	FCIERR_READ_SRC         FCIERROR = 2
	FCIERR_ALLOC_FAIL       FCIERROR = 3
	FCIERR_TEMP_FILE        FCIERROR = 4
	FCIERR_BAD_COMPR_TYPE   FCIERROR = 5
	FCIERR_CAB_FILE         FCIERROR = 6
	FCIERR_USER_ABORT       FCIERROR = 7
	FCIERR_MCI_FAIL         FCIERROR = 8
	FCIERR_CAB_FORMAT_LIMIT FCIERROR = 9
)

type FDIERROR int32

const (
	FDIERROR_NONE                    FDIERROR = 0
	FDIERROR_CABINET_NOT_FOUND       FDIERROR = 1
	FDIERROR_NOT_A_CABINET           FDIERROR = 2
	FDIERROR_UNKNOWN_CABINET_VERSION FDIERROR = 3
	FDIERROR_CORRUPT_CABINET         FDIERROR = 4
	FDIERROR_ALLOC_FAIL              FDIERROR = 5
	FDIERROR_BAD_COMPR_TYPE          FDIERROR = 6
	FDIERROR_MDI_FAIL                FDIERROR = 7
	FDIERROR_TARGET_FILE             FDIERROR = 8
	FDIERROR_RESERVE_MISMATCH        FDIERROR = 9
	FDIERROR_WRONG_CABINET           FDIERROR = 10
	FDIERROR_USER_ABORT              FDIERROR = 11
	FDIERROR_EOF                     FDIERROR = 12
)

type FDIDECRYPTTYPE int32

const (
	FDIDTNEW_CABINET FDIDECRYPTTYPE = 0
	FDIDTNEW_FOLDER  FDIDECRYPTTYPE = 1
	FDIDTDECRYPT     FDIDECRYPTTYPE = 2
)

type FDINOTIFICATIONTYPE int32

const (
	FDINTCABINET_INFO    FDINOTIFICATIONTYPE = 0
	FDINTPARTIAL_FILE    FDINOTIFICATIONTYPE = 1
	FDINTCOPY_FILE       FDINOTIFICATIONTYPE = 2
	FDINTCLOSE_FILE_INFO FDINOTIFICATIONTYPE = 3
	FDINTNEXT_CABINET    FDINOTIFICATIONTYPE = 4
	FDINTENUMERATE       FDINOTIFICATIONTYPE = 5
)

type FEATURE_CHANGE_TIME int32

const (
	FEATURE_CHANGE_TIME_READ          FEATURE_CHANGE_TIME = 0
	FEATURE_CHANGE_TIME_MODULE_RELOAD FEATURE_CHANGE_TIME = 1
	FEATURE_CHANGE_TIME_SESSION       FEATURE_CHANGE_TIME = 2
	FEATURE_CHANGE_TIME_REBOOT        FEATURE_CHANGE_TIME = 3
)

type FEATURE_ENABLED_STATE int32

const (
	FEATURE_ENABLED_STATE_DEFAULT  FEATURE_ENABLED_STATE = 0
	FEATURE_ENABLED_STATE_DISABLED FEATURE_ENABLED_STATE = 1
	FEATURE_ENABLED_STATE_ENABLED  FEATURE_ENABLED_STATE = 2
)

type FH_TARGET_PROPERTY_TYPE int32

const (
	FH_TARGET_NAME       FH_TARGET_PROPERTY_TYPE = 0
	FH_TARGET_URL        FH_TARGET_PROPERTY_TYPE = 1
	FH_TARGET_DRIVE_TYPE FH_TARGET_PROPERTY_TYPE = 2
	MAX_TARGET_PROPERTY  FH_TARGET_PROPERTY_TYPE = 3
)

type FH_TARGET_DRIVE_TYPES int32

const (
	FH_DRIVE_UNKNOWN   FH_TARGET_DRIVE_TYPES = 0
	FH_DRIVE_REMOVABLE FH_TARGET_DRIVE_TYPES = 2
	FH_DRIVE_FIXED     FH_TARGET_DRIVE_TYPES = 3
	FH_DRIVE_REMOTE    FH_TARGET_DRIVE_TYPES = 4
)

type FH_PROTECTED_ITEM_CATEGORY int32

const (
	FH_FOLDER                   FH_PROTECTED_ITEM_CATEGORY = 0
	FH_LIBRARY                  FH_PROTECTED_ITEM_CATEGORY = 1
	MAX_PROTECTED_ITEM_CATEGORY FH_PROTECTED_ITEM_CATEGORY = 2
)

type FH_LOCAL_POLICY_TYPE int32

const (
	FH_FREQUENCY      FH_LOCAL_POLICY_TYPE = 0
	FH_RETENTION_TYPE FH_LOCAL_POLICY_TYPE = 1
	FH_RETENTION_AGE  FH_LOCAL_POLICY_TYPE = 2
	MAX_LOCAL_POLICY  FH_LOCAL_POLICY_TYPE = 3
)

type FH_RETENTION_TYPES int32

const (
	FH_RETENTION_DISABLED  FH_RETENTION_TYPES = 0
	FH_RETENTION_UNLIMITED FH_RETENTION_TYPES = 1
	FH_RETENTION_AGE_BASED FH_RETENTION_TYPES = 2
	MAX_RETENTION_TYPE     FH_RETENTION_TYPES = 3
)

type FH_BACKUP_STATUS int32

const (
	FH_STATUS_DISABLED       FH_BACKUP_STATUS = 0
	FH_STATUS_DISABLED_BY_GP FH_BACKUP_STATUS = 1
	FH_STATUS_ENABLED        FH_BACKUP_STATUS = 2
	FH_STATUS_REHYDRATING    FH_BACKUP_STATUS = 3
	MAX_BACKUP_STATUS        FH_BACKUP_STATUS = 4
)

type FH_DEVICE_VALIDATION_RESULT int32

const (
	FH_ACCESS_DENIED          FH_DEVICE_VALIDATION_RESULT = 0
	FH_INVALID_DRIVE_TYPE     FH_DEVICE_VALIDATION_RESULT = 1
	FH_READ_ONLY_PERMISSION   FH_DEVICE_VALIDATION_RESULT = 2
	FH_CURRENT_DEFAULT        FH_DEVICE_VALIDATION_RESULT = 3
	FH_NAMESPACE_EXISTS       FH_DEVICE_VALIDATION_RESULT = 4
	FH_TARGET_PART_OF_LIBRARY FH_DEVICE_VALIDATION_RESULT = 5
	FH_VALID_TARGET           FH_DEVICE_VALIDATION_RESULT = 6
	MAX_VALIDATION_RESULT     FH_DEVICE_VALIDATION_RESULT = 7
)

type FhBackupStopReason int32

const (
	BackupInvalidStopReason        FhBackupStopReason = 0
	BackupLimitUserBusyMachineOnAC FhBackupStopReason = 1
	BackupLimitUserIdleMachineOnDC FhBackupStopReason = 2
	BackupLimitUserBusyMachineOnDC FhBackupStopReason = 3
	BackupCancelled                FhBackupStopReason = 4
)

type CommandStateChangeConstants int32

const (
	CSC_UPDATECOMMANDS  CommandStateChangeConstants = -1
	CSC_NAVIGATEFORWARD CommandStateChangeConstants = 1
	CSC_NAVIGATEBACK    CommandStateChangeConstants = 2
)

type SecureLockIconConstants int32

const (
	SECURELOCKICONUNSECURE          SecureLockIconConstants = 0
	SECURELOCKICONMIXED             SecureLockIconConstants = 1
	SECURELOCKICONSECUREUNKNOWNBITS SecureLockIconConstants = 2
	SECURELOCKICONSECURE40BIT       SecureLockIconConstants = 3
	SECURELOCKICONSECURE56BIT       SecureLockIconConstants = 4
	SECURELOCKICONSECUREFORTEZZA    SecureLockIconConstants = 5
	SECURELOCKICONSECURE128BIT      SecureLockIconConstants = 6
)

type NewProcessCauseConstants int32

const (
	ProtectedModeRedirect NewProcessCauseConstants = 1
)

type BrowserNavConstants int32

const (
	NAVOPENINNEWWINDOW       BrowserNavConstants = 1
	NAVNOHISTORY             BrowserNavConstants = 2
	NAVNOREADFROMCACHE       BrowserNavConstants = 4
	NAVNOWRITETOCACHE        BrowserNavConstants = 8
	NAVALLOWAUTOSEARCH       BrowserNavConstants = 16
	NAVBROWSERBAR            BrowserNavConstants = 32
	NAVHYPERLINK             BrowserNavConstants = 64
	NAVENFORCERESTRICTED     BrowserNavConstants = 128
	NAVNEWWINDOWSMANAGED     BrowserNavConstants = 256
	NAVUNTRUSTEDFORDOWNLOAD  BrowserNavConstants = 512
	NAVTRUSTEDFORACTIVEX     BrowserNavConstants = 1024
	NAVOPENINNEWTAB          BrowserNavConstants = 2048
	NAVOPENINBACKGROUNDTAB   BrowserNavConstants = 4096
	NAVKEEPWORDWHEELTEXT     BrowserNavConstants = 8192
	NAVVIRTUALTAB            BrowserNavConstants = 16384
	NAVBLOCKREDIRECTSXDOMAIN BrowserNavConstants = 32768
	NAVOPENNEWFOREGROUNDTAB  BrowserNavConstants = 65536
	NAVTRAVELLOGSCREENSHOT   BrowserNavConstants = 131072
	NAVDEFERUNLOAD           BrowserNavConstants = 262144
	NAVSPECULATIVE           BrowserNavConstants = 524288
	NAVSUGGESTNEWWINDOW      BrowserNavConstants = 1048576
	NAVSUGGESTNEWTAB         BrowserNavConstants = 2097152
	NAVRESERVED1             BrowserNavConstants = 4194304
	NAVHOMEPAGENAVIGATE      BrowserNavConstants = 8388608
	NAVREFRESH               BrowserNavConstants = 16777216
	NAVHOSTNAVIGATION        BrowserNavConstants = 33554432
	NAVRESERVED2             BrowserNavConstants = 67108864
	NAVRESERVED3             BrowserNavConstants = 134217728
	NAVRESERVED4             BrowserNavConstants = 268435456
	NAVRESERVED5             BrowserNavConstants = 536870912
	NAVRESERVED6             BrowserNavConstants = 1073741824
	NAVRESERVED7             BrowserNavConstants = -2147483648
)

type RefreshConstants int32

const (
	REFRESH_NORMAL     RefreshConstants = 0
	REFRESH_IFEXPIRED  RefreshConstants = 1
	REFRESH_COMPLETELY RefreshConstants = 3
)

type WSC_SECURITY_PRODUCT_SUBSTATUS int32

const (
	WSC_SECURITY_PRODUCT_SUBSTATUS_NOT_SET            WSC_SECURITY_PRODUCT_SUBSTATUS = 0
	WSC_SECURITY_PRODUCT_SUBSTATUS_NO_ACTION          WSC_SECURITY_PRODUCT_SUBSTATUS = 1
	WSC_SECURITY_PRODUCT_SUBSTATUS_ACTION_RECOMMENDED WSC_SECURITY_PRODUCT_SUBSTATUS = 2
	WSC_SECURITY_PRODUCT_SUBSTATUS_ACTION_NEEDED      WSC_SECURITY_PRODUCT_SUBSTATUS = 3
)

type WSC_SECURITY_PRODUCT_STATE int32

const (
	WSC_SECURITY_PRODUCT_STATE_ON      WSC_SECURITY_PRODUCT_STATE = 0
	WSC_SECURITY_PRODUCT_STATE_OFF     WSC_SECURITY_PRODUCT_STATE = 1
	WSC_SECURITY_PRODUCT_STATE_SNOOZED WSC_SECURITY_PRODUCT_STATE = 2
	WSC_SECURITY_PRODUCT_STATE_EXPIRED WSC_SECURITY_PRODUCT_STATE = 3
)

type SECURITY_PRODUCT_TYPE int32

const (
	SECURITY_PRODUCT_TYPE_ANTIVIRUS   SECURITY_PRODUCT_TYPE = 0
	SECURITY_PRODUCT_TYPE_FIREWALL    SECURITY_PRODUCT_TYPE = 1
	SECURITY_PRODUCT_TYPE_ANTISPYWARE SECURITY_PRODUCT_TYPE = 2
)

type WSC_SECURITY_SIGNATURE_STATUS int32

const (
	WSC_SECURITY_PRODUCT_OUT_OF_DATE WSC_SECURITY_SIGNATURE_STATUS = 0
	WSC_SECURITY_PRODUCT_UP_TO_DATE  WSC_SECURITY_SIGNATURE_STATUS = 1
)

type WSC_SECURITY_PROVIDER int32

const (
	WSC_SECURITY_PROVIDER_FIREWALL             WSC_SECURITY_PROVIDER = 1
	WSC_SECURITY_PROVIDER_AUTOUPDATE_SETTINGS  WSC_SECURITY_PROVIDER = 2
	WSC_SECURITY_PROVIDER_ANTIVIRUS            WSC_SECURITY_PROVIDER = 4
	WSC_SECURITY_PROVIDER_ANTISPYWARE          WSC_SECURITY_PROVIDER = 8
	WSC_SECURITY_PROVIDER_INTERNET_SETTINGS    WSC_SECURITY_PROVIDER = 16
	WSC_SECURITY_PROVIDER_USER_ACCOUNT_CONTROL WSC_SECURITY_PROVIDER = 32
	WSC_SECURITY_PROVIDER_SERVICE              WSC_SECURITY_PROVIDER = 64
	WSC_SECURITY_PROVIDER_NONE                 WSC_SECURITY_PROVIDER = 0
	WSC_SECURITY_PROVIDER_ALL                  WSC_SECURITY_PROVIDER = 127
)

type WSC_SECURITY_PROVIDER_HEALTH int32

const (
	WSC_SECURITY_PROVIDER_HEALTH_GOOD         WSC_SECURITY_PROVIDER_HEALTH = 0
	WSC_SECURITY_PROVIDER_HEALTH_NOTMONITORED WSC_SECURITY_PROVIDER_HEALTH = 1
	WSC_SECURITY_PROVIDER_HEALTH_POOR         WSC_SECURITY_PROVIDER_HEALTH = 2
	WSC_SECURITY_PROVIDER_HEALTH_SNOOZE       WSC_SECURITY_PROVIDER_HEALTH = 3
)

type TDI_TL_IO_CONTROL_TYPE int32

const (
	EndpointIoControlType   TDI_TL_IO_CONTROL_TYPE = 0
	SetSockOptIoControlType TDI_TL_IO_CONTROL_TYPE = 1
	GetSockOptIoControlType TDI_TL_IO_CONTROL_TYPE = 2
	SocketIoControlType     TDI_TL_IO_CONTROL_TYPE = 3
)

type WLDP_HOST int32

const (
	WLDP_HOST_RUNDLL32 WLDP_HOST = 0
	WLDP_HOST_SVCHOST  WLDP_HOST = 1
	WLDP_HOST_MAX      WLDP_HOST = 2
)

type WLDP_HOST_ID int32

const (
	WLDP_HOST_ID_UNKNOWN    WLDP_HOST_ID = 0
	WLDP_HOST_ID_GLOBAL     WLDP_HOST_ID = 1
	WLDP_HOST_ID_VBA        WLDP_HOST_ID = 2
	WLDP_HOST_ID_WSH        WLDP_HOST_ID = 3
	WLDP_HOST_ID_POWERSHELL WLDP_HOST_ID = 4
	WLDP_HOST_ID_IE         WLDP_HOST_ID = 5
	WLDP_HOST_ID_MSI        WLDP_HOST_ID = 6
	WLDP_HOST_ID_ALL        WLDP_HOST_ID = 7
	WLDP_HOST_ID_MAX        WLDP_HOST_ID = 8
)

type DECISION_LOCATION int32

const (
	DECISION_LOCATION_REFRESH_GLOBAL_DATA         DECISION_LOCATION = 0
	DECISION_LOCATION_PARAMETER_VALIDATION        DECISION_LOCATION = 1
	DECISION_LOCATION_AUDIT                       DECISION_LOCATION = 2
	DECISION_LOCATION_FAILED_CONVERT_GUID         DECISION_LOCATION = 3
	DECISION_LOCATION_ENTERPRISE_DEFINED_CLASS_ID DECISION_LOCATION = 4
	DECISION_LOCATION_GLOBAL_BUILT_IN_LIST        DECISION_LOCATION = 5
	DECISION_LOCATION_PROVIDER_BUILT_IN_LIST      DECISION_LOCATION = 6
	DECISION_LOCATION_ENFORCE_STATE_LIST          DECISION_LOCATION = 7
	DECISION_LOCATION_NOT_FOUND                   DECISION_LOCATION = 8
	DECISION_LOCATION_UNKNOWN                     DECISION_LOCATION = 9
)

type WLDP_KEY int32

const (
	KEY_UNKNOWN  WLDP_KEY = 0
	KEY_OVERRIDE WLDP_KEY = 1
	KEY_ALL_KEYS WLDP_KEY = 2
)

type VALUENAME int32

const (
	VALUENAME_UNKNOWN                     VALUENAME = 0
	VALUENAME_ENTERPRISE_DEFINED_CLASS_ID VALUENAME = 1
	VALUENAME_BUILT_IN_LIST               VALUENAME = 2
)

type WLDP_WINDOWS_LOCKDOWN_MODE int32

const (
	WLDP_WINDOWS_LOCKDOWN_MODE_UNLOCKED WLDP_WINDOWS_LOCKDOWN_MODE = 0
	WLDP_WINDOWS_LOCKDOWN_MODE_TRIAL    WLDP_WINDOWS_LOCKDOWN_MODE = 1
	WLDP_WINDOWS_LOCKDOWN_MODE_LOCKED   WLDP_WINDOWS_LOCKDOWN_MODE = 2
	WLDP_WINDOWS_LOCKDOWN_MODE_MAX      WLDP_WINDOWS_LOCKDOWN_MODE = 3
)

type WLDP_WINDOWS_LOCKDOWN_RESTRICTION int32

const (
	WLDP_WINDOWS_LOCKDOWN_RESTRICTION_NONE               WLDP_WINDOWS_LOCKDOWN_RESTRICTION = 0
	WLDP_WINDOWS_LOCKDOWN_RESTRICTION_NOUNLOCK           WLDP_WINDOWS_LOCKDOWN_RESTRICTION = 1
	WLDP_WINDOWS_LOCKDOWN_RESTRICTION_NOUNLOCK_PERMANENT WLDP_WINDOWS_LOCKDOWN_RESTRICTION = 2
	WLDP_WINDOWS_LOCKDOWN_RESTRICTION_MAX                WLDP_WINDOWS_LOCKDOWN_RESTRICTION = 3
)

type XmlNodeType int32

const (
	XmlNodeType_None                  XmlNodeType = 0
	XmlNodeType_Element               XmlNodeType = 1
	XmlNodeType_Attribute             XmlNodeType = 2
	XmlNodeType_Text                  XmlNodeType = 3
	XmlNodeType_CDATA                 XmlNodeType = 4
	XmlNodeType_ProcessingInstruction XmlNodeType = 7
	XmlNodeType_Comment               XmlNodeType = 8
	XmlNodeType_DocumentType          XmlNodeType = 10
	XmlNodeType_Whitespace            XmlNodeType = 13
	XmlNodeType_EndElement            XmlNodeType = 15
	XmlNodeType_XmlDeclaration        XmlNodeType = 17
	XmlNodeType_Last                  XmlNodeType = 17
)

type XmlConformanceLevel int32

const (
	XmlConformanceLevel_Auto     XmlConformanceLevel = 0
	XmlConformanceLevel_Fragment XmlConformanceLevel = 1
	XmlConformanceLevel_Document XmlConformanceLevel = 2
	XmlConformanceLevel_Last     XmlConformanceLevel = 2
)

type DtdProcessing int32

const (
	DtdProcessing_Prohibit DtdProcessing = 0
	DtdProcessing_Parse    DtdProcessing = 1
	DtdProcessing_Last     DtdProcessing = 1
)

type XmlReadState int32

const (
	XmlReadState_Initial     XmlReadState = 0
	XmlReadState_Interactive XmlReadState = 1
	XmlReadState_Error       XmlReadState = 2
	XmlReadState_EndOfFile   XmlReadState = 3
	XmlReadState_Closed      XmlReadState = 4
)

type XmlReaderProperty int32

const (
	XmlReaderProperty_MultiLanguage      XmlReaderProperty = 0
	XmlReaderProperty_ConformanceLevel   XmlReaderProperty = 1
	XmlReaderProperty_RandomAccess       XmlReaderProperty = 2
	XmlReaderProperty_XmlResolver        XmlReaderProperty = 3
	XmlReaderProperty_DtdProcessing      XmlReaderProperty = 4
	XmlReaderProperty_ReadState          XmlReaderProperty = 5
	XmlReaderProperty_MaxElementDepth    XmlReaderProperty = 6
	XmlReaderProperty_MaxEntityExpansion XmlReaderProperty = 7
	XmlReaderProperty_Last               XmlReaderProperty = 7
)

type XmlError int32

const (
	MX_E_MX                     XmlError = -1072894464
	MX_E_INPUTEND               XmlError = -1072894463
	MX_E_ENCODING               XmlError = -1072894462
	MX_E_ENCODINGSWITCH         XmlError = -1072894461
	MX_E_ENCODINGSIGNATURE      XmlError = -1072894460
	WC_E_WC                     XmlError = -1072894432
	WC_E_WHITESPACE             XmlError = -1072894431
	WC_E_SEMICOLON              XmlError = -1072894430
	WC_E_GREATERTHAN            XmlError = -1072894429
	WC_E_QUOTE                  XmlError = -1072894428
	WC_E_EQUAL                  XmlError = -1072894427
	WC_E_LESSTHAN               XmlError = -1072894426
	WC_E_HEXDIGIT               XmlError = -1072894425
	WC_E_DIGIT                  XmlError = -1072894424
	WC_E_LEFTBRACKET            XmlError = -1072894423
	WC_E_LEFTPAREN              XmlError = -1072894422
	WC_E_XMLCHARACTER           XmlError = -1072894421
	WC_E_NAMECHARACTER          XmlError = -1072894420
	WC_E_SYNTAX                 XmlError = -1072894419
	WC_E_CDSECT                 XmlError = -1072894418
	WC_E_COMMENT                XmlError = -1072894417
	WC_E_CONDSECT               XmlError = -1072894416
	WC_E_DECLATTLIST            XmlError = -1072894415
	WC_E_DECLDOCTYPE            XmlError = -1072894414
	WC_E_DECLELEMENT            XmlError = -1072894413
	WC_E_DECLENTITY             XmlError = -1072894412
	WC_E_DECLNOTATION           XmlError = -1072894411
	WC_E_NDATA                  XmlError = -1072894410
	WC_E_PUBLIC                 XmlError = -1072894409
	WC_E_SYSTEM                 XmlError = -1072894408
	WC_E_NAME                   XmlError = -1072894407
	WC_E_ROOTELEMENT            XmlError = -1072894406
	WC_E_ELEMENTMATCH           XmlError = -1072894405
	WC_E_UNIQUEATTRIBUTE        XmlError = -1072894404
	WC_E_TEXTXMLDECL            XmlError = -1072894403
	WC_E_LEADINGXML             XmlError = -1072894402
	WC_E_TEXTDECL               XmlError = -1072894401
	WC_E_XMLDECL                XmlError = -1072894400
	WC_E_ENCNAME                XmlError = -1072894399
	WC_E_PUBLICID               XmlError = -1072894398
	WC_E_PESINTERNALSUBSET      XmlError = -1072894397
	WC_E_PESBETWEENDECLS        XmlError = -1072894396
	WC_E_NORECURSION            XmlError = -1072894395
	WC_E_ENTITYCONTENT          XmlError = -1072894394
	WC_E_UNDECLAREDENTITY       XmlError = -1072894393
	WC_E_PARSEDENTITY           XmlError = -1072894392
	WC_E_NOEXTERNALENTITYREF    XmlError = -1072894391
	WC_E_PI                     XmlError = -1072894390
	WC_E_SYSTEMID               XmlError = -1072894389
	WC_E_QUESTIONMARK           XmlError = -1072894388
	WC_E_CDSECTEND              XmlError = -1072894387
	WC_E_MOREDATA               XmlError = -1072894386
	WC_E_DTDPROHIBITED          XmlError = -1072894385
	WC_E_INVALIDXMLSPACE        XmlError = -1072894384
	NC_E_NC                     XmlError = -1072894368
	NC_E_QNAMECHARACTER         XmlError = -1072894367
	NC_E_QNAMECOLON             XmlError = -1072894366
	NC_E_NAMECOLON              XmlError = -1072894365
	NC_E_DECLAREDPREFIX         XmlError = -1072894364
	NC_E_UNDECLAREDPREFIX       XmlError = -1072894363
	NC_E_EMPTYURI               XmlError = -1072894362
	NC_E_XMLPREFIXRESERVED      XmlError = -1072894361
	NC_E_XMLNSPREFIXRESERVED    XmlError = -1072894360
	NC_E_XMLURIRESERVED         XmlError = -1072894359
	NC_E_XMLNSURIRESERVED       XmlError = -1072894358
	SC_E_SC                     XmlError = -1072894336
	SC_E_MAXELEMENTDEPTH        XmlError = -1072894335
	SC_E_MAXENTITYEXPANSION     XmlError = -1072894334
	WR_E_WR                     XmlError = -1072894208
	WR_E_NONWHITESPACE          XmlError = -1072894207
	WR_E_NSPREFIXDECLARED       XmlError = -1072894206
	WR_E_NSPREFIXWITHEMPTYNSURI XmlError = -1072894205
	WR_E_DUPLICATEATTRIBUTE     XmlError = -1072894204
	WR_E_XMLNSPREFIXDECLARATION XmlError = -1072894203
	WR_E_XMLPREFIXDECLARATION   XmlError = -1072894202
	WR_E_XMLURIDECLARATION      XmlError = -1072894201
	WR_E_XMLNSURIDECLARATION    XmlError = -1072894200
	WR_E_NAMESPACEUNDECLARED    XmlError = -1072894199
	WR_E_INVALIDXMLSPACE        XmlError = -1072894198
	WR_E_INVALIDACTION          XmlError = -1072894197
	WR_E_INVALIDSURROGATEPAIR   XmlError = -1072894196
	XML_E_INVALID_DECIMAL       XmlError = -1072898019
	XML_E_INVALID_HEXIDECIMAL   XmlError = -1072898018
	XML_E_INVALID_UNICODE       XmlError = -1072898017
	XML_E_INVALIDENCODING       XmlError = -1072897938
)

type XmlStandalone int32

const (
	XmlStandalone_Omit XmlStandalone = 0
	XmlStandalone_Yes  XmlStandalone = 1
	XmlStandalone_No   XmlStandalone = 2
	XmlStandalone_Last XmlStandalone = 2
)

type XmlWriterProperty int32

const (
	XmlWriterProperty_MultiLanguage       XmlWriterProperty = 0
	XmlWriterProperty_Indent              XmlWriterProperty = 1
	XmlWriterProperty_ByteOrderMark       XmlWriterProperty = 2
	XmlWriterProperty_OmitXmlDeclaration  XmlWriterProperty = 3
	XmlWriterProperty_ConformanceLevel    XmlWriterProperty = 4
	XmlWriterProperty_CompactEmptyElement XmlWriterProperty = 5
	XmlWriterProperty_Last                XmlWriterProperty = 5
)

type DEVPROP_OPERATOR int32

const (
	DEVPROP_OPERATOR_MODIFIER_NOT                         DEVPROP_OPERATOR = 65536
	DEVPROP_OPERATOR_MODIFIER_IGNORE_CASE                 DEVPROP_OPERATOR = 131072
	DEVPROP_OPERATOR_NONE                                 DEVPROP_OPERATOR = 0
	DEVPROP_OPERATOR_EXISTS                               DEVPROP_OPERATOR = 1
	DEVPROP_OPERATOR_NOT_EXISTS                           DEVPROP_OPERATOR = 65537
	DEVPROP_OPERATOR_EQUALS                               DEVPROP_OPERATOR = 2
	DEVPROP_OPERATOR_NOT_EQUALS                           DEVPROP_OPERATOR = 65538
	DEVPROP_OPERATOR_GREATER_THAN                         DEVPROP_OPERATOR = 3
	DEVPROP_OPERATOR_LESS_THAN                            DEVPROP_OPERATOR = 4
	DEVPROP_OPERATOR_GREATER_THAN_EQUALS                  DEVPROP_OPERATOR = 5
	DEVPROP_OPERATOR_LESS_THAN_EQUALS                     DEVPROP_OPERATOR = 6
	DEVPROP_OPERATOR_EQUALS_IGNORE_CASE                   DEVPROP_OPERATOR = 131074
	DEVPROP_OPERATOR_NOT_EQUALS_IGNORE_CASE               DEVPROP_OPERATOR = 196610
	DEVPROP_OPERATOR_BITWISE_AND                          DEVPROP_OPERATOR = 7
	DEVPROP_OPERATOR_BITWISE_OR                           DEVPROP_OPERATOR = 8
	DEVPROP_OPERATOR_BEGINS_WITH                          DEVPROP_OPERATOR = 9
	DEVPROP_OPERATOR_ENDS_WITH                            DEVPROP_OPERATOR = 10
	DEVPROP_OPERATOR_CONTAINS                             DEVPROP_OPERATOR = 11
	DEVPROP_OPERATOR_BEGINS_WITH_IGNORE_CASE              DEVPROP_OPERATOR = 131081
	DEVPROP_OPERATOR_ENDS_WITH_IGNORE_CASE                DEVPROP_OPERATOR = 131082
	DEVPROP_OPERATOR_CONTAINS_IGNORE_CASE                 DEVPROP_OPERATOR = 131083
	DEVPROP_OPERATOR_LIST_CONTAINS                        DEVPROP_OPERATOR = 4096
	DEVPROP_OPERATOR_LIST_ELEMENT_BEGINS_WITH             DEVPROP_OPERATOR = 8192
	DEVPROP_OPERATOR_LIST_ELEMENT_ENDS_WITH               DEVPROP_OPERATOR = 12288
	DEVPROP_OPERATOR_LIST_ELEMENT_CONTAINS                DEVPROP_OPERATOR = 16384
	DEVPROP_OPERATOR_LIST_CONTAINS_IGNORE_CASE            DEVPROP_OPERATOR = 135168
	DEVPROP_OPERATOR_LIST_ELEMENT_BEGINS_WITH_IGNORE_CASE DEVPROP_OPERATOR = 139264
	DEVPROP_OPERATOR_LIST_ELEMENT_ENDS_WITH_IGNORE_CASE   DEVPROP_OPERATOR = 143360
	DEVPROP_OPERATOR_LIST_ELEMENT_CONTAINS_IGNORE_CASE    DEVPROP_OPERATOR = 147456
	DEVPROP_OPERATOR_AND_OPEN                             DEVPROP_OPERATOR = 1048576
	DEVPROP_OPERATOR_AND_CLOSE                            DEVPROP_OPERATOR = 2097152
	DEVPROP_OPERATOR_OR_OPEN                              DEVPROP_OPERATOR = 3145728
	DEVPROP_OPERATOR_OR_CLOSE                             DEVPROP_OPERATOR = 4194304
	DEVPROP_OPERATOR_NOT_OPEN                             DEVPROP_OPERATOR = 5242880
	DEVPROP_OPERATOR_NOT_CLOSE                            DEVPROP_OPERATOR = 6291456
	DEVPROP_OPERATOR_ARRAY_CONTAINS                       DEVPROP_OPERATOR = 268435456
	DEVPROP_OPERATOR_MASK_EVAL                            DEVPROP_OPERATOR = 4095
	DEVPROP_OPERATOR_MASK_LIST                            DEVPROP_OPERATOR = 61440
	DEVPROP_OPERATOR_MASK_MODIFIER                        DEVPROP_OPERATOR = 983040
	DEVPROP_OPERATOR_MASK_NOT_LOGICAL                     DEVPROP_OPERATOR = -267386881
	DEVPROP_OPERATOR_MASK_LOGICAL                         DEVPROP_OPERATOR = 267386880
	DEVPROP_OPERATOR_MASK_ARRAY                           DEVPROP_OPERATOR = -268435456
)

type DEV_OBJECT_TYPE int32

const (
	DevObjectTypeUnknown                DEV_OBJECT_TYPE = 0
	DevObjectTypeDeviceInterface        DEV_OBJECT_TYPE = 1
	DevObjectTypeDeviceContainer        DEV_OBJECT_TYPE = 2
	DevObjectTypeDevice                 DEV_OBJECT_TYPE = 3
	DevObjectTypeDeviceInterfaceClass   DEV_OBJECT_TYPE = 4
	DevObjectTypeAEP                    DEV_OBJECT_TYPE = 5
	DevObjectTypeAEPContainer           DEV_OBJECT_TYPE = 6
	DevObjectTypeDeviceInstallerClass   DEV_OBJECT_TYPE = 7
	DevObjectTypeDeviceInterfaceDisplay DEV_OBJECT_TYPE = 8
	DevObjectTypeDeviceContainerDisplay DEV_OBJECT_TYPE = 9
	DevObjectTypeAEPService             DEV_OBJECT_TYPE = 10
	DevObjectTypeDevicePanel            DEV_OBJECT_TYPE = 11
)

type DEV_QUERY_FLAGS int32

const (
	DevQueryFlagNone          DEV_QUERY_FLAGS = 0
	DevQueryFlagUpdateResults DEV_QUERY_FLAGS = 1
	DevQueryFlagAllProperties DEV_QUERY_FLAGS = 2
	DevQueryFlagLocalize      DEV_QUERY_FLAGS = 4
	DevQueryFlagAsyncClose    DEV_QUERY_FLAGS = 8
)

type DEV_QUERY_STATE int32

const (
	DevQueryStateInitialized   DEV_QUERY_STATE = 0
	DevQueryStateEnumCompleted DEV_QUERY_STATE = 1
	DevQueryStateAborted       DEV_QUERY_STATE = 2
	DevQueryStateClosed        DEV_QUERY_STATE = 3
)

type DEV_QUERY_RESULT_ACTION int32

const (
	DevQueryResultStateChange DEV_QUERY_RESULT_ACTION = 0
	DevQueryResultAdd         DEV_QUERY_RESULT_ACTION = 1
	DevQueryResultUpdate      DEV_QUERY_RESULT_ACTION = 2
	DevQueryResultRemove      DEV_QUERY_RESULT_ACTION = 3
)

type GlobalFilter int32

const (
	GF_FRAGMENTS  GlobalFilter = 2
	GF_STRONGHOST GlobalFilter = 8
	GF_FRAGCACHE  GlobalFilter = 9
)

type PfForwardAction int32

const (
	PF_ACTION_FORWARD PfForwardAction = 0
	PF_ACTION_DROP    PfForwardAction = 1
)

type PfAddresType int32

const (
	PF_IPV4 PfAddresType = 0
	PF_IPV6 PfAddresType = 1
)

type PfFrameType int32

const (
	PFFT_FILTER PfFrameType = 1
	PFFT_FRAG   PfFrameType = 2
	PFFT_SPOOF  PfFrameType = 3
)

type EXTENDED_NAME_FORMAT int32

const (
	NameUnknown          EXTENDED_NAME_FORMAT = 0
	NameFullyQualifiedDN EXTENDED_NAME_FORMAT = 1
	NameSamCompatible    EXTENDED_NAME_FORMAT = 2
	NameDisplay          EXTENDED_NAME_FORMAT = 3
	NameUniqueId         EXTENDED_NAME_FORMAT = 6
	NameCanonical        EXTENDED_NAME_FORMAT = 7
	NameUserPrincipal    EXTENDED_NAME_FORMAT = 8
	NameCanonicalEx      EXTENDED_NAME_FORMAT = 9
	NameServicePrincipal EXTENDED_NAME_FORMAT = 10
	NameDnsDomain        EXTENDED_NAME_FORMAT = 12
	NameGivenName        EXTENDED_NAME_FORMAT = 13
	NameSurname          EXTENDED_NAME_FORMAT = 14
)
