// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package debug implements the Windows.Win32.Debug namespace.
package debug

type CONTEXT struct {
	ContextFlags      int
	Dr0               int
	Dr1               int
	Dr2               int
	Dr3               int
	Dr6               int
	Dr7               int
	FloatSave         int
	SegGs             int
	SegFs             int
	SegEs             int
	SegDs             int
	Edi               int
	Esi               int
	Ebx               int
	Edx               int
	Ecx               int
	Eax               int
	Ebp               int
	Eip               int
	SegCs             int
	EFlags            int
	Esp               int
	SegSs             int
	ExtendedRegisters int
}

type LDT_ENTRY struct {
	LimitLow int
	BaseLow  int
	HighWord int
}

type WOW64_FLOATING_SAVE_AREA struct {
	ControlWord   int
	StatusWord    int
	TagWord       int
	ErrorOffset   int
	ErrorSelector int
	DataOffset    int
	DataSelector  int
	RegisterArea  int
	Cr0NpxState   int
}

type WOW64_CONTEXT struct {
	ContextFlags      int
	Dr0               int
	Dr1               int
	Dr2               int
	Dr3               int
	Dr6               int
	Dr7               int
	FloatSave         int
	SegGs             int
	SegFs             int
	SegEs             int
	SegDs             int
	Edi               int
	Esi               int
	Ebx               int
	Edx               int
	Ecx               int
	Eax               int
	Ebp               int
	Eip               int
	SegCs             int
	EFlags            int
	Esp               int
	SegSs             int
	ExtendedRegisters int
}

type WOW64_LDT_ENTRY struct {
	LimitLow int
	BaseLow  int
	HighWord int
}

type EXCEPTION_RECORD struct {
	ExceptionCode        int
	ExceptionFlags       int
	ExceptionRecord      int
	ExceptionAddress     int
	NumberParameters     int
	ExceptionInformation int
}

type EXCEPTION_RECORD64 struct {
	ExceptionCode        int
	ExceptionFlags       int
	ExceptionRecord      int
	ExceptionAddress     int
	NumberParameters     int
	UNUSEDALIGNMENT      int
	ExceptionInformation int
}

type EXCEPTION_POINTERS struct {
	ExceptionRecord int
	ContextRecord   int
}

type IMAGE_FILE_HEADER struct {
	Machine              int
	NumberOfSections     int
	TimeDateStamp        int
	PointerToSymbolTable int
	NumberOfSymbols      int
	SizeOfOptionalHeader int
	Characteristics      int
}

type IMAGE_DATA_DIRECTORY struct {
	VirtualAddress int
	Size           int
}

type IMAGE_OPTIONAL_HEADER64 struct {
	Magic                       int
	MajorLinkerVersion          int
	MinorLinkerVersion          int
	SizeOfCode                  int
	SizeOfInitializedData       int
	SizeOfUninitializedData     int
	AddressOfEntryPoint         int
	BaseOfCode                  int
	ImageBase                   int
	SectionAlignment            int
	FileAlignment               int
	MajorOperatingSystemVersion int
	MinorOperatingSystemVersion int
	MajorImageVersion           int
	MinorImageVersion           int
	MajorSubsystemVersion       int
	MinorSubsystemVersion       int
	Win32VersionValue           int
	SizeOfImage                 int
	SizeOfHeaders               int
	CheckSum                    int
	Subsystem                   int
	DllCharacteristics          int
	SizeOfStackReserve          int
	SizeOfStackCommit           int
	SizeOfHeapReserve           int
	SizeOfHeapCommit            int
	LoaderFlags                 int
	NumberOfRvaAndSizes         int
	DataDirectory               int
}

type IMAGE_NT_HEADERS64 struct {
	Signature      int
	FileHeader     int
	OptionalHeader int
}

type IMAGE_SECTION_HEADER struct {
	Name                 int
	Misc                 int
	VirtualAddress       int
	SizeOfRawData        int
	PointerToRawData     int
	PointerToRelocations int
	PointerToLinenumbers int
	NumberOfRelocations  int
	NumberOfLinenumbers  int
	Characteristics      int
}

type IMAGE_LOAD_CONFIG_DIRECTORY32 struct {
	Size                                     int
	TimeDateStamp                            int
	MajorVersion                             int
	MinorVersion                             int
	GlobalFlagsClear                         int
	GlobalFlagsSet                           int
	CriticalSectionDefaultTimeout            int
	DeCommitFreeBlockThreshold               int
	DeCommitTotalFreeThreshold               int
	LockPrefixTable                          int
	MaximumAllocationSize                    int
	VirtualMemoryThreshold                   int
	ProcessHeapFlags                         int
	ProcessAffinityMask                      int
	CSDVersion                               int
	DependentLoadFlags                       int
	EditList                                 int
	SecurityCookie                           int
	SEHandlerTable                           int
	SEHandlerCount                           int
	GuardCFCheckFunctionPointer              int
	GuardCFDispatchFunctionPointer           int
	GuardCFFunctionTable                     int
	GuardCFFunctionCount                     int
	GuardFlags                               int
	CodeIntegrity                            int
	GuardAddressTakenIatEntryTable           int
	GuardAddressTakenIatEntryCount           int
	GuardLongJumpTargetTable                 int
	GuardLongJumpTargetCount                 int
	DynamicValueRelocTable                   int
	CHPEMetadataPointer                      int
	GuardRFFailureRoutine                    int
	GuardRFFailureRoutineFunctionPointer     int
	DynamicValueRelocTableOffset             int
	DynamicValueRelocTableSection            int
	Reserved2                                int
	GuardRFVerifyStackPointerFunctionPointer int
	HotPatchTableOffset                      int
	Reserved3                                int
	EnclaveConfigurationPointer              int
	VolatileMetadataPointer                  int
	GuardEHContinuationTable                 int
	GuardEHContinuationCount                 int
}

