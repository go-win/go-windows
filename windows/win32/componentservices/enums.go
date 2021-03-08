// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package componentservices implements the Windows.Win32.ComponentServices namespace.
package componentservices

type COMAdminInUse int32

const (
	COMAdminNotInUse = 0
	COMAdminInUseByCatalog = 1
	COMAdminInUseByRegistryUnknown = 2
	COMAdminInUseByRegistryProxyStub = 3
	COMAdminInUseByRegistryTypeLib = 4
	COMAdminInUseByRegistryClsid = 5
)

type COMAdminComponentType int32

const (
	COMAdmin32BitComponent = 1
	COMAdmin64BitComponent = 2
)

type COMAdminApplicationInstallOptions int32

const (
	COMAdminInstallNoUsers = 0
	COMAdminInstallUsers = 1
	COMAdminInstallForceOverwriteOfFiles = 2
)

type COMAdminApplicationExportOptions int32

const (
	COMAdminExportNoUsers = 0
	COMAdminExportUsers = 1
	COMAdminExportApplicationProxy = 2
	COMAdminExportForceOverwriteOfFiles = 4
	COMAdminExportIn10Format = 16
)

type COMAdminThreadingModels int32

const (
	COMAdminThreadingModelApartment = 0
	COMAdminThreadingModelFree = 1
	COMAdminThreadingModelMain = 2
	COMAdminThreadingModelBoth = 3
	COMAdminThreadingModelNeutral = 4
	COMAdminThreadingModelNotSpecified = 5
)

type COMAdminTransactionOptions int32

const (
	COMAdminTransactionIgnored = 0
	COMAdminTransactionNone = 1
	COMAdminTransactionSupported = 2
	COMAdminTransactionRequired = 3
	COMAdminTransactionRequiresNew = 4
)

type COMAdminTxIsolationLevelOptions int32

const (
	COMAdminTxIsolationLevelAny = 0
	COMAdminTxIsolationLevelReadUnCommitted = 1
	COMAdminTxIsolationLevelReadCommitted = 2
	COMAdminTxIsolationLevelRepeatableRead = 3
	COMAdminTxIsolationLevelSerializable = 4
)

type COMAdminSynchronizationOptions int32

const (
	COMAdminSynchronizationIgnored = 0
	COMAdminSynchronizationNone = 1
	COMAdminSynchronizationSupported = 2
	COMAdminSynchronizationRequired = 3
	COMAdminSynchronizationRequiresNew = 4
)

type COMAdminActivationOptions int32

const (
	COMAdminActivationInproc = 0
	COMAdminActivationLocal = 1
)

type COMAdminAccessChecksLevelOptions int32

const (
	COMAdminAccessChecksApplicationLevel = 0
	COMAdminAccessChecksApplicationComponentLevel = 1
)

type COMAdminAuthenticationLevelOptions int32

const (
	COMAdminAuthenticationDefault = 0
	COMAdminAuthenticationNone = 1
	COMAdminAuthenticationConnect = 2
	COMAdminAuthenticationCall = 3
	COMAdminAuthenticationPacket = 4
	COMAdminAuthenticationIntegrity = 5
	COMAdminAuthenticationPrivacy = 6
)

type COMAdminImpersonationLevelOptions int32

const (
	COMAdminImpersonationAnonymous = 1
	COMAdminImpersonationIdentify = 2
	COMAdminImpersonationImpersonate = 3
	COMAdminImpersonationDelegate = 4
)

type COMAdminAuthenticationCapabilitiesOptions int32

const (
	COMAdminAuthenticationCapabilitiesNone = 0
	COMAdminAuthenticationCapabilitiesSecureReference = 2
	COMAdminAuthenticationCapabilitiesStaticCloaking = 32
	COMAdminAuthenticationCapabilitiesDynamicCloaking = 64
)

type COMAdminOS int32

