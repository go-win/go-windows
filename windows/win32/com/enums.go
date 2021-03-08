// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package com implements the Windows.Win32.Com namespace.
package com

type MEMCTX int32

const (
	MEMCTX_TASK = 1
	MEMCTX_SHARED = 2
	MEMCTX_MACSYSTEM = 3
	MEMCTX_UNKNOWN = -1
	MEMCTX_SAME = -2
)

type CLSCTX int32

const (
	CLSCTX_INPROC_SERVER = 1
	CLSCTX_INPROC_HANDLER = 2
	CLSCTX_LOCAL_SERVER = 4
	CLSCTX_INPROC_SERVER16 = 8
	CLSCTX_REMOTE_SERVER = 16
	CLSCTX_INPROC_HANDLER16 = 32
	CLSCTX_RESERVED1 = 64
	CLSCTX_RESERVED2 = 128
	CLSCTX_RESERVED3 = 256
	CLSCTX_RESERVED4 = 512
	CLSCTX_NO_CODE_DOWNLOAD = 1024
	CLSCTX_RESERVED5 = 2048
	CLSCTX_NO_CUSTOM_MARSHAL = 4096
	CLSCTX_ENABLE_CODE_DOWNLOAD = 8192
	CLSCTX_NO_FAILURE_LOG = 16384
	CLSCTX_DISABLE_AAA = 32768
	CLSCTX_ENABLE_AAA = 65536
	CLSCTX_FROM_DEFAULT_CONTEXT = 131072
	CLSCTX_ACTIVATE_X86_SERVER = 262144
	CLSCTX_ACTIVATE_32_BIT_SERVER = 262144
	CLSCTX_ACTIVATE_64_BIT_SERVER = 524288
	CLSCTX_ENABLE_CLOAKING = 1048576
	CLSCTX_APPCONTAINER = 4194304
	CLSCTX_ACTIVATE_AAA_AS_IU = 8388608
	CLSCTX_RESERVED6 = 16777216
	CLSCTX_ACTIVATE_ARM32_SERVER = 33554432
	CLSCTX_PS_DLL = -2147483648
)

type MSHLFLAGS int32

const (
	MSHLFLAGS_NORMAL = 0
	MSHLFLAGS_TABLESTRONG = 1
	MSHLFLAGS_TABLEWEAK = 2
	MSHLFLAGS_NOPING = 4
	MSHLFLAGS_RESERVED1 = 8
	MSHLFLAGS_RESERVED2 = 16
	MSHLFLAGS_RESERVED3 = 32
	MSHLFLAGS_RESERVED4 = 64
)

type MSHCTX int32

const (
	MSHCTX_LOCAL = 0
	MSHCTX_NOSHAREDMEM = 1
	MSHCTX_DIFFERENTMACHINE = 2
	MSHCTX_INPROC = 3
	MSHCTX_CROSSCTX = 4
	MSHCTX_RESERVED1 = 5
)

type REGCLS int32

const (
	REGCLS_SINGLEUSE = 0
	REGCLS_MULTIPLEUSE = 1
	REGCLS_MULTI_SEPARATE = 2
	REGCLS_SUSPENDED = 4
	REGCLS_SURROGATE = 8
	REGCLS_AGILE = 16
)

type COINITBASE int32

const (
	COINITBASE_MULTITHREADED = 0
)

type EXTCONN int32

const (
	EXTCONN_STRONG = 1
	EXTCONN_WEAK = 2
	EXTCONN_CALLABLE = 4
)

type EOLE_AUTHENTICATION_CAPABILITIES int32

const (
	EOAC_NONE = 0
	EOAC_MUTUAL_AUTH = 1
	EOAC_STATIC_CLOAKING = 32
	EOAC_DYNAMIC_CLOAKING = 64
	EOAC_ANY_AUTHORITY = 128
	EOAC_MAKE_FULLSIC = 256
	EOAC_DEFAULT = 2048
	EOAC_SECURE_REFS = 2
	EOAC_ACCESS_CONTROL = 4
	EOAC_APPID = 8
	EOAC_DYNAMIC = 16
	EOAC_REQUIRE_FULLSIC = 512
	EOAC_AUTO_IMPERSONATE = 1024
	EOAC_DISABLE_AAA = 4096
	EOAC_NO_CUSTOM_MARSHAL = 8192
	EOAC_RESERVED1 = 16384
)

type RPCOPT_PROPERTIES int32

const (
	COMBND_RPCTIMEOUT = 1
	COMBND_SERVER_LOCALITY = 2
	COMBND_RESERVED1 = 4
	COMBND_RESERVED2 = 5
	COMBND_RESERVED3 = 8
	COMBND_RESERVED4 = 16
)

type RPCOPT_SERVER_LOCALITY_VALUES int32

const (
	SERVER_LOCALITY_PROCESS_LOCAL = 0
	SERVER_LOCALITY_MACHINE_LOCAL = 1
	SERVER_LOCALITY_REMOTE = 2
)

type GLOBALOPT_PROPERTIES int32

const (
	COMGLB_EXCEPTION_HANDLING = 1
	COMGLB_APPID = 2
	COMGLB_RPC_THREADPOOL_SETTING = 3
	COMGLB_RO_SETTINGS = 4
	COMGLB_UNMARSHALING_POLICY = 5
	COMGLB_PROPERTIES_RESERVED1 = 6
	COMGLB_PROPERTIES_RESERVED2 = 7
	COMGLB_PROPERTIES_RESERVED3 = 8
)

type GLOBALOPT_EH_VALUES int32

const (
	COMGLB_EXCEPTION_HANDLE = 0
	COMGLB_EXCEPTION_DONOT_HANDLE_FATAL = 1
	COMGLB_EXCEPTION_DONOT_HANDLE = 1
	COMGLB_EXCEPTION_DONOT_HANDLE_ANY = 2
)

type GLOBALOPT_RPCTP_VALUES int32

const (
	COMGLB_RPC_THREADPOOL_SETTING_DEFAULT_POOL = 0
	COMGLB_RPC_THREADPOOL_SETTING_PRIVATE_POOL = 1
)

type GLOBALOPT_RO_FLAGS int32

const (
	COMGLB_STA_MODALLOOP_REMOVE_TOUCH_MESSAGES = 1
	COMGLB_STA_MODALLOOP_SHARED_QUEUE_REMOVE_INPUT_MESSAGES = 2
	COMGLB_STA_MODALLOOP_SHARED_QUEUE_DONOT_REMOVE_INPUT_MESSAGES = 4
	COMGLB_FAST_RUNDOWN = 8
	COMGLB_RESERVED1 = 16
	COMGLB_RESERVED2 = 32
	COMGLB_RESERVED3 = 64
	COMGLB_STA_MODALLOOP_SHARED_QUEUE_REORDER_POINTER_MESSAGES = 128
	COMGLB_RESERVED4 = 256
	COMGLB_RESERVED5 = 512
	COMGLB_RESERVED6 = 1024
)

type GLOBALOPT_UNMARSHALING_POLICY_VALUES int32

const (
	COMGLB_UNMARSHALING_POLICY_NORMAL = 0
	COMGLB_UNMARSHALING_POLICY_STRONG = 1
	COMGLB_UNMARSHALING_POLICY_HYBRID = 2
)

type DCOM_CALL_STATE int32

const (
	DCOM_NONE = 0
	DCOM_CALL_COMPLETE = 1
	DCOM_CALL_CANCELED = 2
)

type APTTYPEQUALIFIER int32