type IMAGE_LOAD_CONFIG_DIRECTORY64 struct {
	Size                                     int
	TimeDateStamp                            int
	MajorVersion                             int
	MinorVersion                             int
	GlobalFlagsClear                         int
	GlobalFlagsSet                           int
	CriticalSectionDefaultTimeout            int
	DeCommitFreeBlockThreshold               int
	DeCommitTotalFreeThreshold               int
	LockPrefixTable                          int
	MaximumAllocationSize                    int
	VirtualMemoryThreshold                   int
	ProcessAffinityMask                      int
	ProcessHeapFlags                         int
	CSDVersion                               int
	DependentLoadFlags                       int
	EditList                                 int
	SecurityCookie                           int
	SEHandlerTable                           int
	SEHandlerCount                           int
	GuardCFCheckFunctionPointer              int
	GuardCFDispatchFunctionPointer           int
	GuardCFFunctionTable                     int
	GuardCFFunctionCount                     int
	GuardFlags                               int
	CodeIntegrity                            int
	GuardAddressTakenIatEntryTable           int
	GuardAddressTakenIatEntryCount           int
	GuardLongJumpTargetTable                 int
	GuardLongJumpTargetCount                 int
	DynamicValueRelocTable                   int
	CHPEMetadataPointer                      int
	GuardRFFailureRoutine                    int
	GuardRFFailureRoutineFunctionPointer     int
	DynamicValueRelocTableOffset             int
	DynamicValueRelocTableSection            int
	Reserved2                                int
	GuardRFVerifyStackPointerFunctionPointer int
	HotPatchTableOffset                      int
	Reserved3                                int
	EnclaveConfigurationPointer              int
	VolatileMetadataPointer                  int
	GuardEHContinuationTable                 int
	GuardEHContinuationCount                 int
}

type IMAGE_DEBUG_DIRECTORY struct {
	Characteristics  int
	TimeDateStamp    int
	MajorVersion     int
	MinorVersion     int
	Type             int
	SizeOfData       int
	AddressOfRawData int
	PointerToRawData int
}

type IMAGE_COFF_SYMBOLS_HEADER struct {
	NumberOfSymbols      int
	LvaToFirstSymbol     int
	NumberOfLinenumbers  int
	LvaToFirstLinenumber int
	RvaToFirstByteOfCode int
	RvaToLastByteOfCode  int
	RvaToFirstByteOfData int
	RvaToLastByteOfData  int
}

type FPO_DATA struct {
	ULOFFSTART int
	CBPROCSIZE int
	CDWLOCALS  int
	CDWPARAMS  int
	BITFIELD   int
}

type IMAGE_FUNCTION_ENTRY struct {
	StartingAddress int
	EndingAddress   int
	EndOfPrologue   int
}

type IMAGE_FUNCTION_ENTRY64 struct {
	StartingAddress int
	EndingAddress   int
	Anonymous       int
}

type FLASHWINFO struct {
	CBSIZE    int
	HWND      int
	DWFLAGS   int
	UCOUNT    int
	DWTIMEOUT int
}

type EXCEPTION_DEBUG_INFO struct {
	ExceptionRecord int
	DWFIRSTCHANCE   int
}

type CREATE_THREAD_DEBUG_INFO struct {
	HTHREAD           int
	LPTHREADLOCALBASE int
	LPSTARTADDRESS    int
}

type CREATE_PROCESS_DEBUG_INFO struct {
	HFILE                 int
	HPROCESS              int
	HTHREAD               int
	LPBASEOFIMAGE         int
	DWDEBUGINFOFILEOFFSET int
	NDEBUGINFOSIZE        int
	LPTHREADLOCALBASE     int
	LPSTARTADDRESS        int
	LPIMAGENAME           int
	FUNICODE              int
}

type EXIT_THREAD_DEBUG_INFO struct {
	DWEXITCODE int
}

type EXIT_PROCESS_DEBUG_INFO struct {
	DWEXITCODE int
}

type LOAD_DLL_DEBUG_INFO struct {
	HFILE                 int
	LPBASEOFDLL           int
	DWDEBUGINFOFILEOFFSET int
	NDEBUGINFOSIZE        int
	LPIMAGENAME           int
	FUNICODE              int
}

type UNLOAD_DLL_DEBUG_INFO struct {
	LPBASEOFDLL int
}

type OUTPUT_DEBUG_STRING_INFO struct {
	LPDEBUGSTRINGDATA  int
	FUNICODE           int
	NDEBUGSTRINGLENGTH int
}

type RIP_INFO struct {
	DWERROR int
	DWTYPE  int
}

type DEBUG_EVENT struct {
	DWDEBUGEVENTCODE int
	DWPROCESSID      int
	DWTHREADID       int
	U                int
}

type WAITCHAIN_NODE_INFO struct {
	ObjectType   int
	ObjectStatus int
	Anonymous    int
}

