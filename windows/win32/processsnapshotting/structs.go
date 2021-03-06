// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package processsnapshotting implements the Windows.Win32.ProcessSnapshotting namespace.
package processsnapshotting

type HPSS__ struct {
	UNUSED int
}

type HPSSWALK__ struct {
	UNUSED int
}

type PSS_PROCESS_INFORMATION struct {
	ExitStatus                 int
	PebBaseAddress             int
	AffinityMask               int
	BasePriority               int
	ProcessId                  int
	ParentProcessId            int
	Flags                      int
	CreateTime                 int
	ExitTime                   int
	KernelTime                 int
	UserTime                   int
	PriorityClass              int
	PeakVirtualSize            int
	VirtualSize                int
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
	ExecuteFlags               int
	ImageFileName              int
}

type PSS_VA_CLONE_INFORMATION struct {
	VaCloneHandle int
}

type PSS_AUXILIARY_PAGES_INFORMATION struct {
	AuxPagesCaptured int
}

type PSS_VA_SPACE_INFORMATION struct {
	RegionCount int
}

type PSS_HANDLE_INFORMATION struct {
	HandlesCaptured int
}

type PSS_THREAD_INFORMATION struct {
	ThreadsCaptured int
	ContextLength   int
}

type PSS_HANDLE_TRACE_INFORMATION struct {
	SectionHandle int
	Size          int
}

type PSS_PERFORMANCE_COUNTERS struct {
	TotalCycleCount         int
	TotalWallClockPeriod    int
	VaCloneCycleCount       int
	VaCloneWallClockPeriod  int
	VaSpaceCycleCount       int
	VaSpaceWallClockPeriod  int
	AuxPagesCycleCount      int
	AuxPagesWallClockPeriod int
	HandlesCycleCount       int
	HandlesWallClockPeriod  int
	ThreadsCycleCount       int
	ThreadsWallClockPeriod  int
}

type PSS_AUXILIARY_PAGE_ENTRY struct {
	Address          int
	BasicInformation int
	CaptureTime      int
	PageContents     int
	PageSize         int
}

type PSS_VA_SPACE_ENTRY struct {
	BaseAddress          int
	AllocationBase       int
	AllocationProtect    int
	RegionSize           int
	State                int
	Protect              int
	Type                 int
	TimeDateStamp        int
	SizeOfImage          int
	ImageBase            int
	CheckSum             int
	MappedFileNameLength int
	MappedFileName       int
}

type PSS_HANDLE_ENTRY struct {
	Handle                  int
	Flags                   int
	ObjectType              int
	CaptureTime             int
	Attributes              int
	GrantedAccess           int
	HandleCount             int
	PointerCount            int
	PagedPoolCharge         int
	NonPagedPoolCharge      int
	CreationTime            int
	TypeNameLength          int
	TypeName                int
	ObjectNameLength        int
	ObjectName              int
	TypeSpecificInformation int
}

type PSS_THREAD_ENTRY struct {
	ExitStatus               int
	TebBaseAddress           int
	ProcessId                int
	ThreadId                 int
	AffinityMask             int
	Priority                 int
	BasePriority             int
	LastSyscallFirstArgument int
	LastSyscallNumber        int
	CreateTime               int
	ExitTime                 int
	KernelTime               int
	UserTime                 int
	Win32StartAddress        int
	CaptureTime              int
	Flags                    int
	SuspendCount             int
	SizeOfContextRecord      int
	ContextRecord            int
}

type PSS_ALLOCATOR struct {
	Context      int
	AllocRoutine int
	FreeRoutine  int
}
