// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package networklistmanager implements the Windows.Win32.NetworkListManager namespace.
package networklistmanager

type NetworkListManager struct {
}

type NLM_USAGE_DATA struct {
	UsageInMegabytes int
	LastSyncTime     int
}

type NLM_DATAPLAN_STATUS struct {
	InterfaceGuid              int
	UsageData                  int
	DataLimitInMegabytes       int
	InboundBandwidthInKbps     int
	OutboundBandwidthInKbps    int
	NextBillingCycle           int
	MaxTransferSizeInMegabytes int
	Reserved                   int
}

type NLM_SOCKADDR struct {
	DATA int
}

type NLM_SIMULATED_PROFILE_INFO struct {
	ProfileName          int
	COST                 int
	UsageInMegabytes     int
	DataLimitInMegabytes int
}

type NET_INTERFACE_CONTEXT struct {
	InterfaceIndex    int
	ConfigurationName int
}

type NET_INTERFACE_CONTEXT_TABLE struct {
	InterfaceContextHandle int
	NumberOfEntries        int
	InterfaceContextArray  int
}