type MINIDUMP_LOCATION_DESCRIPTOR struct {
	DataSize int
	Rva      int
}

type MINIDUMP_LOCATION_DESCRIPTOR64 struct {
	DataSize int
	Rva      int
}

type MINIDUMP_MEMORY_DESCRIPTOR struct {
	StartOfMemoryRange int
	Memory             int
}

type MINIDUMP_MEMORY_DESCRIPTOR64 struct {
	StartOfMemoryRange int
	DataSize           int
}

type MINIDUMP_HEADER struct {
	Signature          int
	Version            int
	NumberOfStreams    int
	StreamDirectoryRva int
	CheckSum           int
	Anonymous          int
	Flags              int
}

type MINIDUMP_DIRECTORY struct {
	StreamType int
	Location   int
}

type MINIDUMP_STRING struct {
	Length int
	Buffer int
}

type CPU_INFORMATION struct {
	X86CpuInfo   int
	OtherCpuInfo int
}

type MINIDUMP_SYSTEM_INFO struct {
	ProcessorArchitecture int
	ProcessorLevel        int
	ProcessorRevision     int
	Anonymous1            int
	MajorVersion          int
	MinorVersion          int
	BuildNumber           int
	PlatformId            int
	CSDVersionRva         int
	Anonymous2            int
	Cpu                   int
}

type MINIDUMP_THREAD struct {
	ThreadId      int
	SuspendCount  int
	PriorityClass int
	Priority      int
	Teb           int
	Stack         int
	ThreadContext int
}

type MINIDUMP_THREAD_LIST struct {
	NumberOfThreads int
	Threads         int
}

type MINIDUMP_THREAD_EX struct {
	ThreadId      int
	SuspendCount  int
	PriorityClass int
	Priority      int
	Teb           int
	Stack         int
	ThreadContext int
	BackingStore  int
}

type MINIDUMP_THREAD_EX_LIST struct {
	NumberOfThreads int
	Threads         int
}

type MINIDUMP_EXCEPTION struct {
	ExceptionCode        int
	ExceptionFlags       int
	ExceptionRecord      int
	ExceptionAddress     int
	NumberParameters     int
	UNUSEDALIGNMENT      int
	ExceptionInformation int
}

type MINIDUMP_EXCEPTION_STREAM struct {
	ThreadId        int
	ALIGNMENT       int
	ExceptionRecord int
	ThreadContext   int
}

type MINIDUMP_MODULE struct {
	BaseOfImage   int
	SizeOfImage   int
	CheckSum      int
	TimeDateStamp int
	ModuleNameRva int
	VersionInfo   int
	CvRecord      int
	MiscRecord    int
	Reserved0     int
	Reserved1     int
}

type MINIDUMP_MODULE_LIST struct {
	NumberOfModules int
	Modules         int
}

type MINIDUMP_MEMORY_LIST struct {
	NumberOfMemoryRanges int
	MemoryRanges         int
}

type MINIDUMP_MEMORY64_LIST struct {
	NumberOfMemoryRanges int
	BaseRva              int
	MemoryRanges         int
}

type MINIDUMP_EXCEPTION_INFORMATION struct {
	ThreadId          int
	ExceptionPointers int
	ClientPointers    int
}

type MINIDUMP_EXCEPTION_INFORMATION64 struct {
	ThreadId        int
	ExceptionRecord int
	ContextRecord   int
	ClientPointers  int
}

type MINIDUMP_HANDLE_OBJECT_INFORMATION struct {
	NextInfoRva int
	InfoType    int
	SizeOfInfo  int
}

type MINIDUMP_HANDLE_DESCRIPTOR struct {
	Handle        int
	TypeNameRva   int
	ObjectNameRva int
	Attributes    int
	GrantedAccess int
	HandleCount   int
	PointerCount  int
}

type MINIDUMP_HANDLE_DESCRIPTOR_2 struct {
	Handle        int
	TypeNameRva   int
	ObjectNameRva int
	Attributes    int
	GrantedAccess int
	HandleCount   int
	PointerCount  int
	ObjectInfoRva int
	Reserved0     int
}

type MINIDUMP_HANDLE_DATA_STREAM struct {
	SizeOfHeader        int
	SizeOfDescriptor    int
	NumberOfDescriptors int
	Reserved            int
}

type MINIDUMP_HANDLE_OPERATION_LIST struct {
	SizeOfHeader    int
	SizeOfEntry     int
	NumberOfEntries int
	Reserved        int
}

type MINIDUMP_FUNCTION_TABLE_DESCRIPTOR struct {
	MinimumAddress int
	MaximumAddress int
	BaseAddress    int
	EntryCount     int
	SizeOfAlignPad int
}

type MINIDUMP_FUNCTION_TABLE_STREAM struct {
	SizeOfHeader           int
	SizeOfDescriptor       int
	SizeOfNativeDescriptor int
	SizeOfFunctionEntry    int
	NumberOfDescriptors    int
	SizeOfAlignPad         int
}

type MINIDUMP_UNLOADED_MODULE struct {
	BaseOfImage   int
	SizeOfImage   int
	CheckSum      int
	TimeDateStamp int
	ModuleNameRva int
}

type MINIDUMP_UNLOADED_MODULE_LIST struct {
	SizeOfHeader    int
	SizeOfEntry     int
	NumberOfEntries int
}

