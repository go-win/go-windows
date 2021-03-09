// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsclustering implements the Windows.Win32.WindowsClustering namespace.
package windowsclustering

type HCLUSTER struct {
}

type HNODE struct {
}

type HRESOURCE struct {
}

type HGROUP struct {
}

type HNETWORK struct {
}

type HNETINTERFACE struct {
}

type HCHANGE struct {
}

type HCLUSENUM struct {
}

type HGROUPENUM struct {
}

type HRESENUM struct {
}

type HNETWORKENUM struct {
}

type HNODEENUM struct {
}

type HNETINTERFACEENUM struct {
}

type HRESTYPEENUM struct {
}

type HREGBATCH struct {
}

type HREGBATCHPORT struct {
}

type HREGBATCHNOTIFICATION struct {
}

type HREGREADBATCH struct {
}

type HREGREADBATCHREPLY struct {
}

type HNODEENUMEX struct {
}

type HCLUSENUMEX struct {
}

type HGROUPENUMEX struct {
}

type HRESENUMEX struct {
}

type HGROUPSET struct {
}

type HGROUPSETENUM struct {
}

type CLUSTERVERSIONINFO_NT4 struct {
	DWVERSIONINFOSIZE int
	MajorVersion      int
	MinorVersion      int
	BuildNumber       int
	SZVENDORID        int
	SZCSDVERSION      int
}

type CLUSTERVERSIONINFO struct {
	DWVERSIONINFOSIZE       int
	MajorVersion            int
	MinorVersion            int
	BuildNumber             int
	SZVENDORID              int
	SZCSDVERSION            int
	DWCLUSTERHIGHESTVERSION int
	DWCLUSTERLOWESTVERSION  int
	DWFLAGS                 int
	DWRESERVED              int
}

type CLUS_STARTING_PARAMS struct {
	DWSIZE int
	BFORM  int
	BFIRST int
}

type CLUSCTL_RESOURCE_STATE_CHANGE_REASON_STRUCT struct {
	DWSIZE    int
	DWVERSION int
	EREASON   int
}

type CLUSTER_BATCH_COMMAND struct {
	Command   int
	DWOPTIONS int
	WZNAME    int
	LPDATA    int
	CBDATA    int
}

type CLUSTER_READ_BATCH_COMMAND struct {
	Command      int
	DWOPTIONS    int
	WZSUBKEYNAME int
	WZVALUENAME  int
	LPDATA       int
	CBDATA       int
}

type CLUSTER_ENUM_ITEM struct {
	DWVERSION int
	DWTYPE    int
	CBID      int
	LPSZID    int
	CBNAME    int
	LPSZNAME  int
}

type CLUSTER_CREATE_GROUP_INFO struct {
	DWVERSION int
	GROUPTYPE int
}

type CLUSTER_VALIDATE_PATH struct {
	SZPATH int
}

type CLUSTER_VALIDATE_DIRECTORY struct {
	SZPATH int
}

type CLUSTER_VALIDATE_NETNAME struct {
	SZNETWORKNAME int
}

type CLUSTER_VALIDATE_CSV_FILENAME struct {
	SZFILENAME int
}

type CLUSTER_SET_PASSWORD_STATUS struct {
	NodeId       int
	SetAttempted int
	ReturnStatus int
}

type CLUSTER_IP_ENTRY struct {
	LPSZIPADDRESS  int
	DWPREFIXLENGTH int
}

type CREATE_CLUSTER_CONFIG struct {
	DWVERSION              int
	LPSZCLUSTERNAME        int
	CNODES                 int
	PPSZNODENAMES          int
	CIPENTRIES             int
	PIPENTRIES             int
	FEMPTYCLUSTER          int
	MANAGEMENTPOINTTYPE    int
	MANAGEMENTPOINTRESTYPE int
}

type CREATE_CLUSTER_NAME_ACCOUNT struct {
	DWVERSION              int
	LPSZCLUSTERNAME        int
	DWFLAGS                int
	PSZUSERNAME            int
	PSZPASSWORD            int
	PSZDOMAIN              int
	MANAGEMENTPOINTTYPE    int
	MANAGEMENTPOINTRESTYPE int
	BUPGRADEVCOS           int
}