const (
	COMAdminOSNotInitialized = 0
	COMAdminOSWindows3_1 = 1
	COMAdminOSWindows9x = 2
	COMAdminOSWindows2000 = 3
	COMAdminOSWindows2000AdvancedServer = 4
	COMAdminOSWindows2000Unknown = 5
	COMAdminOSUnknown = 6
	COMAdminOSWindowsXPPersonal = 11
	COMAdminOSWindowsXPProfessional = 12
	COMAdminOSWindowsNETStandardServer = 13
	COMAdminOSWindowsNETEnterpriseServer = 14
	COMAdminOSWindowsNETDatacenterServer = 15
	COMAdminOSWindowsNETWebServer = 16
	COMAdminOSWindowsLonghornPersonal = 17
	COMAdminOSWindowsLonghornProfessional = 18
	COMAdminOSWindowsLonghornStandardServer = 19
	COMAdminOSWindowsLonghornEnterpriseServer = 20
	COMAdminOSWindowsLonghornDatacenterServer = 21
	COMAdminOSWindowsLonghornWebServer = 22
	COMAdminOSWindows7Personal = 23
	COMAdminOSWindows7Professional = 24
	COMAdminOSWindows7StandardServer = 25
	COMAdminOSWindows7EnterpriseServer = 26
	COMAdminOSWindows7DatacenterServer = 27
	COMAdminOSWindows7WebServer = 28
	COMAdminOSWindows8Personal = 29
	COMAdminOSWindows8Professional = 30
	COMAdminOSWindows8StandardServer = 31
	COMAdminOSWindows8EnterpriseServer = 32
	COMAdminOSWindows8DatacenterServer = 33
	COMAdminOSWindows8WebServer = 34
	COMAdminOSWindowsBluePersonal = 35
	COMAdminOSWindowsBlueProfessional = 36
	COMAdminOSWindowsBlueStandardServer = 37
	COMAdminOSWindowsBlueEnterpriseServer = 38
	COMAdminOSWindowsBlueDatacenterServer = 39
	COMAdminOSWindowsBlueWebServer = 40
)

type COMAdminServiceOptions int32

const (
	COMAdminServiceLoadBalanceRouter = 1
)

type COMAdminServiceStatusOptions int32

const (
	COMAdminServiceStopped = 0
	COMAdminServiceStartPending = 1
	COMAdminServiceStopPending = 2
	COMAdminServiceRunning = 3
	COMAdminServiceContinuePending = 4
	COMAdminServicePausePending = 5
	COMAdminServicePaused = 6
	COMAdminServiceUnknownState = 7
)

type COMAdminQCMessageAuthenticateOptions int32

const (
	COMAdminQCMessageAuthenticateSecureApps = 0
	COMAdminQCMessageAuthenticateOff = 1
	COMAdminQCMessageAuthenticateOn = 2
)

type COMAdminFileFlags int32

const (
	COMAdminFileFlagLoadable = 1
	COMAdminFileFlagCOM = 2
	COMAdminFileFlagContainsPS = 4
	COMAdminFileFlagContainsComp = 8
	COMAdminFileFlagContainsTLB = 16
	COMAdminFileFlagSelfReg = 32
	COMAdminFileFlagSelfUnReg = 64
	COMAdminFileFlagUnloadableDLL = 128
	COMAdminFileFlagDoesNotExist = 256
	COMAdminFileFlagAlreadyInstalled = 512
	COMAdminFileFlagBadTLB = 1024
	COMAdminFileFlagGetClassObjFailed = 2048
	COMAdminFileFlagClassNotAvailable = 4096
	COMAdminFileFlagRegistrar = 8192
	COMAdminFileFlagNoRegistrar = 16384
	COMAdminFileFlagDLLRegsvrFailed = 32768
	COMAdminFileFlagRegTLBFailed = 65536
	COMAdminFileFlagRegistrarFailed = 131072
	COMAdminFileFlagError = 262144
)

type COMAdminComponentFlags int32

const (
	COMAdminCompFlagTypeInfoFound = 1
	COMAdminCompFlagCOMPlusPropertiesFound = 2
	COMAdminCompFlagProxyFound = 4
	COMAdminCompFlagInterfacesFound = 8
	COMAdminCompFlagAlreadyInstalled = 16
	COMAdminCompFlagNotInApplication = 32
)

type COMAdminErrorCodes int32