type XSTATE_CONFIG_FEATURE_MSC_INFO struct {
	SizeOfInfo      int
	ContextSize     int
	EnabledFeatures int
	Features        int
}

type MINIDUMP_MISC_INFO struct {
	SizeOfInfo        int
	Flags1            int
	ProcessId         int
	ProcessCreateTime int
	ProcessUserTime   int
	ProcessKernelTime int
}

type MINIDUMP_MISC_INFO_2 struct {
	SizeOfInfo                int
	Flags1                    int
	ProcessId                 int
	ProcessCreateTime         int
	ProcessUserTime           int
	ProcessKernelTime         int
	ProcessorMaxMhz           int
	ProcessorCurrentMhz       int
	ProcessorMhzLimit         int
	ProcessorMaxIdleState     int
	ProcessorCurrentIdleState int
}

type MINIDUMP_MISC_INFO_3 struct {
	SizeOfInfo                int
	Flags1                    int
	ProcessId                 int
	ProcessCreateTime         int
	ProcessUserTime           int
	ProcessKernelTime         int
	ProcessorMaxMhz           int
	ProcessorCurrentMhz       int
	ProcessorMhzLimit         int
	ProcessorMaxIdleState     int
	ProcessorCurrentIdleState int
	ProcessIntegrityLevel     int
	ProcessExecuteFlags       int
	ProtectedProcess          int
	TimeZoneId                int
	TimeZone                  int
}

type MINIDUMP_MISC_INFO_4 struct {
	SizeOfInfo                int
	Flags1                    int
	ProcessId                 int
	ProcessCreateTime         int
	ProcessUserTime           int
	ProcessKernelTime         int
	ProcessorMaxMhz           int
	ProcessorCurrentMhz       int
	ProcessorMhzLimit         int
	ProcessorMaxIdleState     int
	ProcessorCurrentIdleState int
	ProcessIntegrityLevel     int
	ProcessExecuteFlags       int
	ProtectedProcess          int
	TimeZoneId                int
	TimeZone                  int
	BuildString               int
	DbgBldStr                 int
}

type MINIDUMP_MISC_INFO_5 struct {
	SizeOfInfo                int
	Flags1                    int
	ProcessId                 int
	ProcessCreateTime         int
	ProcessUserTime           int
	ProcessKernelTime         int
	ProcessorMaxMhz           int
	ProcessorCurrentMhz       int
	ProcessorMhzLimit         int
	ProcessorMaxIdleState     int
	ProcessorCurrentIdleState int
	ProcessIntegrityLevel     int
	ProcessExecuteFlags       int
	ProtectedProcess          int
	TimeZoneId                int
	TimeZone                  int
	BuildString               int
	DbgBldStr                 int
	XStateData                int
	ProcessCookie             int
}

type MINIDUMP_MEMORY_INFO struct {
	BaseAddress       int
	AllocationBase    int
	AllocationProtect int
	ALIGNMENT1        int
	RegionSize        int
	State             int
	Protect           int
	Type              int
	ALIGNMENT2        int
}

type MINIDUMP_MEMORY_INFO_LIST struct {
	SizeOfHeader    int
	SizeOfEntry     int
	NumberOfEntries int
}

type MINIDUMP_THREAD_NAME struct {
	ThreadId        int
	RvaOfThreadName int
}

type MINIDUMP_THREAD_NAME_LIST struct {
	NumberOfThreadNames int
	ThreadNames         int
}

type MINIDUMP_THREAD_INFO struct {
	ThreadId     int
	DumpFlags    int
	DumpError    int
	ExitStatus   int
	CreateTime   int
	ExitTime     int
	KernelTime   int
	UserTime     int
	StartAddress int
	Affinity     int
}

type MINIDUMP_THREAD_INFO_LIST struct {
	SizeOfHeader    int
	SizeOfEntry     int
	NumberOfEntries int
}

type MINIDUMP_TOKEN_INFO_HEADER struct {
	TokenSize   int
	TokenId     int
	TokenHandle int
}

type MINIDUMP_TOKEN_INFO_LIST struct {
	TokenListSize     int
	TokenListEntries  int
	ListHeaderSize    int
	ElementHeaderSize int
}

type MINIDUMP_SYSTEM_BASIC_INFORMATION struct {
	TimerResolution              int
	PageSize                     int
	NumberOfPhysicalPages        int
	LowestPhysicalPageNumber     int
	HighestPhysicalPageNumber    int
	AllocationGranularity        int
	MinimumUserModeAddress       int
	MaximumUserModeAddress       int
	ActiveProcessorsAffinityMask int
	NumberOfProcessors           int
}

type MINIDUMP_SYSTEM_FILECACHE_INFORMATION struct {
	CurrentSize                           int
	PeakSize                              int
	PageFaultCount                        int
	MinimumWorkingSet                     int
	MaximumWorkingSet                     int
	CurrentSizeIncludingTransitionInPages int
	PeakSizeIncludingTransitionInPages    int
	TransitionRePurposeCount              int
	Flags                                 int
}

type MINIDUMP_SYSTEM_BASIC_PERFORMANCE_INFORMATION struct {
	AvailablePages int
	CommittedPages int
	CommitLimit    int
	PeakCommitment int
}