type NOTIFY_FILTER_AND_TYPE struct {
	DWOBJECTTYPE int
	FilterFlags  int
}

type CLUSTER_MEMBERSHIP_INFO struct {
	HasQuorum   int
	UpnodesSize int
	Upnodes     int
}

type CLUSTER_AVAILABILITY_SET_CONFIG struct {
	DWVERSION         int
	DWUPDATEDOMAINS   int
	DWFAULTDOMAINS    int
	BRESERVESPARENODE int
}

type CLUSTER_GROUP_ENUM_ITEM struct {
	DWVERSION      int
	CBID           int
	LPSZID         int
	CBNAME         int
	LPSZNAME       int
	STATE          int
	CBOWNERNODE    int
	LPSZOWNERNODE  int
	DWFLAGS        int
	CBPROPERTIES   int
	PPROPERTIES    int
	CBROPROPERTIES int
	PROPROPERTIES  int
}

type CLUSTER_RESOURCE_ENUM_ITEM struct {
	DWVERSION          int
	CBID               int
	LPSZID             int
	CBNAME             int
	LPSZNAME           int
	CBOWNERGROUPNAME   int
	LPSZOWNERGROUPNAME int
	CBOWNERGROUPID     int
	LPSZOWNERGROUPID   int
	CBPROPERTIES       int
	PPROPERTIES        int
	CBROPROPERTIES     int
	PROPROPERTIES      int
}

type GROUP_FAILURE_INFO struct {
	DWFAILOVERATTEMPTSREMAINING int
	DWFAILOVERPERIODREMAINING   int
}

type GROUP_FAILURE_INFO_BUFFER struct {
	DWVERSION int
	Info      int
}

type RESOURCE_FAILURE_INFO struct {
	DWRESTARTATTEMPTSREMAINING int
	DWRESTARTPERIODREMAINING   int
}

type RESOURCE_FAILURE_INFO_BUFFER struct {
	DWVERSION int
	Info      int
}

type RESOURCE_TERMINAL_FAILURE_INFO_BUFFER struct {
	ISTERMINALFAILURE      int
	RESTARTPERIODREMAINING int
}

type CLUSPROP_SYNTAX struct {
	DW        int
	Anonymous int
}

type CLUSPROP_VALUE struct {
	Syntax   int
	CBLENGTH int
}

type CLUSPROP_BINARY struct {
	AnonymousBase_clusapi_L5092_C41 int
	RGB                             int
}

type CLUSPROP_WORD struct {
	AnonymousBase_clusapi_L5102_C39 int
	W                               int
}

type CLUSPROP_DWORD struct {
	AnonymousBase_clusapi_L5112_C40 int
	DW                              int
}

type CLUSPROP_LONG struct {
	AnonymousBase_clusapi_L5122_C39 int
	L                               int
}

type CLUSPROP_SZ struct {
	AnonymousBase_clusapi_L5132_C37 int
	SZ                              int
}

type CLUSPROP_ULARGE_INTEGER struct {
	AnonymousBase_clusapi_L5149_C14 int
	LI                              int
}

type CLUSPROP_LARGE_INTEGER struct {
	AnonymousBase_clusapi_L5162_C14 int
	LI                              int
}

type CLUSPROP_SECURITY_DESCRIPTOR struct {
	AnonymousBase_clusapi_L5174_C54 int
	Anonymous                       int
}

type CLUSPROP_FILETIME struct {
	AnonymousBase_clusapi_L5188_C14 int
	FT                              int
}

type CLUS_RESOURCE_CLASS_INFO struct {
	Anonymous int
}

type CLUSPROP_RESOURCE_CLASS struct {
	AnonymousBase_clusapi_L5213_C14 int
	RC                              int
}

type CLUSPROP_RESOURCE_CLASS_INFO struct {
	AnonymousBase_clusapi_L5224_C14 int
	AnonymousBase_clusapi_L5225_C14 int
}

type CLUSPROP_REQUIRED_DEPENDENCY struct {
	Value       int
	ResClass    int
	ResTypeName int
}

