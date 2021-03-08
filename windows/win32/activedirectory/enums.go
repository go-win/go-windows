// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package activedirectory implements the Windows.Win32.ActiveDirectory namespace.
package activedirectory

type ADSTYPEENUM int32

const (
	ADSTYPE_INVALID = 0
	ADSTYPE_DN_STRING = 1
	ADSTYPE_CASE_EXACT_STRING = 2
	ADSTYPE_CASE_IGNORE_STRING = 3
	ADSTYPE_PRINTABLE_STRING = 4
	ADSTYPE_NUMERIC_STRING = 5
	ADSTYPE_BOOLEAN = 6
	ADSTYPE_INTEGER = 7
	ADSTYPE_OCTET_STRING = 8
	ADSTYPE_UTC_TIME = 9
	ADSTYPE_LARGE_INTEGER = 10
	ADSTYPE_PROV_SPECIFIC = 11
	ADSTYPE_OBJECT_CLASS = 12
	ADSTYPE_CASEIGNORE_LIST = 13
	ADSTYPE_OCTET_LIST = 14
	ADSTYPE_PATH = 15
	ADSTYPE_POSTALADDRESS = 16
	ADSTYPE_TIMESTAMP = 17
	ADSTYPE_BACKLINK = 18
	ADSTYPE_TYPEDNAME = 19
	ADSTYPE_HOLD = 20
	ADSTYPE_NETADDRESS = 21
	ADSTYPE_REPLICAPOINTER = 22
	ADSTYPE_FAXNUMBER = 23
	ADSTYPE_EMAIL = 24
	ADSTYPE_NT_SECURITY_DESCRIPTOR = 25
	ADSTYPE_UNKNOWN = 26
	ADSTYPE_DN_WITH_BINARY = 27
	ADSTYPE_DN_WITH_STRING = 28
)

type ADS_AUTHENTICATION_ENUM int32

const (
	ADS_SECURE_AUTHENTICATION = 1
	ADS_USE_ENCRYPTION = 2
	ADS_USE_SSL = 2
	ADS_READONLY_SERVER = 4
	ADS_PROMPT_CREDENTIALS = 8
	ADS_NO_AUTHENTICATION = 16
	ADS_FAST_BIND = 32
	ADS_USE_SIGNING = 64
	ADS_USE_SEALING = 128
	ADS_USE_DELEGATION = 256
	ADS_SERVER_BIND = 512
	ADS_NO_REFERRAL_CHASING = 1024
	ADS_AUTH_RESERVED = -2147483648
)

type ADS_STATUSENUM int32

const (
	ADS_STATUS_S_OK = 0
	ADS_STATUS_INVALID_SEARCHPREF = 1
	ADS_STATUS_INVALID_SEARCHPREFVALUE = 2
)

type ADS_DEREFENUM int32

const (
	ADS_DEREF_NEVER = 0
	ADS_DEREF_SEARCHING = 1
	ADS_DEREF_FINDING = 2
	ADS_DEREF_ALWAYS = 3
)

type ADS_SCOPEENUM int32

const (
	ADS_SCOPE_BASE = 0
	ADS_SCOPE_ONELEVEL = 1
	ADS_SCOPE_SUBTREE = 2
)

type ADS_PREFERENCES_ENUM int32

const (
	ADSIPROP_ASYNCHRONOUS = 0
	ADSIPROP_DEREF_ALIASES = 1
	ADSIPROP_SIZE_LIMIT = 2
	ADSIPROP_TIME_LIMIT = 3
	ADSIPROP_ATTRIBTYPES_ONLY = 4
	ADSIPROP_SEARCH_SCOPE = 5
	ADSIPROP_TIMEOUT = 6
	ADSIPROP_PAGESIZE = 7
	ADSIPROP_PAGED_TIME_LIMIT = 8
	ADSIPROP_CHASE_REFERRALS = 9
	ADSIPROP_SORT_ON = 10
	ADSIPROP_CACHE_RESULTS = 11
	ADSIPROP_ADSIFLAG = 12
)