type MINIDUMP_SYSTEM_PERFORMANCE_INFORMATION struct {
	IdleProcessTime           int
	IoReadTransferCount       int
	IoWriteTransferCount      int
	IoOtherTransferCount      int
	IoReadOperationCount      int
	IoWriteOperationCount     int
	IoOtherOperationCount     int
	AvailablePages            int
	CommittedPages            int
	CommitLimit               int
	PeakCommitment            int
	PageFaultCount            int
	CopyOnWriteCount          int
	TransitionCount           int
	CacheTransitionCount      int
	DemandZeroCount           int
	PageReadCount             int
	PageReadIoCount           int
	CacheReadCount            int
	CacheIoCount              int
	DirtyPagesWriteCount      int
	DirtyWriteIoCount         int
	MappedPagesWriteCount     int
	MappedWriteIoCount        int
	PagedPoolPages            int
	NonPagedPoolPages         int
	PagedPoolAllocs           int
	PagedPoolFrees            int
	NonPagedPoolAllocs        int
	NonPagedPoolFrees         int
	FreeSystemPtes            int
	ResidentSystemCodePage    int
	TotalSystemDriverPages    int
	TotalSystemCodePages      int
	NonPagedPoolLookasideHits int
	PagedPoolLookasideHits    int
	AvailablePagedPoolPages   int
	ResidentSystemCachePage   int
	ResidentPagedPoolPage     int
	ResidentSystemDriverPage  int
	CcFastReadNoWait          int
	CcFastReadWait            int
	CcFastReadResourceMiss    int
	CcFastReadNotPossible     int
	CcFastMdlReadNoWait       int
	CcFastMdlReadWait         int
	CcFastMdlReadResourceMiss int
	CcFastMdlReadNotPossible  int
	CcMapDataNoWait           int
	CcMapDataWait             int
	CcMapDataNoWaitMiss       int
	CcMapDataWaitMiss         int
	CcPinMappedDataCount      int
	CcPinReadNoWait           int
	CcPinReadWait             int
	CcPinReadNoWaitMiss       int
	CcPinReadWaitMiss         int
	CcCopyReadNoWait          int
	CcCopyReadWait            int
	CcCopyReadNoWaitMiss      int
	CcCopyReadWaitMiss        int
	CcMdlReadNoWait           int
	CcMdlReadWait             int
	CcMdlReadNoWaitMiss       int
	CcMdlReadWaitMiss         int
	CcReadAheadIos            int
	CcLazyWriteIos            int
	CcLazyWritePages          int
	CcDataFlushes             int
	CcDataPages               int
	ContextSwitches           int
	FirstLevelTbFills         int
	SecondLevelTbFills        int
	SystemCalls               int
	CcTotalDirtyPages         int
	CcDirtyPageThreshold      int
	ResidentAvailablePages    int
	SharedCommittedPages      int
}

type MINIDUMP_SYSTEM_MEMORY_INFO_1 struct {
	Revision      int
	Flags         int
	BasicInfo     int
	FileCacheInfo int
	BasicPerfInfo int
	PerfInfo      int
}

type MINIDUMP_PROCESS_VM_COUNTERS_1 struct {
	Revision                   int
	PageFaultCount             int
	PeakWorkingSetSize         int
	WorkingSetSize             int
	QuotaPeakPagedPoolUsage    int
	QuotaPagedPoolUsage        int
	QuotaPeakNonPagedPoolUsage int
	QuotaNonPagedPoolUsage     int
	PagefileUsage              int
	PeakPagefileUsage          int
	PrivateUsage               int
}

type MINIDUMP_PROCESS_VM_COUNTERS_2 struct {
	Revision                   int
	Flags                      int
	PageFaultCount             int
	PeakWorkingSetSize         int
	WorkingSetSize             int
	QuotaPeakPagedPoolUsage    int
	QuotaPagedPoolUsage        int
	QuotaPeakNonPagedPoolUsage int
	QuotaNonPagedPoolUsage     int
	PagefileUsage              int
	PeakPagefileUsage          int
	PeakVirtualSize            int
	VirtualSize                int
	PrivateUsage               int
	PrivateWorkingSetSize      int
	SharedCommitUsage          int
	JobSharedCommitUsage       int
	JobPrivateCommitUsage      int
	JobPeakPrivateCommitUsage  int
	JobPrivateCommitLimit      int
	JobTotalCommitLimit        int
}

type MINIDUMP_USER_RECORD struct {
	Type   int
	Memory int
}

type MINIDUMP_USER_STREAM struct {
	Type       int
	BufferSize int
	Buffer     int
}

type MINIDUMP_USER_STREAM_INFORMATION struct {
	UserStreamCount int
	UserStreamArray int
}

type MINIDUMP_THREAD_CALLBACK struct {
	ThreadId      int
	ThreadHandle  int
	Context       int
	SizeOfContext int
	StackBase     int
	StackEnd      int
}

type MINIDUMP_THREAD_EX_CALLBACK struct {
	ThreadId         int
	ThreadHandle     int
	Context          int
	SizeOfContext    int
	StackBase        int
	StackEnd         int
	BackingStoreBase int
	BackingStoreEnd  int
}

type MINIDUMP_INCLUDE_THREAD_CALLBACK struct {
	ThreadId int
}

type MINIDUMP_MODULE_CALLBACK struct {
	FullPath         int
	BaseOfImage      int
	SizeOfImage      int
	CheckSum         int
	TimeDateStamp    int
	VersionInfo      int
	CvRecord         int
	SizeOfCvRecord   int
	MiscRecord       int
	SizeOfMiscRecord int
}

