// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package deviceanddriverinstallation implements the Windows.Win32.DeviceAndDriverInstallation namespace.
package deviceanddriverinstallation

type SP_ALTPLATFORM_INFO_V3 struct {
	CBSIZE                     int
	Platform                   int
	MajorVersion               int
	MinorVersion               int
	ProcessorArchitecture      int
	Anonymous                  int
	FirstValidatedMajorVersion int
	FirstValidatedMinorVersion int
	ProductType                int
	SuiteMask                  int
	BuildNumber                int
}

type SP_DEVINFO_DATA struct {
	CBSIZE    int
	ClassGuid int
	DevInst   int
	Reserved  int
}

type SP_DEVICE_INTERFACE_DATA struct {
	CBSIZE             int
	InterfaceClassGuid int
	Flags              int
	Reserved           int
}

type SP_DEVICE_INTERFACE_DETAIL_DATA_A struct {
	CBSIZE     int
	DevicePath int
}

type SP_DEVICE_INTERFACE_DETAIL_DATA_W struct {
	CBSIZE     int
	DevicePath int
}

type SP_DEVINFO_LIST_DETAIL_DATA_A struct {
	CBSIZE              int
	ClassGuid           int
	RemoteMachineHandle int
	RemoteMachineName   int
}

type SP_DEVINFO_LIST_DETAIL_DATA_W struct {
	CBSIZE              int
	ClassGuid           int
	RemoteMachineHandle int
	RemoteMachineName   int
}

type SP_DEVINSTALL_PARAMS_A struct {
	CBSIZE                   int
	Flags                    int
	FlagsEx                  int
	HWNDPARENT               int
	InstallMsgHandler        int
	InstallMsgHandlerContext int
	FileQueue                int
	ClassInstallReserved     int
	Reserved                 int
	DriverPath               int
}

type SP_DEVINSTALL_PARAMS_W struct {
	CBSIZE                   int
	Flags                    int
	FlagsEx                  int
	HWNDPARENT               int
	InstallMsgHandler        int
	InstallMsgHandlerContext int
	FileQueue                int
	ClassInstallReserved     int
	Reserved                 int
	DriverPath               int
}

type SP_CLASSINSTALL_HEADER struct {
	CBSIZE          int
	InstallFunction int
}

type SP_ENABLECLASS_PARAMS struct {
	ClassInstallHeader int
	ClassGuid          int
	EnableMessage      int
}

type SP_PROPCHANGE_PARAMS struct {
	ClassInstallHeader int
	StateChange        int
	Scope              int
	HwProfile          int
}

type SP_REMOVEDEVICE_PARAMS struct {
	ClassInstallHeader int
	Scope              int
	HwProfile          int
}

type SP_UNREMOVEDEVICE_PARAMS struct {
	ClassInstallHeader int
	Scope              int
	HwProfile          int
}

type SP_SELECTDEVICE_PARAMS_A struct {
	ClassInstallHeader int
	Title              int
	Instructions       int
	ListLabel          int
	SubTitle           int
	Reserved           int
}

type SP_SELECTDEVICE_PARAMS_W struct {
	ClassInstallHeader int
	Title              int
	Instructions       int
	ListLabel          int
	SubTitle           int
}

type SP_DETECTDEVICE_PARAMS struct {
	ClassInstallHeader   int
	DetectProgressNotify int
	ProgressNotifyParam  int
}

type SP_INSTALLWIZARD_DATA struct {
	ClassInstallHeader int
	Flags              int
	DynamicPages       int
	NumDynamicPages    int
	DynamicPageFlags   int
	PrivateFlags       int
	PrivateData        int
	HWNDWIZARDDLG      int
}

type SP_NEWDEVICEWIZARD_DATA struct {
	ClassInstallHeader int
	Flags              int
	DynamicPages       int
	NumDynamicPages    int
	HWNDWIZARDDLG      int
}

type SP_TROUBLESHOOTER_PARAMS_A struct {
	ClassInstallHeader int
	ChmFile            int
	HtmlTroubleShooter int
}

type SP_TROUBLESHOOTER_PARAMS_W struct {
	ClassInstallHeader int
	ChmFile            int
	HtmlTroubleShooter int
}

type SP_POWERMESSAGEWAKE_PARAMS_A struct {
	ClassInstallHeader int
	PowerMessageWake   int
}

