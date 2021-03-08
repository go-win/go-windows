// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package filesystem implements the Windows.Win32.FileSystem namespace.
package filesystem

type FIND_FIRST_EX_FLAGS uint32

const (
	FIND_FIRST_EX_CASE_SENSITIVE = 1
	FIND_FIRST_EX_LARGE_FETCH = 2
	FIND_FIRST_EX_ON_DISK_ENTRIES_ONLY = 4
)

type FILE_NOTIFY_CHANGE uint32

const (
	FILE_NOTIFY_CHANGE_FILE_NAME = 1
	FILE_NOTIFY_CHANGE_DIR_NAME = 2
	FILE_NOTIFY_CHANGE_ATTRIBUTES = 4
	FILE_NOTIFY_CHANGE_SIZE = 8
	FILE_NOTIFY_CHANGE_LAST_WRITE = 16
	FILE_NOTIFY_CHANGE_LAST_ACCESS = 32
	FILE_NOTIFY_CHANGE_CREATION = 64
	FILE_NOTIFY_CHANGE_SECURITY = 256
)

type DEFINE_DOS_DEVICE_FLAGS uint32

const (
	DDD_RAW_TARGET_PATH = 1
	DDD_REMOVE_DEFINITION = 2
	DDD_EXACT_MATCH_ON_REMOVE = 4
	DDD_NO_BROADCAST_SYSTEM = 8
	DDD_LUID_BROADCAST_DRIVE = 16
)

type FILE_CREATE_FLAGS uint32

const (
	CREATE_NEW = 1
	CREATE_ALWAYS = 2
	OPEN_EXISTING = 3
	OPEN_ALWAYS = 4
	TRUNCATE_EXISTING = 5
)

type FILE_SHARE_FLAGS uint32

const (
	FILE_SHARE_NONE = 0
	FILE_SHARE_DELETE = 4
	FILE_SHARE_READ = 1
	FILE_SHARE_WRITE = 2
)

type FILE_FLAGS_AND_ATTRIBUTES uint32

const (
	FILE_ATTRIBUTE_READONLY = 1
	FILE_ATTRIBUTE_HIDDEN = 2
	FILE_ATTRIBUTE_SYSTEM = 4
	FILE_ATTRIBUTE_DIRECTORY = 16
	FILE_ATTRIBUTE_ARCHIVE = 32
	FILE_ATTRIBUTE_DEVICE = 64
	FILE_ATTRIBUTE_NORMAL = 128
	FILE_ATTRIBUTE_TEMPORARY = 256
	FILE_ATTRIBUTE_SPARSE_FILE = 512
	FILE_ATTRIBUTE_REPARSE_POINT = 1024
	FILE_ATTRIBUTE_COMPRESSED = 2048
	FILE_ATTRIBUTE_OFFLINE = 4096
	FILE_ATTRIBUTE_NOT_CONTENT_INDEXED = 8192
	FILE_ATTRIBUTE_ENCRYPTED = 16384
	FILE_ATTRIBUTE_INTEGRITY_STREAM = 32768
	FILE_ATTRIBUTE_VIRTUAL = 65536
	FILE_ATTRIBUTE_NO_SCRUB_DATA = 131072
	FILE_ATTRIBUTE_EA = 262144
	FILE_ATTRIBUTE_PINNED = 524288
	FILE_ATTRIBUTE_UNPINNED = 1048576
	FILE_ATTRIBUTE_RECALL_ON_OPEN = 262144
	FILE_ATTRIBUTE_RECALL_ON_DATA_ACCESS = 4194304
)

type FILE_ACCESS_FLAGS uint32