type ADSI_DIALECT_ENUM int32

const (
	ADSI_DIALECT_LDAP = 0
	ADSI_DIALECT_SQL = 1
)

type ADS_CHASE_REFERRALS_ENUM int32

const (
	ADS_CHASE_REFERRALS_NEVER = 0
	ADS_CHASE_REFERRALS_SUBORDINATE = 32
	ADS_CHASE_REFERRALS_EXTERNAL = 64
	ADS_CHASE_REFERRALS_ALWAYS = 96
)

type ADS_SEARCHPREF_ENUM int32

const (
	ADS_SEARCHPREF_ASYNCHRONOUS = 0
	ADS_SEARCHPREF_DEREF_ALIASES = 1
	ADS_SEARCHPREF_SIZE_LIMIT = 2
	ADS_SEARCHPREF_TIME_LIMIT = 3
	ADS_SEARCHPREF_ATTRIBTYPES_ONLY = 4
	ADS_SEARCHPREF_SEARCH_SCOPE = 5
	ADS_SEARCHPREF_TIMEOUT = 6
	ADS_SEARCHPREF_PAGESIZE = 7
	ADS_SEARCHPREF_PAGED_TIME_LIMIT = 8
	ADS_SEARCHPREF_CHASE_REFERRALS = 9
	ADS_SEARCHPREF_SORT_ON = 10
	ADS_SEARCHPREF_CACHE_RESULTS = 11
	ADS_SEARCHPREF_DIRSYNC = 12
	ADS_SEARCHPREF_TOMBSTONE = 13
	ADS_SEARCHPREF_VLV = 14
	ADS_SEARCHPREF_ATTRIBUTE_QUERY = 15
	ADS_SEARCHPREF_SECURITY_MASK = 16
	ADS_SEARCHPREF_DIRSYNC_FLAG = 17
	ADS_SEARCHPREF_EXTENDED_DN = 18
)

type ADS_PASSWORD_ENCODING_ENUM int32

const (
	ADS_PASSWORD_ENCODE_REQUIRE_SSL = 0
	ADS_PASSWORD_ENCODE_CLEAR = 1
)

type ADS_PROPERTY_OPERATION_ENUM int32

const (
	ADS_PROPERTY_CLEAR = 1
	ADS_PROPERTY_UPDATE = 2
	ADS_PROPERTY_APPEND = 3
	ADS_PROPERTY_DELETE = 4
)

type ADS_SYSTEMFLAG_ENUM int32

const (
	ADS_SYSTEMFLAG_DISALLOW_DELETE = -2147483648
	ADS_SYSTEMFLAG_CONFIG_ALLOW_RENAME = 1073741824
	ADS_SYSTEMFLAG_CONFIG_ALLOW_MOVE = 536870912
	ADS_SYSTEMFLAG_CONFIG_ALLOW_LIMITED_MOVE = 268435456
	ADS_SYSTEMFLAG_DOMAIN_DISALLOW_RENAME = 134217728
	ADS_SYSTEMFLAG_DOMAIN_DISALLOW_MOVE = 67108864
	ADS_SYSTEMFLAG_CR_NTDS_NC = 1
	ADS_SYSTEMFLAG_CR_NTDS_DOMAIN = 2
	ADS_SYSTEMFLAG_ATTR_NOT_REPLICATED = 1
	ADS_SYSTEMFLAG_ATTR_IS_CONSTRUCTED = 4
)

type ADS_GROUP_TYPE_ENUM int32

const (
	ADS_GROUP_TYPE_GLOBAL_GROUP = 2
	ADS_GROUP_TYPE_DOMAIN_LOCAL_GROUP = 4
	ADS_GROUP_TYPE_LOCAL_GROUP = 4
	ADS_GROUP_TYPE_UNIVERSAL_GROUP = 8
	ADS_GROUP_TYPE_SECURITY_ENABLED = -2147483648
)

type ADS_USER_FLAG_ENUM int32