type SP_POWERMESSAGEWAKE_PARAMS_W struct {
	ClassInstallHeader int
	PowerMessageWake   int
}

type SP_DRVINFO_DATA_V2_A struct {
	CBSIZE        int
	DriverType    int
	Reserved      int
	Description   int
	MfgName       int
	ProviderName  int
	DriverDate    int
	DriverVersion int
}

type SP_DRVINFO_DATA_V2_W struct {
	CBSIZE        int
	DriverType    int
	Reserved      int
	Description   int
	MfgName       int
	ProviderName  int
	DriverDate    int
	DriverVersion int
}

type SP_DRVINFO_DATA_V1_A struct {
	CBSIZE       int
	DriverType   int
	Reserved     int
	Description  int
	MfgName      int
	ProviderName int
}

type SP_DRVINFO_DATA_V1_W struct {
	CBSIZE       int
	DriverType   int
	Reserved     int
	Description  int
	MfgName      int
	ProviderName int
}

type SP_DRVINFO_DETAIL_DATA_A struct {
	CBSIZE          int
	InfDate         int
	CompatIDsOffset int
	CompatIDsLength int
	Reserved        int
	SectionName     int
	InfFileName     int
	DrvDescription  int
	HardwareID      int
}

type SP_DRVINFO_DETAIL_DATA_W struct {
	CBSIZE          int
	InfDate         int
	CompatIDsOffset int
	CompatIDsLength int
	Reserved        int
	SectionName     int
	InfFileName     int
	DrvDescription  int
	HardwareID      int
}

type SP_DRVINSTALL_PARAMS struct {
	CBSIZE      int
	Rank        int
	Flags       int
	PrivateData int
	Reserved    int
}

type COINSTALLER_CONTEXT_DATA struct {
	PostProcessing int
	InstallResult  int
	PrivateData    int
}

type SP_CLASSIMAGELIST_DATA struct {
	CBSIZE    int
	ImageList int
	Reserved  int
}

type SP_PROPSHEETPAGE_REQUEST struct {
	CBSIZE         int
	PageRequested  int
	DeviceInfoSet  int
	DeviceInfoData int
}

type SP_BACKUP_QUEUE_PARAMS_V2_A struct {
	CBSIZE            int
	FullInfPath       int
	FilenameOffset    int
	ReinstallInstance int
}

type SP_BACKUP_QUEUE_PARAMS_V2_W struct {
	CBSIZE            int
	FullInfPath       int
	FilenameOffset    int
	ReinstallInstance int
}

type SP_BACKUP_QUEUE_PARAMS_V1_A struct {
	CBSIZE         int
	FullInfPath    int
	FilenameOffset int
}

type SP_BACKUP_QUEUE_PARAMS_V1_W struct {
	CBSIZE         int
	FullInfPath    int
	FilenameOffset int
}

type CONFLICT_DETAILS_A struct {
	CD_ulSize        int
	CD_ulMask        int
	CD_dnDevInst     int
	CD_rdResDes      int
	CD_ulFlags       int
	CD_szDescription int
}

type CONFLICT_DETAILS_W struct {
	CD_ulSize        int
	CD_ulMask        int
	CD_dnDevInst     int
	CD_rdResDes      int
	CD_ulFlags       int
	CD_szDescription int
}

type MEM_RANGE struct {
	MR_Align    int
	MR_nBytes   int
	MR_Min      int
	MR_Max      int
	MR_Flags    int
	MR_Reserved int
}

type MEM_DES struct {
	MD_Count      int
	MD_Type       int
	MD_Alloc_Base int
	MD_Alloc_End  int
	MD_Flags      int
	MD_Reserved   int
}

type MEM_RESOURCE struct {
	MEM_Header int
	MEM_Data   int
}

type Mem_Large_Range_s struct {
	MLR_Align    int
	MLR_nBytes   int
	MLR_Min      int
	MLR_Max      int
	MLR_Flags    int
	MLR_Reserved int
}

type Mem_Large_Des_s struct {
	MLD_Count      int
	MLD_Type       int
	MLD_Alloc_Base int
	MLD_Alloc_End  int
	MLD_Flags      int
	MLD_Reserved   int
}

type Mem_Large_Resource_s struct {
	MEM_LARGE_Header int
	MEM_LARGE_Data   int
}