const (
	FILE_READ_DATA = 1
	FILE_LIST_DIRECTORY = 1
	FILE_WRITE_DATA = 2
	FILE_ADD_FILE = 2
	FILE_APPEND_DATA = 4
	FILE_ADD_SUBDIRECTORY = 4
	FILE_CREATE_PIPE_INSTANCE = 4
	FILE_READ_EA = 8
	FILE_WRITE_EA = 16
	FILE_EXECUTE = 32
	FILE_TRAVERSE = 32
	FILE_DELETE_CHILD = 64
	FILE_READ_ATTRIBUTES = 128
	FILE_WRITE_ATTRIBUTES = 256
	READ_CONTROL = 131072
	SYNCHRONIZE = 1048576
	STANDARD_RIGHTS_REQUIRED = 983040
	STANDARD_RIGHTS_READ = 131072
	STANDARD_RIGHTS_WRITE = 131072
	STANDARD_RIGHTS_EXECUTE = 131072
	STANDARD_RIGHTS_ALL = 2031616
	SPECIFIC_RIGHTS_ALL = 65535
	FILE_ALL_ACCESS = 2032127
	FILE_GENERIC_READ = 1179785
	FILE_GENERIC_WRITE = 1179926
	FILE_GENERIC_EXECUTE = 1179808
)

type TRANSACTION_OUTCOME int32

const (
	TransactionOutcomeUndetermined = 1
	TransactionOutcomeCommitted = 2
	TransactionOutcomeAborted = 3
)

type STREAM_INFO_LEVELS int32

const (
	FindStreamInfoStandard = 0
	FindStreamInfoMaxInfoLevel = 1
)

type NtmsObjectsTypes int32

const (
	NTMS_UNKNOWN = 0
	NTMS_OBJECT = 1
	NTMS_CHANGER = 2
	NTMS_CHANGER_TYPE = 3
	NTMS_COMPUTER = 4
	NTMS_DRIVE = 5
	NTMS_DRIVE_TYPE = 6
	NTMS_IEDOOR = 7
	NTMS_IEPORT = 8
	NTMS_LIBRARY = 9
	NTMS_LIBREQUEST = 10
	NTMS_LOGICAL_MEDIA = 11
	NTMS_MEDIA_POOL = 12
	NTMS_MEDIA_TYPE = 13
	NTMS_PARTITION = 14
	NTMS_PHYSICAL_MEDIA = 15
	NTMS_STORAGESLOT = 16
	NTMS_OPREQUEST = 17
	NTMS_UI_DESTINATION = 18
	NTMS_NUMBER_OF_OBJECT_TYPES = 19
)

type NtmsAsyncStatus int32

const (
	NTMS_ASYNCSTATE_QUEUED = 0
	NTMS_ASYNCSTATE_WAIT_RESOURCE = 1
	NTMS_ASYNCSTATE_WAIT_OPERATOR = 2
	NTMS_ASYNCSTATE_INPROCESS = 3
	NTMS_ASYNCSTATE_COMPLETE = 4
)

type NtmsAsyncOperations int32

const (
	NTMS_ASYNCOP_MOUNT = 1
)

type NtmsSessionOptions int32

const (
	NTMS_SESSION_QUERYEXPEDITE = 1
)

type NtmsMountOptions int32

const (
	NTMS_MOUNT_READ = 1
	NTMS_MOUNT_WRITE = 2
	NTMS_MOUNT_ERROR_NOT_AVAILABLE = 4
	NTMS_MOUNT_ERROR_IF_UNAVAILABLE = 4
	NTMS_MOUNT_ERROR_OFFLINE = 8
	NTMS_MOUNT_ERROR_IF_OFFLINE = 8
	NTMS_MOUNT_SPECIFIC_DRIVE = 16
	NTMS_MOUNT_NOWAIT = 32
)

type NtmsDismountOptions int32

const (
	NTMS_DISMOUNT_DEFERRED = 1
	NTMS_DISMOUNT_IMMEDIATE = 2
)

type NtmsMountPriority int32

const (
	NTMS_PRIORITY_DEFAULT = 0
	NTMS_PRIORITY_HIGHEST = 15
	NTMS_PRIORITY_HIGH = 7
	NTMS_PRIORITY_NORMAL = 0
	NTMS_PRIORITY_LOW = -7
	NTMS_PRIORITY_LOWEST = -15
)

type NtmsAllocateOptions int32

const (
	NTMS_ALLOCATE_NEW = 1
	NTMS_ALLOCATE_NEXT = 2
	NTMS_ALLOCATE_ERROR_IF_UNAVAILABLE = 4
)