type CLUS_FORCE_QUORUM_INFO struct {
	DWSIZE             int
	DWNODEBITMASK      int
	DWMAXNUMBEROFNODES int
	MULTISZNODELIST    int
}

type CLUS_PARTITION_INFO struct {
	DWFLAGS                    int
	SZDEVICENAME               int
	SZVOLUMELABEL              int
	DWSERIALNUMBER             int
	RGDWMAXIMUMCOMPONENTLENGTH int
	DWFILESYSTEMFLAGS          int
	SZFILESYSTEM               int
}

type CLUS_PARTITION_INFO_EX struct {
	DWFLAGS                    int
	SZDEVICENAME               int
	SZVOLUMELABEL              int
	DWSERIALNUMBER             int
	RGDWMAXIMUMCOMPONENTLENGTH int
	DWFILESYSTEMFLAGS          int
	SZFILESYSTEM               int
	TotalSizeInBytes           int
	FreeSizeInBytes            int
	DeviceNumber               int
	PartitionNumber            int
	VolumeGuid                 int
}

type CLUS_PARTITION_INFO_EX2 struct {
	GptPartitionId  int
	SZPARTITIONNAME int
	EncryptionFlags int
}

type CLUS_CSV_VOLUME_INFO struct {
	VolumeOffset         int
	PartitionNumber      int
	FaultState           int
	BackupState          int
	SZVOLUMEFRIENDLYNAME int
	SZVOLUMENAME         int
}

type CLUS_CSV_VOLUME_NAME struct {
	VolumeOffset int
	SZVOLUMENAME int
	SZROOTPATH   int
}

type CLUSTER_SHARED_VOLUME_STATE_INFO struct {
	SZVOLUMENAME int
	SZNODENAME   int
	VolumeState  int
}

type CLUSTER_SHARED_VOLUME_STATE_INFO_EX struct {
	SZVOLUMENAME             int
	SZNODENAME               int
	VolumeState              int
	SZVOLUMEFRIENDLYNAME     int
	RedirectedIOReason       int
	VolumeRedirectedIOReason int
}

type CLUSTER_SHARED_VOLUME_RENAME_INPUT_VOLUME struct {
	InputType int
	Anonymous int
}

type CLUSTER_SHARED_VOLUME_RENAME_INPUT_NAME struct {
	NewVolumeName int
}

type CLUSTER_SHARED_VOLUME_RENAME_INPUT_GUID_NAME struct {
	NewVolumeName int
	NewVolumeGuid int
}

type CLUSTER_SHARED_VOLUME_RENAME_INPUT struct {
	AnonymousBase_clusapi_L5427_C14 int
	AnonymousBase_clusapi_L5428_C14 int
}

type CLUSTER_SHARED_VOLUME_RENAME_GUID_INPUT struct {
	AnonymousBase_clusapi_L5438_C14 int
	AnonymousBase_clusapi_L5439_C14 int
}

type CLUS_CHKDSK_INFO struct {
	PartitionNumber int
	ChkdskState     int
	FileIdCount     int
	FileIdList      int
}

type CLUS_DISK_NUMBER_INFO struct {
	DiskNumber     int
	BytesPerSector int
}

type CLUS_SHARED_VOLUME_BACKUP_MODE struct {
	BackupState      int
	DelayTimerInSecs int
	VolumeName       int
}

type CLUSPROP_PARTITION_INFO struct {
	AnonymousBase_clusapi_L5470_C14 int
	AnonymousBase_clusapi_L5471_C14 int
}

type CLUSPROP_PARTITION_INFO_EX struct {
	AnonymousBase_clusapi_L5482_C14 int
	AnonymousBase_clusapi_L5483_C14 int
}

type CLUSPROP_PARTITION_INFO_EX2 struct {
	AnonymousBase_clusapi_L5496_C14 int
	AnonymousBase_clusapi_L5497_C14 int
}

type CLUS_FTSET_INFO struct {
	DWROOTSIGNATURE int
	DWFTTYPE        int
}

type CLUSPROP_FTSET_INFO struct {
	AnonymousBase_clusapi_L5518_C14 int
	AnonymousBase_clusapi_L5519_C14 int
}