type IO_RANGE struct {
	IOR_Align      int
	IOR_nPorts     int
	IOR_Min        int
	IOR_Max        int
	IOR_RangeFlags int
	IOR_Alias      int
}

type IO_DES struct {
	IOD_Count      int
	IOD_Type       int
	IOD_Alloc_Base int
	IOD_Alloc_End  int
	IOD_DesFlags   int
}

type IO_RESOURCE struct {
	IO_Header int
	IO_Data   int
}

type DMA_RANGE struct {
	DR_Min   int
	DR_Max   int
	DR_Flags int
}

type DMA_DES struct {
	DD_Count      int
	DD_Type       int
	DD_Flags      int
	DD_Alloc_Chan int
}

type DMA_RESOURCE struct {
	DMA_Header int
	DMA_Data   int
}

type IRQ_RANGE struct {
	IRQR_Min   int
	IRQR_Max   int
	IRQR_Flags int
}

type IRQ_DES_32 struct {
	IRQD_Count     int
	IRQD_Type      int
	IRQD_Flags     int
	IRQD_Alloc_Num int
	IRQD_Affinity  int
}

type IRQ_DES_64 struct {
	IRQD_Count     int
	IRQD_Type      int
	IRQD_Flags     int
	IRQD_Alloc_Num int
	IRQD_Affinity  int
}

type IRQ_RESOURCE_32 struct {
	IRQ_Header int
	IRQ_Data   int
}

type IRQ_RESOURCE_64 struct {
	IRQ_Header int
	IRQ_Data   int
}

type DevPrivate_Range_s struct {
	PR_Data1 int
	PR_Data2 int
	PR_Data3 int
}

type DevPrivate_Des_s struct {
	PD_Count int
	PD_Type  int
	PD_Data1 int
	PD_Data2 int
	PD_Data3 int
	PD_Flags int
}

type DevPrivate_Resource_s struct {
	PRV_Header int
	PRV_Data   int
}

type CS_DES struct {
	CSD_SignatureLength  int
	CSD_LegacyDataOffset int
	CSD_LegacyDataSize   int
	CSD_Flags            int
	CSD_ClassGuid        int
	CSD_Signature        int
}

type CS_RESOURCE struct {
	CS_Header int
}

type PCCARD_DES struct {
	PCD_Count           int
	PCD_Type            int
	PCD_Flags           int
	PCD_ConfigIndex     int
	PCD_Reserved        int
	PCD_MemoryCardBase1 int
	PCD_MemoryCardBase2 int
	PCD_MemoryCardBase  int
	PCD_MemoryFlags     int
	PCD_IoFlags         int
}

type PCCARD_RESOURCE struct {
	PcCard_Header int
}

type MFCARD_DES struct {
	PMF_Count              int
	PMF_Type               int
	PMF_Flags              int
	PMF_ConfigOptions      int
	PMF_IoResourceIndex    int
	PMF_Reserved           int
	PMF_ConfigRegisterBase int
}

type MFCARD_RESOURCE struct {
	MfCard_Header int
}

type BUSNUMBER_RANGE struct {
	BUSR_Min         int
	BUSR_Max         int
	BUSR_nBusNumbers int
	BUSR_Flags       int
}

type BUSNUMBER_DES struct {
	BUSD_Count      int
	BUSD_Type       int
	BUSD_Flags      int
	BUSD_Alloc_Base int
	BUSD_Alloc_End  int
}

type BUSNUMBER_RESOURCE struct {
	BusNumber_Header int
	BusNumber_Data   int
}

type Connection_Des_s struct {
	COND_Type      int
	COND_Flags     int
	COND_Class     int
	COND_ClassType int
	COND_Reserved1 int
	COND_Reserved2 int
	COND_Id        int
}

type Connection_Resource_s struct {
	Connection_Header int
}

type HWProfileInfo_sA struct {
	HWPI_ulHWProfile    int
	HWPI_szFriendlyName int
	HWPI_dwFlags        int
}

type HWProfileInfo_sW struct {
	HWPI_ulHWProfile    int
	HWPI_szFriendlyName int
	HWPI_dwFlags        int
}

type HCMNOTIFICATION__ struct {
	UNUSED int
}

type CM_NOTIFY_FILTER struct {
	CBSIZE     int
	Flags      int
	FilterType int
	Reserved   int
	U          int
}

type CM_NOTIFY_EVENT_DATA struct {
	FilterType int
	Reserved   int
	U          int
}