type NtmsCreateOptions int32

const (
	NTMS_OPEN_EXISTING = 1
	NTMS_CREATE_NEW = 2
	NTMS_OPEN_ALWAYS = 3
)

type NtmsDriveState int32

const (
	NTMS_DRIVESTATE_DISMOUNTED = 0
	NTMS_DRIVESTATE_MOUNTED = 1
	NTMS_DRIVESTATE_LOADED = 2
	NTMS_DRIVESTATE_UNLOADED = 5
	NTMS_DRIVESTATE_BEING_CLEANED = 6
	NTMS_DRIVESTATE_DISMOUNTABLE = 7
)

type NtmsLibraryType int32

const (
	NTMS_LIBRARYTYPE_UNKNOWN = 0
	NTMS_LIBRARYTYPE_OFFLINE = 1
	NTMS_LIBRARYTYPE_ONLINE = 2
	NTMS_LIBRARYTYPE_STANDALONE = 3
)

type NtmsLibraryFlags int32

const (
	NTMS_LIBRARYFLAG_FIXEDOFFLINE = 1
	NTMS_LIBRARYFLAG_CLEANERPRESENT = 2
	NTMS_LIBRARYFLAG_AUTODETECTCHANGE = 4
	NTMS_LIBRARYFLAG_IGNORECLEANERUSESREMAINING = 8
	NTMS_LIBRARYFLAG_RECOGNIZECLEANERBARCODE = 16
)

type NtmsInventoryMethod int32

const (
	NTMS_INVENTORY_NONE = 0
	NTMS_INVENTORY_FAST = 1
	NTMS_INVENTORY_OMID = 2
	NTMS_INVENTORY_DEFAULT = 3
	NTMS_INVENTORY_SLOT = 4
	NTMS_INVENTORY_STOP = 5
	NTMS_INVENTORY_MAX = 6
)

type NtmsSlotState int32

const (
	NTMS_SLOTSTATE_UNKNOWN = 0
	NTMS_SLOTSTATE_FULL = 1
	NTMS_SLOTSTATE_EMPTY = 2
	NTMS_SLOTSTATE_NOTPRESENT = 3
	NTMS_SLOTSTATE_NEEDSINVENTORY = 4
)

type NtmsDoorState int32

const (
	NTMS_DOORSTATE_UNKNOWN = 0
	NTMS_DOORSTATE_CLOSED = 1
	NTMS_DOORSTATE_OPEN = 2
)

type NtmsPortPosition int32

const (
	NTMS_PORTPOSITION_UNKNOWN = 0
	NTMS_PORTPOSITION_EXTENDED = 1
	NTMS_PORTPOSITION_RETRACTED = 2
)

type NtmsPortContent int32

const (
	NTMS_PORTCONTENT_UNKNOWN = 0
	NTMS_PORTCONTENT_FULL = 1
	NTMS_PORTCONTENT_EMPTY = 2
)

type NtmsBarCodeState int32

const (
	NTMS_BARCODESTATE_OK = 1
	NTMS_BARCODESTATE_UNREADABLE = 2
)

type NtmsMediaState int32

const (
	NTMS_MEDIASTATE_IDLE = 0
	NTMS_MEDIASTATE_INUSE = 1
	NTMS_MEDIASTATE_MOUNTED = 2
	NTMS_MEDIASTATE_LOADED = 3
	NTMS_MEDIASTATE_UNLOADED = 4
	NTMS_MEDIASTATE_OPERROR = 5
	NTMS_MEDIASTATE_OPREQ = 6
)

type NtmsPartitionState int32

const (
	NTMS_PARTSTATE_UNKNOWN = 0
	NTMS_PARTSTATE_UNPREPARED = 1
	NTMS_PARTSTATE_INCOMPATIBLE = 2
	NTMS_PARTSTATE_DECOMMISSIONED = 3
	NTMS_PARTSTATE_AVAILABLE = 4
	NTMS_PARTSTATE_ALLOCATED = 5
	NTMS_PARTSTATE_COMPLETE = 6
	NTMS_PARTSTATE_FOREIGN = 7
	NTMS_PARTSTATE_IMPORT = 8
	NTMS_PARTSTATE_RESERVED = 9
)