const (
	COMAdminErrObjectErrors = -2146368511
	COMAdminErrObjectInvalid = -2146368510
	COMAdminErrKeyMissing = -2146368509
	COMAdminErrAlreadyInstalled = -2146368508
	COMAdminErrAppFileWriteFail = -2146368505
	COMAdminErrAppFileReadFail = -2146368504
	COMAdminErrAppFileVersion = -2146368503
	COMAdminErrBadPath = -2146368502
	COMAdminErrApplicationExists = -2146368501
	COMAdminErrRoleExists = -2146368500
	COMAdminErrCantCopyFile = -2146368499
	COMAdminErrNoUser = -2146368497
	COMAdminErrInvalidUserids = -2146368496
	COMAdminErrNoRegistryCLSID = -2146368495
	COMAdminErrBadRegistryProgID = -2146368494
	COMAdminErrAuthenticationLevel = -2146368493
	COMAdminErrUserPasswdNotValid = -2146368492
	COMAdminErrCLSIDOrIIDMismatch = -2146368488
	COMAdminErrRemoteInterface = -2146368487
	COMAdminErrDllRegisterServer = -2146368486
	COMAdminErrNoServerShare = -2146368485
	COMAdminErrDllLoadFailed = -2146368483
	COMAdminErrBadRegistryLibID = -2146368482
	COMAdminErrAppDirNotFound = -2146368481
	COMAdminErrRegistrarFailed = -2146368477
	COMAdminErrCompFileDoesNotExist = -2146368476
	COMAdminErrCompFileLoadDLLFail = -2146368475
	COMAdminErrCompFileGetClassObj = -2146368474
	COMAdminErrCompFileClassNotAvail = -2146368473
	COMAdminErrCompFileBadTLB = -2146368472
	COMAdminErrCompFileNotInstallable = -2146368471
	COMAdminErrNotChangeable = -2146368470
	COMAdminErrNotDeletable = -2146368469
	COMAdminErrSession = -2146368468
	COMAdminErrCompMoveLocked = -2146368467
	COMAdminErrCompMoveBadDest = -2146368466
	COMAdminErrRegisterTLB = -2146368464
	COMAdminErrSystemApp = -2146368461
	COMAdminErrCompFileNoRegistrar = -2146368460
	COMAdminErrCoReqCompInstalled = -2146368459
	COMAdminErrServiceNotInstalled = -2146368458
	COMAdminErrPropertySaveFailed = -2146368457
	COMAdminErrObjectExists = -2146368456
	COMAdminErrComponentExists = -2146368455
	COMAdminErrRegFileCorrupt = -2146368453
	COMAdminErrPropertyOverflow = -2146368452
	COMAdminErrNotInRegistry = -2146368450
	COMAdminErrObjectNotPoolable = -2146368449
	COMAdminErrApplidMatchesClsid = -2146368442
	COMAdminErrRoleDoesNotExist = -2146368441
	COMAdminErrStartAppNeedsComponents = -2146368440
	COMAdminErrRequiresDifferentPlatform = -2146368439
	COMAdminErrQueuingServiceNotAvailable = -2146367998
	COMAdminErrObjectParentMissing = -2146367480
	COMAdminErrObjectDoesNotExist = -2146367479
	COMAdminErrCanNotExportAppProxy = -2146368438
	COMAdminErrCanNotStartApp = -2146368437
	COMAdminErrCanNotExportSystemApp = -2146368436
	COMAdminErrCanNotSubscribeToComponent = -2146368435
	COMAdminErrAppNotRunning = -2146367478
	COMAdminErrEventClassCannotBeSubscriber = -2146368434
	COMAdminErrLibAppProxyIncompatible = -2146368433
	COMAdminErrBasePartitionOnly = -2146368432
	COMAdminErrDuplicatePartitionName = -2146368425
	COMAdminErrPartitionInUse = -2146368423
	COMAdminErrImportedComponentsNotAllowed = -2146368421
	COMAdminErrRegdbNotInitialized = -2146368398
	COMAdminErrRegdbNotOpen = -2146368397
	COMAdminErrRegdbSystemErr = -2146368396
	COMAdminErrRegdbAlreadyRunning = -2146368395
	COMAdminErrMigVersionNotSupported = -2146368384
	COMAdminErrMigSchemaNotFound = -2146368383
	COMAdminErrCatBitnessMismatch = -2146368382
	COMAdminErrCatUnacceptableBitness = -2146368381
	COMAdminErrCatWrongAppBitnessBitness = -2146368380
	COMAdminErrCatPauseResumeNotSupported = -2146368379
	COMAdminErrCatServerFault = -2146368378
	COMAdminErrCantRecycleLibraryApps = -2146367473
	COMAdminErrCantRecycleServiceApps = -2146367471
	COMAdminErrProcessAlreadyRecycled = -2146367470
	COMAdminErrPausedProcessMayNotBeRecycled = -2146367469
	COMAdminErrInvalidPartition = -2146367477
	COMAdminErrPartitionMsiOnly = -2146367463
	COMAdminErrStartAppDisabled = -2146368431
	COMAdminErrCompMoveSource = -2146367460
	COMAdminErrCompMoveDest = -2146367459
	COMAdminErrCompMovePrivate = -2146367458
	COMAdminErrCannotCopyEventClass = -2146367456
)