type CLUS_SCSI_ADDRESS struct {
	Anonymous int
}

type CLUSPROP_SCSI_ADDRESS struct {
	AnonymousBase_clusapi_L5546_C14 int
	AnonymousBase_clusapi_L5547_C14 int
}

type CLUS_NETNAME_VS_TOKEN_INFO struct {
	ProcessID     int
	DesiredAccess int
	InheritHandle int
}

type CLUS_NETNAME_PWD_INFO struct {
	Flags      int
	Password   int
	CreatingDC int
	ObjectGuid int
}

type CLUS_NETNAME_PWD_INFOEX struct {
	Flags      int
	Password   int
	CreatingDC int
	ObjectGuid int
}

type CLUS_DNN_LEADER_STATUS struct {
	IsOnline            int
	IsFileServerPresent int
}

type CLUS_DNN_SODAFS_CLONE_STATUS struct {
	NodeId int
	Status int
}

type CLUS_NETNAME_IP_INFO_ENTRY struct {
	NodeId      int
	AddressSize int
	Address     int
}

type CLUS_NETNAME_IP_INFO_FOR_MULTICHANNEL struct {
	SZNAME     int
	NumEntries int
	IpInfo     int
}

type CLUS_MAINTENANCE_MODE_INFO struct {
	InMaintenance int
}

type CLUS_CSV_MAINTENANCE_MODE_INFO struct {
	InMaintenance int
	VolumeName    int
}

type CLUS_MAINTENANCE_MODE_INFOEX struct {
	InMaintenance        int
	MaintainenceModeType int
	InternalState        int
	Signature            int
}

type CLUS_SET_MAINTENANCE_MODE_INPUT struct {
	InMaintenance      int
	ExtraParameterSize int
	ExtraParameter     int
}

type CLUS_STORAGE_SET_DRIVELETTER struct {
	PartitionNumber int
	DriveLetterMask int
}

type CLUS_STORAGE_GET_AVAILABLE_DRIVELETTERS struct {
	AvailDrivelettersMask int
}

type CLUS_STORAGE_REMAP_DRIVELETTER struct {
	CurrentDriveLetterMask int
	TargetDriveLetterMask  int
}

type CLUS_PROVIDER_STATE_CHANGE_INFO struct {
	DWSIZE        int
	RESOURCESTATE int
	SZPROVIDERID  int
}

type CLUS_CREATE_INFRASTRUCTURE_FILESERVER_INPUT struct {
	FileServerName int
}

type CLUS_CREATE_INFRASTRUCTURE_FILESERVER_OUTPUT struct {
	FileServerName int
}

type CLUSPROP_LIST struct {
	NPROPERTYCOUNT int
	PropertyName   int
}

type FILESHARE_CHANGE struct {
	Change    int
	ShareName int
}

type FILESHARE_CHANGE_LIST struct {
	NumEntries  int
	ChangeEntry int
}

type CLUSCTL_GROUP_GET_LAST_MOVE_TIME_OUTPUT struct {
	GetTickCount64 int
	GetSystemTime  int
	NodeId         int
}

type CLUSPROP_BUFFER_HELPER struct {
	PB                       int
	PW                       int
	PDW                      int
	PL                       int
	PSZ                      int
	PLIST                    int
	PSYNTAX                  int
	PNAME                    int
	PVALUE                   int
	PBINARYVALUE             int
	PWORDVALUE               int
	PDWORDVALUE              int
	PLONGVALUE               int
	PULARGEINTEGERVALUE      int
	PLARGEINTEGERVALUE       int
	PSTRINGVALUE             int
	PMULTISZVALUE            int
	PSECURITYDESCRIPTOR      int
	PRESOURCECLASSVALUE      int
	PRESOURCECLASSINFOVALUE  int
	PDISKSIGNATUREVALUE      int
	PSCSIADDRESSVALUE        int
	PDISKNUMBERVALUE         int
	PPARTITIONINFOVALUE      int
	PREQUIREDDEPENDENCYVALUE int
	PPARTITIONINFOVALUEEX    int
	PPARTITIONINFOVALUEEX2   int
	PFILETIMEVALUE           int
}