type NtmsPoolType int32

const (
	NTMS_POOLTYPE_UNKNOWN = 0
	NTMS_POOLTYPE_SCRATCH = 1
	NTMS_POOLTYPE_FOREIGN = 2
	NTMS_POOLTYPE_IMPORT = 3
	NTMS_POOLTYPE_APPLICATION = 1000
)

type NtmsAllocationPolicy int32

const (
	NTMS_ALLOCATE_FROMSCRATCH = 1
)

type NtmsDeallocationPolicy int32

const (
	NTMS_DEALLOCATE_TOSCRATCH = 1
)

type NtmsReadWriteCharacteristics int32

const (
	NTMS_MEDIARW_UNKNOWN = 0
	NTMS_MEDIARW_REWRITABLE = 1
	NTMS_MEDIARW_WRITEONCE = 2
	NTMS_MEDIARW_READONLY = 3
)

type NtmsLmOperation int32

const (
	NTMS_LM_REMOVE = 0
	NTMS_LM_DISABLECHANGER = 1
	NTMS_LM_DISABLELIBRARY = 1
	NTMS_LM_ENABLECHANGER = 2
	NTMS_LM_ENABLELIBRARY = 2
	NTMS_LM_DISABLEDRIVE = 3
	NTMS_LM_ENABLEDRIVE = 4
	NTMS_LM_DISABLEMEDIA = 5
	NTMS_LM_ENABLEMEDIA = 6
	NTMS_LM_UPDATEOMID = 7
	NTMS_LM_INVENTORY = 8
	NTMS_LM_DOORACCESS = 9
	NTMS_LM_EJECT = 10
	NTMS_LM_EJECTCLEANER = 11
	NTMS_LM_INJECT = 12
	NTMS_LM_INJECTCLEANER = 13
	NTMS_LM_PROCESSOMID = 14
	NTMS_LM_CLEANDRIVE = 15
	NTMS_LM_DISMOUNT = 16
	NTMS_LM_MOUNT = 17
	NTMS_LM_WRITESCRATCH = 18
	NTMS_LM_CLASSIFY = 19
	NTMS_LM_RESERVECLEANER = 20
	NTMS_LM_RELEASECLEANER = 21
	NTMS_LM_MAXWORKITEM = 22
)

type NtmsLmState int32

const (
	NTMS_LM_QUEUED = 0
	NTMS_LM_INPROCESS = 1
	NTMS_LM_PASSED = 2
	NTMS_LM_FAILED = 3
	NTMS_LM_INVALID = 4
	NTMS_LM_WAITING = 5
	NTMS_LM_DEFERRED = 6
	NTMS_LM_DEFFERED = 6
	NTMS_LM_CANCELLED = 7
	NTMS_LM_STOPPED = 8
)

type NtmsOpreqCommand int32

const (
	NTMS_OPREQ_UNKNOWN = 0
	NTMS_OPREQ_NEWMEDIA = 1
	NTMS_OPREQ_CLEANER = 2
	NTMS_OPREQ_DEVICESERVICE = 3
	NTMS_OPREQ_MOVEMEDIA = 4
	NTMS_OPREQ_MESSAGE = 5
)

type NtmsOpreqState int32

const (
	NTMS_OPSTATE_UNKNOWN = 0
	NTMS_OPSTATE_SUBMITTED = 1
	NTMS_OPSTATE_ACTIVE = 2
	NTMS_OPSTATE_INPROGRESS = 3
	NTMS_OPSTATE_REFUSED = 4
	NTMS_OPSTATE_COMPLETE = 5
)

type NtmsLibRequestFlags int32

const (
	NTMS_LIBREQFLAGS_NOAUTOPURGE = 1
	NTMS_LIBREQFLAGS_NOFAILEDPURGE = 2
)

type NtmsOpRequestFlags int32