const (
	ADS_UF_SCRIPT = 1
	ADS_UF_ACCOUNTDISABLE = 2
	ADS_UF_HOMEDIR_REQUIRED = 8
	ADS_UF_LOCKOUT = 16
	ADS_UF_PASSWD_NOTREQD = 32
	ADS_UF_PASSWD_CANT_CHANGE = 64
	ADS_UF_ENCRYPTED_TEXT_PASSWORD_ALLOWED = 128
	ADS_UF_TEMP_DUPLICATE_ACCOUNT = 256
	ADS_UF_NORMAL_ACCOUNT = 512
	ADS_UF_INTERDOMAIN_TRUST_ACCOUNT = 2048
	ADS_UF_WORKSTATION_TRUST_ACCOUNT = 4096
	ADS_UF_SERVER_TRUST_ACCOUNT = 8192
	ADS_UF_DONT_EXPIRE_PASSWD = 65536
	ADS_UF_MNS_LOGON_ACCOUNT = 131072
	ADS_UF_SMARTCARD_REQUIRED = 262144
	ADS_UF_TRUSTED_FOR_DELEGATION = 524288
	ADS_UF_NOT_DELEGATED = 1048576
	ADS_UF_USE_DES_KEY_ONLY = 2097152
	ADS_UF_DONT_REQUIRE_PREAUTH = 4194304
	ADS_UF_PASSWORD_EXPIRED = 8388608
	ADS_UF_TRUSTED_TO_AUTHENTICATE_FOR_DELEGATION = 16777216
)

type ADS_RIGHTS_ENUM int32

const (
	ADS_RIGHT_DELETE = 65536
	ADS_RIGHT_READ_CONTROL = 131072
	ADS_RIGHT_WRITE_DAC = 262144
	ADS_RIGHT_WRITE_OWNER = 524288
	ADS_RIGHT_SYNCHRONIZE = 1048576
	ADS_RIGHT_ACCESS_SYSTEM_SECURITY = 16777216
	ADS_RIGHT_GENERIC_READ = -2147483648
	ADS_RIGHT_GENERIC_WRITE = 1073741824
	ADS_RIGHT_GENERIC_EXECUTE = 536870912
	ADS_RIGHT_GENERIC_ALL = 268435456
	ADS_RIGHT_DS_CREATE_CHILD = 1
	ADS_RIGHT_DS_DELETE_CHILD = 2
	ADS_RIGHT_ACTRL_DS_LIST = 4
	ADS_RIGHT_DS_SELF = 8
	ADS_RIGHT_DS_READ_PROP = 16
	ADS_RIGHT_DS_WRITE_PROP = 32
	ADS_RIGHT_DS_DELETE_TREE = 64
	ADS_RIGHT_DS_LIST_OBJECT = 128
	ADS_RIGHT_DS_CONTROL_ACCESS = 256
)

type ADS_ACETYPE_ENUM int32

const (
	ADS_ACETYPE_ACCESS_ALLOWED = 0
	ADS_ACETYPE_ACCESS_DENIED = 1
	ADS_ACETYPE_SYSTEM_AUDIT = 2
	ADS_ACETYPE_ACCESS_ALLOWED_OBJECT = 5
	ADS_ACETYPE_ACCESS_DENIED_OBJECT = 6
	ADS_ACETYPE_SYSTEM_AUDIT_OBJECT = 7
	ADS_ACETYPE_SYSTEM_ALARM_OBJECT = 8
	ADS_ACETYPE_ACCESS_ALLOWED_CALLBACK = 9
	ADS_ACETYPE_ACCESS_DENIED_CALLBACK = 10
	ADS_ACETYPE_ACCESS_ALLOWED_CALLBACK_OBJECT = 11
	ADS_ACETYPE_ACCESS_DENIED_CALLBACK_OBJECT = 12
	ADS_ACETYPE_SYSTEM_AUDIT_CALLBACK = 13
	ADS_ACETYPE_SYSTEM_ALARM_CALLBACK = 14
	ADS_ACETYPE_SYSTEM_AUDIT_CALLBACK_OBJECT = 15
	ADS_ACETYPE_SYSTEM_ALARM_CALLBACK_OBJECT = 16
)

