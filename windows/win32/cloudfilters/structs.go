// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package cloudfilters implements the Windows.Win32.CloudFilters namespace.
package cloudfilters

type CF_CONNECTION_KEY__ struct {
	Internal int
}

type CF_FS_METADATA struct {
	BasicInfo int
	FileSize  int
}

type CF_PLACEHOLDER_CREATE_INFO struct {
	RelativeFileName   int
	FsMetadata         int
	FileIdentity       int
	FileIdentityLength int
	Flags              int
	Result             int
	CreateUsn          int
}

type CF_PROCESS_INFO struct {
	StructSize    int
	ProcessId     int
	ImagePath     int
	PackageName   int
	ApplicationId int
	CommandLine   int
	SessionId     int
}

type CF_PLATFORM_INFO struct {
	BuildNumber       int
	RevisionNumber    int
	IntegrationNumber int
}

type CF_HYDRATION_POLICY_PRIMARY_USHORT struct {
	US int
}

type CF_HYDRATION_POLICY_MODIFIER_USHORT struct {
	US int
}

type CF_HYDRATION_POLICY struct {
	Primary  int
	Modifier int
}

type CF_POPULATION_POLICY_PRIMARY_USHORT struct {
	US int
}

type CF_POPULATION_POLICY_MODIFIER_USHORT struct {
	US int
}

type CF_POPULATION_POLICY struct {
	Primary  int
	Modifier int
}

type CF_SYNC_POLICIES struct {
	StructSize            int
	Hydration             int
	Population            int
	InSync                int
	HardLink              int
	PlaceholderManagement int
}

type CF_SYNC_REGISTRATION struct {
	StructSize             int
	ProviderName           int
	ProviderVersion        int
	SyncRootIdentity       int
	SyncRootIdentityLength int
	FileIdentity           int
	FileIdentityLength     int
	ProviderId             int
}

type CF_CALLBACK_INFO struct {
	StructSize             int
	ConnectionKey          int
	CallbackContext        int
	VolumeGuidName         int
	VolumeDosName          int
	VolumeSerialNumber     int
	SyncRootFileId         int
	SyncRootIdentity       int
	SyncRootIdentityLength int
	FileId                 int
	FileSize               int
	FileIdentity           int
	FileIdentityLength     int
	NormalizedPath         int
	TransferKey            int
	PriorityHint           int
	CorrelationVector      int
	ProcessInfo            int
	RequestKey             int
}

type CF_CALLBACK_PARAMETERS struct {
	ParamSize int
	Anonymous int
}

type CF_CALLBACK_REGISTRATION struct {
	Type     int
	Callback int
}

type CF_SYNC_STATUS struct {
	StructSize        int
	Code              int
	DescriptionOffset int
	DescriptionLength int
	DeviceIdOffset    int
	DeviceIdLength    int
}

type CF_OPERATION_INFO struct {
	StructSize        int
	Type              int
	ConnectionKey     int
	TransferKey       int
	CorrelationVector int
	SyncStatus        int
	RequestKey        int
}

type CF_OPERATION_PARAMETERS struct {
	ParamSize int
	Anonymous int
}

type CF_FILE_RANGE struct {
	StartingOffset int
	Length         int
}

type CF_PLACEHOLDER_BASIC_INFO struct {
	PinState           int
	InSyncState        int
	FileId             int
	SyncRootFileId     int
	FileIdentityLength int
	FileIdentity       int
}

type CF_PLACEHOLDER_STANDARD_INFO struct {
	OnDiskDataSize     int
	ValidatedDataSize  int
	ModifiedDataSize   int
	PropertiesSize     int
	PinState           int
	InSyncState        int
	FileId             int
	SyncRootFileId     int
	FileIdentityLength int
	FileIdentity       int
}

type CF_SYNC_ROOT_BASIC_INFO struct {
	SyncRootFileId int
}

type CF_SYNC_ROOT_PROVIDER_INFO struct {
	ProviderStatus  int
	ProviderName    int
	ProviderVersion int
}

type CF_SYNC_ROOT_STANDARD_INFO struct {
	SyncRootFileId         int
	HydrationPolicy        int
	PopulationPolicy       int
	InSyncPolicy           int
	HardLinkPolicy         int
	ProviderStatus         int
	ProviderName           int
	ProviderVersion        int
	SyncRootIdentityLength int
	SyncRootIdentity       int
}