type MINIDUMP_INCLUDE_MODULE_CALLBACK struct {
	BaseOfImage int
}

type MINIDUMP_IO_CALLBACK struct {
	Handle      int
	Offset      int
	Buffer      int
	BufferBytes int
}

type MINIDUMP_READ_MEMORY_FAILURE_CALLBACK struct {
	Offset        int
	Bytes         int
	FailureStatus int
}

type MINIDUMP_VM_QUERY_CALLBACK struct {
	Offset int
}

type MINIDUMP_VM_PRE_READ_CALLBACK struct {
	Offset int
	Buffer int
	Size   int
}

type MINIDUMP_VM_POST_READ_CALLBACK struct {
	Offset    int
	Buffer    int
	Size      int
	Completed int
	Status    int
}

type MINIDUMP_CALLBACK_INPUT struct {
	ProcessId     int
	ProcessHandle int
	CallbackType  int
	Anonymous     int
}

type MINIDUMP_CALLBACK_OUTPUT struct {
	Anonymous int
}

type MINIDUMP_CALLBACK_INFORMATION struct {
	CallbackRoutine int
	CallbackParam   int
}

type ProcessDebugManager struct {
}

type DebugHelper struct {
}

type CDebugDocumentHelper struct {
}

type MachineDebugManager_RETAIL struct {
}

type MachineDebugManager_DEBUG struct {
}

type DefaultDebugSessionProvider struct {
}

type DebugPropertyInfo struct {
	M_DWVALIDFIELDS int
	M_BSTRNAME      int
	M_BSTRTYPE      int
	M_BSTRVALUE     int
	M_BSTRFULLNAME  int
	M_DWATTRIB      int
	M_PDEBUGPROP    int
}

type ExtendedDebugPropertyInfo struct {
	DWVALIDFIELDS int
	PSZNAME       int
	PSZTYPE       int
	PSZVALUE      int
	PSZFULLNAME   int
	DWATTRIB      int
	PDEBUGPROP    int
	NDISPID       int
	NTYPE         int
	VARVALUE      int
	PLBVALUE      int
	PDEBUGEXTPROP int
}

type DebugStackFrameDescriptor struct {
	PDSF      int
	DWMIN     int
	DWLIM     int
	FFINAL    int
	PUNKFINAL int
}

type DebugStackFrameDescriptor64 struct {
	PDSF      int
	DWMIN     int
	DWLIM     int
	FFINAL    int
	PUNKFINAL int
}

type PROFILER_HEAP_OBJECT_SCOPE_LIST struct {
	COUNT  int
	SCOPES int
}

type PROFILER_PROPERTY_TYPE_SUBSTRING_INFO struct {
	LENGTH int
	VALUE  int
}

type PROFILER_HEAP_OBJECT_RELATIONSHIP struct {
	RELATIONSHIPID   int
	RELATIONSHIPINFO int
	Anonymous        int
}

type PROFILER_HEAP_OBJECT_RELATIONSHIP_LIST struct {
	COUNT    int
	ELEMENTS int
}

type PROFILER_HEAP_OBJECT_OPTIONAL_INFO struct {
	INFOTYPE  int
	Anonymous int
}

type PROFILER_HEAP_OBJECT struct {
	SIZE              int
	Anonymous         int
	TYPENAMEID        int
	FLAGS             int
	UNUSED            int
	OPTIONALINFOCOUNT int
}

type PROFILER_HEAP_SUMMARY struct {
	VERSION       int
	TOTALHEAPSIZE int
}

type HTMLCSSStyleDeclaration struct {
}

type HTMLStyle struct {
}

type HTMLRuleStyle struct {
}

type HTMLCSSRule struct {
}

type HTMLCSSImportRule struct {
}

type HTMLCSSMediaRule struct {
}

type HTMLCSSMediaList struct {
}

type HTMLCSSNamespaceRule struct {
}

type HTMLMSCSSKeyframeRule struct {
}

type HTMLMSCSSKeyframesRule struct {
}

type HTMLRenderStyle struct {
}

type HTMLCurrentStyle struct {
}

type HTMLDOMAttribute struct {
}

type HTMLDOMTextNode struct {
}

type HTMLDOMImplementation struct {
}

type HTMLAttributeCollection struct {
}

type StaticNodeList struct {
}

type DOMChildrenCollection struct {
}

type HTMLDefaults struct {
}

type HTCDefaultDispatch struct {
}

type HTCPropertyBehavior struct {
}

type HTCMethodBehavior struct {
}

type HTCEventBehavior struct {
}

type HTCAttachBehavior struct {
}

type HTCDescBehavior struct {
}

type HTMLUrnCollection struct {
}

type HTMLGenericElement struct {
}

type HTMLStyleSheetRule struct {
}

type HTMLStyleSheetRulesCollection struct {
}

type HTMLStyleSheetPage struct {
}

type HTMLStyleSheetPagesCollection struct {
}

type HTMLStyleSheet struct {
}

type HTMLStyleSheetsCollection struct {
}

type HTMLLinkElement struct {
}

type HTMLDOMRange struct {
}

type HTMLFormElement struct {
}

type HTMLTextElement struct {
}

type HTMLImg struct {
}

type HTMLImageElementFactory struct {
}

type HTMLBody struct {
}

type HTMLFontElement struct {
}

type HTMLAnchorElement struct {
}