type TX_MISC_CONSTANTS int32

const (
	MAX_TRAN_DESC = 40
)

type ISOLATIONLEVEL int32

const (
	ISOLATIONLEVEL_UNSPECIFIED = -1
	ISOLATIONLEVEL_CHAOS = 16
	ISOLATIONLEVEL_READUNCOMMITTED = 256
	ISOLATIONLEVEL_BROWSE = 256
	ISOLATIONLEVEL_CURSORSTABILITY = 4096
	ISOLATIONLEVEL_READCOMMITTED = 4096
	ISOLATIONLEVEL_REPEATABLEREAD = 65536
	ISOLATIONLEVEL_SERIALIZABLE = 1048576
	ISOLATIONLEVEL_ISOLATED = 1048576
)

type ISOFLAG int32

const (
	ISOFLAG_RETAIN_COMMIT_DC = 1
	ISOFLAG_RETAIN_COMMIT = 2
	ISOFLAG_RETAIN_COMMIT_NO = 3
	ISOFLAG_RETAIN_ABORT_DC = 4
	ISOFLAG_RETAIN_ABORT = 8
	ISOFLAG_RETAIN_ABORT_NO = 12
	ISOFLAG_RETAIN_DONTCARE = 5
	ISOFLAG_RETAIN_BOTH = 10
	ISOFLAG_RETAIN_NONE = 15
	ISOFLAG_OPTIMISTIC = 16
	ISOFLAG_READONLY = 32
)

type XACTTC int32

const (
	XACTTC_NONE = 0
	XACTTC_SYNC_PHASEONE = 1
	XACTTC_SYNC_PHASETWO = 2
	XACTTC_SYNC = 2
	XACTTC_ASYNC_PHASEONE = 4
	XACTTC_ASYNC = 4
)

type XACTRM int32

const (
	XACTRM_OPTIMISTICLASTWINS = 1
	XACTRM_NOREADONLYPREPARES = 2
)

type XACTCONST int32

const (
	XACTCONST_TIMEOUTINFINITE = 0
)

type XACTHEURISTIC int32

const (
	XACTHEURISTIC_ABORT = 1
	XACTHEURISTIC_COMMIT = 2
	XACTHEURISTIC_DAMAGE = 3
	XACTHEURISTIC_DANGER = 4
)

type XACTSTAT int32

const (
	XACTSTAT_NONE = 0
	XACTSTAT_OPENNORMAL = 1
	XACTSTAT_OPENREFUSED = 2
	XACTSTAT_PREPARING = 4
	XACTSTAT_PREPARED = 8
	XACTSTAT_PREPARERETAINING = 16
	XACTSTAT_PREPARERETAINED = 32
	XACTSTAT_COMMITTING = 64
	XACTSTAT_COMMITRETAINING = 128
	XACTSTAT_ABORTING = 256
	XACTSTAT_ABORTED = 512
	XACTSTAT_COMMITTED = 1024
	XACTSTAT_HEURISTIC_ABORT = 2048
	XACTSTAT_HEURISTIC_COMMIT = 4096
	XACTSTAT_HEURISTIC_DAMAGE = 8192
	XACTSTAT_HEURISTIC_DANGER = 16384
	XACTSTAT_FORCED_ABORT = 32768
	XACTSTAT_FORCED_COMMIT = 65536
	XACTSTAT_INDOUBT = 131072
	XACTSTAT_CLOSED = 262144
	XACTSTAT_OPEN = 3
	XACTSTAT_NOTPREPARED = 524227
	XACTSTAT_ALL = 524287
)

type AUTHENTICATION_LEVEL int32

const (
	NO_AUTHENTICATION_REQUIRED = 0
	INCOMING_AUTHENTICATION_REQUIRED = 1
	MUTUAL_AUTHENTICATION_REQUIRED = 2
)

type XACT_DTC_CONSTANTS int32