type ADS_ACEFLAG_ENUM int32

const (
	ADS_ACEFLAG_INHERIT_ACE = 2
	ADS_ACEFLAG_NO_PROPAGATE_INHERIT_ACE = 4
	ADS_ACEFLAG_INHERIT_ONLY_ACE = 8
	ADS_ACEFLAG_INHERITED_ACE = 16
	ADS_ACEFLAG_VALID_INHERIT_FLAGS = 31
	ADS_ACEFLAG_SUCCESSFUL_ACCESS = 64
	ADS_ACEFLAG_FAILED_ACCESS = 128
)

type ADS_FLAGTYPE_ENUM int32

const (
	ADS_FLAG_OBJECT_TYPE_PRESENT = 1
	ADS_FLAG_INHERITED_OBJECT_TYPE_PRESENT = 2
)

type ADS_SD_CONTROL_ENUM int32

const (
	ADS_SD_CONTROL_SE_OWNER_DEFAULTED = 1
	ADS_SD_CONTROL_SE_GROUP_DEFAULTED = 2
	ADS_SD_CONTROL_SE_DACL_PRESENT = 4
	ADS_SD_CONTROL_SE_DACL_DEFAULTED = 8
	ADS_SD_CONTROL_SE_SACL_PRESENT = 16
	ADS_SD_CONTROL_SE_SACL_DEFAULTED = 32
	ADS_SD_CONTROL_SE_DACL_AUTO_INHERIT_REQ = 256
	ADS_SD_CONTROL_SE_SACL_AUTO_INHERIT_REQ = 512
	ADS_SD_CONTROL_SE_DACL_AUTO_INHERITED = 1024
	ADS_SD_CONTROL_SE_SACL_AUTO_INHERITED = 2048
	ADS_SD_CONTROL_SE_DACL_PROTECTED = 4096
	ADS_SD_CONTROL_SE_SACL_PROTECTED = 8192
	ADS_SD_CONTROL_SE_SELF_RELATIVE = 32768
)

type ADS_SD_REVISION_ENUM int32

const (
	ADS_SD_REVISION_DS = 4
)

type ADS_NAME_TYPE_ENUM int32

const (
	ADS_NAME_TYPE_1779 = 1
	ADS_NAME_TYPE_CANONICAL = 2
	ADS_NAME_TYPE_NT4 = 3
	ADS_NAME_TYPE_DISPLAY = 4
	ADS_NAME_TYPE_DOMAIN_SIMPLE = 5
	ADS_NAME_TYPE_ENTERPRISE_SIMPLE = 6
	ADS_NAME_TYPE_GUID = 7
	ADS_NAME_TYPE_UNKNOWN = 8
	ADS_NAME_TYPE_USER_PRINCIPAL_NAME = 9
	ADS_NAME_TYPE_CANONICAL_EX = 10
	ADS_NAME_TYPE_SERVICE_PRINCIPAL_NAME = 11
	ADS_NAME_TYPE_SID_OR_SID_HISTORY_NAME = 12
)

type ADS_NAME_INITTYPE_ENUM int32

const (
	ADS_NAME_INITTYPE_DOMAIN = 1
	ADS_NAME_INITTYPE_SERVER = 2
	ADS_NAME_INITTYPE_GC = 3
)

type ADS_OPTION_ENUM int32

const (
	ADS_OPTION_SERVERNAME = 0
	ADS_OPTION_REFERRALS = 1
	ADS_OPTION_PAGE_SIZE = 2
	ADS_OPTION_SECURITY_MASK = 3
	ADS_OPTION_MUTUAL_AUTH_STATUS = 4
	ADS_OPTION_QUOTA = 5
	ADS_OPTION_PASSWORD_PORTNUMBER = 6
	ADS_OPTION_PASSWORD_METHOD = 7
	ADS_OPTION_ACCUMULATIVE_MODIFICATION = 8
	ADS_OPTION_SKIP_SID_LOOKUP = 9
)