type HTMLLabelElement struct {
}

type HTMLListElement struct {
}

type HTMLUListElement struct {
}

type HTMLOListElement struct {
}

type HTMLLIElement struct {
}

type HTMLBlockElement struct {
}

type HTMLDivElement struct {
}

type HTMLDDElement struct {
}

type HTMLDTElement struct {
}

type HTMLBRElement struct {
}

type HTMLDListElement struct {
}

type HTMLHRElement struct {
}

type HTMLParaElement struct {
}

type HTMLElementCollection struct {
}

type HTMLHeaderElement struct {
}

type HTMLSelectElement struct {
}

type HTMLWndSelectElement struct {
}

type HTMLOptionElement struct {
}

type HTMLOptionElementFactory struct {
}

type HTMLWndOptionElement struct {
}

type HTMLInputElement struct {
}

type HTMLTextAreaElement struct {
}

type HTMLRichtextElement struct {
}

type HTMLButtonElement struct {
}

type HTMLMarqueeElement struct {
}

type HTMLHtmlElement struct {
}

type HTMLHeadElement struct {
}

type HTMLTitleElement struct {
}

type HTMLMetaElement struct {
}

type HTMLBaseElement struct {
}

type HTMLIsIndexElement struct {
}

type HTMLNextIdElement struct {
}

type HTMLBaseFontElement struct {
}

type HTMLUnknownElement struct {
}

type HTMLHistory struct {
}

type COpsProfile struct {
}

type HTMLNavigator struct {
}

type HTMLLocation struct {
}

type CMimeTypes struct {
}

type CPlugins struct {
}

type CEventObj struct {
}

type HTMLStyleMedia struct {
}

type FramesCollection struct {
}

type HTMLScreen struct {
}

type HTMLWindow2 struct {
}

type HTMLWindowProxy struct {
}

type HTMLDocumentCompatibleInfo struct {
}

type HTMLDocumentCompatibleInfoCollection struct {
}

type HTMLDocument struct {
}

type Scriptlet struct {
}

type HTMLEmbed struct {
}

type HTMLAreasCollection struct {
}

type HTMLMapElement struct {
}

type HTMLAreaElement struct {
}

type HTMLTableCaption struct {
}

type HTMLCommentElement struct {
}

type HTMLPhraseElement struct {
}

type HTMLSpanElement struct {
}

type HTMLTable struct {
}

type HTMLTableCol struct {
}

type HTMLTableSection struct {
}

type HTMLTableRow struct {
}

type HTMLTableCell struct {
}

type HTMLScriptElement struct {
}

type HTMLNoShowElement struct {
}

type HTMLObjectElement struct {
}

type HTMLParamElement struct {
}

type HTMLFrameBase struct {
}

type HTMLFrameElement struct {
}

type HTMLIFrame struct {
}

type HTMLDivPosition struct {
}

type HTMLFieldSetElement struct {
}

type HTMLLegendElement struct {
}

type HTMLSpanFlow struct {
}

type HTMLFrameSetSite struct {
}

type HTMLBGsound struct {
}

type HTMLStyleElement struct {
}

type HTMLStyleFontFace struct {
}

type XDomainRequest struct {
}

type XDomainRequestFactory struct {
}

type HTMLStorage struct {
}

type DOMEvent struct {
}

type DOMUIEvent struct {
}

type DOMMouseEvent struct {
}

type DOMDragEvent struct {
}

type DOMMouseWheelEvent struct {
}

type DOMWheelEvent struct {
}

type DOMTextEvent struct {
}

type DOMKeyboardEvent struct {
}

type DOMCompositionEvent struct {
}

type DOMMutationEvent struct {
}

type DOMBeforeUnloadEvent struct {
}

type DOMFocusEvent struct {
}

type DOMCustomEvent struct {
}

type CanvasGradient struct {
}

type CanvasPattern struct {
}

type CanvasTextMetrics struct {
}

type CanvasImageData struct {
}

type CanvasRenderingContext2D struct {
}

type HTMLCanvasElement struct {
}

type DOMProgressEvent struct {
}

type DOMMessageEvent struct {
}

type DOMSiteModeEvent struct {
}

type DOMStorageEvent struct {
}

type XMLHttpRequestEventTarget struct {
}

type HTMLXMLHttpRequest struct {
}

type HTMLXMLHttpRequestFactory struct {
}

type SVGAngle struct {
}

type SVGAnimatedAngle struct {
}

type SVGAnimatedTransformList struct {
}

type SVGAnimatedBoolean struct {
}

type SVGAnimatedEnumeration struct {
}

type SVGAnimatedInteger struct {
}

type SVGAnimatedLength struct {
}

type SVGAnimatedLengthList struct {
}

type SVGAnimatedNumber struct {
}

type SVGAnimatedNumberList struct {
}

type SVGAnimatedRect struct {
}

type SVGAnimatedString struct {
}

type SVGClipPathElement struct {
}

type SVGElement struct {
}

type SVGLength struct {
}

type SVGLengthList struct {
}

type SVGMatrix struct {
}

type SVGNumber struct {
}

type SVGNumberList struct {
}

type SVGPatternElement struct {
}

type SVGPathSeg struct {
}

type SVGPathSegArcAbs struct {
}

type SVGPathSegArcRel struct {
}

type SVGPathSegClosePath struct {
}

type SVGPathSegMovetoAbs struct {
}