const (
	APTTYPEQUALIFIER_NONE = 0
	APTTYPEQUALIFIER_IMPLICIT_MTA = 1
	APTTYPEQUALIFIER_NA_ON_MTA = 2
	APTTYPEQUALIFIER_NA_ON_STA = 3
	APTTYPEQUALIFIER_NA_ON_IMPLICIT_MTA = 4
	APTTYPEQUALIFIER_NA_ON_MAINSTA = 5
	APTTYPEQUALIFIER_APPLICATION_STA = 6
	APTTYPEQUALIFIER_RESERVED_1 = 7
)

type APTTYPE int32

const (
	APTTYPE_CURRENT = -1
	APTTYPE_STA = 0
	APTTYPE_MTA = 1
	APTTYPE_NA = 2
	APTTYPE_MAINSTA = 3
)

type THDTYPE int32

const (
	THDTYPE_BLOCKMESSAGES = 0
	THDTYPE_PROCESSMESSAGES = 1
)

type CO_MARSHALING_CONTEXT_ATTRIBUTES int32

const (
	CO_MARSHALING_SOURCE_IS_APP_CONTAINER = 0
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_1 = -2147483648
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_2 = -2147483647
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_3 = -2147483646
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_4 = -2147483645
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_5 = -2147483644
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_6 = -2147483643
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_7 = -2147483642
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_8 = -2147483641
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_9 = -2147483640
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_10 = -2147483639
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_11 = -2147483638
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_12 = -2147483637
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_13 = -2147483636
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_14 = -2147483635
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_15 = -2147483634
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_16 = -2147483633
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_17 = -2147483632
	CO_MARSHALING_CONTEXT_ATTRIBUTE_RESERVED_18 = -2147483631
)

type STDMSHLFLAGS int32

const (
	SMEXF_SERVER = 1
	SMEXF_HANDLER = 2
)

type COWAIT_FLAGS int32

const (
	COWAIT_DEFAULT = 0
	COWAIT_WAITALL = 1
	COWAIT_ALERTABLE = 2
	COWAIT_INPUTAVAILABLE = 4
	COWAIT_DISPATCH_CALLS = 8
	COWAIT_DISPATCH_WINDOW_MESSAGES = 16
)

type CWMO_FLAGS int32

const (
	CWMO_DEFAULT = 0
	CWMO_DISPATCH_CALLS = 1
	CWMO_DISPATCH_WINDOW_MESSAGES = 2
)

type BIND_FLAGS int32

const (
	BIND_MAYBOTHERUSER = 1
	BIND_JUSTTESTEXISTENCE = 2
)

type MKSYS int32

const (
	MKSYS_NONE = 0
	MKSYS_GENERICCOMPOSITE = 1
	MKSYS_FILEMONIKER = 2
	MKSYS_ANTIMONIKER = 3
	MKSYS_ITEMMONIKER = 4
	MKSYS_POINTERMONIKER = 5
	MKSYS_CLASSMONIKER = 7
	MKSYS_OBJREFMONIKER = 8
	MKSYS_SESSIONMONIKER = 9
	MKSYS_LUAMONIKER = 10
)

type MKREDUCE int32

const (
	MKRREDUCE_ONE = 196608
	MKRREDUCE_TOUSER = 131072
	MKRREDUCE_THROUGHUSER = 65536
	MKRREDUCE_ALL = 0
)

type ADVF int32

const (
	ADVF_NODATA = 1
	ADVF_PRIMEFIRST = 2
	ADVF_ONLYONCE = 4
	ADVF_DATAONSTOP = 64
	ADVFCACHE_NOHANDLER = 8
	ADVFCACHE_FORCEBUILTIN = 16
	ADVFCACHE_ONSAVE = 32
)

type TYMED int32

const (
	TYMED_HGLOBAL = 1
	TYMED_FILE = 2
	TYMED_ISTREAM = 4
	TYMED_ISTORAGE = 8
	TYMED_GDI = 16
	TYMED_MFPICT = 32
	TYMED_ENHMF = 64
	TYMED_NULL = 0
)

type DATADIR int32

const (
	DATADIR_GET = 1
	DATADIR_SET = 2
)

type CALLTYPE int32

const (
	CALLTYPE_TOPLEVEL = 1
	CALLTYPE_NESTED = 2
	CALLTYPE_ASYNC = 3
	CALLTYPE_TOPLEVEL_CALLPENDING = 4
	CALLTYPE_ASYNC_CALLPENDING = 5
)

type SERVERCALL int32

const (
	SERVERCALL_ISHANDLED = 0
	SERVERCALL_REJECTED = 1
	SERVERCALL_RETRYLATER = 2
)

type PENDINGTYPE int32

const (
	PENDINGTYPE_TOPLEVEL = 1
	PENDINGTYPE_NESTED = 2
)

type PENDINGMSG int32

const (
	PENDINGMSG_CANCELCALL = 0
	PENDINGMSG_WAITNOPROCESS = 1
	PENDINGMSG_WAITDEFPROCESS = 2
)

type ApplicationType int32

const (
	ServerApplication = 0
	LibraryApplication = 1
)

type ShutdownType int32

const (
	IdleShutdown = 0
	ForcedShutdown = 1
)

type DISCARDCACHE int32

const (
	DISCARDCACHE_SAVEIFDIRTY = 0
	DISCARDCACHE_NOSAVE = 1
)

type OLEGETMONIKER int32

const (
	OLEGETMONIKER_ONLYIFTHERE = 1
	OLEGETMONIKER_FORCEASSIGN = 2
	OLEGETMONIKER_UNASSIGN = 3
	OLEGETMONIKER_TEMPFORUSER = 4
)

type OLEWHICHMK int32

const (
	OLEWHICHMK_CONTAINER = 1
	OLEWHICHMK_OBJREL = 2
	OLEWHICHMK_OBJFULL = 3
)

type USERCLASSTYPE int32

const (
	USERCLASSTYPE_FULL = 1
	USERCLASSTYPE_SHORT = 2
	USERCLASSTYPE_APPNAME = 3
)

type OLEMISC int32

const (
	OLEMISC_RECOMPOSEONRESIZE = 1
	OLEMISC_ONLYICONIC = 2
	OLEMISC_INSERTNOTREPLACE = 4
	OLEMISC_STATIC = 8
	OLEMISC_CANTLINKINSIDE = 16
	OLEMISC_CANLINKBYOLE1 = 32
	OLEMISC_ISLINKOBJECT = 64
	OLEMISC_INSIDEOUT = 128
	OLEMISC_ACTIVATEWHENVISIBLE = 256
	OLEMISC_RENDERINGISDEVICEINDEPENDENT = 512
	OLEMISC_INVISIBLEATRUNTIME = 1024
	OLEMISC_ALWAYSRUN = 2048
	OLEMISC_ACTSLIKEBUTTON = 4096
	OLEMISC_ACTSLIKELABEL = 8192
	OLEMISC_NOUIACTIVATE = 16384
	OLEMISC_ALIGNABLE = 32768
	OLEMISC_SIMPLEFRAME = 65536
	OLEMISC_SETCLIENTSITEFIRST = 131072
	OLEMISC_IMEMODE = 262144
	OLEMISC_IGNOREACTIVATEWHENVISIBLE = 524288
	OLEMISC_WANTSTOMENUMERGE = 1048576
	OLEMISC_SUPPORTSMULTILEVELUNDO = 2097152
)

type OLECLOSE int32

const (
	OLECLOSE_SAVEIFDIRTY = 0
	OLECLOSE_NOSAVE = 1
	OLECLOSE_PROMPTSAVE = 2
)

type OLERENDER int32

const (
	OLERENDER_NONE = 0
	OLERENDER_DRAW = 1
	OLERENDER_FORMAT = 2
	OLERENDER_ASIS = 3
)

type OLEUPDATE int32

const (
	OLEUPDATE_ALWAYS = 1
	OLEUPDATE_ONCALL = 3
)

type OLELINKBIND int32

