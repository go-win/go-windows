// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package iscsidisc implements the Windows.Win32.IScsiDisc namespace.
package iscsidisc

type NV_SEP_WRITE_CACHE_TYPE int32

const (
	NVSEPWriteCacheTypeUnknown = 0
	NVSEPWriteCacheTypeNone = 1
	NVSEPWriteCacheTypeWriteBack = 2
	NVSEPWriteCacheTypeWriteThrough = 3
)

type MP_STORAGE_DIAGNOSTIC_LEVEL int32

const (
	MpStorageDiagnosticLevelDefault = 0
	MpStorageDiagnosticLevelMax = 1
)

type MP_STORAGE_DIAGNOSTIC_TARGET_TYPE int32

const (
	MpStorageDiagnosticTargetTypeUndefined = 0
	MpStorageDiagnosticTargetTypeMiniport = 2
	MpStorageDiagnosticTargetTypeHbaFirmware = 3
	MpStorageDiagnosticTargetTypeMax = 4
)

type NVCACHE_TYPE int32

const (
	NvCacheTypeUnknown = 0
	NvCacheTypeNone = 1
	NvCacheTypeWriteBack = 2
	NvCacheTypeWriteThrough = 3
)

type NVCACHE_STATUS int32

const (
	NvCacheStatusUnknown = 0
	NvCacheStatusDisabling = 1
	NvCacheStatusDisabled = 2
	NvCacheStatusEnabled = 3
)

type ISCSI_DIGEST_TYPES int32

const (
	ISCSI_DIGEST_TYPE_NONE = 0
	ISCSI_DIGEST_TYPE_CRC32C = 1
)

type ISCSI_AUTH_TYPES int32

const (
	ISCSI_NO_AUTH_TYPE = 0
	ISCSI_CHAP_AUTH_TYPE = 1
	ISCSI_MUTUAL_CHAP_AUTH_TYPE = 2
)

type IKE_AUTHENTICATION_METHOD int32

const (
	IKE_AUTHENTICATION_PRESHARED_KEY_METHOD = 1
)

type TARGETPROTOCOLTYPE int32

const (
	ISCSI_TCP_PROTOCOL_TYPE = 0
)

type TARGET_INFORMATION_CLASS int32

const (
	ProtocolType = 0
	TargetAlias = 1
	DiscoveryMechanisms = 2
	PortalGroups = 3
	PersistentTargetMappings = 4
	InitiatorName = 5
	TargetFlags = 6
	LoginOptions = 7
)
