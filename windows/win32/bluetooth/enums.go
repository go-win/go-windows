// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package bluetooth implements the Windows.Win32.Bluetooth namespace.
package bluetooth

type NodeContainerType int32

const (
	NodeContainerTypeSequence = 0
	NodeContainerTypeAlternative = 1
)

type SDP_TYPE int32

const (
	SDP_TYPE_NIL = 0
	SDP_TYPE_UINT = 1
	SDP_TYPE_INT = 2
	SDP_TYPE_UUID = 3
	SDP_TYPE_STRING = 4
	SDP_TYPE_BOOLEAN = 5
	SDP_TYPE_SEQUENCE = 6
	SDP_TYPE_ALTERNATIVE = 7
	SDP_TYPE_URL = 8
	SDP_TYPE_CONTAINER = 32
)

type SDP_SPECIFICTYPE int32

const (
	SDP_ST_NONE = 0
	SDP_ST_UINT8 = 16
	SDP_ST_UINT16 = 272
	SDP_ST_UINT32 = 528
	SDP_ST_UINT64 = 784
	SDP_ST_UINT128 = 1040
	SDP_ST_INT8 = 32
	SDP_ST_INT16 = 288
	SDP_ST_INT32 = 544
	SDP_ST_INT64 = 800
	SDP_ST_INT128 = 1056
	SDP_ST_UUID16 = 304
	SDP_ST_UUID32 = 544
	SDP_ST_UUID128 = 1072
)

type IO_CAPABILITY int32

const (
	IoCaps_DisplayOnly = 0
	IoCaps_DisplayYesNo = 1
	IoCaps_KeyboardOnly = 2
	IoCaps_NoInputNoOutput = 3
	IoCaps_Undefined = 255
)

type AUTHENTICATION_REQUIREMENTS int32

const (
	MITMProtectionNotRequired = 0
	MITMProtectionRequired = 1
	MITMProtectionNotRequiredBonding = 2
	MITMProtectionRequiredBonding = 3
	MITMProtectionNotRequiredGeneralBonding = 4
	MITMProtectionRequiredGeneralBonding = 5
	MITMProtectionNotDefined = 255
)

type BLUETOOTH_AUTHENTICATION_METHOD int32

const (
	BLUETOOTH_AUTHENTICATION_METHOD_LEGACY = 1
	BLUETOOTH_AUTHENTICATION_METHOD_OOB = 2
	BLUETOOTH_AUTHENTICATION_METHOD_NUMERIC_COMPARISON = 3
	BLUETOOTH_AUTHENTICATION_METHOD_PASSKEY_NOTIFICATION = 4
	BLUETOOTH_AUTHENTICATION_METHOD_PASSKEY = 5
)

type BLUETOOTH_IO_CAPABILITY int32

const (
	BLUETOOTH_IO_CAPABILITY_DISPLAYONLY = 0
	BLUETOOTH_IO_CAPABILITY_DISPLAYYESNO = 1
	BLUETOOTH_IO_CAPABILITY_KEYBOARDONLY = 2
	BLUETOOTH_IO_CAPABILITY_NOINPUTNOOUTPUT = 3
	BLUETOOTH_IO_CAPABILITY_UNDEFINED = 255
)

type BLUETOOTH_AUTHENTICATION_REQUIREMENTS int32

const (
	BLUETOOTH_MITM_ProtectionNotRequired = 0
	BLUETOOTH_MITM_ProtectionRequired = 1
	BLUETOOTH_MITM_ProtectionNotRequiredBonding = 2
	BLUETOOTH_MITM_ProtectionRequiredBonding = 3
	BLUETOOTH_MITM_ProtectionNotRequiredGeneralBonding = 4
	BLUETOOTH_MITM_ProtectionRequiredGeneralBonding = 5
	BLUETOOTH_MITM_ProtectionNotDefined = 255
)