const (
	OLELINKBIND_EVENIFCLASSDIFF = 1
)

type BINDSPEED int32

const (
	BINDSPEED_INDEFINITE = 1
	BINDSPEED_MODERATE = 2
	BINDSPEED_IMMEDIATE = 3
)

type OLECONTF int32

const (
	OLECONTF_EMBEDDINGS = 1
	OLECONTF_LINKS = 2
	OLECONTF_OTHERS = 4
	OLECONTF_ONLYUSER = 8
	OLECONTF_ONLYIFRUNNING = 16
)

type OLEVERBATTRIB int32

const (
	OLEVERBATTRIB_NEVERDIRTIES = 1
	OLEVERBATTRIB_ONCONTAINERMENU = 2
)

type IEObjectType int32

const (
	IE_EPM_OBJECT_EVENT = 0
	IE_EPM_OBJECT_MUTEX = 1
	IE_EPM_OBJECT_SEMAPHORE = 2
	IE_EPM_OBJECT_SHARED_MEMORY = 3
	IE_EPM_OBJECT_WAITABLE_TIMER = 4
	IE_EPM_OBJECT_FILE = 5
	IE_EPM_OBJECT_NAMED_PIPE = 6
	IE_EPM_OBJECT_REGISTRY = 7
)

type MONIKERPROPERTY int32

const (
	MIMETYPEPROP = 0
	USE_SRC_URL = 1
	CLASSIDPROP = 2
	TRUSTEDDOWNLOADPROP = 3
	POPUPLEVELPROP = 4
)

type BINDVERB int32

const (
	BINDVERB_GET = 0
	BINDVERB_POST = 1
	BINDVERB_PUT = 2
	BINDVERB_CUSTOM = 3
	BINDVERB_RESERVED1 = 4
)

type BINDINFOF int32

const (
	BINDINFOF_URLENCODESTGMEDDATA = 1
	BINDINFOF_URLENCODEDEXTRAINFO = 2
)

type BINDF int32

const (
	BINDF_ASYNCHRONOUS = 1
	BINDF_ASYNCSTORAGE = 2
	BINDF_NOPROGRESSIVERENDERING = 4
	BINDF_OFFLINEOPERATION = 8
	BINDF_GETNEWESTVERSION = 16
	BINDF_NOWRITECACHE = 32
	BINDF_NEEDFILE = 64
	BINDF_PULLDATA = 128
	BINDF_IGNORESECURITYPROBLEM = 256
	BINDF_RESYNCHRONIZE = 512
	BINDF_HYPERLINK = 1024
	BINDF_NO_UI = 2048
	BINDF_SILENTOPERATION = 4096
	BINDF_PRAGMA_NO_CACHE = 8192
	BINDF_GETCLASSOBJECT = 16384
	BINDF_RESERVED_1 = 32768
	BINDF_FREE_THREADED = 65536
	BINDF_DIRECT_READ = 131072
	BINDF_FORMS_SUBMIT = 262144
	BINDF_GETFROMCACHE_IF_NET_FAIL = 524288
	BINDF_FROMURLMON = 1048576
	BINDF_FWD_BACK = 2097152
	BINDF_PREFERDEFAULTHANDLER = 4194304
	BINDF_ENFORCERESTRICTED = 8388608
	BINDF_RESERVED_2 = -2147483648
	BINDF_RESERVED_3 = 16777216
	BINDF_RESERVED_4 = 33554432
	BINDF_RESERVED_5 = 67108864
	BINDF_RESERVED_6 = 134217728
	BINDF_RESERVED_7 = 1073741824
	BINDF_RESERVED_8 = 536870912
)

type URL_ENCODING int32

const (
	URL_ENCODING_NONE = 0
	URL_ENCODING_ENABLE_UTF8 = 268435456
	URL_ENCODING_DISABLE_UTF8 = 536870912
)

type BINDINFO_OPTIONS int32

const (
	BINDINFO_OPTIONS_WININETFLAG = 65536
	BINDINFO_OPTIONS_ENABLE_UTF8 = 131072
	BINDINFO_OPTIONS_DISABLE_UTF8 = 262144
	BINDINFO_OPTIONS_USE_IE_ENCODING = 524288
	BINDINFO_OPTIONS_BINDTOOBJECT = 1048576
	BINDINFO_OPTIONS_SECURITYOPTOUT = 2097152
	BINDINFO_OPTIONS_IGNOREMIMETEXTPLAIN = 4194304
	BINDINFO_OPTIONS_USEBINDSTRINGCREDS = 8388608
	BINDINFO_OPTIONS_IGNOREHTTPHTTPSREDIRECTS = 16777216
	BINDINFO_OPTIONS_IGNORE_SSLERRORS_ONCE = 33554432
	BINDINFO_WPC_DOWNLOADBLOCKED = 134217728
	BINDINFO_WPC_LOGGING_ENABLED = 268435456
	BINDINFO_OPTIONS_ALLOWCONNECTDATA = 536870912
	BINDINFO_OPTIONS_DISABLEAUTOREDIRECTS = 1073741824
	BINDINFO_OPTIONS_SHDOCVW_NAVIGATE = -2147483648
)

type BSCF int32

const (
	BSCF_FIRSTDATANOTIFICATION = 1
	BSCF_INTERMEDIATEDATANOTIFICATION = 2
	BSCF_LASTDATANOTIFICATION = 4
	BSCF_DATAFULLYAVAILABLE = 8
	BSCF_AVAILABLEDATASIZEUNKNOWN = 16
	BSCF_SKIPDRAINDATAFORFILEURLS = 32
	BSCF_64BITLENGTHDOWNLOAD = 64
)

type BINDSTATUS int32