type SVGPathSegMovetoRel struct {
}

type SVGPathSegLinetoAbs struct {
}

type SVGPathSegLinetoRel struct {
}

type SVGPathSegCurvetoCubicAbs struct {
}

type SVGPathSegCurvetoCubicRel struct {
}

type SVGPathSegCurvetoCubicSmoothAbs struct {
}

type SVGPathSegCurvetoCubicSmoothRel struct {
}

type SVGPathSegCurvetoQuadraticAbs struct {
}

type SVGPathSegCurvetoQuadraticRel struct {
}

type SVGPathSegCurvetoQuadraticSmoothAbs struct {
}

type SVGPathSegCurvetoQuadraticSmoothRel struct {
}

type SVGPathSegLinetoHorizontalAbs struct {
}

type SVGPathSegLinetoHorizontalRel struct {
}

type SVGPathSegLinetoVerticalAbs struct {
}

type SVGPathSegLinetoVerticalRel struct {
}

type SVGPathSegList struct {
}

type SVGPoint struct {
}

type SVGPointList struct {
}

type SVGRect struct {
}

type SVGStringList struct {
}

type SVGTransform struct {
}

type SVGSVGElement struct {
}

type SVGUseElement struct {
}

type HTMLStyleSheetRulesAppliedCollection struct {
}

type RulesApplied struct {
}

type RulesAppliedCollection struct {
}

type HTMLW3CComputedStyle struct {
}

type SVGTransformList struct {
}

type SVGCircleElement struct {
}

type SVGEllipseElement struct {
}

type SVGLineElement struct {
}

type SVGRectElement struct {
}

type SVGPolygonElement struct {
}

type SVGPolylineElement struct {
}

type SVGGElement struct {
}

type SVGSymbolElement struct {
}

type SVGDefsElement struct {
}

type SVGPathElement struct {
}

type SVGPreserveAspectRatio struct {
}

type SVGTextElement struct {
}

type SVGAnimatedPreserveAspectRatio struct {
}

type SVGImageElement struct {
}

type SVGStopElement struct {
}

type SVGGradientElement struct {
}

type SVGLinearGradientElement struct {
}

type SVGRadialGradientElement struct {
}

type SVGMaskElement struct {
}

type SVGMarkerElement struct {
}

type SVGZoomEvent struct {
}

type SVGAElement struct {
}

type SVGViewElement struct {
}

type HTMLMediaError struct {
}

type HTMLTimeRanges struct {
}

type HTMLMediaElement struct {
}

type HTMLSourceElement struct {
}

type HTMLAudioElement struct {
}

type HTMLAudioElementFactory struct {
}

type HTMLVideoElement struct {
}

type SVGSwitchElement struct {
}

type SVGDescElement struct {
}

type SVGTitleElement struct {
}

type SVGMetadataElement struct {
}

type SVGElementInstance struct {
}

type SVGElementInstanceList struct {
}

type DOMException struct {
}

type RangeException struct {
}

type SVGException struct {
}

type EventException struct {
}

type SVGScriptElement struct {
}

type SVGStyleElement struct {
}

type SVGTextContentElement struct {
}

type SVGTextPositioningElement struct {
}

type DOMDocumentType struct {
}

type NodeIterator struct {
}

type TreeWalker struct {
}

type DOMProcessingInstruction struct {
}

type HTMLPerformance struct {
}

type HTMLPerformanceNavigation struct {
}

type HTMLPerformanceTiming struct {
}

type SVGTSpanElement struct {
}

type CTemplatePrinter struct {
}

type CPrintManagerTemplatePrinter struct {
}

type SVGTextPathElement struct {
}

type XMLSerializer struct {
}

type DOMParser struct {
}

type HTMLDOMXmlSerializerFactory struct {
}

type DOMParserFactory struct {
}

type HTMLSemanticElement struct {
}

type HTMLProgressElement struct {
}

type DOMMSTransitionEvent struct {
}

type DOMMSAnimationEvent struct {
}

type WebGeolocation struct {
}

type WebGeocoordinates struct {
}

type WebGeopositionError struct {
}

type WebGeoposition struct {
}

type CClientCaps struct {
}

type DOMMSManipulationEvent struct {
}

type DOMCloseEvent struct {
}

type ApplicationCache struct {
}

type HtmlDlgSafeHelper struct {
}

type BlockFormats struct {
}

type FontNames struct {
}

type HTMLNamespace struct {
}

type HTMLNamespaceCollection struct {
}

type ThreadDialogProcParam struct {
}

type HTMLDialog struct {
}

type HTMLPopup struct {
}

type HTMLAppBehavior struct {
}

type OldHTMLDocument struct {
}

type OldHTMLFormElement struct {
}

type HTMLInputButtonElement struct {
}

type HTMLInputTextElement struct {
}

type HTMLInputFileElement struct {
}

type HTMLOptionButtonElement struct {
}

type HTMLInputImage struct {
}

type HTML_PAINTER_INFO struct {
	LFLAGS        int
	LZORDER       int
	IIDDRAWOBJECT int
	RCEXPAND      int
}

type HTML_PAINT_XFORM struct {
	EM11 int
	EM12 int
	EM21 int
	EM22 int
	EDX  int
	EDY  int
}

type HTML_PAINT_DRAW_INFO struct {
	RCVIEWPORT int
	HRGNUPDATE int
	XFORM      int
}
