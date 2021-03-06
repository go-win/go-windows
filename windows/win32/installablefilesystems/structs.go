// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package installablefilesystems implements the Windows.Win32.InstallableFileSystems namespace.
package installablefilesystems

type FILTER_FULL_INFORMATION struct {
	NextEntryOffset   int
	FrameID           int
	NumberOfInstances int
	FilterNameLength  int
	FilterNameBuffer  int
}

type FILTER_AGGREGATE_BASIC_INFORMATION struct {
	NextEntryOffset int
	Flags           int
	Type            int
}

type FILTER_AGGREGATE_STANDARD_INFORMATION struct {
	NextEntryOffset int
	Flags           int
	Type            int
}

type FILTER_VOLUME_BASIC_INFORMATION struct {
	FilterVolumeNameLength int
	FilterVolumeName       int
}

type FILTER_VOLUME_STANDARD_INFORMATION struct {
	NextEntryOffset        int
	Flags                  int
	FrameID                int
	FileSystemType         int
	FilterVolumeNameLength int
	FilterVolumeName       int
}

type INSTANCE_BASIC_INFORMATION struct {
	NextEntryOffset          int
	InstanceNameLength       int
	InstanceNameBufferOffset int
}

type INSTANCE_PARTIAL_INFORMATION struct {
	NextEntryOffset          int
	InstanceNameLength       int
	InstanceNameBufferOffset int
	AltitudeLength           int
	AltitudeBufferOffset     int
}

type INSTANCE_FULL_INFORMATION struct {
	NextEntryOffset          int
	InstanceNameLength       int
	InstanceNameBufferOffset int
	AltitudeLength           int
	AltitudeBufferOffset     int
	VolumeNameLength         int
	VolumeNameBufferOffset   int
	FilterNameLength         int
	FilterNameBufferOffset   int
}

type INSTANCE_AGGREGATE_STANDARD_INFORMATION struct {
	NextEntryOffset int
	Flags           int
	Type            int
}

type FILTER_MESSAGE_HEADER struct {
	ReplyLength int
	MessageId   int
}

type FILTER_REPLY_HEADER struct {
	Status    int
	MessageId int
}

type FilterFindHandle struct {
	Value int
}

type FilterInstanceFindHandle struct {
	Value int
}

type FilterVolumeFindHandle struct {
	Value int
}

type FilterVolumeInstanceFindHandle struct {
	Value int
}

type HFILTER struct {
	Value int
}

type HFILTER_INSTANCE struct {
	Value int
}