const (
	BINDSTATUS_FINDINGRESOURCE = 1
	BINDSTATUS_CONNECTING = 2
	BINDSTATUS_REDIRECTING = 3
	BINDSTATUS_BEGINDOWNLOADDATA = 4
	BINDSTATUS_DOWNLOADINGDATA = 5
	BINDSTATUS_ENDDOWNLOADDATA = 6
	BINDSTATUS_BEGINDOWNLOADCOMPONENTS = 7
	BINDSTATUS_INSTALLINGCOMPONENTS = 8
	BINDSTATUS_ENDDOWNLOADCOMPONENTS = 9
	BINDSTATUS_USINGCACHEDCOPY = 10
	BINDSTATUS_SENDINGREQUEST = 11
	BINDSTATUS_CLASSIDAVAILABLE = 12
	BINDSTATUS_MIMETYPEAVAILABLE = 13
	BINDSTATUS_CACHEFILENAMEAVAILABLE = 14
	BINDSTATUS_BEGINSYNCOPERATION = 15
	BINDSTATUS_ENDSYNCOPERATION = 16
	BINDSTATUS_BEGINUPLOADDATA = 17
	BINDSTATUS_UPLOADINGDATA = 18
	BINDSTATUS_ENDUPLOADDATA = 19
	BINDSTATUS_PROTOCOLCLASSID = 20
	BINDSTATUS_ENCODING = 21
	BINDSTATUS_VERIFIEDMIMETYPEAVAILABLE = 22
	BINDSTATUS_CLASSINSTALLLOCATION = 23
	BINDSTATUS_DECODING = 24
	BINDSTATUS_LOADINGMIMEHANDLER = 25
	BINDSTATUS_CONTENTDISPOSITIONATTACH = 26
	BINDSTATUS_FILTERREPORTMIMETYPE = 27
	BINDSTATUS_CLSIDCANINSTANTIATE = 28
	BINDSTATUS_IUNKNOWNAVAILABLE = 29
	BINDSTATUS_DIRECTBIND = 30
	BINDSTATUS_RAWMIMETYPE = 31
	BINDSTATUS_PROXYDETECTING = 32
	BINDSTATUS_ACCEPTRANGES = 33
	BINDSTATUS_COOKIE_SENT = 34
	BINDSTATUS_COMPACT_POLICY_RECEIVED = 35
	BINDSTATUS_COOKIE_SUPPRESSED = 36
	BINDSTATUS_COOKIE_STATE_UNKNOWN = 37
	BINDSTATUS_COOKIE_STATE_ACCEPT = 38
	BINDSTATUS_COOKIE_STATE_REJECT = 39
	BINDSTATUS_COOKIE_STATE_PROMPT = 40
	BINDSTATUS_COOKIE_STATE_LEASH = 41
	BINDSTATUS_COOKIE_STATE_DOWNGRADE = 42
	BINDSTATUS_POLICY_HREF = 43
	BINDSTATUS_P3P_HEADER = 44
	BINDSTATUS_SESSION_COOKIE_RECEIVED = 45
	BINDSTATUS_PERSISTENT_COOKIE_RECEIVED = 46
	BINDSTATUS_SESSION_COOKIES_ALLOWED = 47
	BINDSTATUS_CACHECONTROL = 48
	BINDSTATUS_CONTENTDISPOSITIONFILENAME = 49
	BINDSTATUS_MIMETEXTPLAINMISMATCH = 50
	BINDSTATUS_PUBLISHERAVAILABLE = 51
	BINDSTATUS_DISPLAYNAMEAVAILABLE = 52
	BINDSTATUS_SSLUX_NAVBLOCKED = 53
	BINDSTATUS_SERVER_MIMETYPEAVAILABLE = 54
	BINDSTATUS_SNIFFED_CLASSIDAVAILABLE = 55
	BINDSTATUS_64BIT_PROGRESS = 56
	BINDSTATUS_LAST = 56
	BINDSTATUS_RESERVED_0 = 57
	BINDSTATUS_RESERVED_1 = 58
	BINDSTATUS_RESERVED_2 = 59
	BINDSTATUS_RESERVED_3 = 60
	BINDSTATUS_RESERVED_4 = 61
	BINDSTATUS_RESERVED_5 = 62
	BINDSTATUS_RESERVED_6 = 63
	BINDSTATUS_RESERVED_7 = 64
	BINDSTATUS_RESERVED_8 = 65
	BINDSTATUS_RESERVED_9 = 66
	BINDSTATUS_RESERVED_A = 67
	BINDSTATUS_RESERVED_B = 68
	BINDSTATUS_RESERVED_C = 69
	BINDSTATUS_RESERVED_D = 70
	BINDSTATUS_RESERVED_E = 71
	BINDSTATUS_RESERVED_F = 72
	BINDSTATUS_RESERVED_10 = 73
	BINDSTATUS_RESERVED_11 = 74
	BINDSTATUS_RESERVED_12 = 75
	BINDSTATUS_RESERVED_13 = 76
	BINDSTATUS_LAST_PRIVATE = 76
)

type BINDF2 int32

const (
	BINDF2_DISABLEBASICOVERHTTP = 1
	BINDF2_DISABLEAUTOCOOKIEHANDLING = 2
	BINDF2_READ_DATA_GREATER_THAN_4GB = 4
	BINDF2_DISABLE_HTTP_REDIRECT_XSECURITYID = 8
	BINDF2_SETDOWNLOADMODE = 32
	BINDF2_DISABLE_HTTP_REDIRECT_CACHING = 64
	BINDF2_KEEP_CALLBACK_MODULE_LOADED = 128
	BINDF2_ALLOW_PROXY_CRED_PROMPT = 256
	BINDF2_RESERVED_17 = 512
	BINDF2_RESERVED_16 = 1024
	BINDF2_RESERVED_15 = 2048
	BINDF2_RESERVED_14 = 4096
	BINDF2_RESERVED_13 = 8192
	BINDF2_RESERVED_12 = 16384
	BINDF2_RESERVED_11 = 32768
	BINDF2_RESERVED_10 = 65536
	BINDF2_RESERVED_F = 131072
	BINDF2_RESERVED_E = 262144
	BINDF2_RESERVED_D = 524288
	BINDF2_RESERVED_C = 1048576
	BINDF2_RESERVED_B = 2097152
	BINDF2_RESERVED_A = 4194304
	BINDF2_RESERVED_9 = 8388608
	BINDF2_RESERVED_8 = 16777216
	BINDF2_RESERVED_7 = 33554432
	BINDF2_RESERVED_6 = 67108864
	BINDF2_RESERVED_5 = 134217728
	BINDF2_RESERVED_4 = 268435456
	BINDF2_RESERVED_3 = 536870912
	BINDF2_RESERVED_2 = 1073741824
	BINDF2_RESERVED_1 = -2147483648
)

type AUTHENTICATEF int32

const (
	AUTHENTICATEF_PROXY = 1
	AUTHENTICATEF_BASIC = 2
	AUTHENTICATEF_HTTP = 4
)

type CIP_STATUS int32

const (
	CIP_DISK_FULL = 0
	CIP_ACCESS_DENIED = 1
	CIP_NEWER_VERSION_EXISTS = 2
	CIP_OLDER_VERSION_EXISTS = 3
	CIP_NAME_CONFLICT = 4
	CIP_TRUST_VERIFICATION_COMPONENT_MISSING = 5
	CIP_EXE_SELF_REGISTERATION_TIMEOUT = 6
	CIP_UNSAFE_TO_ABORT = 7
	CIP_NEED_REBOOT = 8
	CIP_NEED_REBOOT_UI_PERMISSION = 9
)

type Uri_PROPERTY int32

const (
	Uri_PROPERTY_ABSOLUTE_URI = 0
	Uri_PROPERTY_STRING_START = 0
	Uri_PROPERTY_AUTHORITY = 1
	Uri_PROPERTY_DISPLAY_URI = 2
	Uri_PROPERTY_DOMAIN = 3
	Uri_PROPERTY_EXTENSION = 4
	Uri_PROPERTY_FRAGMENT = 5
	Uri_PROPERTY_HOST = 6
	Uri_PROPERTY_PASSWORD = 7
	Uri_PROPERTY_PATH = 8
	Uri_PROPERTY_PATH_AND_QUERY = 9
	Uri_PROPERTY_QUERY = 10
	Uri_PROPERTY_RAW_URI = 11
	Uri_PROPERTY_SCHEME_NAME = 12
	Uri_PROPERTY_USER_INFO = 13
	Uri_PROPERTY_USER_NAME = 14
	Uri_PROPERTY_STRING_LAST = 14
	Uri_PROPERTY_HOST_TYPE = 15
	Uri_PROPERTY_DWORD_START = 15
	Uri_PROPERTY_PORT = 16
	Uri_PROPERTY_SCHEME = 17
	Uri_PROPERTY_ZONE = 18
	Uri_PROPERTY_DWORD_LAST = 18
)

type Uri_HOST_TYPE int32

const (
	Uri_HOST_UNKNOWN = 0
	Uri_HOST_DNS = 1
	Uri_HOST_IPV4 = 2
	Uri_HOST_IPV6 = 3
	Uri_HOST_IDN = 4
)

type BINDSTRING int32