const (
	NTMS_OPREQFLAGS_NOAUTOPURGE = 1
	NTMS_OPREQFLAGS_NOFAILEDPURGE = 2
	NTMS_OPREQFLAGS_NOALERTS = 16
	NTMS_OPREQFLAGS_NOTRAYICON = 32
)

type NtmsMediaPoolPolicy int32

const (
	NTMS_POOLPOLICY_PURGEOFFLINESCRATCH = 1
	NTMS_POOLPOLICY_KEEPOFFLINEIMPORT = 2
)

type NtmsOperationalState int32

const (
	NTMS_READY = 0
	NTMS_INITIALIZING = 10
	NTMS_NEEDS_SERVICE = 20
	NTMS_NOT_PRESENT = 21
)

type NtmsCreateNtmsMediaOptions int32

const (
	NTMS_ERROR_ON_DUPLICATE = 1
)

type NtmsEnumerateOption int32

const (
	NTMS_ENUM_DEFAULT = 0
	NTMS_ENUM_ROOTPOOL = 1
)

type NtmsEjectOperation int32

const (
	NTMS_EJECT_START = 0
	NTMS_EJECT_STOP = 1
	NTMS_EJECT_QUEUE = 2
	NTMS_EJECT_FORCE = 3
	NTMS_EJECT_IMMEDIATE = 4
	NTMS_EJECT_ASK_USER = 5
)

type NtmsInjectOperation int32

const (
	NTMS_INJECT_START = 0
	NTMS_INJECT_STOP = 1
	NTMS_INJECT_RETRACT = 2
	NTMS_INJECT_STARTMANY = 3
)

type NtmsDriveType int32

const (
	NTMS_UNKNOWN_DRIVE = 0
)

type NtmsAccessMask int32

const (
	NTMS_USE_ACCESS = 1
	NTMS_MODIFY_ACCESS = 2
	NTMS_CONTROL_ACCESS = 4
)

type NtmsUITypes int32

const (
	NTMS_UITYPE_INVALID = 0
	NTMS_UITYPE_INFO = 1
	NTMS_UITYPE_REQ = 2
	NTMS_UITYPE_ERR = 3
	NTMS_UITYPE_MAX = 4
)

type NtmsUIOperations int32

const (
	NTMS_UIDEST_ADD = 1
	NTMS_UIDEST_DELETE = 2
	NTMS_UIDEST_DELETEALL = 3
	NTMS_UIOPERATION_MAX = 4
)

type NtmsNotificationOperations int32

const (
	NTMS_OBJ_UPDATE = 1
	NTMS_OBJ_INSERT = 2
	NTMS_OBJ_DELETE = 3
	NTMS_EVENT_SIGNAL = 4
	NTMS_EVENT_COMPLETE = 5
)

type CLS_CONTEXT_MODE int32

const (
	ClsContextNone = 0
	ClsContextUndoNext = 1
	ClsContextPrevious = 2
	ClsContextForward = 3
)

type CLFS_CONTEXT_MODE int32

const (
	ClfsContextNone = 0
	ClfsContextUndoNext = 1
	ClfsContextPrevious = 2
	ClfsContextForward = 3
)

type CLS_LOG_INFORMATION_CLASS int32

const (
	ClfsLogBasicInformation = 0
	ClfsLogBasicInformationPhysical = 1
	ClfsLogPhysicalNameInformation = 2
	ClfsLogStreamIdentifierInformation = 3
	ClfsLogSystemMarkingInformation = 4
	ClfsLogPhysicalLsnInformation = 5
)

type CLS_IOSTATS_CLASS int32

const (
	ClsIoStatsDefault = 0
	ClsIoStatsMax = 65535
)

type CLFS_IOSTATS_CLASS int32

const (
	ClfsIoStatsDefault = 0
	ClfsIoStatsMax = 65535
)

type CLFS_LOG_ARCHIVE_MODE int32

const (
	ClfsLogArchiveEnabled = 1
	ClfsLogArchiveDisabled = 2
)

type CLFS_MGMT_POLICY_TYPE int32