type ADS_SECURITY_INFO_ENUM int32

const (
	ADS_SECURITY_INFO_OWNER = 1
	ADS_SECURITY_INFO_GROUP = 2
	ADS_SECURITY_INFO_DACL = 4
	ADS_SECURITY_INFO_SACL = 8
)

type ADS_SETTYPE_ENUM int32

const (
	ADS_SETTYPE_FULL = 1
	ADS_SETTYPE_PROVIDER = 2
	ADS_SETTYPE_SERVER = 3
	ADS_SETTYPE_DN = 4
)

type ADS_FORMAT_ENUM int32

const (
	ADS_FORMAT_WINDOWS = 1
	ADS_FORMAT_WINDOWS_NO_SERVER = 2
	ADS_FORMAT_WINDOWS_DN = 3
	ADS_FORMAT_WINDOWS_PARENT = 4
	ADS_FORMAT_X500 = 5
	ADS_FORMAT_X500_NO_SERVER = 6
	ADS_FORMAT_X500_DN = 7
	ADS_FORMAT_X500_PARENT = 8
	ADS_FORMAT_SERVER = 9
	ADS_FORMAT_PROVIDER = 10
	ADS_FORMAT_LEAF = 11
)

type ADS_DISPLAY_ENUM int32

const (
	ADS_DISPLAY_FULL = 1
	ADS_DISPLAY_VALUE_ONLY = 2
)

type ADS_ESCAPE_MODE_ENUM int32

const (
	ADS_ESCAPEDMODE_DEFAULT = 1
	ADS_ESCAPEDMODE_ON = 2
	ADS_ESCAPEDMODE_OFF = 3
	ADS_ESCAPEDMODE_OFF_EX = 4
)

type ADS_PATHTYPE_ENUM int32

const (
	ADS_PATH_FILE = 1
	ADS_PATH_FILESHARE = 2
	ADS_PATH_REGISTRY = 3
)

type ADS_SD_FORMAT_ENUM int32

const (
	ADS_SD_FORMAT_IID = 1
	ADS_SD_FORMAT_RAW = 2
	ADS_SD_FORMAT_HEXSTRING = 3
)

type DS_MANGLE_FOR int32

const (
	DS_MANGLE_UNKNOWN = 0
	DS_MANGLE_OBJECT_RDN_FOR_DELETION = 1
	DS_MANGLE_OBJECT_RDN_FOR_NAME_CONFLICT = 2
)

type DS_NAME_FORMAT int32

const (
	DS_UNKNOWN_NAME = 0
	DS_FQDN_1779_NAME = 1
	DS_NT4_ACCOUNT_NAME = 2
	DS_DISPLAY_NAME = 3
	DS_UNIQUE_ID_NAME = 6
	DS_CANONICAL_NAME = 7
	DS_USER_PRINCIPAL_NAME = 8
	DS_CANONICAL_NAME_EX = 9
	DS_SERVICE_PRINCIPAL_NAME = 10
	DS_SID_OR_SID_HISTORY_NAME = 11
	DS_DNS_DOMAIN_NAME = 12
)

type DS_NAME_FLAGS int32

const (
	DS_NAME_NO_FLAGS = 0
	DS_NAME_FLAG_SYNTACTICAL_ONLY = 1
	DS_NAME_FLAG_EVAL_AT_DC = 2
	DS_NAME_FLAG_GCVERIFY = 4
	DS_NAME_FLAG_TRUST_REFERRAL = 8
)

type DS_NAME_ERROR int32

const (
	DS_NAME_NO_ERROR = 0
	DS_NAME_ERROR_RESOLVING = 1
	DS_NAME_ERROR_NOT_FOUND = 2
	DS_NAME_ERROR_NOT_UNIQUE = 3
	DS_NAME_ERROR_NO_MAPPING = 4
	DS_NAME_ERROR_DOMAIN_ONLY = 5
	DS_NAME_ERROR_NO_SYNTACTICAL_MAPPING = 6
	DS_NAME_ERROR_TRUST_REFERRAL = 7
)