const (
	BINDSTRING_HEADERS = 1
	BINDSTRING_ACCEPT_MIMES = 2
	BINDSTRING_EXTRA_URL = 3
	BINDSTRING_LANGUAGE = 4
	BINDSTRING_USERNAME = 5
	BINDSTRING_PASSWORD = 6
	BINDSTRING_UA_PIXELS = 7
	BINDSTRING_UA_COLOR = 8
	BINDSTRING_OS = 9
	BINDSTRING_USER_AGENT = 10
	BINDSTRING_ACCEPT_ENCODINGS = 11
	BINDSTRING_POST_COOKIE = 12
	BINDSTRING_POST_DATA_MIME = 13
	BINDSTRING_URL = 14
	BINDSTRING_IID = 15
	BINDSTRING_FLAG_BIND_TO_OBJECT = 16
	BINDSTRING_PTR_BIND_CONTEXT = 17
	BINDSTRING_XDR_ORIGIN = 18
	BINDSTRING_DOWNLOADPATH = 19
	BINDSTRING_ROOTDOC_URL = 20
	BINDSTRING_INITIAL_FILENAME = 21
	BINDSTRING_PROXY_USERNAME = 22
	BINDSTRING_PROXY_PASSWORD = 23
	BINDSTRING_ENTERPRISE_ID = 24
	BINDSTRING_DOC_URL = 25
	BINDSTRING_SAMESITE_COOKIE_LEVEL = 26
)

type PI_FLAGS int32

const (
	PI_PARSE_URL = 1
	PI_FILTER_MODE = 2
	PI_FORCE_ASYNC = 4
	PI_USE_WORKERTHREAD = 8
	PI_MIMEVERIFICATION = 16
	PI_CLSIDLOOKUP = 32
	PI_DATAPROGRESS = 64
	PI_SYNCHRONOUS = 128
	PI_APARTMENTTHREADED = 256
	PI_CLASSINSTALL = 512
	PI_PASSONBINDCTX = 8192
	PI_NOMIMEHANDLER = 32768
	PI_LOADAPPDIRECT = 16384
	PD_FORCE_SWITCH = 65536
	PI_PREFERDEFAULTHANDLER = 131072
)

type OIBDG_FLAGS int32

const (
	OIBDG_APARTMENTTHREADED = 256
	OIBDG_DATAONLY = 4096
)

type PARSEACTION int32

const (
	PARSE_CANONICALIZE = 1
	PARSE_FRIENDLY = 2
	PARSE_SECURITY_URL = 3
	PARSE_ROOTDOCUMENT = 4
	PARSE_DOCUMENT = 5
	PARSE_ANCHOR = 6
	PARSE_ENCODE_IS_UNESCAPE = 7
	PARSE_DECODE_IS_ESCAPE = 8
	PARSE_PATH_FROM_URL = 9
	PARSE_URL_FROM_PATH = 10
	PARSE_MIME = 11
	PARSE_SERVER = 12
	PARSE_SCHEMA = 13
	PARSE_SITE = 14
	PARSE_DOMAIN = 15
	PARSE_LOCATION = 16
	PARSE_SECURITY_DOMAIN = 17
	PARSE_ESCAPE = 18
	PARSE_UNESCAPE = 19
)

type PSUACTION int32

const (
	PSU_DEFAULT = 1
	PSU_SECURITY_URL_ONLY = 2
)

type QUERYOPTION int32

const (
	QUERY_EXPIRATION_DATE = 1
	QUERY_TIME_OF_LAST_CHANGE = 2
	QUERY_CONTENT_ENCODING = 3
	QUERY_CONTENT_TYPE = 4
	QUERY_REFRESH = 5
	QUERY_RECOMBINE = 6
	QUERY_CAN_NAVIGATE = 7
	QUERY_USES_NETWORK = 8
	QUERY_IS_CACHED = 9
	QUERY_IS_INSTALLEDENTRY = 10
	QUERY_IS_CACHED_OR_MAPPED = 11
	QUERY_USES_CACHE = 12
	QUERY_IS_SECURE = 13
	QUERY_IS_SAFE = 14
	QUERY_USES_HISTORYFOLDER = 15
	QUERY_IS_CACHED_AND_USABLE_OFFLINE = 16
)

type INTERNETFEATURELIST int32

const (
	FEATURE_OBJECT_CACHING = 0
	FEATURE_ZONE_ELEVATION = 1
	FEATURE_MIME_HANDLING = 2
	FEATURE_MIME_SNIFFING = 3
	FEATURE_WINDOW_RESTRICTIONS = 4
	FEATURE_WEBOC_POPUPMANAGEMENT = 5
	FEATURE_BEHAVIORS = 6
	FEATURE_DISABLE_MK_PROTOCOL = 7
	FEATURE_LOCALMACHINE_LOCKDOWN = 8
	FEATURE_SECURITYBAND = 9
	FEATURE_RESTRICT_ACTIVEXINSTALL = 10
	FEATURE_VALIDATE_NAVIGATE_URL = 11
	FEATURE_RESTRICT_FILEDOWNLOAD = 12
	FEATURE_ADDON_MANAGEMENT = 13
	FEATURE_PROTOCOL_LOCKDOWN = 14
	FEATURE_HTTP_USERNAME_PASSWORD_DISABLE = 15
	FEATURE_SAFE_BINDTOOBJECT = 16
	FEATURE_UNC_SAVEDFILECHECK = 17
	FEATURE_GET_URL_DOM_FILEPATH_UNENCODED = 18
	FEATURE_TABBED_BROWSING = 19
	FEATURE_SSLUX = 20
	FEATURE_DISABLE_NAVIGATION_SOUNDS = 21
	FEATURE_DISABLE_LEGACY_COMPRESSION = 22
	FEATURE_FORCE_ADDR_AND_STATUS = 23
	FEATURE_XMLHTTP = 24
	FEATURE_DISABLE_TELNET_PROTOCOL = 25
	FEATURE_FEEDS = 26
	FEATURE_BLOCK_INPUT_PROMPTS = 27
	FEATURE_ENTRY_COUNT = 28
)

type PUAF int32

const (
	PUAF_DEFAULT = 0
	PUAF_NOUI = 1
	PUAF_ISFILE = 2
	PUAF_WARN_IF_DENIED = 4
	PUAF_FORCEUI_FOREGROUND = 8
	PUAF_CHECK_TIFS = 16
	PUAF_DONTCHECKBOXINDIALOG = 32
	PUAF_TRUSTED = 64
	PUAF_ACCEPT_WILDCARD_SCHEME = 128
	PUAF_ENFORCERESTRICTED = 256
	PUAF_NOSAVEDFILECHECK = 512
	PUAF_REQUIRESAVEDFILECHECK = 1024
	PUAF_DONT_USE_CACHE = 4096
	PUAF_RESERVED1 = 8192
	PUAF_RESERVED2 = 16384
	PUAF_LMZ_UNLOCKED = 65536
	PUAF_LMZ_LOCKED = 131072
	PUAF_DEFAULTZONEPOL = 262144
	PUAF_NPL_USE_LOCKED_IF_RESTRICTED = 524288
	PUAF_NOUIIFLOCKED = 1048576
	PUAF_DRAGPROTOCOLCHECK = 2097152
)

type PUAFOUT int32

const (
	PUAFOUT_DEFAULT = 0
	PUAFOUT_ISLOCKZONEPOLICY = 1
)

type SZM_FLAGS int32

const (
	SZM_CREATE = 0
	SZM_DELETE = 1
)

type URLZONE int32

const (
	URLZONE_INVALID = -1
	URLZONE_PREDEFINED_MIN = 0
	URLZONE_LOCAL_MACHINE = 0
	URLZONE_INTRANET = 1
	URLZONE_TRUSTED = 2
	URLZONE_INTERNET = 3
	URLZONE_UNTRUSTED = 4
	URLZONE_PREDEFINED_MAX = 999
	URLZONE_USER_MIN = 1000
	URLZONE_USER_MAX = 10000
)

type URLTEMPLATE int32