type SR_RESOURCE_TYPE_REPLICATED_PARTITION_INFO struct {
	PartitionOffset int
	Capabilities    int
}

type SR_RESOURCE_TYPE_REPLICATED_PARTITION_ARRAY struct {
	Count          int
	PartitionArray int
}

type SR_RESOURCE_TYPE_QUERY_ELIGIBLE_LOGDISKS struct {
	DataDiskGuid        int
	IncludeOfflineDisks int
}

type SR_RESOURCE_TYPE_QUERY_ELIGIBLE_TARGET_DATADISKS struct {
	SourceDataDiskGuid         int
	TargetReplicationGroupGuid int
	SkipConnectivityCheck      int
	IncludeOfflineDisks        int
}

type SR_RESOURCE_TYPE_QUERY_ELIGIBLE_SOURCE_DATADISKS struct {
	DataDiskGuid                 int
	IncludeAvailableStoargeDisks int
}

type SR_RESOURCE_TYPE_DISK_INFO struct {
	Reason   int
	DiskGuid int
}

type SR_RESOURCE_TYPE_ELIGIBLE_DISKS_RESULT struct {
	Count    int
	DiskInfo int
}

type SR_RESOURCE_TYPE_REPLICATED_DISK struct {
	Type                    int
	ClusterDiskResourceGuid int
	ReplicationGroupId      int
	ReplicationGroupName    int
}

type SR_RESOURCE_TYPE_REPLICATED_DISKS_RESULT struct {
	Count           int
	ReplicatedDisks int
}

type CLUSCTL_RESOURCE_TYPE_STORAGE_GET_AVAILABLE_DISKS_EX2_INPUT struct {
	DWFLAGS        int
	GUIDPOOLFILTER int
}

type RESOURCE_STATUS struct {
	ResourceState int
	CheckPoint    int
	WaitHint      int
	EventHandle   int
}

type NodeUtilizationInfoElement struct {
	Id                              int
	AvailableMemory                 int
	AvailableMemoryAfterReclamation int
}

type ResourceUtilizationInfoElement struct {
	PhysicalNumaId int
	CurrentMemory  int
}

type GET_OPERATION_CONTEXT_PARAMS struct {
	Size     int
	Version  int
	Type     int
	Priority int
}

type RESOURCE_STATUS_EX struct {
	ResourceState                int
	CheckPoint                   int
	EventHandle                  int
	ApplicationSpecificErrorCode int
	Flags                        int
	WaitHint                     int
}

type CLRES_V1_FUNCTIONS struct {
	Open                int
	Close               int
	Online              int
	Offline             int
	Terminate           int
	LooksAlive          int
	IsAlive             int
	Arbitrate           int
	Release             int
	ResourceControl     int
	ResourceTypeControl int
}

type CLRES_V2_FUNCTIONS struct {
	Open                int
	Close               int
	Online              int
	Offline             int
	Terminate           int
	LooksAlive          int
	IsAlive             int
	Arbitrate           int
	Release             int
	ResourceControl     int
	ResourceTypeControl int
	Cancel              int
}

type CLRES_V3_FUNCTIONS struct {
	Open                     int
	Close                    int
	Online                   int
	Offline                  int
	Terminate                int
	LooksAlive               int
	IsAlive                  int
	Arbitrate                int
	Release                  int
	BeginResourceControl     int
	BeginResourceTypeControl int
	Cancel                   int
}

type CLRES_V4_FUNCTIONS struct {
	Open                           int
	Close                          int
	Online                         int
	Offline                        int
	Terminate                      int
	LooksAlive                     int
	IsAlive                        int
	Arbitrate                      int
	Release                        int
	BeginResourceControl           int
	BeginResourceTypeControl       int
	Cancel                         int
	BeginResourceControlAsUser     int
	BeginResourceTypeControlAsUser int
}

type CLRES_FUNCTION_TABLE struct {
	TableSize int
	Version   int
	Anonymous int
}

type RESUTIL_LARGEINT_DATA struct {
	Default int
	Minimum int
	Maximum int
}

type RESUTIL_ULARGEINT_DATA struct {
	Default int
	Minimum int
	Maximum int
}