type DS_SPN_NAME_TYPE int32

const (
	DS_SPN_DNS_HOST = 0
	DS_SPN_DN_HOST = 1
	DS_SPN_NB_HOST = 2
	DS_SPN_DOMAIN = 3
	DS_SPN_NB_DOMAIN = 4
	DS_SPN_SERVICE = 5
)

type DS_SPN_WRITE_OP int32

const (
	DS_SPN_ADD_SPN_OP = 0
	DS_SPN_REPLACE_SPN_OP = 1
	DS_SPN_DELETE_SPN_OP = 2
)

type DS_REPSYNCALL_ERROR int32

const (
	DS_REPSYNCALL_WIN32_ERROR_CONTACTING_SERVER = 0
	DS_REPSYNCALL_WIN32_ERROR_REPLICATING = 1
	DS_REPSYNCALL_SERVER_UNREACHABLE = 2
)

type DS_REPSYNCALL_EVENT int32

const (
	DS_REPSYNCALL_EVENT_ERROR = 0
	DS_REPSYNCALL_EVENT_SYNC_STARTED = 1
	DS_REPSYNCALL_EVENT_SYNC_COMPLETED = 2
	DS_REPSYNCALL_EVENT_FINISHED = 3
)

type DS_KCC_TASKID int32

const (
	DS_KCC_TASKID_UPDATE_TOPOLOGY = 0
)

type DS_REPL_INFO_TYPE int32

const (
	DS_REPL_INFO_NEIGHBORS = 0
	DS_REPL_INFO_CURSORS_FOR_NC = 1
	DS_REPL_INFO_METADATA_FOR_OBJ = 2
	DS_REPL_INFO_KCC_DSA_CONNECT_FAILURES = 3
	DS_REPL_INFO_KCC_DSA_LINK_FAILURES = 4
	DS_REPL_INFO_PENDING_OPS = 5
	DS_REPL_INFO_METADATA_FOR_ATTR_VALUE = 6
	DS_REPL_INFO_CURSORS_2_FOR_NC = 7
	DS_REPL_INFO_CURSORS_3_FOR_NC = 8
	DS_REPL_INFO_METADATA_2_FOR_OBJ = 9
	DS_REPL_INFO_METADATA_2_FOR_ATTR_VALUE = 10
	DS_REPL_INFO_METADATA_EXT_FOR_ATTR_VALUE = 11
	DS_REPL_INFO_TYPE_MAX = 12
)

type DS_REPL_OP_TYPE int32

const (
	DS_REPL_OP_TYPE_SYNC = 0
	DS_REPL_OP_TYPE_ADD = 1
	DS_REPL_OP_TYPE_DELETE = 2
	DS_REPL_OP_TYPE_MODIFY = 3
	DS_REPL_OP_TYPE_UPDATE_REFS = 4
)

type DSROLE_MACHINE_ROLE int32

const (
	DsRole_RoleStandaloneWorkstation = 0
	DsRole_RoleMemberWorkstation = 1
	DsRole_RoleStandaloneServer = 2
	DsRole_RoleMemberServer = 3
	DsRole_RoleBackupDomainController = 4
	DsRole_RolePrimaryDomainController = 5
)

type DSROLE_SERVER_STATE int32

const (
	DsRoleServerUnknown = 0
	DsRoleServerPrimary = 1
	DsRoleServerBackup = 2
)

type DSROLE_PRIMARY_DOMAIN_INFO_LEVEL int32

const (
	DsRolePrimaryDomainInfoBasic = 1
	DsRoleUpgradeStatus = 2
	DsRoleOperationState = 3
)

type DSROLE_OPERATION_STATE int32

const (
	DsRoleOperationIdle = 0
	DsRoleOperationActive = 1
	DsRoleOperationNeedReboot = 2
)