const (
	URLTEMPLATE_CUSTOM = 0
	URLTEMPLATE_PREDEFINED_MIN = 65536
	URLTEMPLATE_LOW = 65536
	URLTEMPLATE_MEDLOW = 66816
	URLTEMPLATE_MEDIUM = 69632
	URLTEMPLATE_MEDHIGH = 70912
	URLTEMPLATE_HIGH = 73728
	URLTEMPLATE_PREDEFINED_MAX = 131072
)

type __MIDL_IInternetZoneManager_0001 int32

const (
	MAX_ZONE_PATH = 260
	MAX_ZONE_DESCRIPTION = 200
)

type ZAFLAGS int32

const (
	ZAFLAGS_CUSTOM_EDIT = 1
	ZAFLAGS_ADD_SITES = 2
	ZAFLAGS_REQUIRE_VERIFICATION = 4
	ZAFLAGS_INCLUDE_PROXY_OVERRIDE = 8
	ZAFLAGS_INCLUDE_INTRANET_SITES = 16
	ZAFLAGS_NO_UI = 32
	ZAFLAGS_SUPPORTS_VERIFICATION = 64
	ZAFLAGS_UNC_AS_INTRANET = 128
	ZAFLAGS_DETECT_INTRANET = 256
	ZAFLAGS_USE_LOCKED_ZONES = 65536
	ZAFLAGS_VERIFY_TEMPLATE_SETTINGS = 131072
	ZAFLAGS_NO_CACHE = 262144
)

type URLZONEREG int32

const (
	URLZONEREG_DEFAULT = 0
	URLZONEREG_HKLM = 1
	URLZONEREG_HKCU = 2
)

type BINDHANDLETYPES int32

const (
	BINDHANDLETYPES_APPCACHE = 0
	BINDHANDLETYPES_DEPENDENCY = 1
	BINDHANDLETYPES_COUNT = 2
)

type UASFLAGS int32

const (
	UAS_NORMAL = 0
	UAS_BLOCKED = 1
	UAS_NOPARENTENABLE = 2
	UAS_MASK = 3
)

type GUIDKIND int32

const (
	GUIDKIND_DEFAULT_SOURCE_DISP_IID = 1
)

type CTRLINFO int32

const (
	CTRLINFO_EATS_RETURN = 1
	CTRLINFO_EATS_ESCAPE = 2
)

type XFORMCOORDS int32

const (
	XFORMCOORDS_POSITION = 1
	XFORMCOORDS_SIZE = 2
	XFORMCOORDS_HIMETRICTOCONTAINER = 4
	XFORMCOORDS_CONTAINERTOHIMETRIC = 8
	XFORMCOORDS_EVENTCOMPAT = 16
)

type PROPPAGESTATUS int32

const (
	PROPPAGESTATUS_DIRTY = 1
	PROPPAGESTATUS_VALIDATE = 2
	PROPPAGESTATUS_CLEAN = 4
)

type PictureAttributes int32

const (
	PICTURE_SCALABLE = 1
	PICTURE_TRANSPARENT = 2
)

type ACTIVATEFLAGS int32

const (
	ACTIVATE_WINDOWLESS = 1
)

type OLEDCFLAGS int32

const (
	OLEDC_NODRAW = 1
	OLEDC_PAINTBKGND = 2
	OLEDC_OFFSCREEN = 4
)

type VIEWSTATUS int32

const (
	VIEWSTATUS_OPAQUE = 1
	VIEWSTATUS_SOLIDBKGND = 2
	VIEWSTATUS_DVASPECTOPAQUE = 4
	VIEWSTATUS_DVASPECTTRANSPARENT = 8
	VIEWSTATUS_SURFACE = 16
	VIEWSTATUS_3DSURFACE = 32
)

type HITRESULT int32

const (
	HITRESULT_OUTSIDE = 0
	HITRESULT_TRANSPARENT = 1
	HITRESULT_CLOSE = 2
	HITRESULT_HIT = 3
)

type DVASPECT2 int32

const (
	DVASPECT_OPAQUE = 16
	DVASPECT_TRANSPARENT = 32
)

type ExtentMode int32

const (
	DVEXTENT_CONTENT = 0
	DVEXTENT_INTEGRAL = 1
)

type AspectInfoFlag int32

const (
	DVASPECTINFOFLAG_CANOPTIMIZE = 1
)

type POINTERINACTIVE int32

const (
	POINTERINACTIVE_ACTIVATEONENTRY = 1
	POINTERINACTIVE_DEACTIVATEONLEAVE = 2
	POINTERINACTIVE_ACTIVATEONDRAG = 4
)

type PROPBAG2_TYPE int32

const (
	PROPBAG2_TYPE_UNDEFINED = 0
	PROPBAG2_TYPE_DATA = 1
	PROPBAG2_TYPE_URL = 2
	PROPBAG2_TYPE_OBJECT = 3
	PROPBAG2_TYPE_STREAM = 4
	PROPBAG2_TYPE_STORAGE = 5
	PROPBAG2_TYPE_MONIKER = 6
)

type QACONTAINERFLAGS int32

const (
	QACONTAINER_SHOWHATCHING = 1
	QACONTAINER_SHOWGRABHANDLES = 2
	QACONTAINER_USERMODE = 4
	QACONTAINER_DISPLAYASDEFAULT = 8
	QACONTAINER_UIDEAD = 16
	QACONTAINER_AUTOCLIP = 32
	QACONTAINER_MESSAGEREFLECT = 64
	QACONTAINER_SUPPORTSMNEMONICS = 128
)

type OLE_TRISTATE int32

const (
	triUnchecked = 0
	triChecked = 1
	triGray = 2
)

type DOCMISC int32

const (
	DOCMISC_CANCREATEMULTIPLEVIEWS = 1
	DOCMISC_SUPPORTCOMPLEXRECTANGLES = 2
	DOCMISC_CANTOPENEDIT = 4
	DOCMISC_NOFILESUPPORT = 8
)

type __MIDL_IPrint_0001 int32

const (
	PRINTFLAG_MAYBOTHERUSER = 1
	PRINTFLAG_PROMPTUSER = 2
	PRINTFLAG_USERMAYCHANGEPRINTER = 4
	PRINTFLAG_RECOMPOSETODEVICE = 8
	PRINTFLAG_DONTACTUALLYPRINT = 16
	PRINTFLAG_FORCEPROPERTIES = 32
	PRINTFLAG_PRINTTOFILE = 64
)

type OLECMDF int32

const (
	OLECMDF_SUPPORTED = 1
	OLECMDF_ENABLED = 2
	OLECMDF_LATCHED = 4
	OLECMDF_NINCHED = 8
	OLECMDF_INVISIBLE = 16
	OLECMDF_DEFHIDEONCTXTMENU = 32
)

type OLECMDTEXTF int32

const (
	OLECMDTEXTF_NONE = 0
	OLECMDTEXTF_NAME = 1
	OLECMDTEXTF_STATUS = 2
)

type OLECMDEXECOPT int32

const (
	OLECMDEXECOPT_DODEFAULT = 0
	OLECMDEXECOPT_PROMPTUSER = 1
	OLECMDEXECOPT_DONTPROMPTUSER = 2
	OLECMDEXECOPT_SHOWHELP = 3
)

type OLECMDID int32