type RESUTIL_FILETIME_DATA struct {
	Default int
	Minimum int
	Maximum int
}

type RESUTIL_PROPERTY_ITEM struct {
	Name      int
	KeyName   int
	Format    int
	Anonymous int
	Minimum   int
	Maximum   int
	Flags     int
	Offset    int
}

type CLRES_CALLBACK_FUNCTION_TABLE struct {
	LogEvent                               int
	SetResourceStatusEx                    int
	SetResourceLockedMode                  int
	SignalFailure                          int
	SetResourceInMemoryNodeLocalProperties int
	EndControlCall                         int
	EndTypeControlCall                     int
	ExtendControlCall                      int
	ExtendTypeControlCall                  int
	RaiseResTypeNotification               int
	ChangeResourceProcessForDumps          int
	ChangeResTypeProcessForDumps           int
	SetInternalState                       int
	SetResourceLockedModeEx                int
}

type MONITOR_STATE struct {
	LastUpdate     int
	State          int
	ActiveResource int
	ResmonStop     int
}

type POST_UPGRADE_VERSION_INFO struct {
	NEWMAJORVERSION   int
	NEWUPGRADEVERSION int
	OLDMAJORVERSION   int
	OLDUPGRADEVERSION int
	RESERVED          int
}

type CLUSTER_HEALTH_FAULT struct {
	Id          int
	ErrorType   int
	ErrorCode   int
	Description int
	Provider    int
	Flags       int
	Reserved    int
}

type CLUSTER_HEALTH_FAULT_ARRAY struct {
	NUMFAULTS int
	FAULTS    int
}

type CLUS_WORKER struct {
	HTHREAD   int
	Terminate int
}

type HCLUSCRYPTPROVIDER struct {
}

type PaxosTagCStruct struct {
	PADDING__PAXOSTAGVTABLE           int
	PADDING__NEXTEPOCHVTABLE          int
	PADDING__NEXTEPOCH_DATETIMEVTABLE int
	NextEpoch_DateTime_ticks          int
	NextEpoch_Value                   int
	PADDING__BOUNDRYNEXTEPOCH         int
	PADDING__EPOCHVTABLE              int
	PADDING__EPOCH_DATETIMEVTABLE     int
	Epoch_DateTime_ticks              int
	Epoch_Value                       int
	PADDING__BOUNDRYEPOCH             int
	Sequence                          int
	PADDING__BOUNDRYSEQUENCE          int
}

type WitnessTagUpdateHelper struct {
	Version         int
	PAXOSTOSET      int
	PAXOSTOVALIDATE int
}

type WitnessTagHelper struct {
	Version         int
	PAXOSTOVALIDATE int
}

type ClusApplication struct {
}

type Cluster struct {
}

type ClusVersion struct {
}

type ClusResType struct {
}

type ClusProperty struct {
}

type ClusProperties struct {
}

type DomainNames struct {
}

type ClusNetwork struct {
}

type ClusNetInterface struct {
}

type ClusNetInterfaces struct {
}

type ClusResDependencies struct {
}

type ClusResGroupResources struct {
}

type ClusResTypeResources struct {
}

type ClusResGroupPreferredOwnerNodes struct {
}

type ClusResPossibleOwnerNodes struct {
}

type ClusNetworks struct {
}

type ClusNetworkNetInterfaces struct {
}

type ClusNodeNetInterfaces struct {
}

type ClusRefObject struct {
}

type ClusterNames struct {
}

type ClusNode struct {
}

type ClusNodes struct {
}

type ClusResGroup struct {
}

type ClusResGroups struct {
}

type ClusResource struct {
}

type ClusResources struct {
}

type ClusResTypes struct {
}

type ClusResTypePossibleOwnerNodes struct {
}

type ClusPropertyValue struct {
}

type ClusPropertyValues struct {
}

type ClusPropertyValueData struct {
}

type ClusPartition struct {
}

type ClusPartitionEx struct {
}

type ClusPartitions struct {
}

type ClusDisk struct {
}

type ClusDisks struct {
}

type ClusScsiAddress struct {
}

type ClusRegistryKeys struct {
}

type ClusCryptoKeys struct {
}

type ClusResDependents struct {
}