const (
	ClfsMgmtPolicyMaximumSize = 0
	ClfsMgmtPolicyMinimumSize = 1
	ClfsMgmtPolicyNewContainerSize = 2
	ClfsMgmtPolicyGrowthRate = 3
	ClfsMgmtPolicyLogTail = 4
	ClfsMgmtPolicyAutoShrink = 5
	ClfsMgmtPolicyAutoGrow = 6
	ClfsMgmtPolicyNewContainerPrefix = 7
	ClfsMgmtPolicyNewContainerSuffix = 8
	ClfsMgmtPolicyNewContainerExtension = 9
	ClfsMgmtPolicyInvalid = 10
)

type CLFS_MGMT_NOTIFICATION_TYPE int32

const (
	ClfsMgmtAdvanceTailNotification = 0
	ClfsMgmtLogFullHandlerNotification = 1
	ClfsMgmtLogUnpinnedNotification = 2
	ClfsMgmtLogWriteNotification = 3
)

type SERVER_CERTIFICATE_TYPE int32

const (
	QUIC = 0
)

type FINDEX_INFO_LEVELS int32

const (
	FindExInfoStandard = 0
	FindExInfoBasic = 1
	FindExInfoMaxInfoLevel = 2
)

type FINDEX_SEARCH_OPS int32

const (
	FindExSearchNameMatch = 0
	FindExSearchLimitToDirectories = 1
	FindExSearchLimitToDevices = 2
	FindExSearchMaxSearchOp = 3
)

type READ_DIRECTORY_NOTIFY_INFORMATION_CLASS int32

const (
	ReadDirectoryNotifyInformation = 1
	ReadDirectoryNotifyExtendedInformation = 2
)

type GET_FILEEX_INFO_LEVELS int32

const (
	GetFileExInfoStandard = 0
	GetFileExMaxInfoLevel = 1
)

type FILE_INFO_BY_HANDLE_CLASS int32

const (
	FileBasicInfo = 0
	FileStandardInfo = 1
	FileNameInfo = 2
	FileRenameInfo = 3
	FileDispositionInfo = 4
	FileAllocationInfo = 5
	FileEndOfFileInfo = 6
	FileStreamInfo = 7
	FileCompressionInfo = 8
	FileAttributeTagInfo = 9
	FileIdBothDirectoryInfo = 10
	FileIdBothDirectoryRestartInfo = 11
	FileIoPriorityHintInfo = 12
	FileRemoteProtocolInfo = 13
	FileFullDirectoryInfo = 14
	FileFullDirectoryRestartInfo = 15
	FileStorageInfo = 16
	FileAlignmentInfo = 17
	FileIdInfo = 18
	FileIdExtdDirectoryInfo = 19
	FileIdExtdDirectoryRestartInfo = 20
	FileDispositionInfoEx = 21
	FileRenameInfoEx = 22
	FileCaseSensitiveInfo = 23
	FileNormalizedNameInfo = 24
	MaximumFileInfoByHandleClass = 25
)

type STORAGE_QUERY_TYPE int32

const (
	PropertyStandardQuery = 0
	PropertyExistsQuery = 1
	PropertyMaskQuery = 2
	PropertyQueryMaxDefined = 3
)

type STORAGE_PROPERTY_ID int32

const (
	StorageDeviceProperty = 0
	StorageAdapterProperty = 1
	StorageDeviceIdProperty = 2
	StorageDeviceUniqueIdProperty = 3
	StorageDeviceWriteCacheProperty = 4
	StorageMiniportProperty = 5
	StorageAccessAlignmentProperty = 6
	StorageDeviceSeekPenaltyProperty = 7
	StorageDeviceTrimProperty = 8
	StorageDeviceWriteAggregationProperty = 9
	StorageDeviceDeviceTelemetryProperty = 10
	StorageDeviceLBProvisioningProperty = 11
	StorageDevicePowerProperty = 12
	StorageDeviceCopyOffloadProperty = 13
	StorageDeviceResiliencyProperty = 14
	StorageDeviceMediumProductType = 15
	StorageAdapterRpmbProperty = 16
	StorageAdapterCryptoProperty = 17
	StorageDeviceIoCapabilityProperty = 48
	StorageAdapterProtocolSpecificProperty = 49
	StorageDeviceProtocolSpecificProperty = 50
	StorageAdapterTemperatureProperty = 51
	StorageDeviceTemperatureProperty = 52
	StorageAdapterPhysicalTopologyProperty = 53
	StorageDevicePhysicalTopologyProperty = 54
	StorageDeviceAttributesProperty = 55
	StorageDeviceManagementStatus = 56
	StorageAdapterSerialNumberProperty = 57
	StorageDeviceLocationProperty = 58
	StorageDeviceNumaProperty = 59
	StorageDeviceZonedDeviceProperty = 60
	StorageDeviceUnsafeShutdownCount = 61
	StorageDeviceEnduranceProperty = 62
)