const (
	OLECMDID_OPEN = 1
	OLECMDID_NEW = 2
	OLECMDID_SAVE = 3
	OLECMDID_SAVEAS = 4
	OLECMDID_SAVECOPYAS = 5
	OLECMDID_PRINT = 6
	OLECMDID_PRINTPREVIEW = 7
	OLECMDID_PAGESETUP = 8
	OLECMDID_SPELL = 9
	OLECMDID_PROPERTIES = 10
	OLECMDID_CUT = 11
	OLECMDID_COPY = 12
	OLECMDID_PASTE = 13
	OLECMDID_PASTESPECIAL = 14
	OLECMDID_UNDO = 15
	OLECMDID_REDO = 16
	OLECMDID_SELECTALL = 17
	OLECMDID_CLEARSELECTION = 18
	OLECMDID_ZOOM = 19
	OLECMDID_GETZOOMRANGE = 20
	OLECMDID_UPDATECOMMANDS = 21
	OLECMDID_REFRESH = 22
	OLECMDID_STOP = 23
	OLECMDID_HIDETOOLBARS = 24
	OLECMDID_SETPROGRESSMAX = 25
	OLECMDID_SETPROGRESSPOS = 26
	OLECMDID_SETPROGRESSTEXT = 27
	OLECMDID_SETTITLE = 28
	OLECMDID_SETDOWNLOADSTATE = 29
	OLECMDID_STOPDOWNLOAD = 30
	OLECMDID_ONTOOLBARACTIVATED = 31
	OLECMDID_FIND = 32
	OLECMDID_DELETE = 33
	OLECMDID_HTTPEQUIV = 34
	OLECMDID_HTTPEQUIV_DONE = 35
	OLECMDID_ENABLE_INTERACTION = 36
	OLECMDID_ONUNLOAD = 37
	OLECMDID_PROPERTYBAG2 = 38
	OLECMDID_PREREFRESH = 39
	OLECMDID_SHOWSCRIPTERROR = 40
	OLECMDID_SHOWMESSAGE = 41
	OLECMDID_SHOWFIND = 42
	OLECMDID_SHOWPAGESETUP = 43
	OLECMDID_SHOWPRINT = 44
	OLECMDID_CLOSE = 45
	OLECMDID_ALLOWUILESSSAVEAS = 46
	OLECMDID_DONTDOWNLOADCSS = 47
	OLECMDID_UPDATEPAGESTATUS = 48
	OLECMDID_PRINT2 = 49
	OLECMDID_PRINTPREVIEW2 = 50
	OLECMDID_SETPRINTTEMPLATE = 51
	OLECMDID_GETPRINTTEMPLATE = 52
	OLECMDID_PAGEACTIONBLOCKED = 55
	OLECMDID_PAGEACTIONUIQUERY = 56
	OLECMDID_FOCUSVIEWCONTROLS = 57
	OLECMDID_FOCUSVIEWCONTROLSQUERY = 58
	OLECMDID_SHOWPAGEACTIONMENU = 59
	OLECMDID_ADDTRAVELENTRY = 60
	OLECMDID_UPDATETRAVELENTRY = 61
	OLECMDID_UPDATEBACKFORWARDSTATE = 62
	OLECMDID_OPTICAL_ZOOM = 63
	OLECMDID_OPTICAL_GETZOOMRANGE = 64
	OLECMDID_WINDOWSTATECHANGED = 65
	OLECMDID_ACTIVEXINSTALLSCOPE = 66
	OLECMDID_UPDATETRAVELENTRY_DATARECOVERY = 67
	OLECMDID_SHOWTASKDLG = 68
	OLECMDID_POPSTATEEVENT = 69
	OLECMDID_VIEWPORT_MODE = 70
	OLECMDID_LAYOUT_VIEWPORT_WIDTH = 71
	OLECMDID_VISUAL_VIEWPORT_EXCLUDE_BOTTOM = 72
	OLECMDID_USER_OPTICAL_ZOOM = 73
	OLECMDID_PAGEAVAILABLE = 74
	OLECMDID_GETUSERSCALABLE = 75
	OLECMDID_UPDATE_CARET = 76
	OLECMDID_ENABLE_VISIBILITY = 77
	OLECMDID_MEDIA_PLAYBACK = 78
	OLECMDID_SETFAVICON = 79
	OLECMDID_SET_HOST_FULLSCREENMODE = 80
	OLECMDID_EXITFULLSCREEN = 81
	OLECMDID_SCROLLCOMPLETE = 82
	OLECMDID_ONBEFOREUNLOAD = 83
	OLECMDID_SHOWMESSAGE_BLOCKABLE = 84
	OLECMDID_SHOWTASKDLG_BLOCKABLE = 85
)

type MEDIAPLAYBACK_STATE int32

const (
	MEDIAPLAYBACK_RESUME = 0
	MEDIAPLAYBACK_PAUSE = 1
	MEDIAPLAYBACK_PAUSE_AND_SUSPEND = 2
	MEDIAPLAYBACK_RESUME_FROM_SUSPEND = 3
)

type IGNOREMIME int32

const (
	IGNOREMIME_PROMPT = 1
	IGNOREMIME_TEXT = 2
)

type WPCSETTING int32

const (
	WPCSETTING_LOGGING_ENABLED = 1
	WPCSETTING_FILEDOWNLOAD_BLOCKED = 2
)

type OLECMDID_REFRESHFLAG int32

const (
	OLECMDIDF_REFRESH_NORMAL = 0
	OLECMDIDF_REFRESH_IFEXPIRED = 1
	OLECMDIDF_REFRESH_CONTINUE = 2
	OLECMDIDF_REFRESH_COMPLETELY = 3
	OLECMDIDF_REFRESH_NO_CACHE = 4
	OLECMDIDF_REFRESH_RELOAD = 5
	OLECMDIDF_REFRESH_LEVELMASK = 255
	OLECMDIDF_REFRESH_CLEARUSERINPUT = 4096
	OLECMDIDF_REFRESH_PROMPTIFOFFLINE = 8192
	OLECMDIDF_REFRESH_THROUGHSCRIPT = 16384
	OLECMDIDF_REFRESH_SKIPBEFOREUNLOADEVENT = 32768
	OLECMDIDF_REFRESH_PAGEACTION_ACTIVEXINSTALL = 65536
	OLECMDIDF_REFRESH_PAGEACTION_FILEDOWNLOAD = 131072
	OLECMDIDF_REFRESH_PAGEACTION_LOCALMACHINE = 262144
	OLECMDIDF_REFRESH_PAGEACTION_POPUPWINDOW = 524288
	OLECMDIDF_REFRESH_PAGEACTION_PROTLOCKDOWNLOCALMACHINE = 1048576
	OLECMDIDF_REFRESH_PAGEACTION_PROTLOCKDOWNTRUSTED = 2097152
	OLECMDIDF_REFRESH_PAGEACTION_PROTLOCKDOWNINTRANET = 4194304
	OLECMDIDF_REFRESH_PAGEACTION_PROTLOCKDOWNINTERNET = 8388608
	OLECMDIDF_REFRESH_PAGEACTION_PROTLOCKDOWNRESTRICTED = 16777216
	OLECMDIDF_REFRESH_PAGEACTION_MIXEDCONTENT = 33554432
	OLECMDIDF_REFRESH_PAGEACTION_INVALID_CERT = 67108864
	OLECMDIDF_REFRESH_PAGEACTION_ALLOW_VERSION = 134217728
)

type OLECMDID_PAGEACTIONFLAG int32