const (
	XACT_E_CONNECTION_REQUEST_DENIED = -2147168000
	XACT_E_TOOMANY_ENLISTMENTS = -2147167999
	XACT_E_DUPLICATE_GUID = -2147167998
	XACT_E_NOTSINGLEPHASE = -2147167997
	XACT_E_RECOVERYALREADYDONE = -2147167996
	XACT_E_PROTOCOL = -2147167995
	XACT_E_RM_FAILURE = -2147167994
	XACT_E_RECOVERY_FAILED = -2147167993
	XACT_E_LU_NOT_FOUND = -2147167992
	XACT_E_DUPLICATE_LU = -2147167991
	XACT_E_LU_NOT_CONNECTED = -2147167990
	XACT_E_DUPLICATE_TRANSID = -2147167989
	XACT_E_LU_BUSY = -2147167988
	XACT_E_LU_NO_RECOVERY_PROCESS = -2147167987
	XACT_E_LU_DOWN = -2147167986
	XACT_E_LU_RECOVERING = -2147167985
	XACT_E_LU_RECOVERY_MISMATCH = -2147167984
	XACT_E_RM_UNAVAILABLE = -2147167983
	XACT_E_LRMRECOVERYALREADYDONE = -2147167982
	XACT_E_NOLASTRESOURCEINTERFACE = -2147167981
	XACT_S_NONOTIFY = 315648
	XACT_OK_NONOTIFY = 315649
	dwUSER_MS_SQLSERVER = 65535
)

type _DtcLu_LocalRecovery_Work int32

const (
	DTCINITIATEDRECOVERYWORK_CHECKLUSTATUS = 1
	DTCINITIATEDRECOVERYWORK_TRANS = 2
	DTCINITIATEDRECOVERYWORK_TMDOWN = 3
)

type _DtcLu_Xln int32

const (
	DTCLUXLN_COLD = 1
	DTCLUXLN_WARM = 2
)

type _DtcLu_Xln_Confirmation int32

const (
	DTCLUXLNCONFIRMATION_CONFIRM = 1
	DTCLUXLNCONFIRMATION_LOGNAMEMISMATCH = 2
	DTCLUXLNCONFIRMATION_COLDWARMMISMATCH = 3
	DTCLUXLNCONFIRMATION_OBSOLETE = 4
)

type _DtcLu_Xln_Response int32

const (
	DTCLUXLNRESPONSE_OK_SENDOURXLNBACK = 1
	DTCLUXLNRESPONSE_OK_SENDCONFIRMATION = 2
	DTCLUXLNRESPONSE_LOGNAMEMISMATCH = 3
	DTCLUXLNRESPONSE_COLDWARMMISMATCH = 4
)

type _DtcLu_Xln_Error int32

const (
	DTCLUXLNERROR_PROTOCOL = 1
	DTCLUXLNERROR_LOGNAMEMISMATCH = 2
	DTCLUXLNERROR_COLDWARMMISMATCH = 3
)

type _DtcLu_CompareState int32

const (
	DTCLUCOMPARESTATE_COMMITTED = 1
	DTCLUCOMPARESTATE_HEURISTICCOMMITTED = 2
	DTCLUCOMPARESTATE_HEURISTICMIXED = 3
	DTCLUCOMPARESTATE_HEURISTICRESET = 4
	DTCLUCOMPARESTATE_INDOUBT = 5
	DTCLUCOMPARESTATE_RESET = 6
)

type _DtcLu_CompareStates_Confirmation int32

const (
	DTCLUCOMPARESTATESCONFIRMATION_CONFIRM = 1
	DTCLUCOMPARESTATESCONFIRMATION_PROTOCOL = 2
)

type _DtcLu_CompareStates_Error int32

const (
	DTCLUCOMPARESTATESERROR_PROTOCOL = 1
)

type _DtcLu_CompareStates_Response int32

const (
	DTCLUCOMPARESTATESRESPONSE_OK = 1
	DTCLUCOMPARESTATESRESPONSE_PROTOCOL = 2
)

type TRACKING_COLL_TYPE int32

const (
	TRKCOLL_PROCESSES = 0
	TRKCOLL_APPLICATIONS = 1
	TRKCOLL_COMPONENTS = 2
)

type DUMPTYPE int32

const (
	DUMPTYPE_FULL = 0
	DUMPTYPE_MINI = 1
	DUMPTYPE_NONE = 2
)

type COMPLUS_APPTYPE int32

const (
	APPTYPE_UNKNOWN = -1
	APPTYPE_SERVER = 1
	APPTYPE_LIBRARY = 0
	APPTYPE_SWC = 2
)

type GetAppTrackerDataFlags int32

const (
	GATD_INCLUDE_PROCESS_EXE_NAME = 1
	GATD_INCLUDE_LIBRARY_APPS = 2
	GATD_INCLUDE_SWC = 4
	GATD_INCLUDE_CLASS_NAME = 8
	GATD_INCLUDE_APPLICATION_NAME = 16
)

type TransactionVote int32

const (
	TxCommit = 0
	TxAbort = 1
)

type CrmTransactionState int32

