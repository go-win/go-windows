// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package menusandresources implements the Windows.Win32.MenusAndResources namespace.
package menusandresources

type DI_FLAGS uint32

const (
	DI_MASK = 1
	DI_IMAGE = 2
	DI_NORMAL = 3
	DI_COMPAT = 4
	DI_DEFAULTSIZE = 8
	DI_NOMIRROR = 16
)

type POINTER_INPUT_TYPE int32

const (
	PT_POINTER = 1
	PT_TOUCH = 2
	PT_PEN = 3
	PT_MOUSE = 4
	PT_TOUCHPAD = 5
)

type EDIT_CONTROL_FEATURE int32

const (
	EDIT_CONTROL_FEATURE_ENTERPRISE_DATA_PROTECTION_PASTE_SUPPORT = 0
	EDIT_CONTROL_FEATURE_PASTE_NOTIFICATIONS = 1
)

type HANDEDNESS int32

const (
	HANDEDNESS_LEFT = 0
	HANDEDNESS_RIGHT = 1
)

type MrmPlatformVersion int32

const (
	MrmPlatformVersion_Default = 0
	MrmPlatformVersion_Windows10_0_0_0 = 17432576
	MrmPlatformVersion_Windows10_0_0_5 = 17432581
)

type MrmPackagingMode int32

const (
	MrmPackagingModeStandaloneFile = 0
	MrmPackagingModeAutoSplit = 1
	MrmPackagingModeResourcePack = 2
)

type MrmPackagingOptions int32

const (
	MrmPackagingOptionsNone = 0
	MrmPackagingOptionsOmitSchemaFromResourcePacks = 1
	MrmPackagingOptionsSplitLanguageVariants = 2
)

type MrmDumpType int32

const (
	MrmDumpType_Basic = 0
	MrmDumpType_Detailed = 1
	MrmDumpType_Schema = 2
)

type MrmResourceIndexerMessageSeverity int32

const (
	MrmResourceIndexerMessageSeverityVerbose = 0
	MrmResourceIndexerMessageSeverityInfo = 1
	MrmResourceIndexerMessageSeverityWarning = 2
	MrmResourceIndexerMessageSeverityError = 3
)