const (
	OLECMDIDF_PAGEACTION_FILEDOWNLOAD = 1
	OLECMDIDF_PAGEACTION_ACTIVEXINSTALL = 2
	OLECMDIDF_PAGEACTION_ACTIVEXTRUSTFAIL = 4
	OLECMDIDF_PAGEACTION_ACTIVEXUSERDISABLE = 8
	OLECMDIDF_PAGEACTION_ACTIVEXDISALLOW = 16
	OLECMDIDF_PAGEACTION_ACTIVEXUNSAFE = 32
	OLECMDIDF_PAGEACTION_POPUPWINDOW = 64
	OLECMDIDF_PAGEACTION_LOCALMACHINE = 128
	OLECMDIDF_PAGEACTION_MIMETEXTPLAIN = 256
	OLECMDIDF_PAGEACTION_SCRIPTNAVIGATE = 512
	OLECMDIDF_PAGEACTION_SCRIPTNAVIGATE_ACTIVEXINSTALL = 512
	OLECMDIDF_PAGEACTION_PROTLOCKDOWNLOCALMACHINE = 1024
	OLECMDIDF_PAGEACTION_PROTLOCKDOWNTRUSTED = 2048
	OLECMDIDF_PAGEACTION_PROTLOCKDOWNINTRANET = 4096
	OLECMDIDF_PAGEACTION_PROTLOCKDOWNINTERNET = 8192
	OLECMDIDF_PAGEACTION_PROTLOCKDOWNRESTRICTED = 16384
	OLECMDIDF_PAGEACTION_PROTLOCKDOWNDENY = 32768
	OLECMDIDF_PAGEACTION_POPUPALLOWED = 65536
	OLECMDIDF_PAGEACTION_SCRIPTPROMPT = 131072
	OLECMDIDF_PAGEACTION_ACTIVEXUSERAPPROVAL = 262144
	OLECMDIDF_PAGEACTION_MIXEDCONTENT = 524288
	OLECMDIDF_PAGEACTION_INVALID_CERT = 1048576
	OLECMDIDF_PAGEACTION_INTRANETZONEREQUEST = 2097152
	OLECMDIDF_PAGEACTION_XSSFILTERED = 4194304
	OLECMDIDF_PAGEACTION_SPOOFABLEIDNHOST = 8388608
	OLECMDIDF_PAGEACTION_ACTIVEX_EPM_INCOMPATIBLE = 16777216
	OLECMDIDF_PAGEACTION_SCRIPTNAVIGATE_ACTIVEXUSERAPPROVAL = 33554432
	OLECMDIDF_PAGEACTION_WPCBLOCKED = 67108864
	OLECMDIDF_PAGEACTION_WPCBLOCKED_ACTIVEX = 134217728
	OLECMDIDF_PAGEACTION_EXTENSION_COMPAT_BLOCKED = 268435456
	OLECMDIDF_PAGEACTION_NORESETACTIVEX = 536870912
	OLECMDIDF_PAGEACTION_GENERIC_STATE = 1073741824
	OLECMDIDF_PAGEACTION_RESET = -2147483648
)

type OLECMDID_BROWSERSTATEFLAG int32

const (
	OLECMDIDF_BROWSERSTATE_EXTENSIONSOFF = 1
	OLECMDIDF_BROWSERSTATE_IESECURITY = 2
	OLECMDIDF_BROWSERSTATE_PROTECTEDMODE_OFF = 4
	OLECMDIDF_BROWSERSTATE_RESET = 8
	OLECMDIDF_BROWSERSTATE_REQUIRESACTIVEX = 16
	OLECMDIDF_BROWSERSTATE_DESKTOPHTMLDIALOG = 32
	OLECMDIDF_BROWSERSTATE_BLOCKEDVERSION = 64
)

type OLECMDID_OPTICAL_ZOOMFLAG int32

const (
	OLECMDIDF_OPTICAL_ZOOM_NOPERSIST = 1
	OLECMDIDF_OPTICAL_ZOOM_NOLAYOUT = 16
	OLECMDIDF_OPTICAL_ZOOM_NOTRANSIENT = 32
	OLECMDIDF_OPTICAL_ZOOM_RELOADFORNEWTAB = 64
)

type PAGEACTION_UI int32

const (
	PAGEACTION_UI_DEFAULT = 0
	PAGEACTION_UI_MODAL = 1
	PAGEACTION_UI_MODELESS = 2
	PAGEACTION_UI_SILENT = 3
)

type OLECMDID_WINDOWSTATE_FLAG int32

const (
	OLECMDIDF_WINDOWSTATE_USERVISIBLE = 1
	OLECMDIDF_WINDOWSTATE_ENABLED = 2
	OLECMDIDF_WINDOWSTATE_USERVISIBLE_VALID = 65536
	OLECMDIDF_WINDOWSTATE_ENABLED_VALID = 131072
)

type OLECMDID_VIEWPORT_MODE_FLAG int32

const (
	OLECMDIDF_VIEWPORTMODE_FIXED_LAYOUT_WIDTH = 1
	OLECMDIDF_VIEWPORTMODE_EXCLUDE_VISUAL_BOTTOM = 2
	OLECMDIDF_VIEWPORTMODE_FIXED_LAYOUT_WIDTH_VALID = 65536
	OLECMDIDF_VIEWPORTMODE_EXCLUDE_VISUAL_BOTTOM_VALID = 131072
)

type OLEUIPASTEFLAG int32

const (
	OLEUIPASTE_ENABLEICON = 2048
	OLEUIPASTE_PASTEONLY = 0
	OLEUIPASTE_PASTE = 512
	OLEUIPASTE_LINKANYTYPE = 1024
	OLEUIPASTE_LINKTYPE1 = 1
	OLEUIPASTE_LINKTYPE2 = 2
	OLEUIPASTE_LINKTYPE3 = 4
	OLEUIPASTE_LINKTYPE4 = 8
	OLEUIPASTE_LINKTYPE5 = 16
	OLEUIPASTE_LINKTYPE6 = 32
	OLEUIPASTE_LINKTYPE7 = 64
	OLEUIPASTE_LINKTYPE8 = 128
)

type CALLFRAME_COPY int32

const (
	CALLFRAME_COPY_NESTED = 1
	CALLFRAME_COPY_INDEPENDENT = 2
)

type CALLFRAME_FREE int32

const (
	CALLFRAME_FREE_NONE = 0
	CALLFRAME_FREE_IN = 1
	CALLFRAME_FREE_INOUT = 2
	CALLFRAME_FREE_OUT = 4
	CALLFRAME_FREE_TOP_INOUT = 8
	CALLFRAME_FREE_TOP_OUT = 16
	CALLFRAME_FREE_ALL = 31
)

type CALLFRAME_NULL int32

const (
	CALLFRAME_NULL_NONE = 0
	CALLFRAME_NULL_INOUT = 2
	CALLFRAME_NULL_OUT = 4
	CALLFRAME_NULL_ALL = 6
)

type CALLFRAME_WALK int32

const (
	CALLFRAME_WALK_IN = 1
	CALLFRAME_WALK_INOUT = 2
	CALLFRAME_WALK_OUT = 4
)

type RECORD_READING_POLICY int32

const (
	RECORD_READING_POLICY_FORWARD = 1
	RECORD_READING_POLICY_BACKWARD = 2
	RECORD_READING_POLICY_RANDOM = 3
)

type COINIT int32

const (
	COINIT_APARTMENTTHREADED = 2
	COINIT_MULTITHREADED = 0
	COINIT_DISABLE_OLE1DDE = 4
	COINIT_SPEED_OVER_MEMORY = 8
)

type COMSD int32

const (
	SD_LAUNCHPERMISSIONS = 0
	SD_ACCESSPERMISSIONS = 1
	SD_LAUNCHRESTRICTIONS = 2
	SD_ACCESSRESTRICTIONS = 3
)

type EOC_ChangeType int32

const (
	EOC_NewObject = 0
	EOC_ModifiedObject = 1
	EOC_DeletedObject = 2
)

type DVASPECT int32

const (
	DVASPECT_CONTENT = 1
	DVASPECT_THUMBNAIL = 2
	DVASPECT_ICON = 4
	DVASPECT_DOCPRINT = 8
)

type TYSPEC int32

const (
	TYSPEC_CLSID = 0
	TYSPEC_FILEEXT = 1
	TYSPEC_MIMETYPE = 2
	TYSPEC_FILENAME = 3
	TYSPEC_PROGID = 4
	TYSPEC_PACKAGENAME = 5
	TYSPEC_OBJECTID = 6
)