const (
	TxState_Active = 0
	TxState_Committed = 1
	TxState_Aborted = 2
	TxState_Indoubt = 3
)

type CSC_InheritanceConfig int32

const (
	CSC_Inherit = 0
	CSC_Ignore = 1
)

type CSC_ThreadPool int32

const (
	CSC_ThreadPoolNone = 0
	CSC_ThreadPoolInherit = 1
	CSC_STAThreadPool = 2
	CSC_MTAThreadPool = 3
)

type CSC_Binding int32

const (
	CSC_NoBinding = 0
	CSC_BindToPoolThread = 1
)

type CSC_TransactionConfig int32

const (
	CSC_NoTransaction = 0
	CSC_IfContainerIsTransactional = 1
	CSC_CreateTransactionIfNecessary = 2
	CSC_NewTransaction = 3
)

type CSC_SynchronizationConfig int32

const (
	CSC_NoSynchronization = 0
	CSC_IfContainerIsSynchronized = 1
	CSC_NewSynchronizationIfNecessary = 2
	CSC_NewSynchronization = 3
)

type CSC_TrackerConfig int32

const (
	CSC_DontUseTracker = 0
	CSC_UseTracker = 1
)

type CSC_PartitionConfig int32

const (
	CSC_NoPartition = 0
	CSC_InheritPartition = 1
	CSC_NewPartition = 2
)

type CSC_IISIntrinsicsConfig int32

const (
	CSC_NoIISIntrinsics = 0
	CSC_InheritIISIntrinsics = 1
)

type CSC_COMTIIntrinsicsConfig int32

const (
	CSC_NoCOMTIIntrinsics = 0
	CSC_InheritCOMTIIntrinsics = 1
)

type CSC_SxsConfig int32

const (
	CSC_NoSxs = 0
	CSC_InheritSxs = 1
	CSC_NewSxs = 2
)

type __MIDL___MIDL_itf_autosvcs_0001_0150_0001 int32

const (
	mtsErrCtxAborted = -2147164158
	mtsErrCtxAborting = -2147164157
	mtsErrCtxNoContext = -2147164156
	mtsErrCtxNotRegistered = -2147164155
	mtsErrCtxSynchTimeout = -2147164154
	mtsErrCtxOldReference = -2147164153
	mtsErrCtxRoleNotFound = -2147164148
	mtsErrCtxNoSecurity = -2147164147
	mtsErrCtxWrongThread = -2147164146
	mtsErrCtxTMNotAvailable = -2147164145
	comQCErrApplicationNotQueued = -2146368000
	comQCErrNoQueueableInterfaces = -2146367999
	comQCErrQueuingServiceNotAvailable = -2146367998
	comQCErrQueueTransactMismatch = -2146367997
	comqcErrRecorderMarshalled = -2146367996
	comqcErrOutParam = -2146367995
	comqcErrRecorderNotTrusted = -2146367994
	comqcErrPSLoad = -2146367993
	comqcErrMarshaledObjSameTxn = -2146367992
	comqcErrInvalidMessage = -2146367920
	comqcErrMsmqSidUnavailable = -2146367919
	comqcErrWrongMsgExtension = -2146367918
	comqcErrMsmqServiceUnavailable = -2146367917
	comqcErrMsgNotAuthenticated = -2146367916
	comqcErrMsmqConnectorUsed = -2146367915
	comqcErrBadMarshaledObject = -2146367914
)

type __MIDL___MIDL_itf_autosvcs_0001_0159_0001 int32

const (
	LockSetGet = 0
	LockMethod = 1
)

type __MIDL___MIDL_itf_autosvcs_0001_0159_0002 int32

const (
	Standard = 0
	Process = 1
)

type CRMFLAGS int32

const (
	CRMFLAG_FORGETTARGET = 1
	CRMFLAG_WRITTENDURINGPREPARE = 2
	CRMFLAG_WRITTENDURINGCOMMIT = 4
	CRMFLAG_WRITTENDURINGABORT = 8
	CRMFLAG_WRITTENDURINGRECOVERY = 16
	CRMFLAG_WRITTENDURINGREPLAY = 32
	CRMFLAG_REPLAYINPROGRESS = 64
)

type CRMREGFLAGS int32

const (
	CRMREGFLAG_PREPAREPHASE = 1
	CRMREGFLAG_COMMITPHASE = 2
	CRMREGFLAG_ABORTPHASE = 4
	CRMREGFLAG_ALLPHASES = 7
	CRMREGFLAG_FAILIFINDOUBTSREMAIN = 16
)