type STORAGE_PORT_CODE_SET int32

const (
	StoragePortCodeSetReserved = 0
	StoragePortCodeSetStorport = 1
	StoragePortCodeSetSCSIport = 2
	StoragePortCodeSetSpaceport = 3
	StoragePortCodeSetATAport = 4
	StoragePortCodeSetUSBport = 5
	StoragePortCodeSetSBP2port = 6
	StoragePortCodeSetSDport = 7
)

type STORAGE_PROTOCOL_TYPE int32

const (
	ProtocolTypeUnknown = 0
	ProtocolTypeScsi = 1
	ProtocolTypeAta = 2
	ProtocolTypeNvme = 3
	ProtocolTypeSd = 4
	ProtocolTypeUfs = 5
	ProtocolTypeProprietary = 126
	ProtocolTypeMaxReserved = 127
)

type STORAGE_PROTOCOL_NVME_DATA_TYPE int32

const (
	NVMeDataTypeUnknown = 0
	NVMeDataTypeIdentify = 1
	NVMeDataTypeLogPage = 2
	NVMeDataTypeFeature = 3
)

type STORAGE_PROTOCOL_ATA_DATA_TYPE int32

const (
	AtaDataTypeUnknown = 0
	AtaDataTypeIdentify = 1
	AtaDataTypeLogPage = 2
)

type STORAGE_DEVICE_FORM_FACTOR int32

const (
	FormFactorUnknown = 0
	FormFactor3_5 = 1
	FormFactor2_5 = 2
	FormFactor1_8 = 3
	FormFactor1_8Less = 4
	FormFactorEmbedded = 5
	FormFactorMemoryCard = 6
	FormFactormSata = 7
	FormFactorM_2 = 8
	FormFactorPCIeBoard = 9
	FormFactorDimm = 10
)

type STORAGE_COMPONENT_HEALTH_STATUS int32

const (
	HealthStatusUnknown = 0
	HealthStatusNormal = 1
	HealthStatusThrottled = 2
	HealthStatusWarning = 3
	HealthStatusDisabled = 4
	HealthStatusFailed = 5
)

type WRITE_CACHE_TYPE int32

const (
	WriteCacheTypeUnknown = 0
	WriteCacheTypeNone = 1
	WriteCacheTypeWriteBack = 2
	WriteCacheTypeWriteThrough = 3
)

type WRITE_CACHE_ENABLE int32

const (
	WriteCacheEnableUnknown = 0
	WriteCacheDisabled = 1
	WriteCacheEnabled = 2
)

type WRITE_CACHE_CHANGE int32

const (
	WriteCacheChangeUnknown = 0
	WriteCacheNotChangeable = 1
	WriteCacheChangeable = 2
)

type WRITE_THROUGH int32

const (
	WriteThroughUnknown = 0
	WriteThroughNotSupported = 1
	WriteThroughSupported = 2
)

type STORAGE_DEVICE_POWER_CAP_UNITS int32

const (
	StorageDevicePowerCapUnitsPercent = 0
	StorageDevicePowerCapUnitsMilliwatts = 1
)

type MEDIA_TYPE int32

const (
	Unknown = 0
	F5_1Pt2_512 = 1
	F3_1Pt44_512 = 2
	F3_2Pt88_512 = 3
	F3_20Pt8_512 = 4
	F3_720_512 = 5
	F5_360_512 = 6
	F5_320_512 = 7
	F5_320_1024 = 8
	F5_180_512 = 9
	F5_160_512 = 10
	RemovableMedia = 11
	FixedMedia = 12
	F3_120M_512 = 13
	F3_640_512 = 14
	F5_640_512 = 15
	F5_720_512 = 16
	F3_1Pt2_512 = 17
	F3_1Pt23_1024 = 18
	F5_1Pt23_1024 = 19
	F3_128Mb_512 = 20
	F3_230Mb_512 = 21
	F8_256_128 = 22
	F3_200Mb_512 = 23
	F3_240M_512 = 24
	F3_32M_512 = 25
)

type PARTITION_STYLE int32

const (
	PARTITION_STYLE_MBR = 0
	PARTITION_STYLE_GPT = 1
	PARTITION_STYLE_RAW = 2
)

type CSV_CONTROL_OP int32

const (
	CsvControlStartRedirectFile = 2
	CsvControlStopRedirectFile = 3
	CsvControlQueryRedirectState = 4
	CsvControlQueryFileRevision = 6
	CsvControlQueryMdsPath = 8
	CsvControlQueryFileRevisionFileId128 = 9
	CsvControlQueryVolumeRedirectState = 10
	CsvControlEnableUSNRangeModificationTracking = 13
	CsvControlMarkHandleLocalVolumeMount = 14
	CsvControlUnmarkHandleLocalVolumeMount = 15
	CsvControlGetCsvFsMdsPathV2 = 18
	CsvControlDisableCaching = 19
	CsvControlEnableCaching = 20
	CsvControlStartForceDFO = 21
	CsvControlStopForceDFO = 22
)

type FILE_STORAGE_TIER_MEDIA_TYPE int32

const (
	FileStorageTierMediaTypeUnspecified = 0
	FileStorageTierMediaTypeDisk = 1
	FileStorageTierMediaTypeSsd = 2
	FileStorageTierMediaTypeScm = 4
	FileStorageTierMediaTypeMax = 5
)

type COPYFILE2_MESSAGE_TYPE int32

const (
	COPYFILE2_CALLBACK_NONE = 0
	COPYFILE2_CALLBACK_CHUNK_STARTED = 1
	COPYFILE2_CALLBACK_CHUNK_FINISHED = 2
	COPYFILE2_CALLBACK_STREAM_STARTED = 3
	COPYFILE2_CALLBACK_STREAM_FINISHED = 4
	COPYFILE2_CALLBACK_POLL_CONTINUE = 5
	COPYFILE2_CALLBACK_ERROR = 6
	COPYFILE2_CALLBACK_MAX = 7
)

type COPYFILE2_MESSAGE_ACTION int32

const (
	COPYFILE2_PROGRESS_CONTINUE = 0
	COPYFILE2_PROGRESS_CANCEL = 1
	COPYFILE2_PROGRESS_STOP = 2
	COPYFILE2_PROGRESS_QUIET = 3
	COPYFILE2_PROGRESS_PAUSE = 4
)

type COPYFILE2_COPY_PHASE int32

const (
	COPYFILE2_PHASE_NONE = 0
	COPYFILE2_PHASE_PREPARE_SOURCE = 1
	COPYFILE2_PHASE_PREPARE_DEST = 2
	COPYFILE2_PHASE_READ_SOURCE = 3
	COPYFILE2_PHASE_WRITE_DESTINATION = 4
	COPYFILE2_PHASE_SERVER_COPY = 5
	COPYFILE2_PHASE_NAMEGRAFT_COPY = 6
	COPYFILE2_PHASE_MAX = 7
)

type PRIORITY_HINT int32

const (
	IoPriorityHintVeryLow = 0
	IoPriorityHintLow = 1
	IoPriorityHintNormal = 2
	MaximumIoPriorityHintType = 3
)

type FILE_ID_TYPE int32

const (
	FileIdType = 0
	ObjectIdType = 1
	ExtendedFileIdType = 2
	MaximumFileIdType = 3
)
